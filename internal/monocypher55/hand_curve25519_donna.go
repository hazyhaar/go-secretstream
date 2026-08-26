// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"encoding/binary"
	"math/bits"
)

// fe51 représente un élément du corps F_(2^255 - 19) en 5 limbs de 51 bits (Radix-2^51).
type fe51 [5]uint64

const mask51 uint64 = 0x7FFFFFFFFFFFF

func fe51_frombytes(out *fe51, in []byte) {
	w0 := binary.LittleEndian.Uint64(in[0:8])
	w1 := binary.LittleEndian.Uint64(in[8:16])
	w2 := binary.LittleEndian.Uint64(in[16:24])
	w3 := binary.LittleEndian.Uint64(in[24:32])

	out[0] = w0 & mask51
	out[1] = (w0>>51 | w1<<13) & mask51
	out[2] = (w1>>38 | w2<<26) & mask51
	out[3] = (w2>>25 | w3<<39) & mask51
	out[4] = (w3 >> 12) & mask51
}

// fe51_tobytes sérialise un élément fe51 en 32 octets Little-Endian.
// Contrat d'entrée : in < 2p (garanti par fe51_carry). Une unique soustraction conditionnelle de p est effectuée.
func fe51_tobytes(out []byte, in *fe51) {
	t := *in
	// Report de retenue 51 bits
	c := t[0] >> 51
	t[0] &= mask51
	t[1] += c
	c = t[1] >> 51
	t[1] &= mask51
	t[2] += c
	c = t[2] >> 51
	t[2] &= mask51
	t[3] += c
	c = t[3] >> 51
	t[3] &= mask51
	t[4] += c
	c = t[4] >> 51
	t[4] &= mask51
	t[0] += c * 19

	c = t[0] >> 51
	t[0] &= mask51
	t[1] += c

	// Soustraction conditionnelle de p = 2^255 - 19
	m0, c0 := bits.Sub64(t[0], mask51-18, 0)
	m1, c1 := bits.Sub64(t[1], mask51, c0)
	m2, c2 := bits.Sub64(t[2], mask51, c1)
	m3, c3 := bits.Sub64(t[3], mask51, c2)
	m4, c4 := bits.Sub64(t[4], mask51, c3)

	mask := uint64(0) - (1 - c4)
	f0 := (t[0] & ^mask) | (m0 & mask)
	f1 := (t[1] & ^mask) | (m1 & mask)
	f2 := (t[2] & ^mask) | (m2 & mask)
	f3 := (t[3] & ^mask) | (m3 & mask)
	f4 := (t[4] & ^mask) | (m4 & mask)

	binary.LittleEndian.PutUint64(out[0:8], f0|(f1<<51))
	binary.LittleEndian.PutUint64(out[8:16], (f1>>13)|(f2<<38))
	binary.LittleEndian.PutUint64(out[16:24], (f2>>26)|(f3<<25))
	binary.LittleEndian.PutUint64(out[24:32], (f3>>39)|(f4<<12))
}

func fe51_0(out *fe51) {
	out[0] = 0
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
}

func fe51_1(out *fe51) {
	out[0] = 1
	out[1] = 0
	out[2] = 0
	out[3] = 0
	out[4] = 0
}

func fe51_copy(out *fe51, in *fe51) {
	*out = *in
}

func fe51_cswap(a, b *fe51, swap int) {
	mask := uint64(0) - uint64(swap)
	for i := 0; i < 5; i++ {
		x := mask & (a[i] ^ b[i])
		a[i] ^= x
		b[i] ^= x
	}
}

func fe51_add(out, a, b *fe51) {
	out[0] = a[0] + b[0]
	out[1] = a[1] + b[1]
	out[2] = a[2] + b[2]
	out[3] = a[3] + b[3]
	out[4] = a[4] + b[4]
}

func fe51_sub(out, a, b *fe51) {
	const bias0 = (1 << 52) - 38
	const bias1 = (1 << 52) - 2
	out[0] = (a[0] + bias0) - b[0]
	out[1] = (a[1] + bias1) - b[1]
	out[2] = (a[2] + bias1) - b[2]
	out[3] = (a[3] + bias1) - b[3]
	out[4] = (a[4] + bias1) - b[4]
}

