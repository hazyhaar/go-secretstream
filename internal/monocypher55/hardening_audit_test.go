// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"bytes"
	"crypto/rand"
	"sync"
	"testing"
)

// TestBoundaryChaCha20DJB valide l'exactitude du chiffrement et du retour de compteur
// sur les frontières critiques : 0, 1, 63, 64, 65, 127, 128, 129, 255, 256, 257, 320, 512, 1024, 1025 octets.
func TestBoundaryChaCha20DJB(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 8)
	for i := range key {
		key[i] = byte(i + 1)
	}
	for i := range nonce {
		nonce[i] = byte(i + 0x42)
	}

	boundaries := []int{0, 1, 63, 64, 65, 127, 128, 129, 255, 256, 257, 320, 512, 1024, 1025, 4096}
	for _, size := range boundaries {
		pt := make([]byte, size)
		for i := range pt {
			pt[i] = byte((i * 17) & 0xff)
		}
		ct := make([]byte, size)
		dec := make([]byte, size)

		initialCtr := uint64(5)
		retCtr := Crypto_chacha20_djb(ct, pt, uint64(size), key, nonce, initialCtr)

		expectedBlocks := uint64((size + 63) / 64)
		expectedCtr := initialCtr + expectedBlocks
		if retCtr != expectedCtr {
			t.Fatalf("size=%d: retCtr=%d attendu=%d", size, retCtr, expectedCtr)
		}

		// Déchiffrement
		Crypto_chacha20_djb(dec, ct, uint64(size), key, nonce, initialCtr)
		if !bytes.Equal(pt, dec) {
			t.Fatalf("size=%d: échec du roundtrip ChaCha20", size)
		}
	}
}

// TestEd25519SignCheckAndMalleability valide le cycle signature/vérification Ed25519
// ainsi que le rejet strict des signatures corrompues et des points hors courbe.
func TestEd25519SignCheckAndMalleability(t *testing.T) {
	seed := make([]byte, 32)
	for i := range seed {
		seed[i] = byte(i + 0x55)
	}

	sk := make([]byte, 64)
	pk := make([]byte, 32)
	Crypto_eddsa_key_pair(sk, pk, seed)

	msg := []byte("message authentifié par monocypher55")
	sig := make([]byte, 64)
	Crypto_eddsa_sign(sig, sk, msg, uint64(len(msg)))

	// Vérification nominale
	if res := Crypto_eddsa_check(sig, pk, msg, uint64(len(msg))); res != 0 {
		t.Fatalf("échec de la vérification de signature nominale (res=%d)", res)
	}

	// Signature corrompue sur R
	corruptedSig := make([]byte, 64)
	copy(corruptedSig, sig)
	corruptedSig[0] ^= 0x01
	if res := Crypto_eddsa_check(corruptedSig, pk, msg, uint64(len(msg))); res == 0 {
		t.Fatalf("Crypto_eddsa_check a validé une signature avec R corrompu !")
	}

	// Signature corrompue sur S
	copy(corruptedSig, sig)
	corruptedSig[63] ^= 0x01
	if res := Crypto_eddsa_check(corruptedSig, pk, msg, uint64(len(msg))); res == 0 {
		t.Fatalf("Crypto_eddsa_check a validé une signature avec S corrompu !")
	}

	// Message altéré
	corruptedMsg := []byte("message authentifie par monocypher55")
	if res := Crypto_eddsa_check(sig, pk, corruptedMsg, uint64(len(corruptedMsg))); res == 0 {
		t.Fatalf("Crypto_eddsa_check a validé un message altéré !")
	}

	// Clé publique hors courbe (point non décodable sur Edwards25519)
	invalidPK := make([]byte, 32)
	for i := range invalidPK {
		invalidPK[i] = 0xFF
	}
	if res := Crypto_eddsa_check(sig, invalidPK, msg, uint64(len(msg))); res == 0 {
		t.Fatalf("Crypto_eddsa_check a validé une clé publique hors courbe !")
	}
}

// TestCryptoWipePoison vérifie qu'un buffer empoisonné est intégralement mis à zéro
func TestCryptoWipePoison(t *testing.T) {
	buf := make([]byte, 256)
	for i := range buf {
		buf[i] = 0xAA
	}
	Crypto_wipe(buf, uint64(len(buf)))
	for i, b := range buf {
		if b != 0 {
			t.Fatalf("Crypto_wipe a laissé l'octet %d non nul (0x%02X)", i, b)
		}
	}
}

// TestX25519DirtySmallConcurrency vérifie que Crypto_x25519_dirty_small
// ne mute aucun état global lors d'appels concurrents.
func TestX25519DirtySmallConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	sk := make([]byte, 32)
	for i := range sk {
		sk[i] = byte(i + 7)
	}

	expectedPk := make([]byte, 32)
	Crypto_x25519_dirty_small(expectedPk, sk)

	for worker := 0; worker < 16; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for iter := 0; iter < 100; iter++ {
				pk := make([]byte, 32)
				Crypto_x25519_dirty_small(pk, sk)
				if !bytes.Equal(pk, expectedPk) {
					t.Errorf("divergence de clé publique sous concurrence")
					return
				}
			}
		}()
	}
	wg.Wait()
}

// TestAEADLockDstCorruptedTagRejection vérifie le rejet systématique des tags altérés
func TestAEADLockDstCorruptedTagRejection(t *testing.T) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	ad := []byte("authentification additionnelle")
	pt := []byte("message secret ultra sensible zero leak")

	rand.Read(key)
	rand.Read(nonce)

	ct := make([]byte, len(pt))
	tag := make([]byte, 16)
	LockDst(ct, tag, key, nonce, ad, pt)

	// Déchiffrement nominal
	dec := make([]byte, len(pt))
	if err := UnlockDst(dec, key, nonce, tag, ad, ct); err != nil {
		t.Fatalf("échec du déchiffrement nominal: %v", err)
	}
	if !bytes.Equal(dec, pt) {
		t.Fatalf("altération du texte clair")
	}

	// Tag altéré
	corruptedTag := make([]byte, 16)
	copy(corruptedTag, tag)
	corruptedTag[15] ^= 0x01
	if err := UnlockDst(dec, key, nonce, corruptedTag, ad, ct); err == nil {
		t.Fatalf("UnlockDst a validé un tag corrompu !")
	}

	// AD altéré
	corruptedAD := []byte("authentification additionnelle altérée")
	if err := UnlockDst(dec, key, nonce, tag, corruptedAD, ct); err == nil {
		t.Fatalf("UnlockDst a validé des données associées altérées !")
	}
}
