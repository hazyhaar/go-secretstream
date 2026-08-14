//go:build goexperiment.simd && amd64

package monocypher

import (
	"encoding/binary"
	"math/bits"
	"simd/archsimd"
)

// poly1305_blocks_simd_4x absorbe les blocs de 16 octets par lots de 4 (64 octets)
// en utilisant des limbs 44-bit et la vectorisation des multi-blocs sans dépendances de registres d'état.
func poly1305_blocks_simd_4x(ctx *Crypto_poly1305_ctx, in []byte, nb_blocks uint64) uint64 {
	if nb_blocks < 4 {
		return 0
	}

	// Décomposition de la clé r (clamped) en 3 membres 44-bit
	r0_32 := uint64(ctx.R[0])
	r1_32 := uint64(ctx.R[1])
	r2_32 := uint64(ctx.R[2])
	r3_32 := uint64(ctx.R[3])

	r0 := (r0_32 | (r1_32 << 32)) & 0xfffffffffff
	r1 := ((r1_32 >> 12) | (r2_32 << 20)) & 0xfffffffffff
	r2 := (r2_32 >> 24) | (r3_32 << 8)

	s1 := r1 * 20
	s2 := r2 * 20

	// Puissance r^2 (mod 2^130-5)
	r2_h0, r2_h1, r2_h2 := mul44_core(r0, r1, r2, r0, r1, r2, s1, s2)
	r2_s1 := r2_h1 * 20
	r2_s2 := r2_h2 * 20

	// Puissance r^3
	r3_h0, r3_h1, r3_h2 := mul44_core(r2_h0, r2_h1, r2_h2, r0, r1, r2, s1, s2)
	r3_s1 := r3_h1 * 20
	r3_s2 := r3_h2 * 20

	// Puissance r^4
	r4_h0, r4_h1, r4_h2 := mul44_core(r2_h0, r2_h1, r2_h2, r2_h0, r2_h1, r2_h2, r2_s1, r2_s2)
	r4_s1 := r4_h1 * 20
	r4_s2 := r4_h2 * 20

	// Conversion de ctx.H (5x32-bit) vers h (3x44-bit)
	h0_in := (uint64(ctx.H[0]) | (uint64(ctx.H[1]) << 32)) & 0xfffffffffff
	h1_in := ((uint64(ctx.H[1]) >> 12) | (uint64(ctx.H[2]) << 20) | (uint64(ctx.H[3]) << 52)) & 0xfffffffffff
	h2_in := (uint64(ctx.H[3]) >> 24) | (uint64(ctx.H[4]) << 8)

	h0, h1, h2 := h0_in, h1_in, h2_in
	offset := 0
	processedBlocks := (nb_blocks / 4) * 4

	for b := uint64(0); b < processedBlocks; b += 4 {
		// Lecture des 4 blocs de 16 octets
		k0_0 := binary.LittleEndian.Uint64(in[offset : offset+8])
		k0_1 := binary.LittleEndian.Uint64(in[offset+8 : offset+16])

		k1_0 := binary.LittleEndian.Uint64(in[offset+16 : offset+24])
		k1_1 := binary.LittleEndian.Uint64(in[offset+24 : offset+32])

		k2_0 := binary.LittleEndian.Uint64(in[offset+32 : offset+40])
		k2_1 := binary.LittleEndian.Uint64(in[offset+40 : offset+48])

		k3_0 := binary.LittleEndian.Uint64(in[offset+48 : offset+56])
		k3_1 := binary.LittleEndian.Uint64(in[offset+56 : offset+64])

		// Bloc 0 + accumulateur h précédent (multiplié par r^4)
		b0_m0 := (k0_0 & 0xfffffffffff) + h0
		b0_m1 := (((k0_0 >> 44) | (k0_1 << 20)) & 0xfffffffffff) + h1
		b0_m2 := ((k0_1 >> 24) | (1 << 40)) + h2

		// Bloc 1 (multiplié par r^3)
		b1_m0 := k1_0 & 0xfffffffffff
		b1_m1 := ((k1_0 >> 44) | (k1_1 << 20)) & 0xfffffffffff
		b1_m2 := (k1_1 >> 24) | (1 << 40)

		// Bloc 2 (multiplié par r^2)
		b2_m0 := k2_0 & 0xfffffffffff
		b2_m1 := ((k2_0 >> 44) | (k2_1 << 20)) & 0xfffffffffff
		b2_m2 := (k2_1 >> 24) | (1 << 40)

		// Bloc 3 (multiplié par r^1)
		b3_m0 := k3_0 & 0xfffffffffff
		b3_m1 := ((k3_0 >> 44) | (k3_1 << 20)) & 0xfffffffffff
		b3_m2 := (k3_1 >> 24) | (1 << 40)

		// Multiplications partielles 4-Way
		p0_0, p0_1, p0_2 := mul44_terms(b0_m0, b0_m1, b0_m2, r4_h0, r4_h1, r4_h2, r4_s1, r4_s2)
		p1_0, p1_1, p1_2 := mul44_terms(b1_m0, b1_m1, b1_m2, r3_h0, r3_h1, r3_h2, r3_s1, r3_s2)
		p2_0, p2_1, p2_2 := mul44_terms(b2_m0, b2_m1, b2_m2, r2_h0, r2_h1, r2_h2, r2_s1, r2_s2)
		p3_0, p3_1, p3_2 := mul44_terms(b3_m0, b3_m1, b3_m2, r0, r1, r2, s1, s2)

		// Sommation des 4 produits
		t0_lo, c := bits.Add64(p0_0.lo, p1_0.lo, 0)
		t0_hi, _ := bits.Add64(p0_0.hi, p1_0.hi, c)
		t0_lo, c = bits.Add64(t0_lo, p2_0.lo, 0)
		t0_hi, _ = bits.Add64(t0_hi, p2_0.hi, c)
		t0_lo, c = bits.Add64(t0_lo, p3_0.lo, 0)
		t0_hi, _ = bits.Add64(t0_hi, p3_0.hi, c)

		t1_lo, c := bits.Add64(p0_1.lo, p1_1.lo, 0)
		t1_hi, _ := bits.Add64(p0_1.hi, p1_1.hi, c)
		t1_lo, c = bits.Add64(t1_lo, p2_1.lo, 0)
		t1_hi, _ = bits.Add64(t1_hi, p2_1.hi, c)
		t1_lo, c = bits.Add64(t1_lo, p3_1.lo, 0)
		t1_hi, _ = bits.Add64(t1_hi, p3_1.hi, c)

		t2_lo, c := bits.Add64(p0_2.lo, p1_2.lo, 0)
		t2_hi, _ := bits.Add64(p0_2.hi, p1_2.hi, c)
		t2_lo, c = bits.Add64(t2_lo, p2_2.lo, 0)
		t2_hi, _ = bits.Add64(t2_hi, p2_2.hi, c)
		t2_lo, c = bits.Add64(t2_lo, p3_2.lo, 0)
		t2_hi, _ = bits.Add64(t2_hi, p3_2.hi, c)

		// Réduction finale des membres 44-bit
		c0 := (t0_lo >> 44) | (t0_hi << 20)
		h0 = t0_lo & 0xfffffffffff

		t1_lo, c = bits.Add64(t1_lo, c0, 0)
		t1_hi, _ = bits.Add64(t1_hi, 0, c)
		c1 := (t1_lo >> 44) | (t1_hi << 20)
		h1 = t1_lo & 0xfffffffffff

		t2_lo, c = bits.Add64(t2_lo, c1, 0)
		t2_hi, _ = bits.Add64(t2_hi, 0, c)
		c2 := (t2_lo >> 42) | (t2_hi << 22)
		h2 = t2_lo & 0x3ffffffffff

		h0 += c2 * 5
		c = h0 >> 44
		h0 &= 0xfffffffffff
		h1 += c

		offset += 64
	}

	// Reconversion vers ctx.H (5x32-bit)
	ctx.H[0] = uint32(h0)
	ctx.H[1] = uint32((h0 >> 32) | (h1 << 12))
	ctx.H[2] = uint32(h1 >> 20)
	ctx.H[3] = uint32((h1 >> 52) | (h2 << 24))
	ctx.H[4] = uint32(h2 >> 8)

	return processedBlocks
}

