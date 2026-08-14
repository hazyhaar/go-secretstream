//go:build goexperiment.simd

package secretstream55_test

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"code.hazyhaar.fr/devhoros/c2simd"
	"github.com/hazyhaar/go-secretstream"
	"golang.org/x/crypto/chacha20poly1305"
)

// BenchmarkSecretStream55_SteadyState_WriteOnly_1MB mesure l'écriture seule en régime permanent (0 allocation hot path)
func BenchmarkSecretStream55_SteadyState_WriteOnly_1MB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 1024*1024)
	rand.Read(payload)

	outBuf := make([]byte, 1024*1024+4096)
	buf := bytes.NewBuffer(outBuf[:0])
	enc, err := secretstream55.NewEncryptor(buf, key)
	if err != nil {
		b.Fatal(err)
	}

	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		_, err := enc.Write(payload)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSecretStream55_SteadyState_ReadOnly_1MB mesure la lecture seule en régime permanent (0 allocation hot path)
func BenchmarkSecretStream55_SteadyState_ReadOnly_1MB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 1024*1024)
	rand.Read(payload)

	var encBuf bytes.Buffer
	enc, err := secretstream55.NewEncryptor(&encBuf, key)
	if err != nil {
		b.Fatal(err)
	}
	if _, err := enc.Write(payload); err != nil {
		b.Fatal(err)
	}
	encryptedBytes := append([]byte(nil), encBuf.Bytes()...)

	plainDst := make([]byte, len(payload))
	b.SetBytes(int64(len(payload)))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(encryptedBytes)
		dec, err := secretstream55.NewDecryptor(r, key)
		if err != nil {
			b.Fatal(err)
		}
		_, err = io.ReadFull(dec, plainDst)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkSecretStream55_FullDuplex_1MB mesure le cycle complet chiffrement + déchiffrement en buffers préalloués
func BenchmarkSecretStream55_FullDuplex_1MB(b *testing.B) {
	key := make([]byte, 32)
	rand.Read(key)
	payload := make([]byte, 1024*1024)
	rand.Read(payload)

	outBuf := make([]byte, 1024*1024+4096)
	plainDst := make([]byte, len(payload))
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
		_, err = io.ReadFull(dec, plainDst)
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
