package monocypher

import (
	"bytes"
	"testing"
)

func FuzzChaCha20DJB(f *testing.F) {
	f.Add([]byte("Hello, World! Fuzzing ChaCha20 SIMD kernel."), []byte("01234567890123456789012345678901"), []byte("nonce123"), uint64(0))
	f.Add(make([]byte, 256), make([]byte, 32), make([]byte, 8), uint64(42))
	f.Add(make([]byte, 1025), make([]byte, 32), make([]byte, 8), uint64(100))

	f.Fuzz(func(t *testing.T, pt, key, nonce []byte, ctr uint64) {
		if len(key) != 32 || len(nonce) != 8 {
			return
		}
		ct := make([]byte, len(pt))
		dec := make([]byte, len(pt))

		retCtr := Crypto_chacha20_djb(ct, pt, uint64(len(pt)), key, nonce, ctr)
		expectedBlocks := uint64((len(pt) + 63) / 64)
		if retCtr != ctr+expectedBlocks {
			t.Fatalf("retCtr=%d attendu=%d", retCtr, ctr+expectedBlocks)
		}

		Crypto_chacha20_djb(dec, ct, uint64(len(pt)), key, nonce, ctr)
		if !bytes.Equal(pt, dec) {
			t.Fatalf("déchiffrement différent du texte clair")
		}
	})
}

func FuzzAEADLockUnlock(f *testing.F) {
	key := make([]byte, 32)
	nonce := make([]byte, 24)
	f.Add(key, nonce, []byte("ad-data"), []byte("secret payload 123"))
	f.Add(key, nonce, []byte(""), []byte(""))

	f.Fuzz(func(t *testing.T, key, nonce, ad, pt []byte) {
		if len(key) != 32 || len(nonce) != 24 {
			return
		}
		ct := make([]byte, len(pt))
		tag := make([]byte, 16)

		if err := LockDst(ct, tag, key, nonce, ad, pt); err != nil {
			t.Fatalf("LockDst failed: %v", err)
		}

		dec := make([]byte, len(pt))
		if err := UnlockDst(dec, key, nonce, tag, ad, ct); err != nil {
			t.Fatalf("UnlockDst failed: %v", err)
		}

		if !bytes.Equal(dec, pt) {
			t.Fatalf("mismatch decrypted vs plainText")
		}
	})
}
