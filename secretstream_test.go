// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

import (
	"bytes"
	"crypto/rand"
	"io"
	"strings"
	"testing"

	secretstream55 "github.com/hazyhaar/go-secretstream"
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
	if err := enc.Close(); err != nil {
		t.Fatalf("enc.Close failed: %v", err)
	}
	dec, err := secretstream55.NewDecryptor(&encryptedBuf, key)
	if err != nil {
		t.Fatalf("NewDecryptor failed: %v", err)
	}
	decrypted, err := io.ReadAll(dec)
	if err != nil {
		t.Fatalf("dec.Read failed: %v", err)
	}
	if !bytes.Equal(decrypted, payload) {
		t.Fatalf("Decrypted stream mismatch!\nExpected: %q\nGot:      %q", payload, decrypted)
	}

	var truncated bytes.Buffer
	enc2, err := secretstream55.NewEncryptor(&truncated, key)
	if err != nil {
		t.Fatalf("NewEncryptor (sans Close): %v", err)
	}
	if _, err := enc2.Write(payload); err != nil {
		t.Fatalf("Write (sans Close): %v", err)
	}
	dec2, err := secretstream55.NewDecryptor(bytes.NewReader(truncated.Bytes()), key)
	if err != nil {
		t.Fatalf("NewDecryptor (sans Close): %v", err)
	}
	_, err = io.ReadAll(dec2)
	if err == nil || err == io.EOF {
		t.Fatalf("sans Close: attendu flux tronqué, obtenu %v", err)
	}
	if !strings.Contains(err.Error(), "flux tronqué") {
		t.Fatalf("sans Close: attendu flux tronqué, obtenu %v", err)
	}
}

func TestSecretStream_UniqueKeystreamPerChunk(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	identicalChunk := bytes.Repeat([]byte("A"), 64*1024)

	var buf bytes.Buffer
	enc, err := secretstream55.NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor failed: %v", err)
	}
	enc.Write(identicalChunk)
	rawBytes := buf.Bytes()
	chunk1Cipher := append([]byte(nil), rawBytes[24:]...)

	buf.Reset()
	enc.Write(identicalChunk)
	chunk2Cipher := append([]byte(nil), buf.Bytes()...)

	if bytes.Equal(chunk1Cipher, chunk2Cipher) {
		t.Fatalf("P0 SÉCURITÉ ÉCHEC: keystream identique entre deux chunks")
	}
}

func TestLibsodiumFramingRoundtrip(t *testing.T) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	payload := []byte("WAL-G archival stream encrypted with Libsodium crypto_secretstream C framing support in Pure Go.")

	var encryptedBuf bytes.Buffer
	enc, err := secretstream55.NewLibsodiumEncryptor(&encryptedBuf, key)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := enc.Write(payload); err != nil {
		t.Fatal(err)
	}
	if err := enc.Close(); err != nil {
		t.Fatal(err)
	}
	dec, err := secretstream55.NewLibsodiumDecryptor(bytes.NewReader(encryptedBuf.Bytes()), key)
	if err != nil {
		t.Fatal(err)
	}
	decrypted, err := io.ReadAll(dec)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(decrypted, payload) {
		t.Fatalf("Libsodium decrypted mismatch\nwant %q\ngot  %q", payload, decrypted)
	}
}

func TestLibsodiumMultiChunkSizes(t *testing.T) {
	key := bytes.Repeat([]byte{0x42}, 32)
	sizes := []int{0, 1, 15, 16, 63, 64, 65, 1000, 8191, 8192, 8193, 20000, 65536, 100000}
	for _, sz := range sizes {
		sz := sz
		t.Run(itoa(sz), func(t *testing.T) {
			plain := make([]byte, sz)
			for i := range plain {
				plain[i] = byte(i % 251)
			}
			var buf bytes.Buffer
			w, err := secretstream55.NewLibsodiumEncryptor(&buf, key)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := w.Write(plain); err != nil {
				t.Fatal(err)
			}
			if err := w.Close(); err != nil {
				t.Fatal(err)
			}
			r, err := secretstream55.NewLibsodiumDecryptor(bytes.NewReader(buf.Bytes()), key)
			if err != nil {
				t.Fatal(err)
			}
			got, err := io.ReadAll(r)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, plain) {
				t.Fatalf("n=%d mismatch", sz)
			}
		})
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b [16]byte
	i := len(b)
	for n > 0 {
		i--
		b[i] = byte('0' + n%10)
		n /= 10
	}
	return string(b[i:])
}

func TestSecretStream_AntiDoS_ExcessiveFrameLength(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	// Construction d'un flux malicieux : Header valide de 24 octets suivi d'un chunkLen forgé de 2 Gio (0x7fffffff)
	var malicious bytes.Buffer
	malicious.Write(bytes.Repeat([]byte{0x42}, 24))  // Header nonce
	malicious.Write([]byte{0x7f, 0xff, 0xff, 0xff})  // 2 Gio
	malicious.Write(bytes.Repeat([]byte{0x00}, 100)) // Début payload

	dec, err := secretstream55.NewDecryptor(&malicious, key)
	if err != nil {
		t.Fatalf("NewDecryptor failed: %v", err)
	}
	var dst [100]byte
	_, err = dec.Read(dst[:])
	if err == nil {
		t.Fatal("attendu échec anti-DoS sur chunkLen excessif, obtenu nil")
	}
	if !bytes.Contains([]byte(err.Error()), []byte("invalid payload length")) {
		t.Fatalf("attendu message invalid payload length, obtenu: %v", err)
	}
}

type shortWriter struct {
	limit int
	buf   bytes.Buffer
}

func (s *shortWriter) Write(p []byte) (int, error) {
	if len(p) > s.limit {
		n, _ := s.buf.Write(p[:s.limit])
		return n, nil // Short write sans erreur explicite
	}
	return s.buf.Write(p)
}

func TestSecretStream_ShortWriteDetected(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	sw := &shortWriter{limit: 10}
	enc, err := secretstream55.NewEncryptor(sw, key)
	if err != nil {
		t.Fatalf("NewEncryptor failed: %v", err)
	}
	_, err = enc.Write([]byte("message de test qui dépasse la limite d'écriture de 10 octets"))
	if err == nil {
		t.Fatal("attendu erreur sur short write non détecté, obtenu nil")
	}
}
