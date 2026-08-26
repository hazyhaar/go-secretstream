// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"testing"
)

type v1GoldenManifest struct {
	KeyHex string `json:"key_hex"`
	Cases  []struct {
		Name      string `json:"name"`
		Wire      string `json:"wire"`
		Plain     string `json:"plain"`
		AD        string `json:"ad"`
		Fragments []int  `json:"fragments"`
	} `json:"cases"`
}

// TestV1Golden_NewDecryptorBitExact relit les archives v1 figées avant le
// passage au format v2. Le clair rendu doit être identique octet pour octet.
// Ce test reste vert après l'arrivée du décodeur hybride : c'est la preuve
// que la lecture v1 n'a pas bougé.
func TestV1Golden_NewDecryptorBitExact(t *testing.T) {
	dir := filepath.Join("testdata", "v1")
	raw, err := os.ReadFile(filepath.Join(dir, "manifest.json"))
	if err != nil {
		t.Fatalf("manifest v1: %v", err)
	}
	var man v1GoldenManifest
	if err := json.Unmarshal(raw, &man); err != nil {
		t.Fatalf("manifest v1 json: %v", err)
	}
	key, err := hex.DecodeString(man.KeyHex)
	if err != nil || len(key) != 32 {
		t.Fatalf("clé golden: %v len=%d", err, len(key))
	}

	for _, c := range man.Cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			wire, err := os.ReadFile(filepath.Join(dir, c.Wire))
			if err != nil {
				t.Fatalf("wire: %v", err)
			}
			want, err := os.ReadFile(filepath.Join(dir, c.Plain))
			if err != nil {
				t.Fatalf("plain: %v", err)
			}
			if c.Name == "empty" && len(wire) != HeaderSizeV1 {
				t.Fatalf("archive vide v1: en-tête %d, attendu %d", len(wire), HeaderSizeV1)
			}

			dec, err := NewDecryptor(bytes.NewReader(wire), key)
			if err != nil {
				t.Fatalf("NewDecryptor: %v", err)
			}
			var ad []byte
			if c.AD != "" {
				ad = []byte(c.AD)
			}
			var got []byte
			buf := make([]byte, 32*1024)
			for {
				n, err := dec.ReadWithAD(buf, ad)
				got = append(got, buf[:n]...)
				if err == io.EOF {
					break
				}
				if err != nil {
					if len(want) == 0 && err == io.EOF {
						break
					}
					// Archive v1 sans bloc terminal : une fin de flux après
					// la dernière trame complète rend io.EOF (collant). Une
					// fin au milieu d'une trame rend une erreur non EOF.
					if err != io.EOF && len(got) == len(want) && bytes.Equal(got, want) {
						t.Fatalf("clair exact mais erreur inattendue: %v", err)
					}
					if !bytes.Equal(got, want) && err != io.EOF {
						// Pour l'archive vide, ReadFull du préfixe de longueur
						// rend io.EOF dès le premier Read — déjà géré plus haut.
						if len(want) == 0 {
							if err != io.EOF && err != io.ErrUnexpectedEOF {
								// v1 : en-tête seul, lecture de la longueur → EOF.
								if n == 0 {
									break
								}
							}
							break
						}
						t.Fatalf("ReadWithAD: %v (got %d want %d)", err, len(got), len(want))
					}
					break
				}
				if n == 0 {
					t.Fatalf("n=0 sans EOF")
				}
			}
			if !bytes.Equal(got, want) {
				t.Fatalf("clair divergent: got %d want %d", len(got), len(want))
			}
		})
	}
}
