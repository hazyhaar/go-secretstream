// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

var inplaceFragmentSizes = []int{1, 63, 64, 65, 4095, ChunkSize}

func inplacePlainOf(n int, seed byte) []byte {
	p := make([]byte, n)
	for i := range p {
		p[i] = seed + byte(i*41)
	}
	return p
}

func encryptFragments(t *testing.T, key []byte, frags [][]byte) []byte {
	t.Helper()
	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	for i, f := range frags {
		if _, err := enc.Write(f); err != nil {
			t.Fatalf("Write fragment %d (n=%d): %v", i, len(f), err)
		}
	}
	if err := enc.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	return buf.Bytes()
}

func inplaceKey() []byte {
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(0x5A ^ i*13)
	}
	return key
}

// TestReadInPlace_NoIntermediateCopy — lecture avec p assez grand pour chaque
// fragment. Après correctif, plainBuf (alloué à zéro, jamais réécrit) reste
// entièrement nul. Avant correctif le clair transite par plainBuf : rouge.
func TestReadInPlace_NoIntermediateCopy(t *testing.T) {
	key := inplaceKey()
	frags := make([][]byte, len(inplaceFragmentSizes))
	for i, n := range inplaceFragmentSizes {
		frags[i] = inplacePlainOf(n, byte(i+1))
	}
	wire := encryptFragments(t, key, frags)

	dec, err := NewDecryptor(bytes.NewReader(wire), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	if !allZero(dec.plainBuf) {
		t.Fatal("précondition: plainBuf doit être nul à la construction")
	}

	for i, want := range frags {
		p := make([]byte, len(want)+8)
		n, err := dec.ReadWithAD(p, nil)
		if err != nil {
			t.Fatalf("fragment %d n=%d: ReadWithAD: %v", i, len(want), err)
		}
		if n != len(want) {
			t.Fatalf("fragment %d: lu %d, attendu %d", i, n, len(want))
		}
		if !bytes.Equal(p[:n], want) {
			t.Fatalf("fragment %d: clair différent de l'origine", i)
		}
		if !allZero(dec.plainBuf) {
			t.Fatalf("fragment %d n=%d: plainBuf a été écrit (copie intermédiaire encore présente)", i, len(want))
		}
		if len(dec.outBuf) != 0 {
			t.Fatalf("fragment %d: outBuf devrait être vide après lecture en place, reste %d", i, len(dec.outBuf))
		}
	}

	var many bytes.Buffer
	enc, err := NewEncryptor(&many, key)
	if err != nil {
		t.Fatalf("NewEncryptor hot: %v", err)
	}
	chunk := inplacePlainOf(64, 9)
	const hotChunks = 200
	for i := 0; i < hotChunks; i++ {
		if _, err := enc.Write(chunk); err != nil {
			t.Fatalf("Write hot %d: %v", i, err)
		}
	}
	hotDec, err := NewDecryptor(bytes.NewReader(many.Bytes()), key)
	if err != nil {
		t.Fatalf("NewDecryptor hot: %v", err)
	}
	hotP := make([]byte, 64)
	if _, err := hotDec.ReadWithAD(hotP, nil); err != nil {
		t.Fatalf("warmup ReadWithAD: %v", err)
	}
	hotAllocs := testing.AllocsPerRun(50, func() {
		n, err := hotDec.ReadWithAD(hotP, nil)
		if err != nil || n != 64 {
			t.Fatalf("hot ReadWithAD n=%d err=%v", n, err)
		}
	})
	t.Logf("ReadWithAD en place: %v allocs/op (chunkNonce12 s'échappe vers le tas via l'interface moteur, préexistant)", hotAllocs)
}

// TestReadInPlace_ParityWithShortP — même flux lu avec p grand et avec p de
// 7 octets (chemin reliquat). Les deux clairs doivent égaler l'origine.
// Vert d'emblée attendu : le fil et le déchiffrement ne changent pas.
func TestReadInPlace_ParityWithShortP(t *testing.T) {
	key := inplaceKey()
	var origin []byte
	frags := make([][]byte, len(inplaceFragmentSizes))
	for i, n := range inplaceFragmentSizes {
		frags[i] = inplacePlainOf(n, byte(i+3))
		origin = append(origin, frags[i]...)
	}
	wire := encryptFragments(t, key, frags)

	readAll := func(pSize int) []byte {
		dec, err := NewDecryptor(bytes.NewReader(wire), key)
		if err != nil {
			t.Fatalf("NewDecryptor pSize=%d: %v", pSize, err)
		}
		var got []byte
		p := make([]byte, pSize)
		for {
			n, err := dec.ReadWithAD(p, nil)
			got = append(got, p[:n]...)
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("ReadWithAD pSize=%d: %v", pSize, err)
			}
			if n == 0 {
				t.Fatalf("ReadWithAD pSize=%d: n=0 sans EOF", pSize)
			}
		}
		return got
	}

	wide := readAll(len(origin) + 32)
	narrow := readAll(7)
	if !bytes.Equal(wide, origin) {
		t.Fatalf("p grand: clair différent de l'origine (lu %d, origine %d)", len(wide), len(origin))
	}
	if !bytes.Equal(narrow, origin) {
		t.Fatalf("p=7: clair différent de l'origine (lu %d, origine %d)", len(narrow), len(origin))
	}
	if !bytes.Equal(wide, narrow) {
		t.Fatal("p grand et p=7 ont rendu des clairs différents")
	}

	var noClose bytes.Buffer
	enc, err := NewEncryptor(&noClose, key)
	if err != nil {
		t.Fatalf("NewEncryptor sans Close: %v", err)
	}
	for i, f := range frags {
		if _, err := enc.Write(f); err != nil {
			t.Fatalf("Write sans Close %d: %v", i, err)
		}
	}
	dec, err := NewDecryptor(bytes.NewReader(noClose.Bytes()), key)
	if err != nil {
		t.Fatalf("NewDecryptor sans Close: %v", err)
	}
	p := make([]byte, 7)
	for {
		n, err := dec.ReadWithAD(p, nil)
		if err == io.EOF {
			t.Fatal("sans Close: io.EOF, attendu flux tronqué")
		}
		if err != nil {
			if !strings.Contains(err.Error(), "flux tronqué") {
				t.Fatalf("sans Close: attendu flux tronqué, obtenu %v", err)
			}
			if n != 0 {
				t.Fatalf("sans Close: n=%d sur flux tronqué", n)
			}
			break
		}
		if n == 0 {
			t.Fatal("sans Close: n=0 sans erreur")
		}
	}
}

