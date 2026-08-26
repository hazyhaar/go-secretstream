// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestAEAD_BoundariesAndUnaligned(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	if _, err := rand.Read(nonce); err != nil {
		t.Fatal(err)
	}

	// Tailles critiques: seuils de blocs 16, 64, 256, 512, 1024, 4096 et désalignements
	testSizes := []int{
		0, 1, 2, 15, 16, 17, 31, 32, 33, 63, 64, 65, 127, 128, 129,
		255, 256, 257, 511, 512, 513, 1023, 1024, 1025, 4095, 4096, 4097, 65535, 65536, 65537,
	}
	adSizes := []int{0, 1, 15, 16, 17, 32, 64, 128, 512, 1024}
	offsets := []int{0, 1, 2, 3, 7, 11, 15}

	for _, size := range testSizes {
		for _, adSize := range adSizes {
			for _, offset := range offsets {
				plainBuffer := make([]byte, size+offset+64)
				adBuffer := make([]byte, adSize+offset+64)
				cipherBuffer := make([]byte, size+offset+64)
				decBuffer := make([]byte, size+offset+64)
				var mac [16]byte

				plain := plainBuffer[offset : offset+size]
				ad := adBuffer[offset : offset+adSize]
				cipher := cipherBuffer[offset : offset+size]
				dec := decBuffer[offset : offset+size]

				if _, err := rand.Read(plain); err != nil {
					t.Fatal(err)
				}
				if _, err := rand.Read(ad); err != nil {
					t.Fatal(err)
				}

				// Chiffrement
				Crypto_aead_lock(cipher, mac[:], key, nonce, ad, uint64(len(ad)), plain, uint64(len(plain)))

				// Déchiffrement
				res := Crypto_aead_unlock(dec, mac[:], key, nonce, ad, uint64(len(ad)), cipher, uint64(len(cipher)))
				if res != 0 {
					t.Fatalf("Crypto_aead_unlock failed for size=%d adSize=%d offset=%d res=%d", size, adSize, offset, res)
				}
				if !bytes.Equal(plain, dec) {
					t.Fatalf("Plaintext mismatch for size=%d adSize=%d offset=%d", size, adSize, offset)
				}

				// Test d'altération du texte chiffré (Tampering)
				if size > 0 {
					corruptCipher := make([]byte, len(cipher))
					copy(corruptCipher, cipher)
					corruptCipher[size/2] ^= 0x55 // altération d'un octet

					tRes := Crypto_aead_unlock(dec, mac[:], key, nonce, ad, uint64(len(ad)), corruptCipher, uint64(len(corruptCipher)))
					if tRes == 0 {
						t.Fatalf("Tampered cipher accepted for size=%d adSize=%d", size, adSize)
					}
				}

				// Test d'altération de la MAC
				corruptMac := mac
				corruptMac[0] ^= 0x01
				mRes := Crypto_aead_unlock(dec, corruptMac[:], key, nonce, ad, uint64(len(ad)), cipher, uint64(len(cipher)))
				if mRes == 0 {
					t.Fatalf("Tampered MAC accepted for size=%d adSize=%d", size, adSize)
				}
			}
		}
	}
}

func TestCryptoWipe_Zeroization(t *testing.T) {
	secret := make([]byte, 128)
	if _, err := rand.Read(secret); err != nil {
		t.Fatal(err)
	}
	Crypto_wipe(secret, uint64(len(secret)))
	for i, b := range secret {
		if b != 0 {
			t.Fatalf("Secret byte at index %d not wiped (got %x)", i, b)
		}
	}
}

