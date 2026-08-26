// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

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
// 1. Benchmark Monocypher Transpiled (C via ccgo v4 / Pure Go)
// -----------------------------------------------------------------------------

func BenchmarkMonocypher_64KB(b *testing.B) {
	key, nonce, ad, payload := generateTestData(64 * 1024)
	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipherText, mac, err := monocypher.AEADLock(key, nonce, ad, payload)
		if err != nil {
			b.Fatal(err)
		}
		_, err = monocypher.AEADUnlock(key, nonce, mac, ad, cipherText)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMonocypher_1MB(b *testing.B) {
	key, nonce, ad, payload := generateTestData(1024 * 1024)
	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipherText, mac, err := monocypher.AEADLock(key, nonce, ad, payload)
		if err != nil {
			b.Fatal(err)
		}
		_, err = monocypher.AEADUnlock(key, nonce, mac, ad, cipherText)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMonocypher_10MB(b *testing.B) {
	key, nonce, ad, payload := generateTestData(10 * 1024 * 1024)
	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipherText, mac, err := monocypher.AEADLock(key, nonce, ad, payload)
		if err != nil {
			b.Fatal(err)
		}
		_, err = monocypher.AEADUnlock(key, nonce, mac, ad, cipherText)
		if err != nil {
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

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipherText := aead.Seal(nil, nonce, payload, ad)
		_, err := aead.Open(nil, nonce, cipherText, ad)
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

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipherText := aead.Seal(nil, nonce, payload, ad)
		_, err := aead.Open(nil, nonce, cipherText, ad)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNativeGo_10MB(b *testing.B) {
	key, nonce, ad, payload := generateTestData(10 * 1024 * 1024)
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		b.Fatal(err)
	}

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cipherText := aead.Seal(nil, nonce, payload, ad)
		_, err := aead.Open(nil, nonce, cipherText, ad)
		if err != nil {
			b.Fatal(err)
		}
	}
}
