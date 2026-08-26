// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"crypto/rand"
	"strings"
	"testing"
)

func TestEncryptorSeqOverflow(t *testing.T) {
	var buf bytes.Buffer
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	headerLen := buf.Len()
	if headerLen != HeaderSize {
		t.Fatalf("en-tête: obtenu %d, attendu %d", headerLen, HeaderSize)
	}

	enc.seq = ^uint64(0)
	n, err := enc.Write([]byte("x"))
	if n != 0 {
		t.Fatalf("Write sur débordement a déclaré %d octets", n)
	}
	if err == nil || !strings.Contains(err.Error(), "sequence number overflow") {
		t.Fatalf("erreur de débordement attendue, obtenu %v", err)
	}
	if buf.Len() != headerLen {
		t.Fatalf("octets écrits sur le fil pour ce fragment: %d (en-tête seul attendu: %d)", buf.Len(), headerLen)
	}

	n2, err2 := enc.Write([]byte("y"))
	if n2 != 0 {
		t.Fatalf("second Write a déclaré %d octets", n2)
	}
	if err2 == nil || err2.Error() != err.Error() {
		t.Fatalf("stickyErr attendu identique (%v), obtenu %v", err, err2)
	}
	if buf.Len() != headerLen {
		t.Fatalf("second Write a émis des octets: fil=%d", buf.Len())
	}
}
