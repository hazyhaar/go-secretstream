// SPDX-License-Identifier: Apache-2.0 OR MIT

package engine

import (
	"bytes"
	"testing"
)

func testKeyNonce() (key, nonce24, subkey, nonce12 []byte) {
	key = make([]byte, 32)
	nonce24 = make([]byte, 24)
	for i := range key {
		key[i] = byte(0xA5 ^ i*17)
	}
	for i := range nonce24 {
		nonce24[i] = byte(0x3C ^ i*11)
	}
	subkey = make([]byte, 32)
	Default().HChaCha20(subkey, key, nonce24[:16])
	nonce12 = make([]byte, 12)
	copy(nonce12[4:], nonce24[16:24])
	return key, nonce24, subkey, nonce12
}

func TestAliasGuard_LockDst_InexactOverlap(t *testing.T) {
	eng := Default()
	key, nonce, _, _ := testKeyNonce()
	const n = 64
	ad := []byte("ad-lock")
	var mac [16]byte

	refPT := make([]byte, n)
	for i := range refPT {
		refPT[i] = byte(i*3 + 1)
	}
	refCT := make([]byte, n)
	if err := eng.LockDst(refCT, &mac, key, nonce, ad, refPT); err != nil {
		t.Fatalf("out-of-place lock: %v", err)
	}

	buf := make([]byte, n+1)
	copy(buf[0:n], refPT)
	dst := buf[1 : 1+n]
	plain := buf[0:n]
	if !inexactOverlap(dst, plain) {
		t.Fatal("précondition: dst/plain doivent se recouvrir inexactement")
	}

	var mac2 [16]byte
	err := eng.LockDst(dst, &mac2, key, nonce, ad, plain)
	if err == nil {
		if bytes.Equal(dst, refCT) {
			t.Fatal("attendu: erreur de recouvrement ; obtenu: nil et chiffré identique au hors-place")
		}
		t.Fatalf("attendu: erreur de recouvrement ; obtenu: nil et chiffré faux (diffère du hors-place)")
	}
}

func TestAliasGuard_UnlockDst_InexactOverlap(t *testing.T) {
	eng := Default()
	key, nonce, _, _ := testKeyNonce()
	const n = 64
	ad := []byte("ad-unlock")
	pt := make([]byte, n)
	for i := range pt {
		pt[i] = byte(i*5 + 2)
	}
	ct := make([]byte, n)
	var mac [16]byte
	if err := eng.LockDst(ct, &mac, key, nonce, ad, pt); err != nil {
		t.Fatalf("lock: %v", err)
	}

	buf := make([]byte, n+1)
	copy(buf[0:n], ct)
	dst := buf[1 : 1+n]
	cipher := buf[0:n]
	if !inexactOverlap(dst, cipher) {
		t.Fatal("précondition: dst/cipher doivent se recouvrir inexactement")
	}

	_, err := eng.UnlockDst(dst, key, nonce, ad, cipher, mac[:])
	if err == nil {
		t.Fatal("attendu: erreur de recouvrement ; obtenu: nil")
	}
}

func TestAliasGuard_LockSubkeyDst_InexactOverlap(t *testing.T) {
	eng := Default()
	_, _, subkey, nonce12 := testKeyNonce()
	const n = 64
	ad := []byte("ad-sub-lock")
	var mac [16]byte

	refPT := make([]byte, n)
	for i := range refPT {
		refPT[i] = byte(i*7 + 3)
	}
	refCT := make([]byte, n)
	if err := eng.LockSubkeyDst(refCT, &mac, subkey, nonce12, ad, refPT); err != nil {
		t.Fatalf("out-of-place subkey lock: %v", err)
	}

	buf := make([]byte, n+1)
	copy(buf[0:n], refPT)
	dst := buf[1 : 1+n]
	plain := buf[0:n]
	if !inexactOverlap(dst, plain) {
		t.Fatal("précondition: dst/plain doivent se recouvrir inexactement")
	}

	var mac2 [16]byte
	err := eng.LockSubkeyDst(dst, &mac2, subkey, nonce12, ad, plain)
	if err == nil {
		if bytes.Equal(dst, refCT) {
			t.Fatal("attendu: erreur de recouvrement ; obtenu: nil et chiffré identique au hors-place")
		}
		t.Fatalf("attendu: erreur de recouvrement ; obtenu: nil et chiffré faux (diffère du hors-place)")
	}
}

func TestAliasGuard_UnlockSubkeyDst_InexactOverlap(t *testing.T) {
	eng := Default()
	_, _, subkey, nonce12 := testKeyNonce()
	const n = 64
	ad := []byte("ad-sub-unlock")
	pt := make([]byte, n)
	for i := range pt {
		pt[i] = byte(i*11 + 4)
	}
	ct := make([]byte, n)
	var mac [16]byte
	if err := eng.LockSubkeyDst(ct, &mac, subkey, nonce12, ad, pt); err != nil {
		t.Fatalf("subkey lock: %v", err)
	}

	buf := make([]byte, n+1)
	copy(buf[0:n], ct)
	dst := buf[1 : 1+n]
	cipher := buf[0:n]
	if !inexactOverlap(dst, cipher) {
		t.Fatal("précondition: dst/cipher doivent se recouvrir inexactement")
	}

	_, err := eng.UnlockSubkeyDst(dst, subkey, nonce12, ad, cipher, mac[:])
	if err == nil {
		t.Fatal("attendu: erreur de recouvrement ; obtenu: nil")
	}
}

