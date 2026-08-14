//go:build goexperiment.simd && amd64

package monocypher

import (
	"encoding/binary"
	"math/bits"
	"simd/archsimd"
)

func hasAVX2() bool {
	return archsimd.X86.AVX2()
}

const (
	c0_djb uint32 = 0x61707865 // "expa"
	c1_djb uint32 = 0x3320646e // "nd 3"
	c2_djb uint32 = 0x79622d32 // "2-by"
	c3_djb uint32 = 0x6b206574 // "te k"
)

var (
	// rot16Mask : Permutation d'octets vpshufb pour rotation gauche de 16 bits sur entiers 32 bits
	rot16Mask = archsimd.LoadInt8x32Array(&[32]int8{
		2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
		2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
	})

	// rot8Mask : Permutation d'octets vpshufb pour rotation gauche de 8 bits sur entiers 32 bits
	rot8Mask = archsimd.LoadInt8x32Array(&[32]int8{
		3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
		3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
	})

	// ctrInc4 : Incrément vectoriel pour le compteur de blocs ChaCha20 (4 blocs = +4 sur les lanes basses)
	ctrInc4 = archsimd.LoadUint32x8Array(&[8]uint32{
		4, 0, 0, 0,
		4, 0, 0, 0,
	})
)

// chacha20DoubleBlockSIMD256 effectue 1 double-tour (Column Round + Diagonal Round) sur 2 blocs ChaCha20 en parallèle (256-bit Uint32x8).
// Utilise vpshufb (PermuteOrZeroGrouped) pour les rotations 16 et 8 bits (1 instruction au lieu de 3 sans AVX-512).
func chacha20DoubleBlockSIMD256(v0, v1, v2, v3 archsimd.Uint32x8) (archsimd.Uint32x8, archsimd.Uint32x8, archsimd.Uint32x8, archsimd.Uint32x8) {
	// 1. Column Round
	v0 = v0.Add(v1)
	v3 = v3.Xor(v0).AsUint8x32().PermuteOrZeroGrouped(rot16Mask).AsUint32x8()
	v2 = v2.Add(v3)
	v1 = v1.Xor(v2).RotateAllLeft(12)
	v0 = v0.Add(v1)
	v3 = v3.Xor(v0).AsUint8x32().PermuteOrZeroGrouped(rot8Mask).AsUint32x8()
	v2 = v2.Add(v3)
	v1 = v1.Xor(v2).RotateAllLeft(7)

	// 2. Diagonal Round
	v1 = v1.PermuteScalarsGrouped(1, 2, 3, 0)
	v2 = v2.PermuteScalarsGrouped(2, 3, 0, 1)
	v3 = v3.PermuteScalarsGrouped(3, 0, 1, 2)

	v0 = v0.Add(v1)
	v3 = v3.Xor(v0).AsUint8x32().PermuteOrZeroGrouped(rot16Mask).AsUint32x8()
	v2 = v2.Add(v3)
	v1 = v1.Xor(v2).RotateAllLeft(12)
	v0 = v0.Add(v1)
	v3 = v3.Xor(v0).AsUint8x32().PermuteOrZeroGrouped(rot8Mask).AsUint32x8()
	v2 = v2.Add(v3)
	v1 = v1.Xor(v2).RotateAllLeft(7)

	// Inverse Shuffles
	v1 = v1.PermuteScalarsGrouped(3, 0, 1, 2)
	v2 = v2.PermuteScalarsGrouped(2, 3, 0, 1)
	v3 = v3.PermuteScalarsGrouped(1, 2, 3, 0)

	return v0, v1, v2, v3
}