func fe51_mul121666(out, in *fe51) {
	const k = 121666
	hi0, lo0 := bits.Mul64(in[0], k)
	hi1, lo1 := bits.Mul64(in[1], k)
	hi2, lo2 := bits.Mul64(in[2], k)
	hi3, lo3 := bits.Mul64(in[3], k)
	hi4, lo4 := bits.Mul64(in[4], k)

	c_0 := (hi0 << 13) | (lo0 >> 51)
	out0 := lo0 & mask51

	var c uint64
	lo1, c = bits.Add64(lo1, c_0, 0)
	hi1, _ = bits.Add64(hi1, 0, c)
	c_1 := (hi1 << 13) | (lo1 >> 51)
	out1 := lo1 & mask51

	lo2, c = bits.Add64(lo2, c_1, 0)
	hi2, _ = bits.Add64(hi2, 0, c)
	c_2 := (hi2 << 13) | (lo2 >> 51)
	out2 := lo2 & mask51

	lo3, c = bits.Add64(lo3, c_2, 0)
	hi3, _ = bits.Add64(hi3, 0, c)
	c_3 := (hi3 << 13) | (lo3 >> 51)
	out3 := lo3 & mask51

	lo4, c = bits.Add64(lo4, c_3, 0)
	hi4, _ = bits.Add64(hi4, 0, c)
	c_4 := (hi4 << 13) | (lo4 >> 51)
	out4 := lo4 & mask51

	hi_c4, lo_c4 := bits.Mul64(c_4, 19)
	out0, c = bits.Add64(out0, lo_c4, 0)
	hi_0 := hi_c4 + c
	c_0_final := (hi_0 << 13) | (out0 >> 51)

	out[0] = out0 & mask51
	out[1] = out1 + c_0_final
	out[2] = out2
	out[3] = out3
	out[4] = out4
}

