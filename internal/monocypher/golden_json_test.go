package monocypher_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
)

type goldenFile struct {
	KeyHex   string `json:"key_hex"`
	NonceHex string `json:"nonce_hex"`
	Cases    []struct {
		Name         string `json:"name"`
		ADHex        string `json:"ad_hex"`
		PTHex        string `json:"pt_hex"`
		PTLen        int    `json:"pt_len"`
		PTPattern    string `json:"pt_pattern"`
		CTHex        string `json:"ct_hex"`
		MACHex       string `json:"mac_hex"`
		CTHead32Hex  string `json:"ct_head32_hex"`
	} `json:"cases"`
}

func TestGoldenJSON_KnownVectors(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	// monocypher_sgoiter -> devhoros/c2simd/sgoiter/testdata/kat
	p := filepath.Join(filepath.Dir(file), "..", "..", "c2simd", "sgoiter", "testdata", "kat", "aead_sizes.json")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Skip(err)
	}
	var g goldenFile
	if err := json.Unmarshal(b, &g); err != nil {
		t.Fatal(err)
	}
	key, _ := hex.DecodeString(g.KeyHex)
	nonce, _ := hex.DecodeString(g.NonceHex)
	// fix nonce length: json has 25 bytes hex? check - 24 bytes = 48 hex chars
	if len(nonce) > 24 {
		nonce = nonce[:24]
	}
	for _, c := range g.Cases {
		if c.CTHex == "" && c.MACHex == "" && c.CTHead32Hex == "" {
			continue // template only
		}
		c := c
		t.Run(c.Name, func(t *testing.T) {
			ad, _ := hex.DecodeString(c.ADHex)
			var pt []byte
			if c.PTHex != "" {
				pt, _ = hex.DecodeString(c.PTHex)
			} else if c.PTLen > 0 {
				pt = mkPT(c.PTLen, c.PTPattern)
			}
			ct, mac, err := sgoi.AEADLock(key, nonce, ad, pt)
			if err != nil {
				t.Fatal(err)
			}
			if c.MACHex != "" {
				want, _ := hex.DecodeString(c.MACHex)
				if !bytes.Equal(mac, want) {
					t.Fatalf("mac got %x want %s", mac, c.MACHex)
				}
			}
			if c.CTHex != "" {
				want, _ := hex.DecodeString(c.CTHex)
				if !bytes.Equal(ct, want) {
					t.Fatalf("ct mismatch")
				}
			}
			if c.CTHead32Hex != "" && len(ct) >= 32 {
				want, _ := hex.DecodeString(c.CTHead32Hex)
				if !bytes.Equal(ct[:32], want) {
					t.Fatalf("ct head got %x want %s", ct[:32], c.CTHead32Hex)
				}
			}
		})
	}
}
