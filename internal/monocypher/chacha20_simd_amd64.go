//go:build goexperiment.simd && amd64

package monocypher

import (
	"encoding/binary"
	"simd/archsimd"
	"unsafe"
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

// ChaCha20DoubleBlockSIMD256 effectue 1 tour (Column Round + Diagonal Round) sur 2 blocs ChaCha20 en parallèle (256-bit Uint32x8).
func chacha20DoubleBlockSIMD256(v0, v1, v2, v3 archsimd.Uint32x8) (archsimd.Uint32x8, archsimd.Uint32x8, archsimd.Uint32x8, archsimd.Uint32x8) {
	// 1. Column Round
	v0 = v0.Add(v1)
	v3 = v3.Xor(v0).RotateAllLeft(16)
	v2 = v2.Add(v3)
	v1 = v1.Xor(v2).RotateAllLeft(12)
	v0 = v0.Add(v1)
	v3 = v3.Xor(v0).RotateAllLeft(8)
	v2 = v2.Add(v3)
	v1 = v1.Xor(v2).RotateAllLeft(7)

	// 2. Diagonal Round
	v1 = v1.PermuteScalarsGrouped(1, 2, 3, 0)
	v2 = v2.PermuteScalarsGrouped(2, 3, 0, 1)
	v3 = v3.PermuteScalarsGrouped(3, 0, 1, 2)

	v0 = v0.Add(v1)
	v3 = v3.Xor(v0).RotateAllLeft(16)
	v2 = v2.Add(v3)
	v1 = v1.Xor(v2).RotateAllLeft(12)
	v0 = v0.Add(v1)
	v3 = v3.Xor(v0).RotateAllLeft(8)
	v2 = v2.Add(v3)
	v1 = v1.Xor(v2).RotateAllLeft(7)

	// Inverse Shuffles
	v1 = v1.PermuteScalarsGrouped(3, 0, 1, 2)
	v2 = v2.PermuteScalarsGrouped(2, 3, 0, 1)
	v3 = v3.PermuteScalarsGrouped(1, 2, 3, 0)

	return v0, v1, v2, v3
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

	offset := 0
	currCtr := ctr

	for bIdx := uint64(0); bIdx+4 <= numBlocks; bIdx += 4 {
		ctr0_lo, ctr0_hi := uint32(currCtr), uint32(currCtr>>32)
		ctr1_lo, ctr1_hi := uint32(currCtr+1), uint32((currCtr+1)>>32)
		ctr2_lo, ctr2_hi := uint32(currCtr+2), uint32((currCtr+2)>>32)
		ctr3_lo, ctr3_hi := uint32(currCtr+3), uint32((currCtr+3)>>32)

		st3_A := archsimd.LoadUint32x8Array(&[8]uint32{
			ctr0_lo, ctr0_hi, n0, n1,
			ctr1_lo, ctr1_hi, n0, n1,
		})
		st3_B := archsimd.LoadUint32x8Array(&[8]uint32{
			ctr2_lo, ctr2_hi, n0, n1,
			ctr3_lo, ctr3_hi, n0, n1,
		})

		v0_A, v1_A, v2_A, v3_A := st0, st1, st2, st3_A
		v0_B, v1_B, v2_B, v3_B := st0, st1, st2, st3_B

		// 10 doubles-rounds (20 rounds au total)
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

		var buf0_A, buf1_A, buf2_A, buf3_A [8]uint32
		var buf0_B, buf1_B, buf2_B, buf3_B [8]uint32
		v0_A.StoreArray(&buf0_A)
		v1_A.StoreArray(&buf1_A)
		v2_A.StoreArray(&buf2_A)
		v3_A.StoreArray(&buf3_A)

		v0_B.StoreArray(&buf0_B)
		v1_B.StoreArray(&buf1_B)
		v2_B.StoreArray(&buf2_B)
		v3_B.StoreArray(&buf3_B)

		// Écriture / Masquage XOR pour les 4 blocs (256 octets)
		for blk := 0; blk < 4; blk++ {
			bOff := offset + blk*64
			var ks [16]uint32
			if blk < 2 {
				base := blk * 4
				ks[0], ks[1], ks[2], ks[3] = buf0_A[base+0], buf0_A[base+1], buf0_A[base+2], buf0_A[base+3]
				ks[4], ks[5], ks[6], ks[7] = buf1_A[base+0], buf1_A[base+1], buf1_A[base+2], buf1_A[base+3]
				ks[8], ks[9], ks[10], ks[11] = buf2_A[base+0], buf2_A[base+1], buf2_A[base+2], buf2_A[base+3]
				ks[12], ks[13], ks[14], ks[15] = buf3_A[base+0], buf3_A[base+1], buf3_A[base+2], buf3_A[base+3]
			} else {
				base := (blk - 2) * 4
				ks[0], ks[1], ks[2], ks[3] = buf0_B[base+0], buf0_B[base+1], buf0_B[base+2], buf0_B[base+3]
				ks[4], ks[5], ks[6], ks[7] = buf1_B[base+0], buf1_B[base+1], buf1_B[base+2], buf1_B[base+3]
				ks[8], ks[9], ks[10], ks[11] = buf2_B[base+0], buf2_B[base+1], buf2_B[base+2], buf2_B[base+3]
				ks[12], ks[13], ks[14], ks[15] = buf3_B[base+0], buf3_B[base+1], buf3_B[base+2], buf3_B[base+3]
			}

			if plain_text != nil {
				sPtr := (*[8]uint64)(unsafe.Pointer(&plain_text[bOff]))
				dPtr := (*[8]uint64)(unsafe.Pointer(&cipher_text[bOff]))
				kPtr := (*[8]uint64)(unsafe.Pointer(&ks[0]))
				dPtr[0] = sPtr[0] ^ kPtr[0]
				dPtr[1] = sPtr[1] ^ kPtr[1]
				dPtr[2] = sPtr[2] ^ kPtr[2]
				dPtr[3] = sPtr[3] ^ kPtr[3]
				dPtr[4] = sPtr[4] ^ kPtr[4]
				dPtr[5] = sPtr[5] ^ kPtr[5]
				dPtr[6] = sPtr[6] ^ kPtr[6]
				dPtr[7] = sPtr[7] ^ kPtr[7]
			} else {
				dPtr := (*[8]uint64)(unsafe.Pointer(&cipher_text[bOff]))
				kPtr := (*[8]uint64)(unsafe.Pointer(&ks[0]))
				dPtr[0] = kPtr[0]
				dPtr[1] = kPtr[1]
				dPtr[2] = kPtr[2]
				dPtr[3] = kPtr[3]
				dPtr[4] = kPtr[4]
				dPtr[5] = kPtr[5]
				dPtr[6] = kPtr[6]
				dPtr[7] = kPtr[7]
			}
		}

		offset += 256
		currCtr += 4
	}

	return currCtr
}