func fe51_mul(out, a, b *fe51) {
	a0, a1, a2, a3, a4 := a[0], a[1], a[2], a[3], a[4]
	b0, b1, b2, b3, b4 := b[0], b[1], b[2], b[3], b[4]

	a1_19 := a1 * 19
	a2_19 := a2 * 19
	a3_19 := a3 * 19
	a4_19 := a4 * 19

	// r0
	hi0_0, lo0_0 := bits.Mul64(a0, b0)
	hi0_1, lo0_1 := bits.Mul64(a4_19, b1)
	hi0_2, lo0_2 := bits.Mul64(a3_19, b2)
	hi0_3, lo0_3 := bits.Mul64(a2_19, b3)
	hi0_4, lo0_4 := bits.Mul64(a1_19, b4)

	var c, c0, c1 uint64
	var lo0, hi0, lo1, hi1, lo2, hi2, lo3, hi3, lo4, hi4 uint64
	lo0, c = bits.Add64(lo0_0, lo0_1, 0)
	hi0, _ = bits.Add64(hi0_0, hi0_1, c)
	lo0, c = bits.Add64(lo0, lo0_2, 0)
	hi0, _ = bits.Add64(hi0, hi0_2, c)
	lo0, c = bits.Add64(lo0, lo0_3, 0)
	hi0, _ = bits.Add64(hi0, hi0_3, c)
	lo0, c = bits.Add64(lo0, lo0_4, 0)
	hi0, _ = bits.Add64(hi0, hi0_4, c)

	// r1
	hi1_0, lo1_0 := bits.Mul64(a0, b1)
	hi1_1, lo1_1 := bits.Mul64(a1, b0)
	hi1_2, lo1_2 := bits.Mul64(a4_19, b2)
	hi1_3, lo1_3 := bits.Mul64(a3_19, b3)
	hi1_4, lo1_4 := bits.Mul64(a2_19, b4)

	lo1, c = bits.Add64(lo1_0, lo1_1, 0)
	hi1, _ = bits.Add64(hi1_0, hi1_1, c)
	lo1, c = bits.Add64(lo1, lo1_2, 0)
	hi1, _ = bits.Add64(hi1, hi1_2, c)
	lo1, c = bits.Add64(lo1, lo1_3, 0)
	hi1, _ = bits.Add64(hi1, hi1_3, c)
	lo1, c = bits.Add64(lo1, lo1_4, 0)
	hi1, _ = bits.Add64(hi1, hi1_4, c)

	// r2
	hi2_0, lo2_0 := bits.Mul64(a0, b2)
	hi2_1, lo2_1 := bits.Mul64(a1, b1)
	hi2_2, lo2_2 := bits.Mul64(a2, b0)
	hi2_3, lo2_3 := bits.Mul64(a4_19, b3)
	hi2_4, lo2_4 := bits.Mul64(a3_19, b4)

	lo2, c = bits.Add64(lo2_0, lo2_1, 0)
	hi2, _ = bits.Add64(hi2_0, hi2_1, c)
	lo2, c = bits.Add64(lo2, lo2_2, 0)
	hi2, _ = bits.Add64(hi2, hi2_2, c)
	lo2, c = bits.Add64(lo2, lo2_3, 0)
	hi2, _ = bits.Add64(hi2, hi2_3, c)
	lo2, c = bits.Add64(lo2, lo2_4, 0)
	hi2, _ = bits.Add64(hi2, hi2_4, c)

	// r3
	hi3_0, lo3_0 := bits.Mul64(a0, b3)
	hi3_1, lo3_1 := bits.Mul64(a1, b2)
	hi3_2, lo3_2 := bits.Mul64(a2, b1)
	hi3_3, lo3_3 := bits.Mul64(a3, b0)
	hi3_4, lo3_4 := bits.Mul64(a4_19, b4)

	lo3, c = bits.Add64(lo3_0, lo3_1, 0)
	hi3, _ = bits.Add64(hi3_0, hi3_1, c)
	lo3, c = bits.Add64(lo3, lo3_2, 0)
	hi3, _ = bits.Add64(hi3, hi3_2, c)
	lo3, c = bits.Add64(lo3, lo3_3, 0)
	hi3, _ = bits.Add64(hi3, hi3_3, c)
	lo3, c = bits.Add64(lo3, lo3_4, 0)
	hi3, _ = bits.Add64(hi3, hi3_4, c)

	// r4
	hi4_0, lo4_0 := bits.Mul64(a0, b4)
	hi4_1, lo4_1 := bits.Mul64(a1, b3)
	hi4_2, lo4_2 := bits.Mul64(a2, b2)
	hi4_3, lo4_3 := bits.Mul64(a3, b1)
	hi4_4, lo4_4 := bits.Mul64(a4, b0)

	lo4, c = bits.Add64(lo4_0, lo4_1, 0)
	hi4, _ = bits.Add64(hi4_0, hi4_1, c)
	lo4, c = bits.Add64(lo4, lo4_2, 0)
	hi4, _ = bits.Add64(hi4, hi4_2, c)
	lo4, c = bits.Add64(lo4, lo4_3, 0)
	hi4, _ = bits.Add64(hi4, hi4_3, c)
	lo4, c = bits.Add64(lo4, lo4_4, 0)
	hi4, _ = bits.Add64(hi4, hi4_4, c)

	// Propager les reports 51-bit
	c_0 := (hi0 << 13) | (lo0 >> 51)
	out0 := lo0 & mask51

	lo1, c0 = bits.Add64(lo1, c_0, 0)
	hi1, _ = bits.Add64(hi1, 0, c0)
	c_1 := (hi1 << 13) | (lo1 >> 51)
	out1 := lo1 & mask51

	lo2, c0 = bits.Add64(lo2, c_1, 0)
	hi2, _ = bits.Add64(hi2, 0, c0)
	c_2 := (hi2 << 13) | (lo2 >> 51)
	out2 := lo2 & mask51

	lo3, c0 = bits.Add64(lo3, c_2, 0)
	hi3, _ = bits.Add64(hi3, 0, c0)
	c_3 := (hi3 << 13) | (lo3 >> 51)
	out3 := lo3 & mask51

	lo4, c0 = bits.Add64(lo4, c_3, 0)
	hi4, _ = bits.Add64(hi4, 0, c0)
	c_4 := (hi4 << 13) | (lo4 >> 51)
	out4 := lo4 & mask51

	hi_c4, lo_c4 := bits.Mul64(c_4, 19)
	out0, c1 = bits.Add64(out0, lo_c4, 0)
	hi_0 := hi_c4 + c1
	c_0_final := (hi_0 << 13) | (out0 >> 51)

	out[0] = out0 & mask51
	out[1] = out1 + c_0_final
	out[2] = out2
	out[3] = out3
	out[4] = out4
}

