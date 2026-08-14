package engine

import (
	"fmt"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
)

type AEAD interface {
	LockDst(dstCipher []byte, mac *[16]byte, key, nonce, ad, plain []byte) error
	UnlockDst(dstPlain []byte, key, nonce, ad, cipher, mac []byte) ([]byte, error)
	HChaCha20(out, key, in []byte)
}

type monocypherEngine struct{}

func Default() AEAD { return monocypherEngine{} }

func (monocypherEngine) LockDst(dstCipher []byte, mac *[16]byte, key, nonce, ad, plain []byte) error {
	if len(dstCipher) != len(plain) {
		return fmt.Errorf("engine: dst len %d != plain len %d", len(dstCipher), len(plain))
	}
	return sgoi.LockDst(dstCipher, mac[:], key, nonce, ad, plain)
}

func (monocypherEngine) UnlockDst(dstPlain []byte, key, nonce, ad, cipher, mac []byte) ([]byte, error) {
	if len(dstPlain) < len(cipher) {
		return nil, fmt.Errorf("engine: dst len %d < cipher len %d", len(dstPlain), len(cipher))
	}
	pt := dstPlain[:len(cipher)]
	if err := sgoi.UnlockDst(pt, key, nonce, mac, ad, cipher); err != nil {
		return nil, err
	}
	return pt, nil
}

func (monocypherEngine) HChaCha20(out, key, in []byte) {
	sgoi.Crypto_chacha20_h(out, key, in)
}
