//go:build aead_c2simd && !goexperiment.simd

package engine

import (
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

// Fallback scalaire (build sans GOEXPERIMENT=simd) : ChaCha20-IETF-Poly1305
// standard sur la sous-clé pré-dérivée. Même fil que le chemin SIMD.
func (c2simdEngine) LockSubkeyDst(dstCipher []byte, mac *[16]byte, subkey, nonce12, ad, plain []byte) error {
	aead, err := chacha20poly1305.New(subkey)
	if err != nil {
		return err
	}
	if len(dstCipher) < len(plain) {
		return fmt.Errorf("engine: dstCipher shorter than plaintext")
	}
	sealed := aead.Seal(nil, nonce12, plain, ad)
	copy(dstCipher[:len(plain)], sealed[:len(plain)])
	copy(mac[:], sealed[len(plain):])
	return nil
}

func (c2simdEngine) UnlockSubkeyDst(dstPlain []byte, subkey, nonce12, ad, cipher, mac []byte) ([]byte, error) {
	aead, err := chacha20poly1305.New(subkey)
	if err != nil {
		return nil, err
	}
	if len(dstPlain) < len(cipher) {
		return nil, fmt.Errorf("engine: dstPlain shorter than ciphertext")
	}
	sealed := make([]byte, 0, len(cipher)+16)
	sealed = append(sealed, cipher...)
	sealed = append(sealed, mac...)
	pt, err := aead.Open(dstPlain[:0], nonce12, sealed, ad)
	if err != nil {
		return nil, err
	}
	return pt, nil
}
