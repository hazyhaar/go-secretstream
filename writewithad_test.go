// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/engine"
)

func fixedKeyAD(t *testing.T) []byte {
	t.Helper()
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(0xA5 ^ i*17)
	}
	return key
}

func encryptWithFixedNonce(t *testing.T, key []byte, nonce [24]byte, p, ad []byte) []byte {
	t.Helper()
	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	enc.nonce = nonce
	enc.eng.HChaCha20(enc.subkey[:], key[:], enc.nonce[0:16])
	enc.seq = 0
	buf.Reset()
	var hdr [HeaderSize]byte
	writeHeaderV2(hdr[:], &nonce)
	if _, err := buf.Write(hdr[:]); err != nil {
		t.Fatalf("écriture de l'en-tête: %v", err)
	}
	if _, err := enc.WriteWithAD(p, ad); err != nil {
		t.Fatalf("WriteWithAD: %v", err)
	}
	if err := enc.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	return buf.Bytes()
}

func splitMaisonFrames(t *testing.T, wire []byte) (nonce []byte, frames [][]byte) {
	t.Helper()
	hs := HeaderSizeV1
	if isMagicV2(wire) {
		hs = HeaderSize
	}
	if len(wire) < hs {
		t.Fatalf("fil trop court: %d", len(wire))
	}
	nonce = append([]byte(nil), wire[:hs]...)
	rest := wire[hs:]
	for len(rest) > 0 {
		if len(rest) < 4 {
			t.Fatalf("préfixe de longueur tronqué, reste %d", len(rest))
		}
		n := int(binary.BigEndian.Uint32(rest[:4]))
		total := 4 + n
		if n < 0 || len(rest) < total {
			t.Fatalf("frame tronquée: besoin %d, reste %d", total, len(rest))
		}
		frames = append(frames, append([]byte(nil), rest[:total]...))
		rest = rest[total:]
	}
	return nonce, frames
}

