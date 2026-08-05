package monocypher_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	"code.hazyhaar.fr/devhoros/pkg/secretstream55/internal/monocypher"
)

func TestTranspiledMonocypherAEAD(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	if _, err := rand.Read(key); err != nil {
		t.Fatalf("rand.Read key failed: %v", err)
	}
	if _, err := rand.Read(nonce); err != nil {
		t.Fatalf("rand.Read nonce failed: %v", err)
	}

	ad := []byte("Header Metadata AD")
	message := []byte("Top secret payload encrypted via transpiled Monocypher in Pure Go!")

	cipherText, mac, err := monocypher.AEADLock(key, nonce, ad, message)
	if err != nil {
		t.Fatalf("AEADLock failed: %v", err)
	}

	decrypted, err := monocypher.AEADUnlock(key, nonce, mac, ad, cipherText)
	if err != nil {
		t.Fatalf("AEADUnlock failed: %v", err)
	}

	if !bytes.Equal(decrypted, message) {
		t.Fatalf("Decrypted message mismatch!\nExpected: %q\nGot:      %q", string(message), string(decrypted))
	}

	// Verify authentication tag failure on tampered ciphertext
	tamperedCipher := append([]byte(nil), cipherText...)
	tamperedCipher[0] ^= 0xFF
	_, err = monocypher.AEADUnlock(key, nonce, mac, ad, tamperedCipher)
	if err == nil {
		t.Fatal("Expected authentication failure on tampered ciphertext, but got success!")
	}

	t.Logf("Successfully verified XChaCha20-Poly1305 AEAD lock/unlock using transpiled Monocypher in Pure Go (CGO_ENABLED=0)")
}
