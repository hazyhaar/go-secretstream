package monocypher55

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func keyedHash32(key, msg []byte) []byte {
	out := make([]byte, 32)
	var ctx Crypto_blake2b_ctx
	Crypto_blake2b_keyed_init(&ctx, 32, key, uint64(len(key)))
	Crypto_blake2b_update(&ctx, msg, uint64(len(msg)))
	Crypto_blake2b_final(&ctx, out)
	return out
}

// TestBlake2bKeyedNoResidue verrouille le défaut mesuré le 2026-08-15 : le tampon
// de clé était une globale jamais nettoyée ; un hachage keyed à clé courte,
// exécuté après un keyed à clé longue, hachait les résidus de la clé longue dans
// son padding et rendait un condensat FAUX. Le C d'origine zérote ce tampon à
// chaque appel (uint8_t key_block[128] = {0}).
func TestBlake2bKeyedNoResidue(t *testing.T) {
	longKey := bytes.Repeat([]byte{0xAA}, 64)
	shortKey := []byte("cle-courte-16-oc")
	msg := []byte("message")

	cold := keyedHash32(shortKey, msg)

	_ = keyedHash32(longKey, msg)
	after := keyedHash32(shortKey, msg)

	if !bytes.Equal(cold, after) {
		t.Fatalf("résidu de clé longue dans le padding : froid %x, après clé longue %x", cold, after)
	}

	// Ancrage externe : condensat calculé par golang.org/x/crypto/blake2b.New256
	// (même clé, même message) le 2026-08-15 — fige la valeur absolue, pas
	// seulement l'égalité interne.
	want, _ := hex.DecodeString("75513d30293425730ed8e58ee474ead4f12c26f2e2755dbe0e9ebe74bdf29051")
	if !bytes.Equal(cold, want) {
		t.Fatalf("condensat keyed divergent de la référence x/crypto : got %x want %x", cold, want)
	}
}
