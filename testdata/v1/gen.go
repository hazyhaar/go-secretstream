//go:build ignore

// SPDX-License-Identifier: Apache-2.0 OR MIT

// Génère les archives golden v1 à partir de l'Encryptor courant (format v1).
// À exécuter une seule fois AVANT toute modification de secretstream.go :
//
//	env GOWORK=/devhoros/c2simd/go.work GOTOOLCHAIN=go1.27.0 GOEXPERIMENT=simd \
//	  go run ./testdata/v1/gen.go
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	secretstream55 "github.com/hazyhaar/go-secretstream"
)

type manifest struct {
	KeyHex string     `json:"key_hex"`
	Cases  []caseMeta `json:"cases"`
}

type caseMeta struct {
	Name      string `json:"name"`
	Wire      string `json:"wire"`
	Plain     string `json:"plain"`
	AD        string `json:"ad"`
	Fragments []int  `json:"fragments,omitempty"`
	Note      string `json:"note"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func writeFile(path string, b []byte) {
	must(os.WriteFile(path, b, 0o644))
}

func main() {
	dir, err := os.Getwd()
	must(err)
	if filepath.Base(dir) != "v1" {
		dir = filepath.Join(dir, "testdata", "v1")
	}
	must(os.MkdirAll(dir, 0o755))

	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(0x51 ^ i*13)
	}

	// (a) vide : 0 fragment, pas de Close — en-tête nonce seul.
	var emptyWire []byte
	{
		f, err := os.Create(filepath.Join(dir, "empty.wire"))
		must(err)
		enc, err := secretstream55.NewEncryptor(f, key)
		must(err)
		_ = enc
		must(f.Close())
		emptyWire, err = os.ReadFile(filepath.Join(dir, "empty.wire"))
		must(err)
		writeFile(filepath.Join(dir, "empty.plain"), nil)
	}

	// (b) un fragment de 1 000 octets avec AD "ad-1".
	plainB := make([]byte, 1000)
	for i := range plainB {
		plainB[i] = byte(i*41 + 3)
	}
	adB := []byte("ad-1")
	{
		f, err := os.Create(filepath.Join(dir, "ad1000.wire"))
		must(err)
		enc, err := secretstream55.NewEncryptor(f, key)
		must(err)
		_, err = enc.WriteWithAD(plainB, adB)
		must(err)
		must(f.Close())
		writeFile(filepath.Join(dir, "ad1000.plain"), plainB)
	}

	// (c) trois fragments dont un de ChunkSize+1, sans AD.
	frag1 := make([]byte, 100)
	for i := range frag1 {
		frag1[i] = byte(0x11 + i)
	}
	frag2 := make([]byte, secretstream55.ChunkSize+1)
	for i := range frag2 {
		frag2[i] = byte(0x22 + i*3)
	}
	frag3 := make([]byte, 50)
	for i := range frag3 {
		frag3[i] = byte(0x33 + i*5)
	}
	plainC := append(append(append([]byte{}, frag1...), frag2...), frag3...)
	{
		f, err := os.Create(filepath.Join(dir, "three.wire"))
		must(err)
		enc, err := secretstream55.NewEncryptor(f, key)
		must(err)
		_, err = enc.Write(frag1)
		must(err)
		_, err = enc.Write(frag2)
		must(err)
		_, err = enc.Write(frag3)
		must(err)
		must(f.Close())
		writeFile(filepath.Join(dir, "three.plain"), plainC)
	}

	m := manifest{
		KeyHex: hex.EncodeToString(key),
		Cases: []caseMeta{
			{
				Name:  "empty",
				Wire:  "empty.wire",
				Plain: "empty.plain",
				AD:    "",
				Note:  "0 fragment ; en-tête nonce v1 seul ; pas de Close",
			},
			{
				Name:      "ad1000",
				Wire:      "ad1000.wire",
				Plain:     "ad1000.plain",
				AD:        "ad-1",
				Fragments: []int{1000},
				Note:      "un fragment de 1000 octets lié à AD ad-1",
			},
			{
				Name:      "three",
				Wire:      "three.wire",
				Plain:     "three.plain",
				AD:        "",
				Fragments: []int{100, secretstream55.ChunkSize + 1, 50},
				Note:      "trois Write sans AD ; le second dépasse ChunkSize et se coupe en deux trames",
			},
		},
	}
	raw, err := json.MarshalIndent(m, "", "  ")
	must(err)
	writeFile(filepath.Join(dir, "manifest.json"), append(raw, '\n'))

	fmt.Printf("golden v1 écrits sous %s (empty.wire=%d octets)\n", dir, len(emptyWire))
}
