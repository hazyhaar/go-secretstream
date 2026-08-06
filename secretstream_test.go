package secretstream55_test

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"github.com/hazyhaar/go-secretstream"
)

func TestSecretStreamRoundtrip(t *testing.T) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatalf("rand.Read failed: %v", err)
	}

	payload := []byte("Highly confidential data stream encrypted using SIMD-accelerated XChaCha20-Poly1305 in Pure Go (CGO_ENABLED=0).")

	var encryptedBuf bytes.Buffer
	enc, err := secretstream55.NewEncryptor(&encryptedBuf, key)
	if err != nil {
		t.Fatalf("NewEncryptor failed: %v", err)
	}

	if _, err := enc.Write(payload); err != nil {
		t.Fatalf("enc.Write failed: %v", err)
	}

	dec, err := secretstream55.NewDecryptor(&encryptedBuf, key)
	if err != nil {
		t.Fatalf("NewDecryptor failed: %v", err)
	}

	decryptedBuf := make([]byte, len(payload)*2)
	n, err := dec.Read(decryptedBuf)
	if err != nil && err != io.EOF {
		t.Fatalf("dec.Read failed: %v", err)
	}
	decrypted := decryptedBuf[:n]

	if !bytes.Equal(decrypted, payload) {
		t.Fatalf("Decrypted stream mismatch!\nExpected: %q\nGot:      %q", string(payload), string(decrypted))
	}

	t.Logf("Successfully encrypted and decrypted %d bytes of stream payload in Pure Go (CGO_ENABLED=0)", n)
}

func TestLibsodiumFramingRoundtrip(t *testing.T) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatalf("rand.Read failed: %v", err)
	}

	payload := []byte("WAL-G archival stream encrypted with Libsodium crypto_secretstream C framing support in Pure Go.")

	var encryptedBuf bytes.Buffer
	enc, err := secretstream55.NewLibsodiumEncryptor(&encryptedBuf, key)
	if err != nil {
		t.Fatalf("NewLibsodiumEncryptor failed: %v", err)
	}

	if _, err := enc.Write(payload); err != nil {
		t.Fatalf("enc.Write failed: %v", err)
	}

	dec, err := secretstream55.NewLibsodiumDecryptor(&encryptedBuf, key)
	if err != nil {
		t.Fatalf("NewLibsodiumDecryptor failed: %v", err)
	}

	decryptedBuf := make([]byte, len(payload)*2)
	n, err := dec.Read(decryptedBuf)
	if err != nil && err != io.EOF {
		t.Fatalf("dec.Read failed: %v", err)
	}
	decrypted := decryptedBuf[:n]

	if !bytes.Equal(decrypted, payload) {
		t.Fatalf("Libsodium decrypted stream mismatch!\nExpected: %q\nGot:      %q", string(payload), string(decrypted))
	}

	t.Logf("Successfully verified Libsodium C framing compatibility (%d bytes) in Pure Go", n)
}
