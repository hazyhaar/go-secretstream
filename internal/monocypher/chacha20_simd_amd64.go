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
// en micro-entrelacement strict au sein des 10 double-rounds ChaCha20.
func aead_interleaved_write_simd(ctx *Crypto_aead_ctx, polyCtx *Crypto_poly1305_ctx, cipher_text, plain_text []byte, numChunks uint64, startCtr uint64) uint64 {
	r0 := uint64(polyCtx.R[0]) | (uint64(polyCtx.R[1]) << 32)
	r1 := uint64(polyCtx.R[2]) | (uint64(polyCtx.R[3]) << 32)

	h0 := uint64(polyCtx.H[0]) | (uint64(polyCtx.H[1]) << 32)
	h1 := uint64(polyCtx.H[2]) | (uint64(polyCtx.H[3]) << 32)
	h2 := uint64(polyCtx.H[4])

	k0 := binary.LittleEndian.Uint32(ctx.Key[0:4])
	k1 := binary.LittleEndian.Uint32(ctx.Key[4:8])
	k2 := binary.LittleEndian.Uint32(ctx.Key[8:12])
	k3 := binary.LittleEndian.Uint32(ctx.Key[12:16])
	k4 := binary.LittleEndian.Uint32(ctx.Key[16:20])
	k5 := binary.LittleEndian.Uint32(ctx.Key[20:24])
	k6 := binary.LittleEndian.Uint32(ctx.Key[24:28])
	k7 := binary.LittleEndian.Uint32(ctx.Key[28:32])

	n0 := binary.LittleEndian.Uint32(ctx.Nonce[0:4])
	n1 := binary.LittleEndian.Uint32(ctx.Nonce[4:8])

	st0 := archsimd.LoadUint32x8Array(&[8]uint32{c0_djb, c1_djb, c2_djb, c3_djb, c0_djb, c1_djb, c2_djb, c3_djb})
	st1 := archsimd.LoadUint32x8Array(&[8]uint32{k0, k1, k2, k3, k0, k1, k2, k3})
	st2 := archsimd.LoadUint32x8Array(&[8]uint32{k4, k5, k6, k7, k4, k5, k6, k7})

	currCtr := startCtr
	st3_A, st3_B := reloadSt3(currCtr, n0, n1)

	// Chiffrement du 1er chunk (Chunk 0)
	{
		chunk0PT := (*[256]byte)(plain_text[0:256])
		chunk0CT := (*[256]byte)(cipher_text[0:256])

		v0_A, v1_A, v2_A, v3_A := st0, st1, st2, st3_A
		v0_B, v1_B, v2_B, v3_B := st0, st1, st2, st3_B

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

		// Store Chunk 0
		k0_0, k0_1, k0_2, k0_3 := v0_A.GetLo().AsUint8x16(), v1_A.GetLo().AsUint8x16(), v2_A.GetLo().AsUint8x16(), v3_A.GetLo().AsUint8x16()
		archsimd.LoadUint8x16(chunk0PT[0:16]).Xor(k0_0).Store(chunk0CT[0:16])
		archsimd.LoadUint8x16(chunk0PT[16:32]).Xor(k0_1).Store(chunk0CT[16:32])
		archsimd.LoadUint8x16(chunk0PT[32:48]).Xor(k0_2).Store(chunk0CT[32:48])
		archsimd.LoadUint8x16(chunk0PT[48:64]).Xor(k0_3).Store(chunk0CT[48:64])

		k1_0, k1_1, k1_2, k1_3 := v0_A.GetHi().AsUint8x16(), v1_A.GetHi().AsUint8x16(), v2_A.GetHi().AsUint8x16(), v3_A.GetHi().AsUint8x16()
		archsimd.LoadUint8x16(chunk0PT[64:80]).Xor(k1_0).Store(chunk0CT[64:80])
		archsimd.LoadUint8x16(chunk0PT[80:96]).Xor(k1_1).Store(chunk0CT[80:96])
		archsimd.LoadUint8x16(chunk0PT[96:112]).Xor(k1_2).Store(chunk0CT[96:112])
		archsimd.LoadUint8x16(chunk0PT[112:128]).Xor(k1_3).Store(chunk0CT[112:128])

		k2_0, k2_1, k2_2, k2_3 := v0_B.GetLo().AsUint8x16(), v1_B.GetLo().AsUint8x16(), v2_B.GetLo().AsUint8x16(), v3_B.GetLo().AsUint8x16()
		archsimd.LoadUint8x16(chunk0PT[128:144]).Xor(k2_0).Store(chunk0CT[128:144])
		archsimd.LoadUint8x16(chunk0PT[144:160]).Xor(k2_1).Store(chunk0CT[144:160])
		archsimd.LoadUint8x16(chunk0PT[160:176]).Xor(k2_2).Store(chunk0CT[160:176])
		archsimd.LoadUint8x16(chunk0PT[176:192]).Xor(k2_3).Store(chunk0CT[176:192])

		k3_0, k3_1, k3_2, k3_3 := v0_B.GetHi().AsUint8x16(), v1_B.GetHi().AsUint8x16(), v2_B.GetHi().AsUint8x16(), v3_B.GetHi().AsUint8x16()
		archsimd.LoadUint8x16(chunk0PT[192:208]).Xor(k3_0).Store(chunk0CT[192:208])
		archsimd.LoadUint8x16(chunk0PT[208:224]).Xor(k3_1).Store(chunk0CT[208:224])
		archsimd.LoadUint8x16(chunk0PT[224:240]).Xor(k3_2).Store(chunk0CT[224:240])
		archsimd.LoadUint8x16(chunk0PT[240:256]).Xor(k3_3).Store(chunk0CT[240:256])

		currCtr += 4
		st3_A = st3_A.Add(ctrInc4)
		st3_B = st3_B.Add(ctrInc4)
		if lo := uint32(currCtr); lo < 4 || lo > 0xFFFFFFFC {
			st3_A, st3_B = reloadSt3(currCtr, n0, n1)
		}
	}

	// Boucle principale entrelacée pour les chunks 1 à numChunks-1 :
	for c := uint64(1); c < numChunks; c++ {
		prevOff := int((c - 1) * 256)
		currOff := int(c * 256)

		prevChunk := (*[256]byte)(cipher_text[prevOff : prevOff+256])
		currPT := (*[256]byte)(plain_text[currOff : currOff+256])
		currCT := (*[256]byte)(cipher_text[currOff : currOff+256])

		v0_A, v1_A, v2_A, v3_A := st0, st1, st2, st3_A
		v0_B, v1_B, v2_B, v3_B := st0, st1, st2, st3_B

		// Tour 0 : Poly blocks 0 & 1
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[0:8]), binary.LittleEndian.Uint64(prevChunk[8:16]))
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[16:24]), binary.LittleEndian.Uint64(prevChunk[24:32]))

		// Tour 1 : Poly blocks 2 & 3
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[32:40]), binary.LittleEndian.Uint64(prevChunk[40:48]))
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[48:56]), binary.LittleEndian.Uint64(prevChunk[56:64]))

		// Tour 2 : Poly blocks 4 & 5
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[64:72]), binary.LittleEndian.Uint64(prevChunk[72:80]))
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[80:88]), binary.LittleEndian.Uint64(prevChunk[88:96]))

		// Tour 3 : Poly blocks 6 & 7
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[96:104]), binary.LittleEndian.Uint64(prevChunk[104:112]))
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[112:120]), binary.LittleEndian.Uint64(prevChunk[120:128]))

		// Tour 4 : Poly blocks 8 & 9
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[128:136]), binary.LittleEndian.Uint64(prevChunk[136:144]))
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[144:152]), binary.LittleEndian.Uint64(prevChunk[152:160]))

		// Tour 5 : Poly blocks 10 & 11
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[160:168]), binary.LittleEndian.Uint64(prevChunk[168:176]))
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[176:184]), binary.LittleEndian.Uint64(prevChunk[184:192]))

		// Tour 6 : Poly blocks 12 & 13
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[192:200]), binary.LittleEndian.Uint64(prevChunk[200:208]))
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[208:216]), binary.LittleEndian.Uint64(prevChunk[216:224]))

		// Tour 7 : Poly blocks 14 & 15
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[224:232]), binary.LittleEndian.Uint64(prevChunk[232:240]))
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(prevChunk[240:248]), binary.LittleEndian.Uint64(prevChunk[248:256]))

		// Tour 8 & 9 : Finalisation tours ChaCha20
		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)

		v0_A, v1_A, v2_A, v3_A = chacha20DoubleBlockSIMD256(v0_A, v1_A, v2_A, v3_A)
		v0_B, v1_B, v2_B, v3_B = chacha20DoubleBlockSIMD256(v0_B, v1_B, v2_B, v3_B)

		v0_A = v0_A.Add(st0)
		v1_A = v1_A.Add(st1)
		v2_A = v2_A.Add(st2)
		v3_A = v3_A.Add(st3_A)

		v0_B = v0_B.Add(st0)
		v1_B = v1_B.Add(st1)
		v2_B = v2_B.Add(st2)
		v3_B = v3_B.Add(st3_B)

		// Store Chunk c
		k0_0, k0_1, k0_2, k0_3 := v0_A.GetLo().AsUint8x16(), v1_A.GetLo().AsUint8x16(), v2_A.GetLo().AsUint8x16(), v3_A.GetLo().AsUint8x16()
		archsimd.LoadUint8x16(currPT[0:16]).Xor(k0_0).Store(currCT[0:16])
		archsimd.LoadUint8x16(currPT[16:32]).Xor(k0_1).Store(currCT[16:32])
		archsimd.LoadUint8x16(currPT[32:48]).Xor(k0_2).Store(currCT[32:48])
		archsimd.LoadUint8x16(currPT[48:64]).Xor(k0_3).Store(currCT[48:64])

		k1_0, k1_1, k1_2, k1_3 := v0_A.GetHi().AsUint8x16(), v1_A.GetHi().AsUint8x16(), v2_A.GetHi().AsUint8x16(), v3_A.GetHi().AsUint8x16()
		archsimd.LoadUint8x16(currPT[64:80]).Xor(k1_0).Store(currCT[64:80])
		archsimd.LoadUint8x16(currPT[80:96]).Xor(k1_1).Store(currCT[80:96])
		archsimd.LoadUint8x16(currPT[96:112]).Xor(k1_2).Store(currCT[96:112])
		archsimd.LoadUint8x16(currPT[112:128]).Xor(k1_3).Store(currCT[112:128])

		k2_0, k2_1, k2_2, k2_3 := v0_B.GetLo().AsUint8x16(), v1_B.GetLo().AsUint8x16(), v2_B.GetLo().AsUint8x16(), v3_B.GetLo().AsUint8x16()
		archsimd.LoadUint8x16(currPT[128:144]).Xor(k2_0).Store(currCT[128:144])
		archsimd.LoadUint8x16(currPT[144:160]).Xor(k2_1).Store(currCT[144:160])
		archsimd.LoadUint8x16(currPT[160:176]).Xor(k2_2).Store(currCT[160:176])
		archsimd.LoadUint8x16(currPT[176:192]).Xor(k2_3).Store(currCT[176:192])

		k3_0, k3_1, k3_2, k3_3 := v0_B.GetHi().AsUint8x16(), v1_B.GetHi().AsUint8x16(), v2_B.GetHi().AsUint8x16(), v3_B.GetHi().AsUint8x16()
		archsimd.LoadUint8x16(currPT[192:208]).Xor(k3_0).Store(currCT[192:208])
		archsimd.LoadUint8x16(currPT[208:224]).Xor(k3_1).Store(currCT[208:224])
		archsimd.LoadUint8x16(currPT[224:240]).Xor(k3_2).Store(currCT[224:240])
		archsimd.LoadUint8x16(currPT[240:256]).Xor(k3_3).Store(currCT[240:256])

		currCtr += 4
		st3_A = st3_A.Add(ctrInc4)
		st3_B = st3_B.Add(ctrInc4)
		if lo := uint32(currCtr); lo < 4 || lo > 0xFFFFFFFC {
			st3_A, st3_B = reloadSt3(currCtr, n0, n1)
		}
	}

	// Absorption du dernier chunk (Chunk numChunks-1) dans Poly1305
	lastOff := int((numChunks - 1) * 256)
	lastChunk := (*[256]byte)(cipher_text[lastOff : lastOff+256])
	for b := 0; b < 16; b++ {
		bOff := b * 16
		h0, h1, h2 = polyStep(h0, h1, h2, r0, r1, binary.LittleEndian.Uint64(lastChunk[bOff:bOff+8]), binary.LittleEndian.Uint64(lastChunk[bOff+8:bOff+16]))
	}

	polyCtx.H[0] = uint32(h0)
	polyCtx.H[1] = uint32(h0 >> 32)
	polyCtx.H[2] = uint32(h1)
	polyCtx.H[3] = uint32(h1 >> 32)
	polyCtx.H[4] = uint32(h2)

	return currCtr
}
