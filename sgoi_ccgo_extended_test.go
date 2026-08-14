package secretstream55_test

import (
	"bytes"
	"testing"

	ccgo "github.com/hazyhaar/go-secretstream/internal/monocypher"
	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
)

func keyNonce() (key, nonce []byte) {
	key = make([]byte, 32)
	nonce = make([]byte, 24)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	return key, nonce
}

func mkPT(n int, pattern string) []byte {
	pt := make([]byte, n)
	for i := range pt {
		switch pattern {
		case "(i*17+3)%251":
			pt[i] = byte((i*17 + 3) % 251)
		default:
			pt[i] = byte(i % 251)
		}
	}
	return pt
}

func assertParity(t *testing.T, name string, ad, pt []byte) {
	t.Helper()
	key, nonce := keyNonce()
	ctS, macS, err := sgoi.AEADLock(key, nonce, ad, pt)
	if err != nil {
		t.Fatalf("%s sgoi lock: %v", name, err)
	}
	ctC, macC, err := ccgo.AEADLock(key, nonce, ad, pt)
	if err != nil {
		t.Fatalf("%s ccgo lock: %v", name, err)
	}
	if !bytes.Equal(macS, macC) {
		t.Fatalf("%s mac mismatch\nsgoi %x\nccgo %x", name, macS, macC)
	}
	if !bytes.Equal(ctS, ctC) {
		t.Fatalf("%s ct mismatch len=%d", name, len(pt))
	}
	out, err := sgoi.AEADUnlock(key, nonce, macS, ad, ctS)
	if err != nil {
		t.Fatalf("%s unlock: %v", name, err)
	}
	if !bytes.Equal(out, pt) {
		t.Fatalf("%s unlock pt mismatch", name)
	}
}

func TestParityVsCCGO_Sizes(t *testing.T) {
	cases := []struct {
		name string
		ad   []byte
		n    int
		pat  string
	}{
		{"pt0_ad_empty", nil, 0, ""},
		{"pt0_ad_H", []byte("H"), 0, ""},
		{"pt1_ad_empty", nil, 1, "i%251"},
		{"pt63_ad_empty", nil, 63, "i%251"},
		{"pt64_ad_empty", nil, 64, "i%251"},
		{"pt65_ad_empty", nil, 65, "i%251"},
		{"pt129_ad_empty", nil, 129, "i%251"},
		{"pt193_ad_empty", nil, 193, "i%251"},
		{"pt4096_ad_empty", nil, 4096, "i%251"},
		{"pt1024_header", []byte("HEADER 1KB"), 1024, "(i*17+3)%251"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var pt []byte
			if c.n > 0 {
				pt = mkPT(c.n, c.pat)
			}
			assertParity(t, c.name, c.ad, pt)
		})
	}
}

func TestParityVsCCGO_Matrix0to300(t *testing.T) {
	for n := 0; n <= 300; n++ {
		pt := mkPT(n, "i%251")
		assertParity(t, t.Name(), nil, pt)
	}
}

// Split AD then PT so poly1305_update runs with c_idx≠0 on real (non-zero) message
// bytes — catches message[0] cursor regression on the align loop.
func TestParityVsCCGO_PolyRemainderThenData(t *testing.T) {
	cases := []struct {
		name string
		ad   []byte
		n    int
	}{
		{"ad5_pt64", []byte{1, 2, 3, 4, 5}, 64},
		{"ad5_pt65", []byte{1, 2, 3, 4, 5}, 65},
		{"ad15_pt100", bytes.Repeat([]byte{0xA5}, 15), 100},
		{"ad17_pt129", bytes.Repeat([]byte{0x3C}, 17), 129},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertParity(t, c.name, c.ad, mkPT(c.n, "i%251"))
		})
	}
}

func TestParityVsCCGO_MACFailClosed(t *testing.T) {
	key, nonce := keyNonce()
	ad := []byte("AD")
	pt := []byte("plaintext-for-mac-fail")
	ct, mac, err := sgoi.AEADLock(key, nonce, ad, pt)
	if err != nil {
		t.Fatal(err)
	}
	bad := append([]byte(nil), mac...)
	bad[0] ^= 0xff
	if _, err := sgoi.AEADUnlock(key, nonce, bad, ad, ct); err == nil {
		t.Fatal("expected mac failure sgoi")
	}
	if _, err := ccgo.AEADUnlock(key, nonce, bad, ad, ct); err == nil {
		t.Fatal("expected mac failure ccgo")
	}
}
