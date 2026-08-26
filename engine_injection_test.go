// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/engine"
)

// Plomberie d'injection : un moteur passé par NewEncryptorWithEngine produit
// un flux relu par NewDecryptorWithEngine avec le même moteur, et un flux
// bit-identique à celui du constructeur par défaut quand le moteur injecté
// est le moteur par défaut. Un moteur nil est refusé.
func TestWithEngine_RoundTripAndDefaultParity(t *testing.T) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	payload := make([]byte, 3*ChunkSize+777)
	if _, err := rand.Read(payload); err != nil {
		t.Fatal(err)
	}

	var injected bytes.Buffer
	enc, err := NewEncryptorWithEngine(&injected, key, engine.Default())
	if err != nil {
		t.Fatal(err)
	}
	if _, err := enc.Write(payload); err != nil {
		t.Fatal(err)
	}
	if err := enc.Close(); err != nil {
		t.Fatal(err)
	}

	dec, err := NewDecryptorWithEngine(bytes.NewReader(injected.Bytes()), key, engine.Default())
	if err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(dec)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, payload) {
		t.Fatal("round-trip avec moteur injecté : clair divergent")
	}

	// Même nonce impossible à forcer (tiré au hasard dans le constructeur) :
	// la parité avec le constructeur par défaut se vérifie sur la relecture
	// croisée, pas sur l'égalité des octets du fil.
	dec2, err := NewDecryptor(bytes.NewReader(injected.Bytes()), key)
	if err != nil {
		t.Fatal(err)
	}
	got2, err := io.ReadAll(dec2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got2, payload) {
		t.Fatal("flux du moteur injecté illisible par le décodeur par défaut")
	}

	if _, err := NewEncryptorWithEngine(io.Discard, key, nil); err == nil {
		t.Fatal("NewEncryptorWithEngine accepte un moteur nil")
	}
	if _, err := NewDecryptorWithEngine(bytes.NewReader(nil), key, nil); err == nil {
		t.Fatal("NewDecryptorWithEngine accepte un moteur nil")
	}
}
