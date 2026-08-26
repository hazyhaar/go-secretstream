// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"crypto/rand"
	"io"
	"strings"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/engine"
)

// TestContract_TruncationRejection vérifie que la troncature du flux libsodium
// (absence du tag final TAG_FINAL) est rejetée avec io.ErrUnexpectedEOF
// et n'est jamais acceptée silencieusement comme un EOF normal.
func TestContract_TruncationRejection(t *testing.T) {
	key := make([]byte, 32)
	_, _ = rand.Read(key)

	var buf bytes.Buffer
	w, err := NewLibsodiumEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("échec NewLibsodiumEncryptor: %v", err)
	}

	// Émettre un flux multi-chunks
	payload := make([]byte, ChunkSize*2)
	for i := range payload {
		payload[i] = byte(i)
	}
	if _, err := w.Write(payload); err != nil {
		t.Fatalf("échec Write: %v", err)
	}
	if err := w.Close(); err != nil {
		t.Fatalf("échec Close: %v", err)
	}

	wire := buf.Bytes()
	// Tronquer après le premier chunk de données (HeaderSize + ChunkSize + ABytes)
	firstChunkWireLen := HeaderSize + ChunkSize + 17
	if len(wire) <= firstChunkWireLen {
		t.Fatalf("fil produit trop court pour le test de troncature: %d", len(wire))
	}
	truncatedWire := wire[:firstChunkWireLen]

	r, err := NewLibsodiumDecryptor(bytes.NewReader(truncatedWire), key)
	if err != nil {
		t.Fatalf("échec NewLibsodiumDecryptor: %v", err)
	}

	out := make([]byte, len(payload))
	totalRead := 0
	for {
		n, err := r.Read(out[totalRead:])
		totalRead += n
		if err != nil {
			if err == io.EOF {
				t.Fatalf("FAIL: Flux tronqué accepté avec io.EOF propre (troncature silencieuse) !")
			}
			// Toute erreur non-EOF (ErrUnexpectedEOF, échec MAC, etc.) prouve le rejet de troncature
			return
		}
	}
}

// TestContract_StickyErrorLibsodium vérifie qu'après un échec de MAC,
// le lecteur verrouille son état et ne panique jamais lors d'appels subséquents à Read().
func TestContract_StickyErrorLibsodium(t *testing.T) {
	key := make([]byte, 32)
	_, _ = rand.Read(key)

	var buf bytes.Buffer
	w, err := NewLibsodiumEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("échec NewLibsodiumEncryptor: %v", err)
	}
	_, _ = w.Write([]byte("message secret à protéger"))
	_ = w.Close()

	wire := buf.Bytes()
	// Altérer le MAC du fragment
	wire[len(wire)-5] ^= 0xFF

	r, err := NewLibsodiumDecryptor(bytes.NewReader(wire), key)
	if err != nil {
		t.Fatalf("échec NewLibsodiumDecryptor: %v", err)
	}

	tmp := make([]byte, 64)
	_, err1 := r.Read(tmp)
	if err1 == nil {
		t.Fatalf("Erreur attendue sur MAC altéré, obtenu nil")
	}

	// Deuxième lecture : doit renvoyer l'erreur mémorisée sans panique
	_, err2 := r.Read(tmp)
	if err2 == nil {
		t.Fatalf("Deuxième lecture : erreur collante attendue, obtenu nil")
	}
}

// TestContract_SubkeyAliasingOverlap vérifie que le recouvrement inexact de mémoire
// entre dst et src rend une erreur (pas une panique) sur LockSubkeyDst et UnlockSubkeyDst.
func TestContract_SubkeyAliasingOverlap(t *testing.T) {
	eng := engine.Default()
	subkey := make([]byte, 32)
	nonce12 := make([]byte, 12)
	var mac [16]byte

	t.Run("LockSubkeyDst", func(t *testing.T) {
		buf := make([]byte, 128)
		dst := buf[0:64]
		plain := buf[1:65]
		if err := eng.LockSubkeyDst(dst, &mac, subkey, nonce12, nil, plain); err == nil {
			t.Fatalf("FAIL: LockSubkeyDst avec recouvrement inexact aurait dû rendre une erreur")
		}
	})

	t.Run("UnlockSubkeyDst", func(t *testing.T) {
		buf := make([]byte, 128)
		dst := buf[0:64]
		cipher := buf[1:65]
		if _, err := eng.UnlockSubkeyDst(dst, subkey, nonce12, nil, cipher, mac[:]); err == nil {
			t.Fatalf("FAIL: UnlockSubkeyDst avec recouvrement inexact aurait dû rendre une erreur")
		}
	})
}

