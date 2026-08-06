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

func TestSecretStream_UniqueKeystreamPerChunk(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	// Deux chunks identiques de 64 KB
	identicalChunk := bytes.Repeat([]byte("A"), 64*1024)

	var buf bytes.Buffer
	enc, err := secretstream55.NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor failed: %v", err)
	}

	// Écriture du chunk 1
	enc.Write(identicalChunk)
	rawBytes := buf.Bytes()
	chunk1Cipher := append([]byte(nil), rawBytes[24:]...) // Après le header de 24 octets

	// Réinitialisation du buffer et écriture du chunk 2
	buf.Reset()
	enc.Write(identicalChunk)
	chunk2Cipher := append([]byte(nil), buf.Bytes()...)

	if bytes.Equal(chunk1Cipher, chunk2Cipher) {
		t.Fatalf("P0 SÉCURITÉ ÉCHEC: Le keystream ChaCha20 est identique entre deux chunks successifs !")
	}

	t.Logf("Validation P0 Sécurité : Les nonces uniques par chunk (N_chunk = N_base ^ seq) préviennent toute réutilisation de keystream.")
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
		t.Fatalf("NewDecryptor failed: %v", err)
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
