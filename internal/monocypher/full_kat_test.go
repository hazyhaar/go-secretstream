package monocypher_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
)

func TestBlake2b_RoundtripSizes(t *testing.T) {
	for _, n := range []int{0, 1, 63, 64, 65, 128, 1024} {
		msg := make([]byte, n)
		rand.Read(msg)
		var h1, h2 [64]byte
		sgoi.Crypto_blake2b(h1[:], 64, msg, uint64(n))
		var ctx sgoi.Crypto_blake2b_ctx
		sgoi.Crypto_blake2b_init(&ctx, 64)
		sgoi.Crypto_blake2b_update(&ctx, msg, uint64(n))
		sgoi.Crypto_blake2b_final(&ctx, h2[:])
		if !bytes.Equal(h1[:], h2[:]) {
			t.Fatalf("n=%d one-shot != stream", n)
		}
	}
}

func TestX25519_SharedSecret(t *testing.T) {
	var ska, skb, pka, pkb, sa, sb [32]byte
	rand.Read(ska[:])
	rand.Read(skb[:])
	sgoi.Crypto_x25519_public_key(pka[:], ska[:])
	sgoi.Crypto_x25519_public_key(pkb[:], skb[:])
	sgoi.Crypto_x25519(sa[:], ska[:], pkb[:])
	sgoi.Crypto_x25519(sb[:], skb[:], pka[:])
	if !bytes.Equal(sa[:], sb[:]) {
		t.Fatal("dh mismatch")
	}
}

func TestElligator_MapNoPanic(t *testing.T) {
	var hidden, curve [32]byte
	rand.Read(hidden[:])
	sgoi.Crypto_elligator_map(curve[:], hidden[:])
}
