// SPDX-License-Identifier: Apache-2.0

package lsstream

import (
	"bytes"
	"errors"
	"testing"
)

type failAfter struct {
	n       int
	calls   int
	written int
	fail    error
}

func (f *failAfter) Write(p []byte) (int, error) {
	f.calls++
	if f.calls >= f.n {
		return 0, f.fail
	}
	f.written += len(p)
	return len(p), nil
}

func TestWriterStickyErrAfterIOFailure(t *testing.T) {
	key := bytes.Repeat([]byte{0x42}, 32)
	inj := errors.New("injected write failure")
	fw := &failAfter{n: 2, fail: inj}
	w := NewWriter(fw, key)

	payload := make([]byte, chunkSize)
	for i := range payload {
		payload[i] = byte(i)
	}
	_, err := w.Write(payload)
	if err == nil {
		t.Fatal("la première écriture de fragment devait échouer")
	}

	writtenAfterFail := fw.written
	callsAfterFail := fw.calls

	n, err2 := w.Write([]byte("encore"))
	if n != 0 {
		t.Fatalf("Write après erreur a déclaré %d octets", n)
	}
	if err2 == nil {
		t.Fatal("Write après erreur devait être refusé")
	}
	if fw.written != writtenAfterFail {
		t.Fatalf("Write après erreur a émis %d octet(s) de plus", fw.written-writtenAfterFail)
	}
	if fw.calls != callsAfterFail {
		t.Fatalf("Write après erreur a rappelé le Writer sous-jacent (%d appels, était %d)", fw.calls, callsAfterFail)
	}

	err3 := w.Close()
	if err3 == nil {
		t.Fatal("Close après erreur devait rendre une erreur")
	}
	if fw.written != writtenAfterFail {
		t.Fatalf("Close après erreur a émis un TAG_FINAL (%d octet(s) de plus)", fw.written-writtenAfterFail)
	}
	if fw.calls != callsAfterFail {
		t.Fatalf("Close après erreur a rappelé le Writer sous-jacent (%d appels, était %d)", fw.calls, callsAfterFail)
	}
}
