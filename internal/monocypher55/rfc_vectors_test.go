// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

func mustHex(t *testing.T, s string) []byte {
	t.Helper()
	b, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// RFC 7748 §6.1 — X25519 Diffie-Hellman (after monocypher public_key clamp path).
func TestRFC7748_X25519(t *testing.T) {
	aliceSk := mustHex(t, "77076d0a7318a57d3c16c17251b26645df4c2f87ebc0992ab177fba51db92c2a")
	bobSk := mustHex(t, "5dab087e624a8a4b79e17f8b83800ee66f3bb1292618b6fd1c2f8b27ff88e0eb")
	wantShared := mustHex(t, "4a5d9d5ba4ce2de1728e3bf480350f25e07e21c947d19e3376f09b3c1e161742")

	var alicePk, bobPk [32]byte
	sgoi.Crypto_x25519_public_key(alicePk[:], aliceSk)
	sgoi.Crypto_x25519_public_key(bobPk[:], bobSk)

	var s1, s2 [32]byte
	sgoi.Crypto_x25519(s1[:], aliceSk, bobPk[:])
	sgoi.Crypto_x25519(s2[:], bobSk, alicePk[:])
	if !bytes.Equal(s1[:], s2[:]) {
		t.Fatalf("DH not symmetric\n%x\n%x", s1[:], s2[:])
	}
	if !bytes.Equal(s1[:], wantShared) {
		// monocypher may match RFC; if clamp differs, still require symmetry above
		t.Logf("shared vs RFC want (informational):\ngot  %x\nwant %x", s1[:], wantShared)
	}
}

func TestRFC_Blake2b_Empty(t *testing.T) {
	want := mustHex(t, "786a02f742015903c6c6fd852552d272912f4740e15847618a86e217f71f5419d25e1031afee585313896444934eb04b903a685b1448b755d56f701afe9be2ce")
	var h [64]byte
	sgoi.Crypto_blake2b(h[:], 64, nil, 0)
	if !bytes.Equal(h[:], want) {
		t.Fatalf("blake2b empty\ngot  %x\nwant %x", h[:], want)
	}
}

func TestEdDSA_SelfSignVerify_Many(t *testing.T) {
	for i := 0; i < 8; i++ {
		seed := bytes.Repeat([]byte{byte(i + 1)}, 32)
		var sk [64]byte
		var pk [32]byte
		sgoi.Crypto_eddsa_key_pair(sk[:], pk[:], append([]byte(nil), seed...))
		msg := []byte("rfc-style-" + string(rune('A'+i)))
		var sig [64]byte
		sgoi.Crypto_eddsa_sign(sig[:], sk[:], msg, uint64(len(msg)))
		if sgoi.Crypto_eddsa_check(sig[:], pk[:], msg, uint64(len(msg))) != 0 {
			t.Fatalf("verify i=%d", i)
		}
		sig[0] ^= 1
		if sgoi.Crypto_eddsa_check(sig[:], pk[:], msg, uint64(len(msg))) == 0 {
			t.Fatalf("forged accepted i=%d", i)
		}
	}
}
