package monocypher55

// Hand override — Invsqrt keeps intermediate Fe_isequal on reused check buffer.

func Invsqrt(isr *[10]int32, x *[10]int32) int {
	var v14 uint64
	var v22 uint64
	var v30 uint64
	var v38 uint64
	var v46 uint64
	var v54 uint64
	var v62 uint64
	var v70 uint64
	var _arr_v3 [10]int32
	v3 := &_arr_v3
	var _arr_v4 [10]int32
	v4 := &_arr_v4
	var _arr_v5 [10]int32
	v5 := &_arr_v5
	Fe_sq(v3, x)
	Fe_sq(v4, v3)
	Fe_sq(v4, v4)
	Fe_mul(v4, x, v4)
	Fe_mul(v3, v3, v4)
	Fe_sq(v3, v3)
	Fe_mul(v3, v4, v3)
	Fe_sq(v4, v3)
	v14 = 1
	for v14 < 5 {
		Fe_sq(v4, v4)
		v14 = v14 + 1
	}
	Fe_mul(v3, v4, v3)
	Fe_sq(v4, v3)
	v22 = 1
	for v22 < 10 {
		Fe_sq(v4, v4)
		v22 = v22 + 1
	}
	Fe_mul(v4, v4, v3)
	Fe_sq(v5, v4)
	v30 = 1
	for v30 < 20 {
		Fe_sq(v5, v5)
		v30 = v30 + 1
	}
	Fe_mul(v4, v5, v4)
	Fe_sq(v4, v4)
	v38 = 1
	for v38 < 10 {
		Fe_sq(v4, v4)
		v38 = v38 + 1
	}
	Fe_mul(v3, v4, v3)
	Fe_sq(v4, v3)
	v46 = 1
	for v46 < 50 {
		Fe_sq(v4, v4)
		v46 = v46 + 1
	}
	Fe_mul(v4, v4, v3)
	Fe_sq(v5, v4)
	v54 = 1
	for v54 < 0x64 {
		Fe_sq(v5, v5)
		v54 = v54 + 1
	}
	Fe_mul(v4, v5, v4)
	Fe_sq(v4, v4)
	v62 = 1
	for v62 < 50 {
		Fe_sq(v4, v4)
		v62 = v62 + 1
	}
	Fe_mul(v3, v4, v3)
	Fe_sq(v3, v3)
	v70 = 1
	for v70 < 2 {
		Fe_sq(v3, v3)
		v70 = v70 + 1
	}
	Fe_mul(v3, v3, x)
	// quartic = x^((p-1)/4) in v4; t0 exponent result in v3
	Fe_sq(v4, v3)
	Fe_mul(v4, v4, x)
	// Hand-fix: sgoiter elided intermediate Fe_isequal on reused check buffer.
	// C monocypher 4.0.2 invsqrt:
	//   z0 = (x==0); p1=(quartic==1); m1=(quartic==-1); ms=(quartic==-sqrtm1)
	//   isr = t0*sqrtm1; ccopy(isr, t0, 1-(m1|ms)); return p1|m1|z0
	Fe_0(v5)
	z0 := Fe_isequal(x, v5)
	Fe_1(v5)
	p1 := Fe_isequal(v4, v5)
	Fe_neg(v5, v5)
	m1 := Fe_isequal(v4, v5)
	Fe_neg(v5, &sqrtm1)
	ms := Fe_isequal(v4, v5)
	Fe_mul(isr, v3, &sqrtm1)
	Fe_ccopy(isr, v3, 1-(m1|ms))
	for _i := range v3 {
		v3[_i] = 0
	}
	for _i := range v4 {
		v4[_i] = 0
	}
	for _i := range v5 {
		v5[_i] = 0
	}
	return p1 | m1 | z0
}