// reloadSt3 initialise et recharge les vecteurs d'état st3_A et st3_B pour 4 blocs ChaCha20
// avec propagation exacte du compteur 64 bits (lanes 0..3).
func reloadSt3(currCtr uint64, n0, n1 uint32) (archsimd.Uint32x8, archsimd.Uint32x8) {
	st3_A := archsimd.LoadUint32x8Array(&[8]uint32{
		uint32(currCtr), uint32(currCtr >> 32), n0, n1,
		uint32(currCtr + 1), uint32((currCtr + 1) >> 32), n0, n1,
	})
	st3_B := archsimd.LoadUint32x8Array(&[8]uint32{
		uint32(currCtr + 2), uint32((currCtr + 2) >> 32), n0, n1,
		uint32(currCtr + 3), uint32((currCtr + 3) >> 32), n0, n1,
	})
	return st3_A, st3_B
}

// chacha20_djb_simd_4x traite 4 blocs de 64 octets (256 octets) simultanément en SIMD 256 bits (AVX2 archsimd.Uint32x8).
func chacha20_djb_simd_4x(cipher_text, plain_text []byte, key, nonce []byte, ctr uint64, numBlocks uint64) uint64 {
	k0 := binary.LittleEndian.Uint32(key[0:4])
	k1 := binary.LittleEndian.Uint32(key[4:8])
	k2 := binary.LittleEndian.Uint32(key[8:12])
	k3 := binary.LittleEndian.Uint32(key[12:16])
	k4 := binary.LittleEndian.Uint32(key[16:20])
	k5 := binary.LittleEndian.Uint32(key[20:24])
	k6 := binary.LittleEndian.Uint32(key[24:28])
	k7 := binary.LittleEndian.Uint32(key[28:32])

	n0 := binary.LittleEndian.Uint32(nonce[0:4])
	n1 := binary.LittleEndian.Uint32(nonce[4:8])

	st0 := archsimd.LoadUint32x8Array(&[8]uint32{c0_djb, c1_djb, c2_djb, c3_djb, c0_djb, c1_djb, c2_djb, c3_djb})
	st1 := archsimd.LoadUint32x8Array(&[8]uint32{k0, k1, k2, k3, k0, k1, k2, k3})
	st2 := archsimd.LoadUint32x8Array(&[8]uint32{k4, k5, k6, k7, k4, k5, k6, k7})

	currCtr := ctr
	st3_A, st3_B := reloadSt3(currCtr, n0, n1)

	offset := 0

	for bIdx := uint64(0); bIdx+4 <= numBlocks; bIdx += 4 {
		v0_A, v1_A, v2_A, v3_A := st0, st1, st2, st3_A
		v0_B, v1_B, v2_B, v3_B := st0, st1, st2, st3_B

		// 10 doubles-rounds (20 rounds au total) avec vpshufb
		for i := 0; i < 10; i++ {
			v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
			v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		}

		v0_A = v0_A.Add(st0)
		v1_A = v1_A.Add(st1)
		v2_A = v2_A.Add(st2)
		v3_A = v3_A.Add(st3_A)

		v0_B = v0_B.Add(st0)
		v1_B = v1_B.Add(st1)
		v2_B = v2_B.Add(st2)
		v3_B = v3_B.Add(st3_B)

		// Écriture / Masquage XOR vectoriel pour les 4 blocs (256 octets)
		// Bloc 0
		bOff0 := offset
		k0_0 := v0_A.GetLo().AsUint8x16()
		k0_1 := v1_A.GetLo().AsUint8x16()
		k0_2 := v2_A.GetLo().AsUint8x16()
		k0_3 := v3_A.GetLo().AsUint8x16()
		if plain_text != nil {
			archsimd.LoadUint8x16(plain_text[bOff0:bOff0+16]).Xor(k0_0).Store(cipher_text[bOff0 : bOff0+16])
			archsimd.LoadUint8x16(plain_text[bOff0+16:bOff0+32]).Xor(k0_1).Store(cipher_text[bOff0+16 : bOff0+32])
			archsimd.LoadUint8x16(plain_text[bOff0+32:bOff0+48]).Xor(k0_2).Store(cipher_text[bOff0+32 : bOff0+48])
			archsimd.LoadUint8x16(plain_text[bOff0+48:bOff0+64]).Xor(k0_3).Store(cipher_text[bOff0+48 : bOff0+64])
		} else {
			k0_0.Store(cipher_text[bOff0 : bOff0+16])
			k0_1.Store(cipher_text[bOff0+16 : bOff0+32])
			k0_2.Store(cipher_text[bOff0+32 : bOff0+48])
			k0_3.Store(cipher_text[bOff0+48 : bOff0+64])
		}

		// Bloc 1
		bOff1 := offset + 64
		k1_0 := v0_A.GetHi().AsUint8x16()
		k1_1 := v1_A.GetHi().AsUint8x16()
		k1_2 := v2_A.GetHi().AsUint8x16()
		k1_3 := v3_A.GetHi().AsUint8x16()
		if plain_text != nil {
			archsimd.LoadUint8x16(plain_text[bOff1:bOff1+16]).Xor(k1_0).Store(cipher_text[bOff1 : bOff1+16])
			archsimd.LoadUint8x16(plain_text[bOff1+16:bOff1+32]).Xor(k1_1).Store(cipher_text[bOff1+16 : bOff1+32])
			archsimd.LoadUint8x16(plain_text[bOff1+32:bOff1+48]).Xor(k1_2).Store(cipher_text[bOff1+32 : bOff1+48])
			archsimd.LoadUint8x16(plain_text[bOff1+48:bOff1+64]).Xor(k1_3).Store(cipher_text[bOff1+48 : bOff1+64])
		} else {
			k1_0.Store(cipher_text[bOff1 : bOff1+16])
			k1_1.Store(cipher_text[bOff1+16 : bOff1+32])
			k1_2.Store(cipher_text[bOff1+32 : bOff1+48])
			k1_3.Store(cipher_text[bOff1+48 : bOff1+64])
		}

		// Bloc 2
		bOff2 := offset + 128
		k2_0 := v0_B.GetLo().AsUint8x16()
		k2_1 := v1_B.GetLo().AsUint8x16()
		k2_2 := v2_B.GetLo().AsUint8x16()
		k2_3 := v3_B.GetLo().AsUint8x16()
		if plain_text != nil {
			archsimd.LoadUint8x16(plain_text[bOff2:bOff2+16]).Xor(k2_0).Store(cipher_text[bOff2 : bOff2+16])
			archsimd.LoadUint8x16(plain_text[bOff2+16:bOff2+32]).Xor(k2_1).Store(cipher_text[bOff2+16 : bOff2+32])
			archsimd.LoadUint8x16(plain_text[bOff2+32:bOff2+48]).Xor(k2_2).Store(cipher_text[bOff2+32 : bOff2+48])
			archsimd.LoadUint8x16(plain_text[bOff2+48:bOff2+64]).Xor(k2_3).Store(cipher_text[bOff2+48 : bOff2+64])
		} else {
			k2_0.Store(cipher_text[bOff2 : bOff2+16])
			k2_1.Store(cipher_text[bOff2+16 : bOff2+32])
			k2_2.Store(cipher_text[bOff2+32 : bOff2+48])
			k2_3.Store(cipher_text[bOff2+48 : bOff2+64])
		}

		// Bloc 3
		bOff3 := offset + 192
		k3_0 := v0_B.GetHi().AsUint8x16()
		k3_1 := v1_B.GetHi().AsUint8x16()
		k3_2 := v2_B.GetHi().AsUint8x16()
		k3_3 := v3_B.GetHi().AsUint8x16()
		if plain_text != nil {
			archsimd.LoadUint8x16(plain_text[bOff3:bOff3+16]).Xor(k3_0).Store(cipher_text[bOff3 : bOff3+16])
			archsimd.LoadUint8x16(plain_text[bOff3+16:bOff3+32]).Xor(k3_1).Store(cipher_text[bOff3+16 : bOff3+32])
			archsimd.LoadUint8x16(plain_text[bOff3+32:bOff3+48]).Xor(k3_2).Store(cipher_text[bOff3+32 : bOff3+48])
			archsimd.LoadUint8x16(plain_text[bOff3+48:bOff3+64]).Xor(k3_3).Store(cipher_text[bOff3+48 : bOff3+64])
		} else {
			k3_0.Store(cipher_text[bOff3 : bOff3+16])
			k3_1.Store(cipher_text[bOff3+16 : bOff3+32])
			k3_2.Store(cipher_text[bOff3+32 : bOff3+48])
			k3_3.Store(cipher_text[bOff3+48 : bOff3+64])
		}

		offset += 256
		currCtr += 4
		st3_A = st3_A.Add(ctrInc4)
		st3_B = st3_B.Add(ctrInc4)
		if lo := uint32(currCtr); lo < 4 || lo > 0xFFFFFFFC {
			st3_A, st3_B = reloadSt3(currCtr, n0, n1)
		}
	}

	return currCtr
}