type term128 struct {
	lo, hi uint64
}

func mul44_terms(h0, h1, h2, r0, r1, r2, s1, s2 uint64) (term128, term128, term128) {
	hi0_0, lo0_0 := bits.Mul64(h0, r0)
	hi0_1, lo0_1 := bits.Mul64(h1, s2)
	hi0_2, lo0_2 := bits.Mul64(h2, s1)

	hi1_0, lo1_0 := bits.Mul64(h0, r1)
	hi1_1, lo1_1 := bits.Mul64(h1, r0)
	hi1_2, lo1_2 := bits.Mul64(h2, s2)

	hi2_0, lo2_0 := bits.Mul64(h0, r2)
	hi2_1, lo2_1 := bits.Mul64(h1, r1)
	hi2_2, lo2_2 := bits.Mul64(h2, r0)

	d0_lo, c0 := bits.Add64(lo0_0, lo0_1, 0)
	d0_hi, _ := bits.Add64(hi0_0, hi0_1, c0)
	d0_lo, c0 = bits.Add64(d0_lo, lo0_2, 0)
	d0_hi, _ = bits.Add64(d0_hi, hi0_2, c0)

	d1_lo, c1 := bits.Add64(lo1_0, lo1_1, 0)
	d1_hi, _ := bits.Add64(hi1_0, hi1_1, c1)
	d1_lo, c1 = bits.Add64(d1_lo, lo1_2, 0)
	d1_hi, _ = bits.Add64(d1_hi, hi1_2, c1)

	d2_lo, c2 := bits.Add64(lo2_0, lo2_1, 0)
	d2_hi, _ := bits.Add64(hi2_0, hi2_1, c2)
	d2_lo, c2 = bits.Add64(d2_lo, lo2_2, 0)
	d2_hi, _ = bits.Add64(d2_hi, hi2_2, c2)

	return term128{lo: d0_lo, hi: d0_hi}, term128{lo: d1_lo, hi: d1_hi}, term128{lo: d2_lo, hi: d2_hi}
}