func fe51_sq(out, in *fe51) {
	a0, a1, a2, a3, a4 := in[0], in[1], in[2], in[3], in[4]

	a0_2 := a0 * 2
	a1_2 := a1 * 2
	a4_19 := a4 * 19
	a4_38 := a4 * 38
	a3_38 := a3 * 38

	// r0 = a0^2 + 38 a1 a4 + 38 a2 a3
	hi0_0, lo0_0 := bits.Mul64(a0, a0)
	hi0_1, lo0_1 := bits.Mul64(a1, a4_38)
	hi0_2, lo0_2 := bits.Mul64(a2, a3_38)

	// r1 = 2 a0 a1 + 38 a2 a4 + 19 a3^2
	hi1_0, lo1_0 := bits.Mul64(a0_2, a1)
	hi1_1, lo1_1 := bits.Mul64(a2, a4_38)
	hi1_2, lo1_2 := bits.Mul64(a3, a3*19)

	// r2 = 2 a0 a2 + a1^2 + 38 a3 a4
	hi2_0, lo2_0 := bits.Mul64(a0_2, a2)
	hi2_1, lo2_1 := bits.Mul64(a1, a1)
	hi2_2, lo2_2 := bits.Mul64(a3, a4_38)

	// r3 = 2 a0 a3 + 2 a1 a2 + 19 a4^2
	hi3_0, lo3_0 := bits.Mul64(a0_2, a3)
	hi3_1, lo3_1 := bits.Mul64(a1_2, a2)
	hi3_2, lo3_2 := bits.Mul64(a4, a4_19)

	// r4 = 2 a0 a4 + 2 a1 a3 + a2^2
	hi4_0, lo4_0 := bits.Mul64(a0_2, a4)
	hi4_1, lo4_1 := bits.Mul64(a1_2, a3)
	hi4_2, lo4_2 := bits.Mul64(a2, a2)

	var c, c0, c1 uint64
	var lo0, hi0, lo1, hi1, lo2, hi2, lo3, hi3, lo4, hi4 uint64

	lo0, c = bits.Add64(lo0_0, lo0_1, 0)
	hi0, _ = bits.Add64(hi0_0, hi0_1, c)
	lo0, c = bits.Add64(lo0, lo0_2, 0)
	hi0, _ = bits.Add64(hi0, hi0_2, c)

	lo1, c = bits.Add64(lo1_0, lo1_1, 0)
	hi1, _ = bits.Add64(hi1_0, hi1_1, c)
	lo1, c = bits.Add64(lo1, lo1_2, 0)
	hi1, _ = bits.Add64(hi1, hi1_2, c)

	lo2, c = bits.Add64(lo2_0, lo2_1, 0)
	hi2, _ = bits.Add64(hi2_0, hi2_1, c)
	lo2, c = bits.Add64(lo2, lo2_2, 0)
	hi2, _ = bits.Add64(hi2, hi2_2, c)

	lo3, c = bits.Add64(lo3_0, lo3_1, 0)
	hi3, _ = bits.Add64(hi3_0, hi3_1, c)
	lo3, c = bits.Add64(lo3, lo3_2, 0)
	hi3, _ = bits.Add64(hi3, hi3_2, c)

	lo4, c = bits.Add64(lo4_0, lo4_1, 0)
	hi4, _ = bits.Add64(hi4_0, hi4_1, c)
	lo4, c = bits.Add64(lo4, lo4_2, 0)
	hi4, _ = bits.Add64(hi4, hi4_2, c)

	// Propager les reports 51-bit
	c_0 := (hi0 << 13) | (lo0 >> 51)
	out0 := lo0 & mask51

	lo1, c0 = bits.Add64(lo1, c_0, 0)
	hi1, _ = bits.Add64(hi1, 0, c0)
	c_1 := (hi1 << 13) | (lo1 >> 51)
	out1 := lo1 & mask51

	lo2, c0 = bits.Add64(lo2, c_1, 0)
	hi2, _ = bits.Add64(hi2, 0, c0)
	c_2 := (hi2 << 13) | (lo2 >> 51)
	out2 := lo2 & mask51

	lo3, c0 = bits.Add64(lo3, c_2, 0)
	hi3, _ = bits.Add64(hi3, 0, c0)
	c_3 := (hi3 << 13) | (lo3 >> 51)
	out3 := lo3 & mask51

	lo4, c0 = bits.Add64(lo4, c_3, 0)
	hi4, _ = bits.Add64(hi4, 0, c0)
	c_4 := (hi4 << 13) | (lo4 >> 51)
	out4 := lo4 & mask51

	hi_c4, lo_c4 := bits.Mul64(c_4, 19)
	out0, c1 = bits.Add64(out0, lo_c4, 0)
	hi_0 := hi_c4 + c1
	c_0_final := (hi_0 << 13) | (out0 >> 51)

	out[0] = out0 & mask51
	out[1] = out1 + c_0_final
	out[2] = out2
	out[3] = out3
	out[4] = out4
}