// TestWriteWithAD_EmptyMatchesHistoricalWire — un flux produit avec une
// donnée associée vide doit être octet pour octet identique à la
// construction indépendante v2, à nonce fixé, et rester lisible par Read.
func TestWriteWithAD_EmptyMatchesHistoricalWire(t *testing.T) {
	eng := engine.Default()
	key := fixedKeyAD(t)
	payload := make([]byte, ChunkSize+97)
	for i := range payload {
		payload[i] = byte(i*41 + 3)
	}

	var viaWrite bytes.Buffer
	enc, err := NewEncryptor(&viaWrite, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	if _, err := enc.Write(payload); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if err := enc.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	var nonce [24]byte
	copy(nonce[:], viaWrite.Bytes()[12:12+24])

	historical := oldPathStream(t, eng, key, nonce, payload)
	if !bytes.Equal(historical, viaWrite.Bytes()) {
		t.Fatalf("Write diverge du chemin historique (Write %d, historique %d)", viaWrite.Len(), len(historical))
	}

	viaAD := encryptWithFixedNonce(t, key, nonce, payload, nil)
	if !bytes.Equal(viaAD, viaWrite.Bytes()) {
		t.Fatalf("WriteWithAD(nil) diverge de Write (AD %d, Write %d)", len(viaAD), viaWrite.Len())
	}
	if !bytes.Equal(viaAD, historical) {
		t.Fatalf("WriteWithAD(nil) diverge de l'historique (AD %d, historique %d)", len(viaAD), len(historical))
	}

	dec, err := NewDecryptor(bytes.NewReader(viaAD), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	got, err := io.ReadAll(dec)
	if err != nil {
		t.Fatalf("Read du format actuel: %v", err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatalf("déchiffrement Read: obtenu %d octets, attendu %d", len(got), len(payload))
	}
}

// TestWriteWithAD_PermuteFails — deux fragments permutés doivent faire
// échouer le déchiffrement : le préfixe de séquence reste lié à chaque
// fragment même lorsqu'une donnée associée d'appelant est concaténée.
func TestWriteWithAD_PermuteFails(t *testing.T) {
	key := fixedKeyAD(t)
	chunk1 := bytes.Repeat([]byte{0x11}, 200)
	chunk2 := bytes.Repeat([]byte{0x22}, 200)
	ad := []byte("contexte-fragment-c5")

	var wire bytes.Buffer
	enc, err := NewEncryptor(&wire, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	if _, err := enc.WriteWithAD(chunk1, ad); err != nil {
		t.Fatalf("WriteWithAD chunk1: %v", err)
	}
	if _, err := enc.WriteWithAD(chunk2, ad); err != nil {
		t.Fatalf("WriteWithAD chunk2: %v", err)
	}

	nonce, frames := splitMaisonFrames(t, wire.Bytes())
	if len(frames) != 2 {
		t.Fatalf("attendu 2 fragments, obtenu %d", len(frames))
	}

	var intact bytes.Buffer
	intact.Write(nonce)
	intact.Write(frames[0])
	intact.Write(frames[1])
	decOK, err := NewDecryptor(&intact, key)
	if err != nil {
		t.Fatalf("NewDecryptor intact: %v", err)
	}
	got1 := make([]byte, len(chunk1))
	if n, err := decOK.ReadWithAD(got1, ad); err != nil || n != len(chunk1) {
		t.Fatalf("lecture intacte fragment 1: n=%d err=%v", n, err)
	}
	got2 := make([]byte, len(chunk2))
	if n, err := decOK.ReadWithAD(got2, ad); err != nil || n != len(chunk2) {
		t.Fatalf("lecture intacte fragment 2: n=%d err=%v", n, err)
	}
	if !bytes.Equal(got1, chunk1) || !bytes.Equal(got2, chunk2) {
		t.Fatal("va-et-vient intact: clair divergé")
	}

	var swapped bytes.Buffer
	swapped.Write(nonce)
	swapped.Write(frames[1])
	swapped.Write(frames[0])
	decBad, err := NewDecryptor(&swapped, key)
	if err != nil {
		t.Fatalf("NewDecryptor permuté: %v", err)
	}
	dst := make([]byte, len(chunk2)+len(chunk1))
	if _, err := decBad.ReadWithAD(dst, ad); err == nil {
		t.Fatal("permutation acceptée — la protection anti-rejeu a disparu")
	}
}

func TestWriteWithAD_RoundtripAndWrongADFails(t *testing.T) {
	key := fixedKeyAD(t)
	payload := []byte("charge utile liée à une donnée associée d'appelant")
	ad := []byte("ad-appelant-c5")

	var wire bytes.Buffer
	enc, err := NewEncryptor(&wire, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	if _, err := enc.WriteWithAD(payload, ad); err != nil {
		t.Fatalf("WriteWithAD: %v", err)
	}

	dec, err := NewDecryptor(bytes.NewReader(wire.Bytes()), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	got := make([]byte, len(payload))
	n, err := dec.ReadWithAD(got, ad)
	if err != nil {
		t.Fatalf("ReadWithAD: %v", err)
	}
	if n != len(payload) || !bytes.Equal(got, payload) {
		t.Fatalf("va-et-vient AD: n=%d got=%q", n, got[:n])
	}

	decWrong, err := NewDecryptor(bytes.NewReader(wire.Bytes()), key)
	if err != nil {
		t.Fatalf("NewDecryptor mauvais AD: %v", err)
	}
	if _, err := decWrong.ReadWithAD(got, []byte("ad-incorrecte")); err == nil {
		t.Fatal("un AD d'appelant incorrect a été accepté")
	}

	decEmpty, err := NewDecryptor(bytes.NewReader(wire.Bytes()), key)
	if err != nil {
		t.Fatalf("NewDecryptor AD vide: %v", err)
	}
	if _, err := decEmpty.Read(got); err == nil {
		t.Fatal("Read sans AD d'appelant a accepté un fragment lié")
	}
}
