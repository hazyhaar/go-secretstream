// Package engine abstracts AEAD backends for the maison (non-libsodium-wire) stream.
package engine

// AEAD is the one-shot AEAD used by secretstream55 standard framing (not libsodium wire).
type AEAD interface {
	LockDst(dstCipher []byte, mac *[16]byte, key, nonce, ad, plain []byte) error
	UnlockDst(dstPlain []byte, key, nonce, ad, cipher, mac []byte) ([]byte, error)
	HChaCha20(out, key, in []byte)
}
