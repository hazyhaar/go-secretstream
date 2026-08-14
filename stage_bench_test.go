package secretstream55_test

import (
	"crypto/rand"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/poly1305"
)

func BenchmarkStage_Decomposition_64K(b *testing.B) {
	const size = 64 * 1024
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	nonce12 := make([]byte, 12)
	polyKey := make([]byte, 32)
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(polyKey)
	copy(nonce12, nonce[:12])

	pt := make([]byte, size)
	ct := make([]byte, size)
	for i := range pt {
		pt[i] = byte(i)
	}

	// --- 1. ChaCha20 Seul ---
	b.Run("1_ChaCha20_Only/x_crypto_asm", func(b *testing.B) {
		b.SetBytes(size)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			c, _ := chacha20.NewUnauthenticatedCipher(key, nonce12)
			c.XORKeyStream(ct, pt)
		}
	})

	b.Run("1_ChaCha20_Only/purego_simd", func(b *testing.B) {
		b.SetBytes(size)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			sgoi.Crypto_chacha20_djb(ct, pt, uint64(size), key, nonce[:8], 1)
		}
	})

	// --- 2. Poly1305 Seul ---
	b.Run("2_Poly1305_Only/x_crypto_asm", func(b *testing.B) {
		var pk [32]byte
		copy(pk[:], polyKey)
		var out [16]byte
		b.SetBytes(size)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			poly1305.Sum(&out, ct, &pk)
		}
	})

	b.Run("2_Poly1305_Only/purego_simd", func(b *testing.B) {
		var ctx sgoi.Crypto_poly1305_ctx
		var mac [16]byte
		b.SetBytes(size)
		b.ReportAllocs()
		b.ResetTimer()
		for b.Loop() {
			sgoi.Crypto_poly1305_init(&ctx, polyKey)
			sgoi.Crypto_poly1305_update(&ctx, ct, uint64(size))
			sgoi.Crypto_poly1305_final(&ctx, mac[:])
		}
	})
}
