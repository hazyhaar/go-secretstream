// SPDX-License-Identifier: Apache-2.0

package lsstream

import (
	"bytes"
	"io"
	"testing"
)

func TestRoundtrip(t *testing.T) {
	key := bytes.Repeat([]byte{0xab}, 32)
	for _, sz := range []int{0, 1, 100, 8192, 8193, 20000} {
		plain := make([]byte, sz)
		for i := range plain {
			plain[i] = byte(i)
		}
		var buf bytes.Buffer
		w := NewWriter(&buf, key)
		if _, err := w.Write(plain); err != nil {
			t.Fatal(err)
		}
		if err := w.Close(); err != nil {
			t.Fatal(err)
		}
		got, err := io.ReadAll(NewReader(bytes.NewReader(buf.Bytes()), key))
		if err != nil {
			t.Fatalf("n=%d %v", sz, err)
		}
		if !bytes.Equal(got, plain) {
			t.Fatalf("n=%d mismatch", sz)
		}
	}
}

func TestAdvanceWrapDoesNotRekey(t *testing.T) {
	key := bytes.Repeat([]byte{0x11}, 32)
	header := bytes.Repeat([]byte{0x22}, 24)
	st, err := initFromHeader(key, header)
	if err != nil {
		t.Fatal(err)
	}
	st.nonce[0], st.nonce[1], st.nonce[2], st.nonce[3] = 0xff, 0xff, 0xff, 0xff
	k0 := st.k
	wrapped, err := st.advance(bytes.Repeat([]byte{0x33}, 16))
	if err != nil {
		t.Fatal(err)
	}
	if !wrapped {
		t.Fatal("compteur 0xffffffff doit wrapping")
	}
	if st.k != k0 {
		t.Fatal("advance ne doit pas ré-encler")
	}
	if st.nonce[0] != 0 || st.nonce[1] != 0 || st.nonce[2] != 0 || st.nonce[3] != 0 {
		t.Fatalf("compteur après wrap: %x", st.nonce[:4])
	}
}
