// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

import (
	"bytes"
	"testing"

	ccgo "github.com/hazyhaar/go-secretstream/internal/monocypher"
	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

func TestParityVsCCGO_36(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	ad := []byte("HEADER")
	pt := []byte("HELLO MONOCYPHER SGOITER AEAD CGO=0!")
	ctS, macS, err := sgoi.AEADLock(key, nonce, ad, pt)
	if err != nil {
		t.Fatal(err)
	}
	ctC, macC, err := ccgo.AEADLock(key, nonce, ad, pt)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(ctS, ctC) {
		t.Fatalf("ct mismatch\nsgoi %x\nccgo %x", ctS, ctC)
	}
	if !bytes.Equal(macS, macC) {
		t.Fatalf("mac mismatch\nsgoi %x\nccgo %x", macS, macC)
	}
	out, err := sgoi.AEADUnlock(key, nonce, macS, ad, ctS)
	if err != nil || !bytes.Equal(out, pt) {
		t.Fatalf("unlock sgoi %v %q", err, out)
	}
}

func TestParityVsCCGO_1KB(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}
	ad := []byte("HEADER 1KB")
	pt := make([]byte, 1024)
	for i := range pt {
		pt[i] = byte((i*17 + 3) % 251)
	}
	ctS, macS, err := sgoi.AEADLock(key, nonce, ad, pt)
	if err != nil {
		t.Fatal(err)
	}
	ctC, macC, err := ccgo.AEADLock(key, nonce, ad, pt)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(macS, macC) || !bytes.Equal(ctS, ctC) {
		t.Fatalf("1k mismatch mac_eq=%v ct_eq=%v", bytes.Equal(macS, macC), bytes.Equal(ctS, ctC))
	}
	out, err := sgoi.AEADUnlock(key, nonce, macS, ad, ctS)
	if err != nil || !bytes.Equal(out, pt) {
		t.Fatal("unlock 1k")
	}
}
