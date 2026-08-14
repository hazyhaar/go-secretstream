//go:build !goexperiment.simd || !amd64

package monocypher

func hasAVX2() bool {
	return false
}

func chacha20_djb_simd_4x(cipher_text, plain_text, key, nonce []byte, ctr uint64, nb_blocks uint64) uint64 {
	return ctr
}

func aead_interleaved_write_simd(ctx *Crypto_aead_ctx, polyCtx *Crypto_poly1305_ctx, cipher_text, plain_text []byte, numChunks uint64, startCtr uint64) uint64 {
	return startCtr
}