func fe51_invert(out, z *fe51) {
	var t0, t1, t2, t3 fe51
	fe51_sq(&t0, z)
	fe51_sq(&t1, &t0)
	fe51_sq(&t1, &t1)
	fe51_mul(&t1, z, &t1)
	fe51_mul(&t0, &t0, &t1)
	fe51_sq(&t2, &t0)
	fe51_mul(&t1, &t1, &t2)
	fe51_sq(&t2, &t1)
	for i := 1; i < 5; i++ {
		fe51_sq(&t2, &t2)
	}
	fe51_mul(&t1, &t2, &t1)
	fe51_sq(&t2, &t1)
	for i := 1; i < 10; i++ {
		fe51_sq(&t2, &t2)
	}
	fe51_mul(&t2, &t2, &t1)
	fe51_sq(&t3, &t2)
	for i := 1; i < 20; i++ {
		fe51_sq(&t3, &t3)
	}
	fe51_mul(&t2, &t3, &t2)
	fe51_sq(&t2, &t2)
	for i := 1; i < 10; i++ {
		fe51_sq(&t2, &t2)
	}
	fe51_mul(&t1, &t2, &t1)
	fe51_sq(&t2, &t1)
	for i := 1; i < 50; i++ {
		fe51_sq(&t2, &t2)
	}
	fe51_mul(&t2, &t2, &t1)
	fe51_sq(&t3, &t2)
	for i := 1; i < 100; i++ {
		fe51_sq(&t3, &t3)
	}
	fe51_mul(&t2, &t3, &t2)
	fe51_sq(&t2, &t2)
	for i := 1; i < 50; i++ {
		fe51_sq(&t2, &t2)
	}
	fe51_mul(&t1, &t2, &t1)
	fe51_sq(&t1, &t1)
	for i := 1; i < 5; i++ {
		fe51_sq(&t1, &t1)
	}
	fe51_mul(out, &t1, &t0)
}

