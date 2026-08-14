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
