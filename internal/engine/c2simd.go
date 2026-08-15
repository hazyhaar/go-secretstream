//go:build aead_c2simd

package engine

import "code.hazyhaar.fr/devhoros/c2simd"

type c2simdEngine struct{}

// Default returns the production AEAD backend (c2simd SIMD).
func Default() AEAD { return c2simdEngine{} }

func (c2simdEngine) LockDst(dstCipher []byte, mac *[16]byte, key, nonce, ad, plain []byte) error {
	_, err := c2simd.AEADLockDst(dstCipher, mac, key, nonce, ad, plain)
	return err
}

func (c2simdEngine) UnlockDst(dstPlain []byte, key, nonce, ad, cipher, mac []byte) ([]byte, error) {
	return c2simd.AEADUnlockDst(dstPlain, key, nonce, ad, cipher, mac)
}

func (c2simdEngine) HChaCha20(out, key, in []byte) {
	// c2simd.HChaCha20 signature is (key, nonce, out).
	c2simd.HChaCha20(key, in, out)
}
