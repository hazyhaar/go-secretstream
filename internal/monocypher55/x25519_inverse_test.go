// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55_test

import (
	"bytes"
	"crypto/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

// TestX25519_Inverse_Property teste la propriété fondamentale de l'inversion X25519 :
// sk * (sk^-1 * P) = P
func TestX25519_Inverse_Property(t *testing.T) {
	for i := 0; i < 20; i++ {
		sk := make([]byte, 32)
		if _, err := rand.Read(sk); err != nil {
			t.Fatal(err)
		}

		var pk [32]byte
		sgoi.Crypto_x25519_public_key(pk[:], sk)

		var blind [32]byte
		sgoi.Crypto_x25519_inverse(blind[:], sk, pk[:])

		// blind ne doit pas être le vecteur nul
		allZero := true
		for _, b := range blind {
			if b != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			t.Fatalf("tour %d: Crypto_x25519_inverse a retourné le vecteur nul", i)
		}

		var reconstructed [32]byte
		sgoi.Crypto_x25519(reconstructed[:], sk, blind[:])

		if !bytes.Equal(reconstructed[:], pk[:]) {
			t.Fatalf("tour %d: reconstruction échouée:\n  got:  %x\n  want: %x", i, reconstructed, pk)
		}
	}
}

// TestX25519_Inverse_VsC valide la sortie de Crypto_x25519_inverse face à l'oracle C GCC.
func TestX25519_Inverse_VsC(t *testing.T) {
	cSrc := `
#include <stdio.h>
#include <stdint.h>
#include "monocypher.h"

int main() {
    uint8_t sk[32] = {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32};
    uint8_t point[32] = {9,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0};
    uint8_t out[32];
    crypto_x25519_inverse(out, sk, point);
    for (int i = 0; i < 32; i++) {
        printf("%02x", out[i]);
    }
    printf("\n");
    return 0;
}
`
	tmpDir, err := os.MkdirTemp("", "x25519_inv_c_*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	cFile := filepath.Join(tmpDir, "main.c")
	binFile := filepath.Join(tmpDir, "c_oracle")
	if err := os.WriteFile(cFile, []byte(cSrc), 0644); err != nil {
		t.Fatal(err)
	}

	amalg, hdr := findMonocypherC(t)
	cmd := exec.Command("gcc", "-O2",
		"-I"+hdr,
		amalg,
		cFile, "-o", binFile)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("gcc compile failed: %v\n%s", err, out)
	}

	out, err := exec.Command(binFile).Output()
	if err != nil {
		t.Fatalf("oracle execution failed: %v", err)
	}
	cHex := string(bytes.TrimSpace(out))

	sk := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	point := make([]byte, 32)
	point[0] = 9

	var goOut [32]byte
	sgoi.Crypto_x25519_inverse(goOut[:], sk, point)
	var hexOut bytes.Buffer
	for _, b := range goOut {
		hexOut.WriteString(string([]byte{
			"0123456789abcdef"[b>>4],
			"0123456789abcdef"[b&0x0f],
		}))
	}

	if hexOut.String() != cHex {
		t.Fatalf("divergence oracle C vs Go:\n  Go: %s\n  C:  %s", hexOut.String(), cHex)
	}
}

// TestDifferential_ScalarMult_Ref10_vs_Fe51 teste l'oracle différentiel continu
// entre l'implémentation ref10 originale et l'implémentation donna fe51.
func TestDifferential_ScalarMult_Ref10_vs_Fe51(t *testing.T) {
	for i := 0; i < 50; i++ {
		scalar := make([]byte, 32)
		point := make([]byte, 32)
		if _, err := rand.Read(scalar); err != nil {
			t.Fatal(err)
		}
		if _, err := rand.Read(point); err != nil {
			t.Fatal(err)
		}

		var outDonna [32]byte
		var outRef10 [32]byte

		var clamped [32]byte
		copy(clamped[:], scalar)
		clamped[0] &= 248
		clamped[31] &= 127
		clamped[31] |= 64

		// Donna 51 bits (voie active)
		sgoi.Scalarmult51(outDonna[:], clamped[:], point, 255)

		// Ref10 10x25.5 bits (oracle interne)
		sgoi.Scalarmult_ref10(outRef10[:], clamped[:], point, 255)

		if !bytes.Equal(outDonna[:], outRef10[:]) {
			t.Fatalf("tour %d: divergence d'oracle entre ref10 et fe51:\n  fe51:  %x\n  ref10: %x",
				i, outDonna, outRef10)
		}
	}
}
