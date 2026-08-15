//go:build !goexperiment.simd || !amd64

package monocypher55

// Stubs de repli : jamais appelés, tout site d'appel est gardé par hasAVX2()
// qui retourne false dans cette configuration de build.

func hchacha20_simd128(out []byte, key []byte, in []byte) {
	crypto_chacha20_h_scalar(out, key, in)
}

func chacha20_deriv2_simd256(dst []byte, key []byte, nonce []byte, ctr uint64) {
	Crypto_chacha20_djb(dst[:128], nil, 128, key, nonce, ctr)
}

func chacha20_xor2_simd256(dst []byte, src []byte, key []byte, nonce []byte, ctr uint64) {
	Crypto_chacha20_djb(dst[:128], src, 128, key, nonce, ctr)
}