func TestAliasGuard_LockDst_InPlaceExactBitExact(t *testing.T) {
	eng := Default()
	key, nonce, _, _ := testKeyNonce()
	const n = 64
	ad := []byte("ad-inplace")
	pt := make([]byte, n)
	for i := range pt {
		pt[i] = byte(i*13 + 5)
	}
	outOfPlace := make([]byte, n)
	var macOut [16]byte
	if err := eng.LockDst(outOfPlace, &macOut, key, nonce, ad, pt); err != nil {
		t.Fatalf("out-of-place: %v", err)
	}

	inPlace := make([]byte, n)
	copy(inPlace, pt)
	var macIn [16]byte
	if err := eng.LockDst(inPlace, &macIn, key, nonce, ad, inPlace); err != nil {
		t.Fatalf("in-place exact refusé: %v", err)
	}
	if !bytes.Equal(inPlace, outOfPlace) {
		t.Fatal("in-place exact n'est pas bit-exact avec le hors-place")
	}
	if macIn != macOut {
		t.Fatal("MAC in-place distinct du MAC hors-place")
	}
}

func TestAliasGuard_LockSubkeyDst_InPlaceExactBitExact(t *testing.T) {
	eng := Default()
	_, _, subkey, nonce12 := testKeyNonce()
	const n = 64
	ad := []byte("ad-sub-inplace")
	pt := make([]byte, n)
	for i := range pt {
		pt[i] = byte(i*19 + 6)
	}
	outOfPlace := make([]byte, n)
	var macOut [16]byte
	if err := eng.LockSubkeyDst(outOfPlace, &macOut, subkey, nonce12, ad, pt); err != nil {
		t.Fatalf("out-of-place: %v", err)
	}

	inPlace := make([]byte, n)
	copy(inPlace, pt)
	var macIn [16]byte
	if err := eng.LockSubkeyDst(inPlace, &macIn, subkey, nonce12, ad, inPlace); err != nil {
		t.Fatalf("in-place exact refusé: %v", err)
	}
	if !bytes.Equal(inPlace, outOfPlace) {
		t.Fatal("in-place exact n'est pas bit-exact avec le hors-place")
	}
	if macIn != macOut {
		t.Fatal("MAC in-place distinct du MAC hors-place")
	}
}

func TestAliasGuard_DstOverlapsAD(t *testing.T) {
	eng := Default()
	key, nonce, subkey, nonce12 := testKeyNonce()
	buf := make([]byte, 80)
	dst := buf[0:64]
	plain := make([]byte, 64)
	ad := buf[48:72]
	var mac [16]byte
	if err := eng.LockDst(dst, &mac, key, nonce, ad, plain); err == nil {
		t.Fatal("LockDst: recouvrement dst/ad doit rendre une erreur")
	}
	if err := eng.LockSubkeyDst(dst, &mac, subkey, nonce12, ad, plain); err == nil {
		t.Fatal("LockSubkeyDst: recouvrement dst/ad doit rendre une erreur")
	}
}

func TestAliasGuard_DstOverlapsMAC(t *testing.T) {
	eng := Default()
	key, nonce, subkey, nonce12 := testKeyNonce()
	buf := make([]byte, 80)
	dst := buf[0:64]
	plain := make([]byte, 64)
	ad := []byte("ad-mac")
	mac := (*[16]byte)(buf[56:72])
	if err := eng.LockDst(dst, mac, key, nonce, ad, plain); err == nil {
		t.Fatal("LockDst: recouvrement dst/mac doit rendre une erreur")
	}
	if err := eng.LockSubkeyDst(dst, mac, subkey, nonce12, ad, plain); err == nil {
		t.Fatal("LockSubkeyDst: recouvrement dst/mac doit rendre une erreur")
	}

	ct := make([]byte, 64)
	var macOK [16]byte
	if err := eng.LockDst(ct, &macOK, key, nonce, ad, plain); err != nil {
		t.Fatalf("lock témoin: %v", err)
	}
	copy(buf[0:64], ct)
	macUnlock := buf[56:72]
	if _, err := eng.UnlockDst(dst, key, nonce, ad, ct, macUnlock); err == nil {
		t.Fatal("UnlockDst: recouvrement dst/mac doit rendre une erreur")
	}
	if _, err := eng.UnlockSubkeyDst(dst, subkey, nonce12, ad, ct, macUnlock); err == nil {
		t.Fatal("UnlockSubkeyDst: recouvrement dst/mac doit rendre une erreur")
	}
}

func TestAliasGuard_LockSubkeyDst_Allocs(t *testing.T) {
	eng := Default()
	_, _, subkey, nonce12 := testKeyNonce()
	pt := make([]byte, 64)
	dst := make([]byte, 64)
	ad := []byte("ad-alloc")
	var mac [16]byte
	allocs := testing.AllocsPerRun(100, func() {
		if err := eng.LockSubkeyDst(dst, &mac, subkey, nonce12, ad, pt); err != nil {
			t.Fatal(err)
		}
	})
	if allocs != 0 {
		t.Fatalf("LockSubkeyDst 64 octets: %v allocs/op, attendu 0", allocs)
	}
}
