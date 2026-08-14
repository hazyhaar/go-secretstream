//go:build !goexperiment.simd || !amd64

package monocypher

func poly1305_blocks_simd_4x(ctx *Crypto_poly1305_ctx, in []byte, nb_blocks uint64) uint64 {
	return 0
}