// TestReadInPlace_NoPlaintextLeakOnBadMAC — fragment à MAC corrompu, p grand
// rempli d'un motif connu. Après l'erreur, aucun clair candidat ne reste
// dans p (motif intact ou dst mis à zéro par le moteur). Verrou collant exigé.
func TestReadInPlace_NoPlaintextLeakOnBadMAC(t *testing.T) {
	key := inplaceKey()
	plain := inplacePlainOf(4095, 0x11)
	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	if _, err := enc.Write(plain); err != nil {
		t.Fatalf("Write: %v", err)
	}
	wire := buf.Bytes()
	if len(wire) < HeaderSize+4+16 {
		t.Fatalf("fil trop court: %d", len(wire))
	}
	wire[len(wire)-1] ^= 0xFF

	dec, err := NewDecryptor(bytes.NewReader(wire), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	const motif byte = 0xA5
	p := bytes.Repeat([]byte{motif}, len(plain)+16)
	initial := append([]byte(nil), p...)

	n, err1 := dec.ReadWithAD(p, nil)
	if err1 == nil {
		t.Fatal("MAC corrompu accepté")
	}
	if n != 0 {
		t.Fatalf("échec MAC a déclaré %d octets", n)
	}
	if bytes.Equal(p[:len(plain)], plain) {
		t.Fatal("clair candidat laissé dans p après échec MAC")
	}
	for i, b := range p[:len(plain)] {
		if b != motif && b != 0 {
			t.Fatalf("p[%d]=0x%02x : ni motif 0x%02x ni zéro (clair non authentifié)", i, b, motif)
		}
	}
	if !bytes.Equal(p[len(plain):], initial[len(plain):]) {
		t.Fatal("octets au-delà de cipherLen mutés après échec MAC")
	}
	if !allZero(p[:len(plain)]) {
		t.Log("dst laissé au motif (moteur n'écrit pas en échec)")
	}

	n2, err2 := dec.ReadWithAD(p, nil)
	if err2 == nil {
		t.Fatal("deuxième lecture: erreur collante attendue, obtenu nil")
	}
	if n2 != 0 {
		t.Fatalf("deuxième lecture a déclaré %d octets", n2)
	}
	if err2.Error() != err1.Error() {
		t.Fatalf("erreur collante différente: première %v, seconde %v", err1, err2)
	}
}
