package secretstream_test

import (
	"crypto/rand"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/monocypher"
	"golang.org/x/crypto/chacha20poly1305"
)

func BenchmarkAEAD_MonocypherSIMD_ZeroAlloc(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	ad := []byte("Header Metadata AD")
	payload := make([]byte, 64*1024)
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(payload)

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

func BenchmarkAEAD_XCrypto_Standard(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	ad := []byte("Header Metadata AD")
	payload := make([]byte, 64*1024)
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(payload)

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		b.Fatal(err)
	}

	dst := make([]byte, 0, len(payload)+16)

	b.SetBytes(int64(len(payload)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sealed := aead.Seal(dst[:0], nonce, payload, ad)
		_, err := aead.Open(dst[:0], nonce, sealed, ad)
		if err != nil {
			b.Fatal(err)
		}
	}
}
