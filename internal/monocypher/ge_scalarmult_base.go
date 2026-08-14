package monocypher

// Hand completion of monocypher ge_scalarmult_base (comb tables not yet harvested as ge_precomp[]).
// Algorithm: Mike Hamburg twin 4-bit signed combs — monocypher 4.0.2.

func combToPrecomp(t *[3][10]int) Ge_precomp {
	var g Ge_precomp
	copy(g.Yp[:], t[0][:])
	copy(g.Ym[:], t[1][:])
	copy(g.T2[:], t[2][:])
	return g
}

func lookup_add(p *Ge, tmp_c *Ge_precomp, tmp_a, tmp_b []int, comb *[8][3][10]int, scalar []byte, i int) {
	teeth := byte(Scalar_bit(scalar, i)) +
		byte(Scalar_bit(scalar, i+32)<<1) +
		byte(Scalar_bit(scalar, i+64)<<2) +
		byte(Scalar_bit(scalar, i+96)<<3)
	high := teeth >> 3
	index := (teeth ^ (high - 1)) & 7
	for j := 0; j < 8; j++ {
		select_ := int(1 & (((int(j) ^ int(index)) - 1) >> 8))
		pc := combToPrecomp(&comb[j])
		Fe_ccopy(tmp_c.Yp[:], pc.Yp[:], select_)
		Fe_ccopy(tmp_c.Ym[:], pc.Ym[:], select_)
		Fe_ccopy(tmp_c.T2[:], pc.T2[:], select_)
	}
	Fe_neg(tmp_a, tmp_c.T2[:])
	Fe_cswap(tmp_c.T2[:], tmp_a, int(high^1))
	Fe_cswap(tmp_c.Yp[:], tmp_c.Ym[:], int(high^1))
	Ge_madd(p, p, tmp_c, tmp_a, tmp_b)
}

// Ge_scalarmult_base sets p = [scalar]B (base point).
func Ge_scalarmult_base(p *Ge, scalar []byte) {
	var s_scalar [32]byte
	Crypto_eddsa_mul_add(s_scalar[:], scalar, half_mod_L[:], half_ones[:])
	var tmp_a, tmp_b [10]int
	var tmp_c Ge_precomp
	var tmp_d Ge
	Fe_1(tmp_c.Yp[:])
	Fe_1(tmp_c.Ym[:])
	Fe_0(tmp_c.T2[:])
	Ge_zero(p)
	lookup_add(p, &tmp_c, tmp_a[:], tmp_b[:], &b_comb_low, s_scalar[:], 31)
	lookup_add(p, &tmp_c, tmp_a[:], tmp_b[:], &b_comb_high, s_scalar[:], 31+128)
	for i := 30; i >= 0; i-- {
		Ge_double(p, p, &tmp_d)
		lookup_add(p, &tmp_c, tmp_a[:], tmp_b[:], &b_comb_low, s_scalar[:], i)
		lookup_add(p, &tmp_c, tmp_a[:], tmp_b[:], &b_comb_high, s_scalar[:], i+128)
	}
}
