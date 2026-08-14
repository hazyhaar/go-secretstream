//go:build !goexperiment.simd || !amd64

package monocypher

func hasAVX2() bool {
	return false
}

func chacha20_djb_simd_4x(cipher_text, plain_text, key, nonce []byte, ctr uint64, nb_blocks uint64) uint64 {
	return ctr
}
