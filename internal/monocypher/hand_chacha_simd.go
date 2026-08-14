package monocypher

// Hand: Crypto_chacha20_djb — accéléré en SIMD 4-blocks (archsimd.Uint32x8) sous Go 1.27 avec fallback scalaire.
func Crypto_chacha20_djb(cipher_text []byte, plain_text []byte, text_size uint64, key []byte, nonce []byte, ctr uint64) uint64 {
	var v6 uint64
	var v8 uint64
	var v35 uint64
	var v41 uint64
	var v59 uint64
	var v73 uint32
	var v81 uint32
	var v91 uint64
	var v104 uint64
	v6 = 0
	v8 = v6
	var _arr_v12 [16]uint32
	v12 := _arr_v12[:]
	Load32_le_buf(v12, chacha20_constant, uint64(4))
	Load32_le_buf(v12[4:], key, uint64(8))
	Load32_le_buf(v12[14:], nonce, uint64(2))
	v12[12] = uint32(ctr)
	v12[13] = uint32(ctr >> 32)
	var _arr_v31 [16]uint32
	v31 := _arr_v31[:]
	v34 := text_size >> 6
	v35 = 0

	// Traitement vectoriel SIMD par lots de 4 blocs (256 octets) si disponible et supporté
	if hasAVX2() && v34 >= 4 {
		simdBlocks := (v34 / 4) * 4
		simdBytes := simdBlocks * 64
		newCtr := chacha20_djb_simd_4x(cipher_text, plain_text, key, nonce, ctr, simdBlocks)
		if newCtr != ctr {
			v6 += simdBytes
			v8 += simdBytes
			v35 += simdBlocks
			ctr = newCtr
			v12[12] = uint32(ctr)
			v12[13] = uint32(ctr >> 32)
		}
	}

	for v35 < v34 {
		Chacha20_rounds(v31, v12)
		if plain_text != nil {
			v41 = 0
			for v41 < 16 {
				Store32_le(cipher_text[int(v6):], ((v31[int(v41)] + v12[int(v41)]) ^ Load32_le(plain_text[int(v8):])))
				v6 = v6 + 4
				v8 = v8 + 4
				v41 = v41 + 1
			}
		} else {
			v59 = 0
			for v59 < 16 {
				Store32_le(cipher_text[int(v6):], (v31[int(v59)] + v12[int(v59)]))
				v6 = v6 + 4
				v59 = v59 + 1
			}
		}
		v73 = v12[12]
		v75 := v73 + 1
		v73 = v75
		v12[12] = v75
		if v12[12] == uint32(0) {
			v81 = v12[13]
			v83 := v81 + 1
			v81 = v83
			v12[13] = v83
		}
		v35 = v35 + 1
	}
	text_size = text_size & 63
	if text_size > 0 {
		if plain_text == nil {
			plain_text = zero
		}
		Chacha20_rounds(v31, v12)
		var _arr_v90 [64]byte
		v90 := _arr_v90[:]
		v91 = 0
		for v91 < 16 {
			Store32_le(v90[int(v91 * 4):], (v31[int(v91)] + v12[int(v91)]))
			v91 = v91 + 1
		}
		v104 = 0
		for v104 < text_size {
			cipher_text[int(v6 + v104)] = v90[int(v104)] ^ plain_text[int(v8 + v104)]
			v104 = v104 + 1
		}
		for _i := range v90 { v90[_i] = 0 }
	}
	resCtr := (uint64(v12[12]) | (uint64(v12[13]) << 32))
	if text_size > 0 {
		resCtr++
	}
	for _i := range v12 { v12[_i] = 0 }
	for _i := range v31 { v31[_i] = 0 }
	return resCtr
}
