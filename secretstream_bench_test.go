//go:build goexperiment.simd

package secretstream55_test

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"github.com/hazyhaar/c2simd"
	"github.com/hazyhaar/go-secretstream"
	"golang.org/x/crypto/chacha20poly1305"
)

func BenchmarkSecretStream55_PureGo_1MB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 1024*1024)
	rand.Read(payload)

	outBuf := make([]byte, 1024*1024+1024)
	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(outBuf[:0])
		enc, err := secretstream55.NewEncryptor(buf, key)
		if err != nil {
			b.Fatal(err)
		}
		_, err = enc.Write(payload)
		if err != nil {
			b.Fatal(err)
		}

		dec, err := secretstream55.NewDecryptor(buf, key)
		if err != nil {
			b.Fatal(err)
		}
		out := make([]byte, len(payload))
		_, err = io.ReadFull(dec, out)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkC2SIMD_Engine_1MB(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	ad := []byte("ad_header")
	payload := make([]byte, 1024*1024)
	rand.Read(key)
	rand.Read(nonce)
	rand.Read(payload)

	dstBuf := make([]byte, 1024*1024)
	var mac [16]byte

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := c2simd.AEADLockDst(dstBuf, &mac, key, nonce, ad, payload)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkXCrypto_Native_1MB(b *testing.B) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	ad := []byte("ad_header")
	payload := make([]byte, 1024*1024)
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
		_ = aead.Seal(nil, nonce, payload, ad)
	}
}
