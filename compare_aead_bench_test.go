// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

import (
	"bytes"
	"fmt"
	"testing"

	ccgo "github.com/hazyhaar/go-secretstream/internal/monocypher"
	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
	"golang.org/x/crypto/chacha20poly1305"
)

func cmpKeyNonce() (key, nonce []byte) {
	key = make([]byte, 32)
	nonce = make([]byte, 24)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	return
}

func cmpPT(n int) []byte {
	pt := make([]byte, n)
	for i := range pt {
		pt[i] = byte((i*17 + 3) % 251)
	}
	return pt
}

var testSizes = []struct {
	name string
	size int
}{
	{"64B", 64},
	{"1KB", 1024},
	{"64KB", 64 * 1024},
	{"1MB", 1024 * 1024},
}

// BenchmarkCompare_AEAD_ZeroAlloc benchmarks Seal & Open without heap allocations.
func BenchmarkCompare_AEAD_ZeroAlloc(b *testing.B) {
	key, nonce := cmpKeyNonce()
	ad := []byte("HEADER_AD")

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		b.Fatal(err)
	}

	for _, tc := range testSizes {
		pt := cmpPT(tc.size)
		dstCT := make([]byte, tc.size)
		dstPT := make([]byte, tc.size)
		var mac [16]byte

		// Pre-seal for open benches
		sealedX := aead.Seal(nil, nonce, pt, ad)
		if err := sgoi.LockDst(dstCT, mac[:], key, nonce, ad, pt); err != nil {
			b.Fatal(err)
		}

		b.Run(fmt.Sprintf("Seal/purego_avx2/%s", tc.name), func(b *testing.B) {
			b.SetBytes(int64(tc.size))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_ = sgoi.LockDst(dstCT, mac[:], key, nonce, ad, pt)
			}
		})

		b.Run(fmt.Sprintf("Seal/xcrypto_asm/%s", tc.name), func(b *testing.B) {
			buf := make([]byte, 0, tc.size+16)
			b.SetBytes(int64(tc.size))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_ = aead.Seal(buf[:0], nonce, pt, ad)
			}
		})

		b.Run(fmt.Sprintf("Open/purego_avx2/%s", tc.name), func(b *testing.B) {
			b.SetBytes(int64(tc.size))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_ = sgoi.UnlockDst(dstPT, key, nonce, mac[:], ad, dstCT)
			}
		})

		b.Run(fmt.Sprintf("Open/xcrypto_asm/%s", tc.name), func(b *testing.B) {
			buf := make([]byte, 0, tc.size)
			b.SetBytes(int64(tc.size))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_, _ = aead.Open(buf[:0], nonce, sealedX, ad)
			}
		})
	}
}

// BenchmarkCompare_AEAD_Allocating benchmarks one-shot Seal & Open with standard return-slice allocations.
func BenchmarkCompare_AEAD_Allocating(b *testing.B) {
	key, nonce := cmpKeyNonce()
	ad := []byte("HEADER_AD")

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		b.Fatal(err)
	}

	for _, tc := range testSizes {
		pt := cmpPT(tc.size)
		sealedX := aead.Seal(nil, nonce, pt, ad)
		ctSgoi, macSgoi, _ := sgoi.AEADLock(key, nonce, ad, pt)

		b.Run(fmt.Sprintf("Lock/purego_avx2/%s", tc.name), func(b *testing.B) {
			b.SetBytes(int64(tc.size))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_, _, _ = sgoi.AEADLock(key, nonce, ad, pt)
			}
		})

		b.Run(fmt.Sprintf("Lock/xcrypto_asm/%s", tc.name), func(b *testing.B) {
			b.SetBytes(int64(tc.size))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_ = aead.Seal(nil, nonce, pt, ad)
			}
		})

		b.Run(fmt.Sprintf("Unlock/purego_avx2/%s", tc.name), func(b *testing.B) {
			b.SetBytes(int64(tc.size))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_, _ = sgoi.AEADUnlock(key, nonce, macSgoi, ad, ctSgoi)
			}
		})

		b.Run(fmt.Sprintf("Unlock/xcrypto_asm/%s", tc.name), func(b *testing.B) {
			b.SetBytes(int64(tc.size))
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_, _ = aead.Open(nil, nonce, sealedX, ad)
			}
		})
	}
}

func BenchmarkCompare_Blake2b_1K(b *testing.B) {
	msg := cmpPT(1024)
	var h [64]byte
	b.Run("sgoiter", func(b *testing.B) {
		b.SetBytes(1024)
		b.ReportAllocs()
		for b.Loop() {
			sgoi.Crypto_blake2b(h[:], 64, msg, 1024)
		}
	})
}

func BenchmarkCompare_X25519(b *testing.B) {
	sk := bytes.Repeat([]byte{7}, 32)
	var pk, out [32]byte
	sgoi.Crypto_x25519_public_key(pk[:], sk)
	b.Run("sgoiter", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			sgoi.Crypto_x25519(out[:], sk, pk[:])
		}
	})
}

func BenchmarkCompare_CCGO_Legacy(b *testing.B) {
	const n = 1024
	key, nonce := cmpKeyNonce()
	pt := cmpPT(n)
	ad := []byte("HDR")
	b.Run("ccgo_1K", func(b *testing.B) {
		b.SetBytes(n)
		b.ReportAllocs()
		for b.Loop() {
			_, _, _ = ccgo.AEADLock(key, nonce, ad, pt)
		}
	})
}

// BenchmarkCompare_AEAD_Parallel évalue la scalabilité multicœur sous b.RunParallel
func BenchmarkCompare_AEAD_Parallel(b *testing.B) {
	key, nonce := cmpKeyNonce()
	ad := []byte("HEADER_AD")

	for _, size := range []int{1024, 64 * 1024} {
		pt := cmpPT(size)
		name := fmt.Sprintf("%dKB", size/1024)

		b.Run("Seal/purego_parallel/"+name, func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ReportAllocs()
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				dstCT := make([]byte, size)
				var mac [16]byte
				for pb.Next() {
					_ = sgoi.LockDst(dstCT, mac[:], key, nonce, ad, pt)
				}
			})
		})
	}
}

// BenchmarkCompare_AEAD_ADHeavy évalue la charge avec métadonnées volumineuses
func BenchmarkCompare_AEAD_ADHeavy(b *testing.B) {
	key, nonce := cmpKeyNonce()
	cases := []struct {
		name   string
		adSize int
		ptSize int
	}{
		{"1KB_AD_64B_PT", 1024, 64},
		{"64KB_AD_64B_PT", 64 * 1024, 64},
		{"64KB_AD_64KB_PT", 64 * 1024, 64 * 1024},
	}

	for _, tc := range cases {
		ad := cmpPT(tc.adSize)
		pt := cmpPT(tc.ptSize)
		dstCT := make([]byte, tc.ptSize)
		var mac [16]byte
		totalBytes := int64(tc.adSize + tc.ptSize)

		b.Run(tc.name, func(b *testing.B) {
			b.SetBytes(totalBytes)
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				_ = sgoi.LockDst(dstCT, mac[:], key, nonce, ad, pt)
			}
		})
	}
}
