package monocypher_test

import (
	"bytes"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
)

func TestLockDst_MatchesAEADLock(t *testing.T) {
	key, nonce := keyNonce()
	cases := []struct {
		name string
		ad   []byte
		pt   []byte
	}{
		{"nil_nil", nil, nil},
		{"empty_empty", []byte{}, []byte{}},
		{"ad_empty_pt1", nil, []byte{0x41}},
		{"header_pt", []byte("HEADER"), []byte("HELLO MONOCYPHER")},
		{"pt65", []byte{}, mkPT(65, "i%251")},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctAlloc, macAlloc, err := sgoi.AEADLock(key, nonce, c.ad, c.pt)
			if err != nil {
				t.Fatal(err)
			}
			dst := make([]byte, len(c.pt)+8)
			var mac [16]byte
			if err := sgoi.LockDst(dst[:len(c.pt)], mac[:], key, nonce, c.ad, c.pt); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(dst[:len(c.pt)], ctAlloc) {
				t.Fatal("ct mismatch")
			}
			if !bytes.Equal(mac[:], macAlloc) {
				t.Fatal("mac mismatch")
			}
			ptOut := make([]byte, len(c.pt)+4)
			if err := sgoi.UnlockDst(ptOut[:len(c.pt)], key, nonce, mac[:], c.ad, dst[:len(c.pt)]); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(ptOut[:len(c.pt)], c.pt) {
				t.Fatal("roundtrip mismatch")
			}
		})
	}
}

func TestLockDst_NilEmptyNoPanic(t *testing.T) {
	key, nonce := keyNonce()
	var mac [16]byte
	if err := sgoi.LockDst(nil, mac[:], key, nonce, nil, nil); err != nil {
		t.Fatal(err)
	}
	if err := sgoi.UnlockDst(nil, key, nonce, mac[:], nil, nil); err != nil {
		t.Fatal(err)
	}
	if err := sgoi.LockDst([]byte{}, mac[:], key, nonce, []byte{}, []byte{}); err != nil {
		t.Fatal(err)
	}
}

func TestLockDst_RejectsShortBuffers(t *testing.T) {
	key, nonce := keyNonce()
	pt := []byte("abc")
	var mac [16]byte
	if err := sgoi.LockDst(make([]byte, 2), mac[:], key, nonce, nil, pt); err == nil {
		t.Fatal("short dstCT accepted")
	}
	if err := sgoi.LockDst(make([]byte, 3), mac[:5], key, nonce, nil, pt); err == nil {
		t.Fatal("short mac accepted")
	}
}

func TestLockDst_ZeroAlloc(t *testing.T) {
	key, nonce := keyNonce()
	pt := mkPT(1024, "i%251")
	ad := []byte("HEADER")
	dst := make([]byte, len(pt))
	var mac [16]byte
	n := testing.AllocsPerRun(200, func() {
		if err := sgoi.LockDst(dst, mac[:], key, nonce, ad, pt); err != nil {
			t.Fatal(err)
		}
	})
	if n != 0 {
		t.Fatalf("LockDst allocs = %v want 0", n)
	}
	ptOut := make([]byte, len(pt))
	n = testing.AllocsPerRun(200, func() {
		if err := sgoi.UnlockDst(ptOut, key, nonce, mac[:], ad, dst); err != nil {
			t.Fatal(err)
		}
	})
	if n != 0 {
		t.Fatalf("UnlockDst allocs = %v want 0", n)
	}
}
