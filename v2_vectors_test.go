// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type v2VectorsFile struct {
	Engine  string         `json:"engine"`
	Vectors []v2VectorCase `json:"vectors"`
}

type v2VectorCase struct {
	Name         string   `json:"name"`
	KeyHex       string   `json:"key_hex"`
	NonceHex     string   `json:"nonce_hex"`
	FragmentsHex []string `json:"fragments_hex"`
	AdsHex       []string `json:"ads_hex"`
	HeaderHex    string   `json:"header_hex"`
	FramesHex    []string `json:"frames_hex"`
}

func TestV2Vectors(t *testing.T) {
	raw, err := os.ReadFile(filepath.Join("testdata", "v2", "vectors.json"))
	if err != nil {
		t.Fatal(err)
	}
	var file v2VectorsFile
	if err := json.Unmarshal(raw, &file); err != nil {
		t.Fatal(err)
	}
	if file.Engine != "default" {
		t.Fatalf("moteur des vecteurs %q, attendu default", file.Engine)
	}
	if len(file.Vectors) == 0 {
		t.Fatal("aucun vecteur")
	}
	for _, v := range file.Vectors {
		t.Run(v.Name, func(t *testing.T) {
			key, err := hex.DecodeString(v.KeyHex)
			if err != nil {
				t.Fatal(err)
			}
			nonce, err := hex.DecodeString(v.NonceHex)
			if err != nil {
				t.Fatal(err)
			}
			wantHdr, err := hex.DecodeString(v.HeaderHex)
			if err != nil {
				t.Fatal(err)
			}
			var buf bytes.Buffer
			enc, err := NewEncryptor(&buf, key)
			if err != nil {
				t.Fatalf("NewEncryptor: %v", err)
			}
			var n24 [24]byte
			copy(n24[:], nonce)
			enc.nonce = n24
			enc.eng.HChaCha20(enc.subkey[:], key, enc.nonce[0:16])
			enc.seq = 0
			buf.Reset()
			var hdr [HeaderSize]byte
			writeHeaderV2(hdr[:], &enc.nonce)
			if _, err := buf.Write(hdr[:]); err != nil {
				t.Fatal(err)
			}
			for i, fh := range v.FragmentsHex {
				frag, err := hex.DecodeString(fh)
				if err != nil {
					t.Fatal(err)
				}
				var ad []byte
				if i < len(v.AdsHex) {
					ad, err = hex.DecodeString(v.AdsHex[i])
					if err != nil {
						t.Fatal(err)
					}
				}
				if _, err := enc.WriteWithAD(frag, ad); err != nil {
					t.Fatalf("WriteWithAD %d: %v", i, err)
				}
			}
			if err := enc.Close(); err != nil {
				t.Fatalf("Close: %v", err)
			}
			wire := buf.Bytes()
			if len(wire) < HeaderSize {
				t.Fatalf("fil trop court: %d", len(wire))
			}
			if !bytes.Equal(wire[:HeaderSize], wantHdr) {
				t.Fatalf("en-tête: obtenu %x, attendu %x", wire[:HeaderSize], wantHdr)
			}
			_, frames := splitV2(t, wire)
			if len(frames) != len(v.FramesHex) {
				t.Fatalf("trames: obtenu %d, attendu %d", len(frames), len(v.FramesHex))
			}
			for i, fh := range v.FramesHex {
				want, err := hex.DecodeString(fh)
				if err != nil {
					t.Fatal(err)
				}
				if !bytes.Equal(frames[i], want) {
					t.Fatalf("trame %d: obtenu %x, attendu %x", i, frames[i], want)
				}
			}
		})
	}
}
