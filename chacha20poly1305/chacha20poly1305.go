// Package chacha20poly1305 implements the ChaCha20-Poly1305 and XChaCha20-Poly1305
// Authenticated Encryption with Associated Data (AEAD) algorithms as specified in RFC 8439.
//
// It is accelerated by Go 1.27 SIMD (AVX2 256-bit) in 100% Pure Go with zero assembly.
package chacha20poly1305

import (
	"crypto/cipher"
	"errors"

	"github.com/hazyhaar/go-secretstream/internal/monocypher"
)

const (
	// KeySize is the size of the key in bytes.
	KeySize = 32

	// NonceSize is the size of the standard RFC 8439 nonce in bytes.
	NonceSize = 12

	// NonceSizeX is the size of the extended XChaCha20-Poly1305 nonce in bytes.
	NonceSizeX = 24

	// Overhead is the size of the authentication tag in bytes.
	Overhead = 16
)

var (
	errInvalidKey   = errors.New("chacha20poly1305: bad key length")
	errInvalidNonce = errors.New("chacha20poly1305: bad nonce length")
	errOpenFailed   = errors.New("chacha20poly1305: message authentication failed")
)

type chacha20poly1305AEAD struct {
	key       [KeySize]byte
	nonceSize int
}

// New returns a standard ChaCha20-Poly1305 AEAD (RFC 8439 with 12-byte nonce).
func New(key []byte) (cipher.AEAD, error) {
	if len(key) != KeySize {
		return nil, errInvalidKey
	}
	var a chacha20poly1305AEAD
	copy(a.key[:], key)
	a.nonceSize = NonceSize
	return &a, nil
}

// NewX returns an XChaCha20-Poly1305 AEAD (24-byte nonce).
func NewX(key []byte) (cipher.AEAD, error) {
	if len(key) != KeySize {
		return nil, errInvalidKey
	}
	var a chacha20poly1305AEAD
	copy(a.key[:], key)
	a.nonceSize = NonceSizeX
	return &a, nil
}

func (a *chacha20poly1305AEAD) NonceSize() int {
	return a.nonceSize
}

func (a *chacha20poly1305AEAD) Overhead() int {
	return Overhead
}

// Seal encrypts and authenticates plaintext, authenticates additionalData and appends the result to dst.
func (a *chacha20poly1305AEAD) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
	if len(nonce) != a.nonceSize {
		panic(errInvalidNonce)
	}

	ret, out := sliceForAppend(dst, len(plaintext)+Overhead)
	ct := out[:len(plaintext)]
	mac := out[len(plaintext):]

	var ctx monocypher.Crypto_aead_ctx
	if a.nonceSize == NonceSize {
		monocypher.Crypto_aead_init_ietf(&ctx, a.key[:], nonce)
	} else {
		monocypher.Crypto_aead_init_x(&ctx, a.key[:], nonce)
	}

	var adPtr, ptPtr []byte
	if len(additionalData) > 0 {
		adPtr = additionalData
	}
	if len(plaintext) > 0 {
		ptPtr = plaintext
	}

	monocypher.Crypto_aead_write(&ctx, ct, mac, adPtr, uint64(len(additionalData)), ptPtr, uint64(len(plaintext)))
	return ret
}

// Open decrypts and authenticates ciphertext, authenticates additionalData and appends the result to dst.
func (a *chacha20poly1305AEAD) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	if len(nonce) != a.nonceSize {
		return nil, errInvalidNonce
	}
	if len(ciphertext) < Overhead {
		return nil, errOpenFailed
	}

	plainLen := len(ciphertext) - Overhead
	ct := ciphertext[:plainLen]
	mac := ciphertext[plainLen:]

	ret, pt := sliceForAppend(dst, plainLen)

	var ctx monocypher.Crypto_aead_ctx
	if a.nonceSize == NonceSize {
		monocypher.Crypto_aead_init_ietf(&ctx, a.key[:], nonce)
	} else {
		monocypher.Crypto_aead_init_x(&ctx, a.key[:], nonce)
	}

	var adPtr, ctPtr []byte
	if len(additionalData) > 0 {
		adPtr = additionalData
	}
	if plainLen > 0 {
		ctPtr = ct
	}

	res := monocypher.Crypto_aead_read(&ctx, pt, mac, adPtr, uint64(len(additionalData)), ctPtr, uint64(plainLen))
	if res != 0 {
		return nil, errOpenFailed
	}
	return ret, nil
}

func sliceForAppend(in []byte, n int) (head, tail []byte) {
	if total := len(in) + n; cap(in) >= total {
		head = in[:total]
		tail = head[len(in):]
		return
	}
	head = make([]byte, len(in)+n)
	copy(head, in)
	tail = head[len(in):]
	return
}