// TestContract_MaisonIntraChunkTruncationRejection vérifie que la coupure physique
// du flux au milieu d'un fragment du format maison est rejetée en erreur dure.
func TestContract_MaisonIntraChunkTruncationRejection(t *testing.T) {
	key := make([]byte, 32)
	_, _ = rand.Read(key)

	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("échec NewEncryptor: %v", err)
	}

	payload := make([]byte, ChunkSize)
	for i := range payload {
		payload[i] = byte(i)
	}
	if _, err := enc.Write(payload); err != nil {
		t.Fatalf("échec Write: %v", err)
	}

	wire := buf.Bytes()
	// Tronquer au milieu du fragment chiffré
	truncatedWire := wire[:HeaderSize+4+100]

	dec, err := NewDecryptor(bytes.NewReader(truncatedWire), key)
	if err != nil {
		t.Fatalf("échec NewDecryptor: %v", err)
	}

	out := make([]byte, len(payload))
	_, err = dec.Read(out)
	if err == nil {
		t.Fatalf("FAIL: Lecture sur fragment tronqué a réussi avec nil !")
	}
	if err == io.EOF {
		t.Fatalf("FAIL: Fragment coupé accepté comme EOF normal !")
	}
}

// TestContract_StickyErrorMaison vérifie qu'après une erreur sur le déchiffreur maison,
// l'état est scellé et un second Read() retourne la même erreur sans panique.
func TestContract_StickyErrorMaison(t *testing.T) {
	key := make([]byte, 32)
	_, _ = rand.Read(key)

	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatalf("échec NewEncryptor: %v", err)
	}
	_, _ = enc.Write([]byte("données à protéger"))

	wire := buf.Bytes()
	// Altérer le MAC
	wire[len(wire)-5] ^= 0xFF

	dec, err := NewDecryptor(bytes.NewReader(wire), key)
	if err != nil {
		t.Fatalf("échec NewDecryptor: %v", err)
	}

	tmp := make([]byte, 64)
	_, err1 := dec.Read(tmp)
	if err1 == nil {
		t.Fatalf("Erreur attendue sur MAC altéré, obtenu nil")
	}

	_, err2 := dec.Read(tmp)
	if err2 == nil {
		t.Fatalf("Deuxième lecture : erreur collante attendue, obtenu nil")
	}
}

// TestContract_MaisonInterChunkTruncation vérifie qu'une coupure sur
// frontière de trame v2 (sans TagFinal) est détectée. En v1 ce trou reste
// ouvert : une archive qui s'arrête après une trame complète se lit comme
// complète (preuve : testdata/v1/).
func TestContract_MaisonInterChunkTruncation(t *testing.T) {
	key := make([]byte, 32)
	_, _ = rand.Read(key)
	var buf bytes.Buffer
	enc, err := NewEncryptor(&buf, key)
	if err != nil {
		t.Fatal(err)
	}
	payload := make([]byte, ChunkSize*2)
	if _, err := enc.Write(payload); err != nil {
		t.Fatal(err)
	}
	wire := buf.Bytes()
	first := HeaderSize + 4 + 1 + ChunkSize + TagSize
	if len(wire) <= first {
		t.Fatalf("fil trop court: %d", len(wire))
	}
	dec, err := NewDecryptor(bytes.NewReader(wire[:first]), key)
	if err != nil {
		t.Fatal(err)
	}
	out := make([]byte, len(payload))
	total := 0
	var last error
	var lastN int
	for {
		n, e := dec.Read(out[total:])
		total += n
		if e != nil {
			last = e
			lastN = n
			break
		}
	}
	if last == io.EOF {
		t.Fatalf("FAIL: troncature sur frontière acceptée comme io.EOF (total=%d)", total)
	}
	if last == nil || !strings.Contains(last.Error(), "flux tronqué") {
		t.Fatalf("attendu flux tronqué, obtenu err=%v total=%d", last, total)
	}
	if lastN != 0 {
		t.Fatalf("n=%d sur l'erreur de troncature", lastN)
	}
}
