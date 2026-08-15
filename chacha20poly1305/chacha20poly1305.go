// Package chacha20poly1305 provides an XChaCha20-Poly1305 AEAD
// (crypto/cipher.AEAD) implemented in pure Go, vectorized with the Go 1.27
// simd/archsimd intrinsics (GOEXPERIMENT=simd) and falling back to pure
// scalar Go elsewhere.
//
// The wire format is byte-for-byte identical to
// golang.org/x/crypto/chacha20poly1305.NewX (RFC 8439 with the XChaCha
// nonce extension); an interoperability test in this package enforces it.
//
// This is the reference implementation for the proposal in
// https://github.com/golang/go/issues/80881.
package chacha20poly1305

import (
	"crypto/cipher"
	"errors"

	mc "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

const (
	// KeySize is the size of the key used by this AEAD, in bytes.
	KeySize = 32
	// NonceSizeX is the size of the nonce used with the XChaCha20-Poly1305
	// variant, in bytes.
	NonceSizeX = 24
	// Overhead is the size of the Poly1305 authentication tag, in bytes.
	Overhead = 16
)

type xchacha struct {
	key [KeySize]byte
}

// NewX returns an XChaCha20-Poly1305 AEAD that uses the given 256-bit key.
func NewX(key []byte) (cipher.AEAD, error) {
	if len(key) != KeySize {
		return nil, errors.New("chacha20poly1305: bad key length")
	}
	var a xchacha
	copy(a.key[:], key)
	return &a, nil
}

func (*xchacha) NonceSize() int { return NonceSizeX }
func (*xchacha) Overhead() int  { return Overhead }

func (a *xchacha) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
	if len(nonce) != NonceSizeX {
		panic("chacha20poly1305: bad nonce length passed to Seal")
	}
	ret, out := sliceForAppend(dst, len(plaintext)+Overhead)
	var mac [16]byte
	if err := mc.LockDst(out[:len(plaintext)], mac[:], a.key[:], nonce, additionalData, plaintext); err != nil {
		panic("chacha20poly1305: " + err.Error())
	}
	copy(out[len(plaintext):], mac[:])
	return ret
}

func (a *xchacha) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	if len(nonce) != NonceSizeX {
		panic("chacha20poly1305: bad nonce length passed to Open")
	}
	if len(ciphertext) < Overhead {
		return nil, errors.New("chacha20poly1305: message authentication failed")
	}
	tag := ciphertext[len(ciphertext)-Overhead:]
	ct := ciphertext[:len(ciphertext)-Overhead]
	ret, out := sliceForAppend(dst, len(ct))
	if err := mc.UnlockDst(out, a.key[:], nonce, tag, additionalData, ct); err != nil {
		// mc.UnlockDst wipes out on failure; make sure nothing leaks.
		clear(out)
		return nil, errors.New("chacha20poly1305: message authentication failed")
	}
	return ret, nil
}

// sliceForAppend extends the input slice by n bytes (same helper shape as in
// golang.org/x/crypto).
func sliceForAppend(in []byte, n int) (head, tail []byte) {
	if total := len(in) + n; cap(in) >= total {
		head = in[:total]
	} else {
		head = make([]byte, total)
		copy(head, in)
	}
	tail = head[len(in):]
	return
}
