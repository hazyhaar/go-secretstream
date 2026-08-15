//go:build goexperiment.simd && amd64

package monocypher55

import (
	"encoding/binary"
	"simd/archsimd"
)

// rotl32x4 effectue une rotation gauche de n bits sur chaque lane 32-bit d'un
// Uint32x4, émulée par 2 shifts + or (RotateAllLeft exige AVX-512, absent du poste).
func rotl32x4(v archsimd.Uint32x4, n uint64) archsimd.Uint32x4 {
	return v.ShiftAllLeft(n).Or(v.ShiftAllRight(32 - n))
}

var (
	// chachaConstRow128 : ligne 0 constante de l'état ChaCha ("expand 32-byte k")
	chachaConstRow128 = archsimd.LoadUint32x4Array(&[4]uint32{c0_djb, c1_djb, c2_djb, c3_djb})

	// rot16Mask128 : vpshufb 128-bit, rotation gauche de 16 bits par lane 32-bit
	rot16Mask128 = archsimd.LoadInt8x16Array(&[16]int8{
		2, 3, 0, 1, 6, 7, 4, 5, 10, 11, 8, 9, 14, 15, 12, 13,
	})
	// rot8Mask128 : vpshufb 128-bit, rotation gauche de 8 bits par lane 32-bit
	rot8Mask128 = archsimd.LoadInt8x16Array(&[16]int8{
		3, 0, 1, 2, 7, 4, 5, 6, 11, 8, 9, 10, 15, 12, 13, 14,
	})
)

// rotl16x4 et rotl8x4 : rotations 16/8 bits en 1 instruction vpshufb.
func rotl16x4(v archsimd.Uint32x4) archsimd.Uint32x4 {
	return v.AsUint8x16().PermuteOrZero(rot16Mask128).AsUint32x4()
}

func rotl8x4(v archsimd.Uint32x4) archsimd.Uint32x4 {
	return v.AsUint8x16().PermuteOrZero(rot8Mask128).AsUint32x4()
}

// hchacha20_simd128 calcule HChaCha20 (dérivation de sous-clé 256-bit) en SIMD
// 128-bit : l'état 4x4 vit dans 4 registres Uint32x4, la diagonalisation passe
// par PermuteScalars (vpshufd), sans aucun aller-retour mémoire dans la boucle.
// Contrat identique au repli scalaire : key=32o, in=16o, out>=32o.
func hchacha20_simd128(out []byte, key []byte, in []byte) {
	// Chargements vectoriels directs : les lanes 32-bit d'un vmovdqu sont
	// exactement les mots little-endian attendus par l'état ChaCha.
	v0 := chachaConstRow128
	v1 := archsimd.LoadUint8x16(key[0:16]).AsUint32x4()
	v2 := archsimd.LoadUint8x16(key[16:32]).AsUint32x4()
	v3 := archsimd.LoadUint8x16(in[0:16]).AsUint32x4()

	for i := 0; i < 10; i++ {
		// Column round
		v0 = v0.Add(v1)
		v3 = rotl16x4(v3.Xor(v0))
		v2 = v2.Add(v3)
		v1 = rotl32x4(v1.Xor(v2), 12)
		v0 = v0.Add(v1)
		v3 = rotl8x4(v3.Xor(v0))
		v2 = v2.Add(v3)
		v1 = rotl32x4(v1.Xor(v2), 7)

		// Diagonalisation
		v1 = v1.PermuteScalars(1, 2, 3, 0)
		v2 = v2.PermuteScalars(2, 3, 0, 1)
		v3 = v3.PermuteScalars(3, 0, 1, 2)

		// Diagonal round
		v0 = v0.Add(v1)
		v3 = rotl16x4(v3.Xor(v0))
		v2 = v2.Add(v3)
		v1 = rotl32x4(v1.Xor(v2), 12)
		v0 = v0.Add(v1)
		v3 = rotl8x4(v3.Xor(v0))
		v2 = v2.Add(v3)
		v1 = rotl32x4(v1.Xor(v2), 7)

		// Dé-diagonalisation
		v1 = v1.PermuteScalars(3, 0, 1, 2)
		v2 = v2.PermuteScalars(2, 3, 0, 1)
		v3 = v3.PermuteScalars(1, 2, 3, 0)
	}

	// HChaCha20 : sortie = lignes 0 et 3 de l'état, sans feed-forward.
	v0.AsUint8x16().Store(out[0:16])
	v3.AsUint8x16().Store(out[16:32])
}

