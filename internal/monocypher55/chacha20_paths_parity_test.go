package monocypher55

import (
	"bytes"
	"testing"
)

// TestChacha20DjbPathsParity : pour toute taille, chiffrer en UN appel doit
// égaler le chiffrement en morceaux à frontières de blocs (le compteur rendu
// enchaîne). Les découpes forcent chaque chemin (scalaire, paire 256-bit,
// kernel 4x) à produire les mêmes octets que les autres sur les mêmes blocs —
// garde du chemin par paires ajouté le 2026-08-15 pour la bande 65..255 o.
func TestChacha20DjbPathsParity(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 8)
	for i := range key {
		key[i] = byte(i*3 + 1)
	}
	for i := range nonce {
		nonce[i] = byte(0x50 + i)
	}
	sizes := []int{1, 63, 64, 65, 127, 128, 129, 191, 192, 193, 255, 256, 257, 319, 320, 384, 448, 511, 512, 576, 640, 1024}
	splits := []int{64, 128, 192, 256, 320}
	const baseCtr = uint64(0xfffffffd) // traverse le wrap 32-bit du compteur

	for _, n := range sizes {
		plain := make([]byte, n)
		for i := range plain {
			plain[i] = byte(i * 7)
		}
		whole := make([]byte, n)
		endCtr := Crypto_chacha20_djb(whole, plain, uint64(n), key, nonce, baseCtr)

		for _, sp := range splits {
			if sp >= n {
				continue
			}
			parts := make([]byte, n)
			c := Crypto_chacha20_djb(parts[:sp], plain[:sp], uint64(sp), key, nonce, baseCtr)
			c = Crypto_chacha20_djb(parts[sp:], plain[sp:], uint64(n-sp), key, nonce, c)
			if !bytes.Equal(parts, whole) {
				t.Fatalf("n=%d split=%d : flux divergent", n, sp)
			}
			if c != endCtr {
				t.Fatalf("n=%d split=%d : compteur final %d != %d", n, sp, c, endCtr)
			}
		}

		// keystream brut (plain nil) : mêmes chemins, même exigence.
		wholeKS := make([]byte, n)
		Crypto_chacha20_djb(wholeKS, nil, uint64(n), key, nonce, baseCtr)
		for i := range wholeKS {
			if wholeKS[i]^0 != whole[i]^plain[i] {
				t.Fatalf("n=%d : keystream brut divergent du XOR à l'octet %d", n, i)
			}
		}
	}
}
