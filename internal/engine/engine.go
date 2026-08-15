// Package engine abstracts AEAD backends for the maison (non-libsodium-wire) stream.
package engine

// AEAD is the one-shot AEAD used by secretstream55 standard framing (not libsodium wire).
type AEAD interface {
	LockDst(dstCipher []byte, mac *[16]byte, key, nonce, ad, plain []byte) error
	UnlockDst(dstPlain []byte, key, nonce, ad, cipher, mac []byte) ([]byte, error)
	HChaCha20(out, key, in []byte)

	// Subkey path: the caller derives the XChaCha20 subkey once via HChaCha20
	// (key, nonce[0:16]) and then seals/opens each chunk with the 12-byte IETF
	// nonce (4 zero bytes || nonce[16:24]). Byte-identical to the full XChaCha
	// path, but skips the per-chunk HChaCha20 derivation.
	LockSubkeyDst(dstCipher []byte, mac *[16]byte, subkey, nonce12, ad, plain []byte) error
	UnlockSubkeyDst(dstPlain []byte, subkey, nonce12, ad, cipher, mac []byte) ([]byte, error)
}
