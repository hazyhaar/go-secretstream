package secretstream55_test

import (
	"crypto/rand"
	"testing"

	"code.hazyhaar.fr/devhoros/pkg/secretstream55/internal/monocypher"
	"golang.org/x/crypto/chacha20poly1305"
)

func BenchmarkAEAD_MonocypherTranspiled(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	ad := []byte("Header Metadata AD")
	payload := make([]byte, 64*1024) // 64 KB chunk
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(payload)

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

func BenchmarkAEAD_OfficialGoCrypto(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	ad := []byte("Header Metadata AD")
	payload := make([]byte, 64*1024) // 64 KB chunk
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(payload)

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
