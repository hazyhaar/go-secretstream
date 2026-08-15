//go:build aead_c2simd && goexperiment.simd

package engine

import "code.hazyhaar.fr/devhoros/c2simd"

func (c2simdEngine) LockSubkeyDst(dstCipher []byte, mac *[16]byte, subkey, nonce12, ad, plain []byte) error {
	_, err := c2simd.AEADSubkeyLockDst(dstCipher, mac, subkey, nonce12, ad, plain)
	return err
}

func (c2simdEngine) UnlockSubkeyDst(dstPlain []byte, subkey, nonce12, ad, cipher, mac []byte) ([]byte, error) {
	return c2simd.AEADSubkeyUnlockDst(dstPlain, subkey, nonce12, ad, cipher, mac)
}
