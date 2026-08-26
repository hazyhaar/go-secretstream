// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

// Hand: Crypto_chacha20_h — HChaCha20 avec voie SIMD 128-bit (archsimd.Uint32x4)
// sous Go 1.27 GOEXPERIMENT=simd et repli scalaire strictement identique au code émis.
func Crypto_chacha20_h(out []byte, key []byte, in []byte) {
	if hasAVX2() {
		hchacha20_simd128(out, key, in)
		return
	}
	crypto_chacha20_h_scalar(out, key, in)
}

// crypto_chacha20_h_scalar est le corps émis d'origine (chacha20.go), conservé
// comme repli et comme oracle de parité pour la voie SIMD.
func crypto_chacha20_h_scalar(out []byte, key []byte, in []byte) {
	var v4 [16]uint32
	load32_le_buf(v4[:], Chacha20_constant, uint64(4))
	load32_le_buf(v4[4:], key, uint64(8))
	v14 := v4[12:]
	load32_le_buf(v14, in, uint64(4))
	chacha20_rounds(v4[:], v4[:])
	store32_le_buf(out, v4[:], uint64(4))
	store32_le_buf(out[16:], v14, uint64(4))
	clear(v4[:])
}
