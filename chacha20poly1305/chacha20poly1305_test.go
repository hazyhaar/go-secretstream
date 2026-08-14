package chacha20poly1305_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/hazyhaar/go-secretstream/chacha20poly1305"
	xchacha "golang.org/x/crypto/chacha20poly1305"
)

func TestChaCha20Poly1305_Roundtrip(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	rand.Read(key)
	rand.Read(nonce)

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		t.Fatalf("New failed: %v", err)
	}

	for _, size := range []int{0, 1, 64, 256, 1024, 65536} {
		pt := make([]byte, size)
		rand.Read(pt)
		ad := []byte("associated data")

		sealed := aead.Seal(nil, nonce, pt, ad)
		opened, err := aead.Open(nil, nonce, sealed, ad)
		if err != nil {
			t.Fatalf("size=%d Open failed: %v", size, err)
		}
		if !bytes.Equal(pt, opened) {
			t.Fatalf("size=%d plaintext mismatch", size)
		}

		// Validation croisée avec x/crypto
		xAead, _ := xchacha.New(key)
		xOpened, err := xAead.Open(nil, nonce, sealed, ad)
		if err != nil {
			t.Fatalf("size=%d x/crypto failed to open our ciphertext: %v", size, err)
		}
		if !bytes.Equal(pt, xOpened) {
			t.Fatalf("size=%d x/crypto decrypted mismatch", size)
		}
	}
}

func BenchmarkCompare_AVX2_vs_XCrypto(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 12)
	rand.Read(key)
	rand.Read(nonce)

	simdAead, _ := chacha20poly1305.New(key)
	xkAead, _ := xchacha.New(key)

	for _, size := range []int{64, 1024, 64 * 1024, 1024 * 1024} {
		buf := make([]byte, size)
		rand.Read(buf)
		dst := make([]byte, 0, size+16)

		b.Run("PureGo_AVX2/size="+formatSize(size), func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = simdAead.Seal(dst[:0], nonce, buf, nil)
			}
		})

		b.Run("XCrypto_ASM/size="+formatSize(size), func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = xkAead.Seal(dst[:0], nonce, buf, nil)
			}
		})
	}
}

func formatSize(s int) string {
	switch {
	case s >= 1024*1024:
		return "1MB"
	case s >= 1024:
		return "1KB"
	default:
		return "64B"
	}
}
