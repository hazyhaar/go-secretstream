//go:build sgoiter_oracle

// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

// Tests optionnels : comparer monocypher_sgoiter à monocypher ccgo
// sur des vecteurs AEAD (oracle pure Go sans bascule runtime stream).
//
//	cd pkg/secretstream55 && go test -tags sgoiter_oracle ./...

import (
	"bytes"
	"testing"

	ccgo "github.com/hazyhaar/go-secretstream/internal/monocypher"
	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher_sgoiter"
)

func TestOracleSgoiterMatchesCCGO_StreamSized(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	for i := range key {
		key[i] = byte(i + 3)
	}
	for i := range nonce {
		nonce[i] = byte(i + 7)
	}
	for _, n := range []int{0, 1, 15, 16, 63, 64, 100, 1024} {
		pt := make([]byte, n)
		for i := range pt {
			pt[i] = byte(i * 9)
		}
		ad := []byte("oracle")
		ctS, macS, err := sgoi.AEADLock(key, nonce, ad, pt)
		if err != nil {
			t.Fatal(err)
		}
		ctC, macC, err := ccgo.AEADLock(key, nonce, ad, pt)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(ctS, ctC) || !bytes.Equal(macS, macC) {
			t.Fatalf("n=%d mismatch", n)
		}
		out, err := sgoi.AEADUnlock(key, nonce, macS, ad, ctS)
		if err != nil || !bytes.Equal(out, pt) {
			t.Fatalf("unlock n=%d", n)
		}
	}
}
