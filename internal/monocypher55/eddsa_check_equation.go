// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

// Crypto_eddsa_check_equation — monocypher 4.0.2 sliding-window check.
func Crypto_eddsa_check_equation(signature, public_key, h []byte) int {
	const pWWidth = 3
	const pWSize = 1 << (pWWidth - 2) // 2
	const bWWidth = 5                 // monocypher B_W_WIDTH

	var minus_A, minus_R Ge
	s := signature[32:64]
	var s32 [8]uint32
	load32_le_buf(s32[:], s, 8)
	if ge_frombytes_neg_vartime(&minus_A, public_key) != 0 ||
		ge_frombytes_neg_vartime(&minus_R, signature) != 0 ||
		is_above_l(s32[:]) != 0 {
		return -1
	}

	var lutA [pWSize]Ge_cached
	var minus_A2, tmp Ge
	ge_double(&minus_A2, &minus_A, &tmp)
	ge_cache(&lutA[0], &minus_A)
	for i := 1; i < pWSize; i++ {
		ge_add(&tmp, &minus_A2, &lutA[i-1])
		ge_cache(&lutA[i], &tmp)
	}

	var h_slide, s_slide Slide_ctx
	Slide_init(&h_slide, h)
	Slide_init(&s_slide, s)
	i := int(h_slide.Next_check)
	if int(s_slide.Next_check) > i {
		i = int(s_slide.Next_check)
	}
	sum := &minus_A
	ge_zero(sum)
	var _arr_t1, _arr_t2 [10]int32
	t1, t2 := &_arr_t1, &_arr_t2
	for i >= 0 {
		ge_double(sum, sum, &tmp)
		h_digit := Slide_step(&h_slide, pWWidth, i, h)
		s_digit := Slide_step(&s_slide, bWWidth, i, s)
		if h_digit > 0 {
			ge_add(sum, sum, &lutA[h_digit/2])
		}
		if h_digit < 0 {
			ge_sub(sum, sum, &lutA[(-h_digit)/2])
		}
		if s_digit > 0 {
			pc := combToPrecomp(&b_window[s_digit/2])
			ge_madd(sum, sum, &pc, t1, t2)
		}
		if s_digit < 0 {
			pc := combToPrecomp(&b_window[(-s_digit)/2])
			ge_msub(sum, sum, &pc, t1, t2)
		}
		i--
	}

	var cached Ge_cached
	var check [32]byte
	zeroPoint := [32]byte{1}
	ge_cache(&cached, &minus_R)
	ge_add(sum, sum, &cached)
	ge_double(sum, sum, &minus_R)
	ge_double(sum, sum, &minus_R)
	ge_double(sum, sum, &minus_R)
	ge_tobytes(check[:], sum)
	return Crypto_verify32(check[:], zeroPoint[:])
}
