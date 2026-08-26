// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55_test

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

func TestRound6_RejectSmallOrderR(t *testing.T) {
	seed := sha256.Sum256([]byte("small-order"))
	var sk [64]byte
	var pk [32]byte
	sgoi.Crypto_eddsa_key_pair(sk[:], pk[:], append([]byte(nil), seed[:]...))
	msg := []byte("order")
	// R = identity (order 1): 0x01 || 0x00*31, S anything under L
	var sig [64]byte
	sig[0] = 1
	// S = 1
	sig[32] = 1
	if sgoi.Crypto_eddsa_check(sig[:], pk[:], msg, uint64(len(msg))) == 0 {
		t.Fatal("accepted small-order R")
	}
}

func TestRound6_SignVerify_AliasBuffers(t *testing.T) {
	seed := sha256.Sum256([]byte("alias"))
	var sk [64]byte
	var pk [32]byte
	sgoi.Crypto_eddsa_key_pair(sk[:], pk[:], append([]byte(nil), seed[:]...))
	buf := make([]byte, 128)
	msg := []byte("alias-msg")
	copy(buf[0:], msg)
	var sig [64]byte
	sgoi.Crypto_eddsa_sign(sig[:], sk[:], buf[:len(msg)], uint64(len(msg)))
	copy(buf[32:], sig[:])
	if sgoi.Crypto_eddsa_check(buf[32:96], pk[:], buf[:len(msg)], uint64(len(msg))) != 0 {
		t.Fatal("verify with adjacent buffers failed")
	}
}

func TestRound6_Scalarbase_TrimmedLike(t *testing.T) {
	// scalars with bit 254 set like trim_scalar
	for n := 0; n < 8; n++ {
		h := sha256.Sum256([]byte(fmt.Sprintf("trim-%d", n)))
		sc := append([]byte(nil), h[:]...)
		sc[0] &= 248
		sc[31] &= 127
		sc[31] |= 64
		var pt [32]byte
		sgoi.Crypto_eddsa_scalarbase(pt[:], sc)
		// self-consistency: 2*(sc/2) not easy; just ensure not all-zero unless sc=0
		if bytes.Equal(pt[:], make([]byte, 32)) {
			t.Fatalf("zero point n=%d", n)
		}
	}
}
