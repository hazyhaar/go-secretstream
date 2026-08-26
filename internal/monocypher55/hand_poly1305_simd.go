// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"encoding/binary"
	"math/bits"
)

// Hand: Poly_blocks — algorithme Poly1305 64-bit (2 limbs saturés de 64 bits + h2).
// Réduit le nombre de multiplications 64-bit de 22 à 6 par bloc de 16 octets.
func Poly_blocks(ctx *Crypto_poly1305_ctx, in []byte, nb_blocks uint64, end uint32) {
	if nb_blocks == 0 {
		return
	}

	r0 := uint64(ctx.R[0]) | (uint64(ctx.R[1]) << 32)
	r1 := uint64(ctx.R[2]) | (uint64(ctx.R[3]) << 32)

	h0 := uint64(ctx.H[0]) | (uint64(ctx.H[1]) << 32)
	h1 := uint64(ctx.H[2]) | (uint64(ctx.H[3]) << 32)
	h2 := uint64(ctx.H[4])

	offset := 0
	end64 := uint64(end)

	for b := uint64(0); b < nb_blocks; b++ {
		m0 := binary.LittleEndian.Uint64(in[offset : offset+8])
		m1 := binary.LittleEndian.Uint64(in[offset+8 : offset+16])
		offset += 16

		var c uint64
		h0, c = bits.Add64(h0, m0, 0)
		h1, c = bits.Add64(h1, m1, c)
		h2 += c + end64

		// 6 multiplications 64-bit (instructions MULQ / MULXQ)
		hi0_0, lo0_0 := bits.Mul64(h0, r0)
		hi1_0, lo1_0 := bits.Mul64(h1, r0)
		hi0_1, lo0_1 := bits.Mul64(h0, r1)
		hi1_1, lo1_1 := bits.Mul64(h1, r1)

		lo2_0 := h2 * r0
		lo2_1 := h2 * r1

		// Sommation en 4 membres de 64 bits t0, t1, t2, t3
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

		// Réduction partielle modulo 2^130 - 5
		h0 = t0
		h1 = t1
		h2 = t2 & 3

		cc_lo := t2 & ^uint64(3)
		cc_hi := t3

		// + cc (c * 4)
		h0, c = bits.Add64(h0, cc_lo, 0)
		h1, c = bits.Add64(h1, cc_hi, c)
		h2 += c

		// + (cc >> 2) (c * 1)
		c_lo := (cc_lo >> 2) | (cc_hi << 62)
		c_hi := cc_hi >> 2

		h0, c = bits.Add64(h0, c_lo, 0)
		h1, c = bits.Add64(h1, c_hi, c)
		h2 += c
		if h2 > 3 {
			h0, c = bits.Add64(h0, (h2>>2)*5, 0)
			h1, c = bits.Add64(h1, 0, c)
			h2 = (h2 & 3) + c
		}
	}

	ctx.H[0] = uint32(h0)
	ctx.H[1] = uint32(h0 >> 32)
	ctx.H[2] = uint32(h1)
	ctx.H[3] = uint32(h1 >> 32)
	ctx.H[4] = uint32(h2)
}