// polyStep absorbe 1 bloc de 16 octets dans l'accumulateur Poly1305 (64-bit saturé)
func polyStep(h0, h1, h2, r0, r1, m0, m1 uint64) (uint64, uint64, uint64) {
	var c uint64
	h0, c = bits.Add64(h0, m0, 0)
	h1, c = bits.Add64(h1, m1, c)
	h2 += c + 1

	hi0_0, lo0_0 := bits.Mul64(h0, r0)
	hi1_0, lo1_0 := bits.Mul64(h1, r0)
	hi0_1, lo0_1 := bits.Mul64(h0, r1)
	hi1_1, lo1_1 := bits.Mul64(h1, r1)

	lo2_0 := h2 * r0
	lo2_1 := h2 * r1

	m0_lo, m0_hi := lo0_0, hi0_0
	m1_lo, c1 := bits.Add64(lo1_0, lo0_1, 0)
	m1_hi, _ := bits.Add64(hi1_0, hi0_1, c1)
	m2_lo, c2 := bits.Add64(lo2_0, lo1_1, 0)
	m2_hi, _ := bits.Add64(0, hi1_1, c2)
	m3_lo := lo2_1

	t0 := m0_lo
	t1, c := bits.Add64(m1_lo, m0_hi, 0)
	t2, c := bits.Add64(m2_lo, m1_hi, c)
	t3, _ := bits.Add64(m3_lo, m2_hi, c)

	h0 = t0
	h1 = t1
	h2 = t2 & 3

	cc_lo := t2 & ^uint64(3)
	cc_hi := t3

	h0, c = bits.Add64(h0, cc_lo, 0)
	h1, c = bits.Add64(h1, cc_hi, c)
	h2 += c

	c_lo := (cc_lo >> 2) | (cc_hi << 62)
	c_hi := cc_hi >> 2

	h0, c = bits.Add64(h0, c_lo, 0)
	h1, c = bits.Add64(h1, c_hi, c)
	h2 += c

	return h0, h1, h2
}

// aead_interleaved_write_simd effectue le chiffrement ChaCha20 et le hachage Poly1305
// par chunks de 256 octets pour une localité de cache L1 optimale sans déversement de registres.
func aead_interleaved_write_simd(ctx *Crypto_aead_ctx, polyCtx *Crypto_poly1305_ctx, cipher_text, plain_text []byte, numChunks uint64, startCtr uint64) uint64 {
	currCtr := startCtr
	offset := 0

	for c := uint64(0); c < numChunks; c++ {
		// 1. Chiffrement de 4 blocs ChaCha20 (256 octets)
		currCtr = chacha20_djb_simd_4x(cipher_text[offset:offset+256], plain_text[offset:offset+256], ctx.Key[:], ctx.Nonce[:], currCtr, 4)

		// 2. Absorption des 16 blocs de 16 octets dans Poly1305 (en cache L1)
		Poly_blocks(polyCtx, cipher_text[offset:offset+256], 16, 1)

		offset += 256
	}

	return currCtr
}
