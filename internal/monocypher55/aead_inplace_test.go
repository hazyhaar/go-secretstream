// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"
)

// aeadInPlaceSizes couvre le fast-path court (<= 64 exclu ici, 48 le couvre),
// le chemin scalaire (< 256), les frontieres du chemin SIMD 256 octets
// (255/256/257), un multiple exact (512) et deux grandes tailles (8192, 65536).
var aeadInPlaceSizes = []int{48, 200, 255, 256, 257, 512, 8192, 65536}

// TestAEADUnlockInPlace verifie le contrat Monocypher : le dechiffrement
// in-place (plain_text et cipher_text sont le MEME slice, alias exact) doit
// produire le meme plaintext que le dechiffrement hors-place.
func TestAEADUnlockInPlace(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	rand.Read(key)
	rand.Read(nonce)

	for _, withAD := range []bool{false, true} {
		var ad []byte
		if withAD {
			ad = make([]byte, 37)
			rand.Read(ad)
		}
		for _, n := range aeadInPlaceSizes {
			t.Run(fmt.Sprintf("size=%d/ad=%v", n, withAD), func(t *testing.T) {
				plain := make([]byte, n)
				rand.Read(plain)

				ct := make([]byte, n)
				mac := make([]byte, 16)
				if err := LockDst(ct, mac, key, nonce, ad, plain); err != nil {
					t.Fatalf("LockDst: %v", err)
				}

				// Reference hors-place.
				ref := make([]byte, n)
				if err := UnlockDst(ref, key, nonce, mac, ad, ct); err != nil {
					t.Fatalf("UnlockDst hors-place: %v", err)
				}
				if !bytes.Equal(ref, plain) {
					t.Fatalf("hors-place: plaintext errone")
				}

				// In-place : dst == ciphertext, meme slice.
				buf := make([]byte, n)
				copy(buf, ct)
				if err := UnlockDst(buf, key, nonce, mac, ad, buf); err != nil {
					t.Fatalf("UnlockDst in-place: %v", err)
				}
				if !bytes.Equal(buf, plain) {
					t.Fatalf("in-place: plaintext different de la reference (taille %d)", n)
				}
			})
		}
	}
}

// TestAEADUnlockInPlaceBadMAC verifie le contrat wipe-on-failure documente :
// un MAC corrompu sur un dechiffrement in-place doit echouer ET laisser le
// buffer entierement efface (l'appelant perd le contenu, comme en C).
func TestAEADUnlockInPlaceBadMAC(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	rand.Read(key)
	rand.Read(nonce)
	ad := []byte("additional data")

	for _, n := range aeadInPlaceSizes {
		t.Run(fmt.Sprintf("size=%d", n), func(t *testing.T) {
			plain := make([]byte, n)
			rand.Read(plain)

			buf := make([]byte, n)
			mac := make([]byte, 16)
			if err := LockDst(buf, mac, key, nonce, ad, plain); err != nil {
				t.Fatalf("LockDst: %v", err)
			}
			mac[0] ^= 0x01 // corruption du tag

			err := UnlockDst(buf, key, nonce, mac, ad, buf)
			if err == nil {
				t.Fatalf("MAC corrompu accepte")
			}
			zero := make([]byte, n)
			if !bytes.Equal(buf, zero) {
				t.Fatalf("wipe-on-failure viole : buffer non efface (taille %d)", n)
			}
		})
	}
}
