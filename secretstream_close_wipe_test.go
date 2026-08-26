// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"crypto/rand"
	"strings"
	"testing"
)

func allZero(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}

func TestCloseWipesOwnedBuffers_Encryptor(t *testing.T) {
	var buf bytes.Buffer
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	if _, err := enc.Write([]byte("payload-non-vide")); err != nil {
		t.Fatal(err)
	}
	if _, err := enc.WriteWithAD([]byte("avec-ad"), []byte("donnee-associee")); err != nil {
		t.Fatal(err)
	}
	if allZero(enc.subkey[:]) {
		t.Fatal("précondition: subkey devrait être non nulle avant Close")
	}
	if allZero(enc.wireBuf) {
		t.Fatal("précondition: wireBuf devrait être sale avant Close")
	}
	if allZero(enc.adBuf[:]) {
		t.Fatal("précondition: adBuf devrait être sale avant Close")
	}
	if len(enc.adExt) == 0 || allZero(enc.adExt) {
		t.Fatal("précondition: adExt devrait être sale avant Close")
	}

	if err := enc.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if !allZero(enc.subkey[:]) {
		t.Fatal("subkey non effacée après Close")
	}
	if !allZero(enc.scratchPayload) {
		t.Fatal("scratchPayload non effacé après Close")
	}
	if !allZero(enc.wireBuf) {
		t.Fatal("wireBuf non effacé après Close")
	}
	if !allZero(enc.adBuf[:]) {
		t.Fatal("adBuf non effacé après Close")
	}
	if !allZero(enc.adExt) {
		t.Fatal("adExt non effacé après Close")
	}
	if err := enc.Close(); err != nil {
		t.Fatalf("Close idempotent a rendu %v", err)
	}
}

func TestCloseWipesOwnedBuffers_Decryptor(t *testing.T) {
	var buf bytes.Buffer
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("NewEncryptor: %v", err)
	}
	first := []byte("premier")
	plain := []byte("clair-a-dechiffrer")
	if _, err := enc.WriteWithAD(first, []byte("ad-lecteur")); err != nil {
		t.Fatal(err)
	}
	if _, err := enc.WriteWithAD(plain, []byte("ad-lecteur")); err != nil {
		t.Fatal(err)
	}

	dec, err := NewDecryptor(bytes.NewReader(buf.Bytes()), key)
	if err != nil {
		t.Fatalf("NewDecryptor: %v", err)
	}
	gotFirst := make([]byte, len(first))
	if _, err := dec.ReadWithAD(gotFirst, []byte("ad-lecteur")); err != nil {
		t.Fatal(err)
	}
	// Le second fragment est lu avec un tampon plus court que le clair : c'est
	// le seul chemin qui passe encore par plainBuf depuis l'item 1.3 (lecture
	// en place quand le tampon de l'appelant suffit). La précondition « plainBuf
	// sale » exige donc le chemin reliquat.
	out := make([]byte, 4)
	if _, err := dec.ReadWithAD(out, []byte("ad-lecteur")); err != nil {
		t.Fatal(err)
	}
	if allZero(dec.subkey[:]) {
		t.Fatal("précondition: subkey devrait être non nulle avant Close")
	}
	if allZero(dec.inBuf) {
		t.Fatal("précondition: inBuf devrait être sale avant Close")
	}
	if allZero(dec.plainBuf) {
		t.Fatal("précondition: plainBuf devrait être sale avant Close")
	}
	if allZero(dec.adBuf[:]) {
		t.Fatal("précondition: adBuf devrait être sale avant Close")
	}
	if len(dec.adExt) == 0 || allZero(dec.adExt) {
		t.Fatal("précondition: adExt devrait être sale avant Close")
	}

	if err := dec.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if !allZero(dec.subkey[:]) {
		t.Fatal("subkey non effacée après Close")
	}
	if !allZero(dec.inBuf) {
		t.Fatal("inBuf non effacé après Close")
	}
	if !allZero(dec.plainBuf) {
		t.Fatal("plainBuf non effacé après Close")
	}
	if !allZero(dec.adBuf[:]) {
		t.Fatal("adBuf non effacé après Close")
	}
	if !allZero(dec.adExt) {
		t.Fatal("adExt non effacé après Close")
	}
	if err := dec.Close(); err != nil {
		t.Fatalf("Close idempotent a rendu %v", err)
	}
}

func TestWriteAfterClose_Encryptor(t *testing.T) {
	var buf bytes.Buffer
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatal(err)
	}
	if err := enc.Close(); err != nil {
		t.Fatal(err)
	}
	n, err := enc.Write([]byte("x"))
	if n != 0 {
		t.Fatalf("Write après Close a déclaré %d octets", n)
	}
	if err == nil || !strings.Contains(err.Error(), "closed") {
		t.Fatalf("erreur de fermeture attendue, obtenu %v", err)
	}
}

func TestReadAfterClose_Decryptor(t *testing.T) {
	var buf bytes.Buffer
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := enc.Write([]byte("z")); err != nil {
		t.Fatal(err)
	}
	dec, err := NewDecryptor(bytes.NewReader(buf.Bytes()), key)
	if err != nil {
		t.Fatal(err)
	}
	if err := dec.Close(); err != nil {
		t.Fatal(err)
	}
	n, err := dec.Read(make([]byte, 8))
	if n != 0 {
		t.Fatalf("Read après Close a déclaré %d octets", n)
	}
	if err == nil || !strings.Contains(err.Error(), "closed") {
		t.Fatalf("erreur de fermeture attendue, obtenu %v", err)
	}
}
