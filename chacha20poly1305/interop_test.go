package chacha20poly1305

import (
	"bytes"
	"testing"

	xref "golang.org/x/crypto/chacha20poly1305"
)

// TestWireCompatXCrypto enforces byte-for-byte wire identity (ciphertext AND
// tag) with golang.org/x/crypto/chacha20poly1305.NewX, both directions.
func TestWireCompatXCrypto(t *testing.T) {
	key := make([]byte, KeySize)
	nonce := make([]byte, NonceSizeX)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(0x30 + i)
	}
	ours, err := NewX(key)
	if err != nil {
		t.Fatal(err)
	}
	ref, err := xref.NewX(key)
	if err != nil {
		t.Fatal(err)
	}
	sizes := []int{0, 1, 48, 64, 200, 255, 256, 257, 512, 65536, 1 << 20}
	ads := [][]byte{nil, []byte("additional data")}
	for _, n := range sizes {
		for _, ad := range ads {
			plain := make([]byte, n)
			for i := range plain {
				plain[i] = byte(i * 13)
			}
			got := ours.Seal(nil, nonce, plain, ad)
			want := ref.Seal(nil, nonce, plain, ad)
			if !bytes.Equal(got, want) {
				t.Fatalf("n=%d ad=%v : wire divergent", n, ad != nil)
			}
			back, err := ours.Open(nil, nonce, want, ad)
			if err != nil || !bytes.Equal(back, plain) {
				t.Fatalf("n=%d : Open(x/crypto wire) : %v", n, err)
			}
			refBack, err := ref.Open(nil, nonce, got, ad)
			if err != nil || !bytes.Equal(refBack, plain) {
				t.Fatalf("n=%d : x/crypto Open(our wire) : %v", n, err)
			}
		}
	}
}

func TestOpenRejectsForgery(t *testing.T) {
	key := make([]byte, KeySize)
	nonce := make([]byte, NonceSizeX)
	a, _ := NewX(key)
	ct := a.Seal(nil, nonce, make([]byte, 512), nil)
	for _, i := range []int{0, 100, len(ct) - 1} {
		bad := append([]byte(nil), ct...)
		bad[i] ^= 1
		if _, err := a.Open(nil, nonce, bad, nil); err == nil {
			t.Fatalf("forgery accepted at byte %d", i)
		}
	}
	if _, err := a.Open(nil, nonce, ct[:10], nil); err == nil {
		t.Fatal("short ciphertext accepted")
	}
}
