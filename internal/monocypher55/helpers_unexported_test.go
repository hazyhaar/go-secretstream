// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"bytes"
	"testing"
)

func TestFeIsOdd_WipeCheck(t *testing.T) {
	f := [10]int32{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	got := fe_isodd(&f)
	if got != 1 {
		t.Fatalf("Fe_isodd(1) = %d, expected 1 (C1 wipe-before-return bug)", got)
	}
}

func TestRound6_ModL_Boundaries(t *testing.T) {
	x := make([]uint32, 16)
	var out [32]byte
	Mod_l(out[:], x)
	if !bytes.Equal(out[:], make([]byte, 32)) {
		t.Fatalf("0 got %x", out[:])
	}
	L := []uint32{0x5cf5d3ed, 0x5812631a, 0xa2f79cd6, 0x14def9de, 0, 0, 0, 0x10000000}
	copy(x, L)
	Mod_l(out[:], x)
	if !bytes.Equal(out[:], make([]byte, 32)) {
		t.Fatalf("L mod L → %x want 0", out[:])
	}
	x[0] = L[0] + 1
	Mod_l(out[:], x)
	want := make([]byte, 32)
	want[0] = 1
	if !bytes.Equal(out[:], want) {
		t.Fatalf("L+1 → %x want 1", out[:])
	}
	for i := range x {
		x[i] = 0xffffffff
	}
	Mod_l(out[:], x)
	var r8 [8]uint32
	load32_le_buf(r8[:], out[:], 8)
	if is_above_l(r8[:]) != 0 {
		t.Fatalf("0xff.. not fully reduced: %x", out[:])
	}
}
