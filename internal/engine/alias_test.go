// SPDX-License-Identifier: Apache-2.0 OR MIT

package engine

import "testing"

func TestOverlapDetection(t *testing.T) {
	buf := make([]byte, 128)

	s1 := buf[0:32]
	s2 := buf[32:64]
	if anyOverlap(s1, s2) {
		t.Fatal("Disjoint buffers must not overlap")
	}
	if inexactOverlap(s1, s2) {
		t.Fatal("Disjoint buffers must not inexact-overlap")
	}

	s3 := buf[0:32]
	if inexactOverlap(s1, s3) {
		t.Fatal("Exact same slices must not inexact-overlap")
	}
	if !anyOverlap(s1, s3) {
		t.Fatal("Identical slices must report anyOverlap")
	}

	s4 := buf[1:33]
	if !anyOverlap(s1, s4) {
		t.Fatal("Overlapping slices must report anyOverlap")
	}
	if !inexactOverlap(s1, s4) {
		t.Fatal("Shifted slices must report inexactOverlap")
	}

	sBefore := buf[16:48]
	sAfter := buf[32:64]
	if !anyOverlap(sBefore, sAfter) {
		t.Fatal("Partial overlap (before/after) must report anyOverlap")
	}
	if !inexactOverlap(sBefore, sAfter) {
		t.Fatal("Partial overlap (before/after) must report inexactOverlap")
	}

	var empty []byte
	if anyOverlap(s1, empty) || anyOverlap(empty, s1) || anyOverlap(empty, empty) {
		t.Fatal("len 0 must not anyOverlap")
	}
	if inexactOverlap(s1, empty) || inexactOverlap(empty, s1) || inexactOverlap(nil, nil) {
		t.Fatal("len 0 must not inexactOverlap")
	}

	sub := buf[0:16]
	if !anyOverlap(s1, sub) {
		t.Fatal("Strict sub-slice must report anyOverlap")
	}
	if inexactOverlap(s1, sub) {
		t.Fatal("Strict sub-slice sharing start must not inexactOverlap (x/crypto)")
	}

	inner := buf[8:24]
	if !anyOverlap(s1, inner) {
		t.Fatal("Interior strict sub-slice must report anyOverlap")
	}
	if !inexactOverlap(s1, inner) {
		t.Fatal("Interior strict sub-slice must report inexactOverlap")
	}
}
