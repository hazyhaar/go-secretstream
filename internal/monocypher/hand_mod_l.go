package monocypher

// Hand overrides — Barrett r + fixed Remove_l mask (C monocypher 4.0.2).

func Remove_l(out []uint32, x []uint32) {
	// Hand-fix: mask must be fixed from is_above_l before the loop (C monocypher 4.0.2).
	carry := uint64(Is_above_l(x))
	mask := ^uint32(carry) + 1 // carry 0 → 0; carry 1 → 0xffffffff
	for i := 0; i < 8; i++ {
		carry += uint64(x[i]) + uint64(^L[i]&mask)
		out[i] = uint32(carry)
		carry >>= 32
	}
}

func Mod_l(reduced []byte, x []uint32) {
	var v7 uint64
	var v12 uint64
	var v14 uint64
	var v34 uint64
	var v40 uint64
	var v45 uint64
	var v47 uint64
	var v67 uint64
	var v69 uint64
	// Hand-fix: use package Barrett constant r (sgoiter emitted zero local v5).
	var _arr_v6 [25]uint32
	v6 := _arr_v6[:]
	v7 = 0
	for v7 < 9 {
		v12 = 0
		v14 = 0
		for v14 < 16 {
			v19 := v7 + v14
			v12 = v12 + (uint64(v6[int(v19)]) + (uint64(r[int(v7)]) * uint64(x[int(v14)])))
			v6[int(v19)] = uint32(v12)
			v12 = v12 >> 32
			v14 = v14 + 1
		}
		v6[int(v7 + 16)] = uint32(v12)
		v7 = v7 + 1
	}
	v34 = 0
	for v34 < 8 {
		v6[int(v34)] = uint32(0)
		v34 = v34 + 1
	}
	v40 = 0
	for v40 < 8 {
		v45 = 0
		v47 = 0
		for v47 < (uint64(8) - v40) {
			v53 := v40 + v47
			v45 = v45 + (uint64(v6[int(v53)]) + (uint64(v6[int(v40 + 16)]) * uint64(L[int(v47)])))
			v6[int(v53)] = uint32(v45)
			v45 = v45 >> 32
			v47 = v47 + 1
		}
		v40 = v40 + 1
	}
	v67 = 1
	v69 = 0
	for v69 < 8 {
		v67 = v67 + (uint64(x[int(v69)]) + uint64(^v6[int(v69)] & uint32(0xffffffff)))
		v6[int(v69)] = uint32(v67)
		v67 = v67 >> 32
		v69 = v69 + 1
	}
	Remove_l(v6, v6)
	Store32_le_buf(reduced, v6, uint64(8))
	for _i := range v6 { v6[_i] = 0 }
}
