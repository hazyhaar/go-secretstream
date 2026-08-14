package monocypher_test

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
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

func TestRound6_ModL_Boundaries(t *testing.T) {
	// x all-zero → 0
	x := make([]uint32, 16)
	var out [32]byte
	sgoi.Mod_l(out[:], x)
	if !bytes.Equal(out[:], make([]byte, 32)) {
		t.Fatalf("0 got %x", out[:])
	}
	// load L as 8 limbs + zeros → reduce to 0
	L := []uint32{0x5cf5d3ed, 0x5812631a, 0xa2f79cd6, 0x14def9de, 0, 0, 0, 0x10000000}
	copy(x, L)
	sgoi.Mod_l(out[:], x)
	if !bytes.Equal(out[:], make([]byte, 32)) {
		t.Fatalf("L mod L → %x want 0", out[:])
	}
	// L+1 → 1
	x[0] = L[0] + 1
	// handle carry if needed - L[0]+1 doesn't overflow
	sgoi.Mod_l(out[:], x)
	want := make([]byte, 32)
	want[0] = 1
	if !bytes.Equal(out[:], want) {
		t.Fatalf("L+1 → %x want 1", out[:])
	}
	// mul_add identity: 1 * half_mod_L + half_ones already tested via scalarbase
	// max-ish: all 0xffffffff in low 16 limbs
	for i := range x {
		x[i] = 0xffffffff
	}
	sgoi.Mod_l(out[:], x)
	// must be < L: compare as LE uint256 roughly via is_above_l on result
	var r8 [8]uint32
	sgoi.Load32_le_buf(r8[:], out[:], 8)
	if sgoi.Is_above_l(r8[:]) != 0 {
		t.Fatalf("0xff.. not fully reduced: %x", out[:])
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
