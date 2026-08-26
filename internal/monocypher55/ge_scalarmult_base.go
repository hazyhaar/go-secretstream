// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

// Hand completion of monocypher ge_scalarmult_base (comb tables not yet harvested as ge_precomp[]).
// Algorithm: Mike Hamburg twin 4-bit signed combs — monocypher 4.0.2.

func combToPrecomp(t *[3][10]int32) Ge_precomp {
	var g Ge_precomp
	copy(g.Yp[:], t[0][:])
	copy(g.Ym[:], t[1][:])
	copy(g.T2[:], t[2][:])
	return g
}

func lookup_add(p *Ge, tmp_c *Ge_precomp, tmp_a, tmp_b *[10]int32, comb *[8][3][10]int32, scalar []byte, i int) {
	teeth := byte(scalar_bit(scalar, i)) +
		byte(scalar_bit(scalar, i+32)<<1) +
		byte(scalar_bit(scalar, i+64)<<2) +
		byte(scalar_bit(scalar, i+96)<<3)
	high := teeth >> 3
	index := (teeth ^ (high - 1)) & 7
	for j := 0; j < 8; j++ {
		select_ := int(1 & (((int(j) ^ int(index)) - 1) >> 8))
		pc := combToPrecomp(&comb[j])
		fe_ccopy(&tmp_c.Yp, &pc.Yp, select_)
		fe_ccopy(&tmp_c.Ym, &pc.Ym, select_)
		fe_ccopy(&tmp_c.T2, &pc.T2, select_)
	}
	fe_neg(tmp_a, &tmp_c.T2)
	fe_cswap(&tmp_c.T2, tmp_a, int(high^1))
	fe_cswap(&tmp_c.Yp, &tmp_c.Ym, int(high^1))
	ge_madd(p, p, tmp_c, tmp_a, tmp_b)
}

// Ge_scalarmult_base sets p = [scalar]B (base point).
func Ge_scalarmult_base(p *Ge, scalar []byte) {
	var s_scalar [32]byte
	Crypto_eddsa_mul_add(s_scalar[:], scalar, Half_mod_L[:], Half_ones[:])
	var tmp_a, tmp_b [10]int32
	var tmp_c Ge_precomp
	var tmp_d Ge
	fe_1(&tmp_c.Yp)
	fe_1(&tmp_c.Ym)
	fe_0(&tmp_c.T2)
	ge_zero(p)
	lookup_add(p, &tmp_c, &tmp_a, &tmp_b, &b_comb_low, s_scalar[:], 31)
	lookup_add(p, &tmp_c, &tmp_a, &tmp_b, &b_comb_high, s_scalar[:], 31+128)
	for i := 30; i >= 0; i-- {
		ge_double(p, p, &tmp_d)
		lookup_add(p, &tmp_c, &tmp_a, &tmp_b, &b_comb_low, s_scalar[:], i)
		lookup_add(p, &tmp_c, &tmp_a, &tmp_b, &b_comb_high, s_scalar[:], i+128)
	}
}
