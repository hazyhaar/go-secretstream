package secretstream_test

import (
	"crypto/rand"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/monocypher"
	"golang.org/x/crypto/chacha20poly1305"
)

func generateTestData(size int) ([]byte, []byte, []byte, []byte) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	ad := []byte("header_metadata")
	payload := make([]byte, size)

	rand.Read(key)
	rand.Read(nonce)
	rand.Read(payload)

	return key, nonce, ad, payload
}

// -----------------------------------------------------------------------------
// 1. Benchmark Monocypher Pure Go SIMD (AVX2 / 0 Alloc)
// -----------------------------------------------------------------------------

func BenchmarkMonocypher_64KB(b *testing.B) {
	key, nonce, ad, payload := generateTestData(64 * 1024)
	dstCT := make([]byte, len(payload))
	mac := make([]byte, 16)
	dstPT := make([]byte, len(payload))

	b.SetBytes(int64(len(payload)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := monocypher.LockDst(dstCT, mac, key, nonce, ad, payload); err != nil {
			b.Fatal(err)
		}
		if err := monocypher.UnlockDst(dstPT, key, nonce, mac, ad, dstCT); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMonocypher_1MB(b *testing.B) {
	key, nonce, ad, payload := generateTestData(1024 * 1024)
	dstCT := make([]byte, len(payload))
	mac := make([]byte, 16)
	dstPT := make([]byte, len(payload))

	b.SetBytes(int64(len(payload)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := monocypher.LockDst(dstCT, mac, key, nonce, ad, payload); err != nil {
			b.Fatal(err)
		}
		if err := monocypher.UnlockDst(dstPT, key, nonce, mac, ad, dstCT); err != nil {
			b.Fatal(err)
		}
	}
}

// -----------------------------------------------------------------------------
// 2. Benchmark Native Go SIMD (golang.org/x/crypto/chacha20poly1305)
// -----------------------------------------------------------------------------

func BenchmarkNativeGo_64KB(b *testing.B) {
	key, nonce, ad, payload := generateTestData(64 * 1024)
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		b.Fatal(err)
	}
	dst := make([]byte, 0, len(payload)+16)

	b.SetBytes(int64(len(payload)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipherText := aead.Seal(dst[:0], nonce, payload, ad)
		_, err := aead.Open(dst[:0], nonce, cipherText, ad)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNativeGo_1MB(b *testing.B) {
	key, nonce, ad, payload := generateTestData(1024 * 1024)
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		b.Fatal(err)
	}
	dst := make([]byte, 0, len(payload)+16)

	b.SetBytes(int64(len(payload)))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipherText := aead.Seal(dst[:0], nonce, payload, ad)
		_, err := aead.Open(dst[:0], nonce, cipherText, ad)
		if err != nil {
			b.Fatal(err)
		}
	}
}
