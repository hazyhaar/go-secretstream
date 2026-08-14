package monocypher

// Hand: Poly_blocks — accéléré avec déroulage 2-Way ILP et réduction 32-bit sans overhead de radix.
func Poly_blocks(ctx *Crypto_poly1305_ctx, in []byte, nb_blocks uint64, end uint32) {
	var v4 uint64
	var v43 uint32
	var v47 uint32
	var v51 uint32
	var v55 uint32
	var v59 uint32
	var v63 uint64
	v4 = 0

	v10 := ctx.R[0]
	v14 := ctx.R[1]
	v18 := ctx.R[2]
	v22 := ctx.R[3]

	v28 := uint32((uint64(v14) >> 2) + uint64(v14))
	v32 := uint32((uint64(v18) >> 2) + uint64(v18))
	v36 := uint32((uint64(v22) >> 2) + uint64(v22))
	v10_5 := uint32((uint64(v10) >> 2) * 5)

	v43 = ctx.H[0]
	v47 = ctx.H[1]
	v51 = ctx.H[2]
	v55 = ctx.H[3]
	v59 = ctx.H[4]
	v63 = 0

	// Déroulage par paires de blocs (2-Way)
	for v63+2 <= nb_blocks {
		// Bloc 0
		v71_0 := uint64(v43) + uint64(Load32_le(in[int(v4):]))
		v78_0 := uint64(v47) + uint64(Load32_le(in[int(v4+4):]))
		v85_0 := uint64(v51) + uint64(Load32_le(in[int(v4+8):]))
		v92_0 := uint64(v55) + uint64(Load32_le(in[int(v4+12):]))
		v96_0 := v59 + end

		v106_0 := ((((v71_0 * uint64(v10)) + (v78_0 * uint64(v36))) + (v85_0 * uint64(v32))) + (v92_0 * uint64(v28))) + uint64(v96_0*v10_5)
		v116_0 := ((((v71_0 * uint64(v14)) + (v78_0 * uint64(v10))) + (v85_0 * uint64(v36))) + (v92_0 * uint64(v32))) + uint64(v96_0*v28)
		v126_0 := ((((v71_0 * uint64(v18)) + (v78_0 * uint64(v14))) + (v85_0 * uint64(v10))) + (v92_0 * uint64(v36))) + uint64(v96_0*v32)
		v136_0 := ((((v71_0 * uint64(v22)) + (v78_0 * uint64(v18))) + (v85_0 * uint64(v14))) + (v92_0 * uint64(v10))) + uint64(v96_0*v36)

		v139_0 := uint32(uint64(v96_0*(v10&3)) + (v136_0 >> 32))
		v150_0 := ((uint64(v139_0) >> 2) * 5) + (v106_0 & 0xffffffff)
		v159_0 := ((v150_0 >> 32) + (v116_0 & 0xffffffff)) + (v106_0 >> 32)
		v168_0 := ((v159_0 >> 32) + (v126_0 & 0xffffffff)) + (v116_0 >> 32)
		v177_0 := ((v168_0 >> 32) + (v136_0 & 0xffffffff)) + (v126_0 >> 32)

		v43_0 := uint32(v150_0)
		v47_0 := uint32(v159_0)
		v51_0 := uint32(v168_0)
		v55_0 := uint32(v177_0)
		v59_0 := uint32((v177_0 >> 32) + uint64(v139_0&3))

		// Bloc 1
		v71_1 := uint64(v43_0) + uint64(Load32_le(in[int(v4+16):]))
		v78_1 := uint64(v47_0) + uint64(Load32_le(in[int(v4+20):]))
		v85_1 := uint64(v51_0) + uint64(Load32_le(in[int(v4+24):]))
		v92_1 := uint64(v55_0) + uint64(Load32_le(in[int(v4+28):]))
		v96_1 := v59_0 + end

		v106_1 := ((((v71_1 * uint64(v10)) + (v78_1 * uint64(v36))) + (v85_1 * uint64(v32))) + (v92_1 * uint64(v28))) + uint64(v96_1*v10_5)
		v116_1 := ((((v71_1 * uint64(v14)) + (v78_1 * uint64(v10))) + (v85_1 * uint64(v36))) + (v92_1 * uint64(v32))) + uint64(v96_1*v28)
		v126_1 := ((((v71_1 * uint64(v18)) + (v78_1 * uint64(v14))) + (v85_1 * uint64(v10))) + (v92_1 * uint64(v36))) + uint64(v96_1*v32)
		v136_1 := ((((v71_1 * uint64(v22)) + (v78_1 * uint64(v18))) + (v85_1 * uint64(v14))) + (v92_1 * uint64(v10))) + uint64(v96_1*v36)

		v139_1 := uint32(uint64(v96_1*(v10&3)) + (v136_1 >> 32))
		v150_1 := ((uint64(v139_1) >> 2) * 5) + (v106_1 & 0xffffffff)
		v159_1 := ((v150_1 >> 32) + (v116_1 & 0xffffffff)) + (v106_1 >> 32)
		v168_1 := ((v159_1 >> 32) + (v126_1 & 0xffffffff)) + (v116_1 >> 32)
		v177_1 := ((v168_1 >> 32) + (v136_1 & 0xffffffff)) + (v126_1 >> 32)

		v43 = uint32(v150_1)
		v47 = uint32(v159_1)
		v51 = uint32(v168_1)
		v55 = uint32(v177_1)
		v59 = uint32((v177_1 >> 32) + uint64(v139_1&3))

		v4 += 32
		v63 += 2
	}

	// Bloc résiduel impair
	for v63 < nb_blocks {
		v71 := uint64(v43) + uint64(Load32_le(in[int(v4):]))
		v4 += 4
		v78 := uint64(v47) + uint64(Load32_le(in[int(v4):]))
		v4 += 4
		v85 := uint64(v51) + uint64(Load32_le(in[int(v4):]))
		v4 += 4
		v92 := uint64(v55) + uint64(Load32_le(in[int(v4):]))
		v4 += 4
		v96 := v59 + end
		v106 := ((((v71 * uint64(v10)) + (v78 * uint64(v36))) + (v85 * uint64(v32))) + (v92 * uint64(v28))) + uint64(v96*v10_5)
		v116 := ((((v71 * uint64(v14)) + (v78 * uint64(v10))) + (v85 * uint64(v36))) + (v92 * uint64(v32))) + uint64(v96*v28)
		v126 := ((((v71 * uint64(v18)) + (v78 * uint64(v14))) + (v85 * uint64(v10))) + (v92 * uint64(v36))) + uint64(v96*v32)
		v136 := ((((v71 * uint64(v22)) + (v78 * uint64(v18))) + (v85 * uint64(v14))) + (v92 * uint64(v10))) + uint64(v96*v36)
		v139 := uint32(uint64(v96*(v10&3)) + (v136 >> 32))
		v150 := ((uint64(v139) >> 2) * 5) + (v106 & 0xffffffff)
		v159 := ((v150 >> 32) + (v116 & 0xffffffff)) + (v106 >> 32)
		v168 := ((v159 >> 32) + (v126 & 0xffffffff)) + (v116 >> 32)
		v177 := ((v168 >> 32) + (v136 & 0xffffffff)) + (v126 >> 32)
		v43 = uint32(v150)
		v47 = uint32(v159)
		v51 = uint32(v168)
		v55 = uint32(v177)
		v59 = uint32((v177 >> 32) + uint64(v139&3))
		v63 += 1
	}

	ctx.H[0] = v43
	ctx.H[1] = v47
	ctx.H[2] = v51
	ctx.H[3] = v55
	ctx.H[4] = v59
}
