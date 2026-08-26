// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"bytes"
	"testing"
)

func FuzzAEADRoundtrip(f *testing.F) {
	// Corpus initial
	f.Add([]byte("hello world"), []byte("ad data"), byte(0))
	f.Add(make([]byte, 64), make([]byte, 32), byte(1))
	f.Add(make([]byte, 256), make([]byte, 16), byte(2))
	f.Add(make([]byte, 1024), make([]byte, 0), byte(3))

	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(i + 1)
	}
	nonce := make([]byte, 24)
	for i := range nonce {
		nonce[i] = byte(i + 10)
	}

	f.Fuzz(func(t *testing.T, plain []byte, ad []byte, seed byte) {
		if len(plain) > 65536 {
			plain = plain[:65536]
		}
		if len(ad) > 4096 {
			ad = ad[:4096]
		}

		cipher := make([]byte, len(plain))
		dec := make([]byte, len(plain))
		var mac [16]byte

		Crypto_aead_lock(cipher, mac[:], key, nonce, ad, uint64(len(ad)), plain, uint64(len(plain)))
		res := Crypto_aead_unlock(dec, mac[:], key, nonce, ad, uint64(len(ad)), cipher, uint64(len(cipher)))
		if res != 0 {
			t.Fatalf("Fuzz unlock failed on len(plain)=%d len(ad)=%d", len(plain), len(ad))
		}
		if !bytes.Equal(plain, dec) {
			t.Fatalf("Fuzz mismatch on len(plain)=%d len(ad)=%d", len(plain), len(ad))
		}
	})
}

// FuzzChaCha20DJB vérifie l'invariant de chiffrement/déchiffrement, le compteur de retour
// et la parité stricte avec l'oracle scalaire sur tout le domaine des compteurs 64-bit,
// en ciblant particulièrement les frontières de franchissement 2^32.
func FuzzChaCha20DJB(f *testing.F) {
	// Graines ciblant les franchissements 2^32 avec tous les alignements possibles
	f.Add(uint64(0), []byte("test message 64 bytes aligned length for initial fuzz test!"), byte(1))
	f.Add(uint64(0xFFFFFFF8), make([]byte, 1024), byte(2))
	f.Add(uint64(0xFFFFFFF9), make([]byte, 1024), byte(3))
	f.Add(uint64(0xFFFFFFFA), make([]byte, 512), byte(4))
	f.Add(uint64(0xFFFFFFFB), make([]byte, 512), byte(5))
	f.Add(uint64(0xFFFFFFFC), make([]byte, 512), byte(6))
	f.Add(uint64(0xFFFFFFFE), make([]byte, 512), byte(7))
	f.Add(uint64(0xFFFFFFFF), make([]byte, 512), byte(8))
	f.Add(uint64(0x100000000), make([]byte, 512), byte(9))

	key := make([]byte, 32)
	nonce := make([]byte, 8)
	for i := range key {
		key[i] = byte(i + 3)
	}
	for i := range nonce {
		nonce[i] = byte(i + 7)
	}

	f.Fuzz(func(t *testing.T, ctr uint64, plain []byte, seed byte) {
		if len(plain) > 65536 {
			plain = plain[:65536]
		}
		if len(plain) == 0 {
			return
		}

		cipher := make([]byte, len(plain))
		dec := make([]byte, len(plain))

		// 1. Chiffrement SIMD
		retCtr := Crypto_chacha20_djb(cipher, plain, uint64(len(plain)), key, nonce, ctr)

		// Vérification de l'invariant de retour du compteur : ctr + (len(plain) + 63)/64
		expectedBlocks := (uint64(len(plain)) + 63) / 64
		if retCtr != ctr+expectedBlocks {
			t.Fatalf("retCtr invariant violation: got %d, want %d (ctr=%d, blocks=%d)", retCtr, ctr+expectedBlocks, ctr, expectedBlocks)
		}

		// 2. Déchiffrement
		Crypto_chacha20_djb(dec, cipher, uint64(len(cipher)), key, nonce, ctr)
		if !bytes.Equal(plain, dec) {
			t.Fatalf("ChaCha20 roundtrip mismatch on len=%d, ctr=0x%x", len(plain), ctr)
		}

		// 3. Parité stricte avec l'oracle scalaire bloc par bloc
		cipherScal := make([]byte, len(plain))
		blocks := uint64(len(plain)) / 64
		for b := uint64(0); b < blocks; b++ {
			bOff := b * 64
			Crypto_chacha20_djb(cipherScal[bOff:bOff+64], plain[bOff:bOff+64], 64, key, nonce, ctr+b)
		}
		rem := uint64(len(plain)) % 64
		if rem > 0 {
			bOff := blocks * 64
			Crypto_chacha20_djb(cipherScal[bOff:], plain[bOff:], rem, key, nonce, ctr+blocks)
		}

		if !bytes.Equal(cipher, cipherScal) {
			t.Fatalf("SIMD vs Scalar parity mismatch on len=%d, ctr=0x%x", len(plain), ctr)
		}
	})
}