// chacha20_deriv2_simd256 dérive exactement 2 blocs consécutifs de keystream
// ChaCha20 (compteurs ctr et ctr+1) en UNE passe de rounds 256-bit : les 2 blocs
// vivent côte à côte dans 4 vecteurs Uint32x8 (lane basse = bloc ctr, lane haute
// = bloc ctr+1). Réutilise le double-round vectoriel chacha20DoubleBlockSIMD256.
// dst doit faire au moins 128 octets ; key=32o, nonce=8o (format djb).
func chacha20_deriv2_simd256(dst []byte, key []byte, nonce []byte, ctr uint64) {
	n0 := binary.LittleEndian.Uint32(nonce[0:4])
	n1 := binary.LittleEndian.Uint32(nonce[4:8])

	ctrB := ctr + 1

	// Duplication 128→256 par vinserti128 : lane basse = bloc ctr, haute = ctr+1.
	kRow0 := archsimd.LoadUint8x16(key[0:16]).AsUint32x4()
	kRow1 := archsimd.LoadUint8x16(key[16:32]).AsUint32x4()
	var z8 archsimd.Uint32x8
	st0 := z8.SetLo(chachaConstRow128).SetHi(chachaConstRow128)
	st1 := z8.SetLo(kRow0).SetHi(kRow0)
	st2 := z8.SetLo(kRow1).SetHi(kRow1)
	st3 := archsimd.LoadUint32x8Array(&[8]uint32{
		uint32(ctr), uint32(ctr >> 32), n0, n1,
		uint32(ctrB), uint32(ctrB >> 32), n0, n1,
	})

	v0, v1, v2, v3 := st0, st1, st2, st3
	for i := 0; i < 10; i++ {
		v0, v1, v2, v3 = chacha20DoubleBlockSIMD256(v0, v1, v2, v3)
	}

	v0 = v0.Add(st0)
	v1 = v1.Add(st1)
	v2 = v2.Add(st2)
	v3 = v3.Add(st3)

	// Bloc ctr (lanes basses)
	v0.GetLo().AsUint8x16().Store(dst[0:16])
	v1.GetLo().AsUint8x16().Store(dst[16:32])
	v2.GetLo().AsUint8x16().Store(dst[32:48])
	v3.GetLo().AsUint8x16().Store(dst[48:64])

	// Bloc ctr+1 (lanes hautes)
	v0.GetHi().AsUint8x16().Store(dst[64:80])
	v1.GetHi().AsUint8x16().Store(dst[80:96])
	v2.GetHi().AsUint8x16().Store(dst[96:112])
	v3.GetHi().AsUint8x16().Store(dst[112:128])
}

// chacha20_xor2_simd256 chiffre exactement 2 blocs (128 octets) en une passe
// de rounds 256-bit : keystream des compteurs ctr/ctr+1 XORé avec src vers dst
// (src nil = keystream brut). Comble la bande 65..255 octets que le kernel 4x
// (seuil 4 blocs) laissait entièrement scalaire — mesure 2026-08-15 :
// 255 o à 645 ns contre 256 o à 475 ns.
func chacha20_xor2_simd256(dst []byte, src []byte, key []byte, nonce []byte, ctr uint64) {
	n0 := binary.LittleEndian.Uint32(nonce[0:4])
	n1 := binary.LittleEndian.Uint32(nonce[4:8])

	ctrB := ctr + 1

	kRow0 := archsimd.LoadUint8x16(key[0:16]).AsUint32x4()
	kRow1 := archsimd.LoadUint8x16(key[16:32]).AsUint32x4()
	var z8 archsimd.Uint32x8
	st0 := z8.SetLo(chachaConstRow128).SetHi(chachaConstRow128)
	st1 := z8.SetLo(kRow0).SetHi(kRow0)
	st2 := z8.SetLo(kRow1).SetHi(kRow1)
	st3 := archsimd.LoadUint32x8Array(&[8]uint32{
		uint32(ctr), uint32(ctr >> 32), n0, n1,
		uint32(ctrB), uint32(ctrB >> 32), n0, n1,
	})

	v0, v1, v2, v3 := st0, st1, st2, st3
	for i := 0; i < 10; i++ {
		v0, v1, v2, v3 = chacha20DoubleBlockSIMD256(v0, v1, v2, v3)
	}

	v0 = v0.Add(st0)
	v1 = v1.Add(st1)
	v2 = v2.Add(st2)
	v3 = v3.Add(st3)

	k0, k1, k2, k3 := v0.GetLo().AsUint8x16(), v1.GetLo().AsUint8x16(), v2.GetLo().AsUint8x16(), v3.GetLo().AsUint8x16()
	k4, k5, k6, k7 := v0.GetHi().AsUint8x16(), v1.GetHi().AsUint8x16(), v2.GetHi().AsUint8x16(), v3.GetHi().AsUint8x16()
	if src == nil {
		k0.Store(dst[0:16])
		k1.Store(dst[16:32])
		k2.Store(dst[32:48])
		k3.Store(dst[48:64])
		k4.Store(dst[64:80])
		k5.Store(dst[80:96])
		k6.Store(dst[96:112])
		k7.Store(dst[112:128])
		return
	}
	_ = src[127]
	_ = dst[127]
	archsimd.LoadUint8x16(src[0:16]).Xor(k0).Store(dst[0:16])
	archsimd.LoadUint8x16(src[16:32]).Xor(k1).Store(dst[16:32])
	archsimd.LoadUint8x16(src[32:48]).Xor(k2).Store(dst[32:48])
	archsimd.LoadUint8x16(src[48:64]).Xor(k3).Store(dst[48:64])
	archsimd.LoadUint8x16(src[64:80]).Xor(k4).Store(dst[64:80])
	archsimd.LoadUint8x16(src[80:96]).Xor(k5).Store(dst[80:96])
	archsimd.LoadUint8x16(src[96:112]).Xor(k6).Store(dst[96:112])
	archsimd.LoadUint8x16(src[112:128]).Xor(k7).Store(dst[112:128])
}