// TestChaCha20_CounterWrap64Bit vérifie la propagation de retenue 64-bit sur le compteur ChaCha20
// pour tous les décalages d'alignement (0xFFFFFFF8 à 0xFFFFFFFF) sur 16 blocs (1024 octets).
func TestChaCha20_CounterWrap64Bit(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 8)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 5)
	}

	testCtrs := []uint64{
		0xFFFFFFF8, 0xFFFFFFF9, 0xFFFFFFFA, 0xFFFFFFFB,
		0xFFFFFFFC, 0xFFFFFFFD, 0xFFFFFFFE, 0xFFFFFFFF,
	}

	for _, ctr := range testCtrs {
		const totalBlocks = 16
		const totalBytes = totalBlocks * 64
		plain := make([]byte, totalBytes)
		for i := range plain {
			plain[i] = byte((i*13 + 7) % 251)
		}

		cipherSimd := make([]byte, totalBytes)
		cipherScal := make([]byte, totalBytes)

		// SIMD (1 appel 1024 octets, 16 blocs)
		Crypto_chacha20_djb(cipherSimd, plain, totalBytes, key, nonce, ctr)

		// Scalaire bloc par bloc (taille 64 octets < 256 octets SIMD)
		for b := uint64(0); b < totalBlocks; b++ {
			bOff := b * 64
			Crypto_chacha20_djb(cipherScal[bOff:bOff+64], plain[bOff:bOff+64], 64, key, nonce, ctr+b)
		}

		if !bytes.Equal(cipherSimd, cipherScal) {
			for b := 0; b < totalBlocks; b++ {
				bOff := b * 64
				if !bytes.Equal(cipherSimd[bOff:bOff+64], cipherScal[bOff:bOff+64]) {
					t.Errorf("Mismatch on ctr=0x%x block %d (blockCtr=0x%x): simd != scalar", ctr, b, ctr+uint64(b))
				}
			}
			t.Fatalf("Counter wrap 64-bit mismatch on ctr=0x%x", ctr)
		}
	}
}

// TestAEAD_WipeOnFailure_Tampering certifie formellement le contrat de sécurité
// WIPE-ON-FAILURE : si le MAC est altéré ou corrompu, le tampon plain_text doit être
// intégralement effacé à zéro (0x00), interdisant toute fuite de plaintext non authentifié.
func TestAEAD_WipeOnFailure_Tampering(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	if _, err := rand.Read(nonce); err != nil {
		t.Fatal(err)
	}

	testSizes := []int{1, 16, 63, 64, 65, 255, 256, 257, 1024, 4096, 65536}
	for _, size := range testSizes {
		plain := make([]byte, size)
		for i := range plain {
			plain[i] = byte(i%251 + 1) // octets strictement non nuls
		}
		cipher := make([]byte, size)
		dec := make([]byte, size)
		var mac [16]byte

		Crypto_aead_lock(cipher, mac[:], key, nonce, nil, 0, plain, uint64(size))

		// 1. Déchiffrement avec MAC altéré
		badMac := mac
		badMac[0] ^= 0xFF
		copy(dec, plain) // initialisé avec des données non nulles

		res := Crypto_aead_unlock(dec, badMac[:], key, nonce, nil, 0, cipher, uint64(size))
		if res == 0 {
			t.Fatalf("Crypto_aead_unlock accepted corrupted MAC for size=%d", size)
		}
		for i, b := range dec {
			if b != 0 {
				t.Fatalf("WIPE-ON-FAILURE violation on bad MAC: dec[%d] = 0x%02x (want 0x00) for size=%d", i, b, size)
			}
		}

		// 2. Déchiffrement avec Ciphertext altéré
		badCipher := make([]byte, size)
		copy(badCipher, cipher)
		badCipher[size/2] ^= 0xAA
		copy(dec, plain)

		res2 := Crypto_aead_unlock(dec, mac[:], key, nonce, nil, 0, badCipher, uint64(size))
		if res2 == 0 {
			t.Fatalf("Crypto_aead_unlock accepted corrupted cipher for size=%d", size)
		}
		for i, b := range dec {
			if b != 0 {
				t.Fatalf("WIPE-ON-FAILURE violation on bad cipher: dec[%d] = 0x%02x (want 0x00) for size=%d", i, b, size)
			}
		}
	}
}