func Scalarmult51(q, scalar, p []byte, nb_bits int) {
	var x1, x2, z2, x3, z3, a, b, c, d, e, aa, bb, da, cb, t0, t1, e121666 fe51

	fe51_frombytes(&x1, p)
	fe51_1(&x2)
	fe51_0(&z2)
	fe51_copy(&x3, &x1)
	fe51_1(&z3)

	swap := 0
	for pos := nb_bits - 1; pos >= 0; pos-- {
		b_bit := int((scalar[pos>>3] >> uint(pos&7)) & 1)
		swap ^= b_bit
		fe51_cswap(&x2, &x3, swap)
		fe51_cswap(&z2, &z3, swap)
		swap = b_bit

		fe51_add(&a, &x2, &z2)
		fe51_sub(&b, &x2, &z2)
		fe51_add(&c, &x3, &z3)
		fe51_sub(&d, &x3, &z3)

		fe51_sq(&aa, &a)
		fe51_sq(&bb, &b)
		fe51_mul(&da, &d, &a)
		fe51_mul(&cb, &c, &b)

		fe51_sub(&e, &aa, &bb)
		fe51_mul121666(&e121666, &e)
		fe51_add(&z2, &bb, &e121666)
		fe51_mul(&z2, &z2, &e)
		fe51_mul(&x2, &aa, &bb)

		fe51_add(&t0, &da, &cb)
		fe51_sq(&x3, &t0)
		fe51_sub(&t1, &da, &cb)
		fe51_sq(&t1, &t1)
		fe51_mul(&z3, &x1, &t1)
	}
	fe51_cswap(&x2, &x3, swap)
	fe51_cswap(&z2, &z3, swap)

	fe51_invert(&z2, &z2)
	fe51_mul(&x2, &x2, &z2)
	fe51_tobytes(q, &x2)
}

func Scalarmult(q []byte, scalar []byte, p []byte, nb_bits int) {
	Scalarmult51(q, scalar, p, nb_bits)
}

func Scalarmult_ref10(q []byte, scalar []byte, p []byte, nb_bits int) {
	var v16 int
	var v18 int
	var _arr_v4 [10]int32
	v4 := &_arr_v4
	var _arr_v5 [10]int32
	v5 := &_arr_v5
	var _arr_v6 [10]int32
	v6 := &_arr_v6
	var _arr_v7 [10]int32
	v7 := &_arr_v7
	var _arr_v8 [10]int32
	v8 := &_arr_v8
	var _arr_v9 [10]int32
	v9 := &_arr_v9
	var _arr_v10 [10]int32
	v10 := &_arr_v10
	var _arr_v11 [10]int32
	v11 := &_arr_v11
	var _arr_v12 [10]int32
	v12 := &_arr_v12
	var _arr_v13 [10]int32
	v13 := &_arr_v13
	var _arr_v14 [10]int32
	v14 := &_arr_v14
	var _arr_v15 [10]int32
	v15 := &_arr_v15
	fe_frombytes(v4, p)
	fe_1(v5)
	fe_0(v6)
	fe_copy(v7, v4)
	fe_1(v8)
	v16 = 0
	v18 = nb_bits - 1
	for v18 >= 0 {
		v19 := scalar_bit(scalar, v18)
		v16 ^= v19
		fe_cswap(v5, v7, v16)
		fe_cswap(v6, v8, v16)
		v16 = v19
		fe_add(v9, v5, v6)
		fe_sub(v6, v5, v6)
		fe_add(v10, v7, v8)
		fe_sub(v7, v7, v8)
		fe_sq(v11, v9)
		fe_sq(v12, v6)
		fe_mul(v13, v7, v9)
		fe_mul(v14, v10, v6)
		fe_sub(v15, v11, v12)
		fe_mul_small(v8, v15, 121666)
		fe_add(v8, v8, v12)
		fe_mul(v6, v8, v15)
		fe_mul(v5, v11, v12)
		fe_add(v8, v13, v14)
		fe_sq(v7, v8)
		fe_sub(v8, v13, v14)
		fe_sq(v8, v8)
		fe_mul(v8, v4, v8)
		v18--
	}
	fe_cswap(v5, v7, v16)
	fe_cswap(v6, v8, v16)
	fe_invert(v6, v6)
	fe_mul(v5, v5, v6)
	fe_tobytes(q, v5)
	clear(v4[:])
	clear(v5[:])
	clear(v6[:])
	clear(v7[:])
	clear(v8[:])
	clear(v9[:])
	clear(v10[:])
	clear(v11[:])
	clear(v12[:])
	clear(v13[:])
	clear(v14[:])
	clear(v15[:])
}