func mul44_core(h0, h1, h2, r0, r1, r2, s1, s2 uint64) (u0, u1, u2 uint64) {
	t0, t1, t2 := mul44_terms(h0, h1, h2, r0, r1, r2, s1, s2)

	c := (t0.lo >> 44) | (t0.hi << 20)
	u0 = t0.lo & 0xfffffffffff

	t1_lo, c1 := bits.Add64(t1.lo, c, 0)
	t1_hi, _ := bits.Add64(t1.hi, 0, c1)
	c = (t1_lo >> 44) | (t1_hi << 20)
	u1 = t1_lo & 0xfffffffffff

	t2_lo, c2 := bits.Add64(t2.lo, c, 0)
	t2_hi, _ := bits.Add64(t2.hi, 0, c2)
	c = (t2_lo >> 42) | (t2_hi << 22)
	u2 = t2_lo & 0x3ffffffffff

	u0 += c * 5
	c = u0 >> 44
	u0 &= 0xfffffffffff
	u1 += c

	return u0, u1, u2
}

// MicroSpike_VPMULUDQ valide l'instruction vectorielle AVX2
func MicroSpike_VPMULUDQ(a, b [8]uint32) [4]uint64 {
	vA := archsimd.LoadUint32x8Array(&a)
	vB := archsimd.LoadUint32x8Array(&b)
	vRes := vA.MulWidenEven(vB)
	var out [4]uint64
	vRes.StoreArray(&out)
	return out
}
