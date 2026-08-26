// SPDX-License-Identifier: Apache-2.0 OR MIT

package engine

import (
	"errors"
	"fmt"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher_sgoiter"
)

var (
	errInexactOverlap = errors.New("engine sgoiter: inexact overlap between dst and payload")
	errOverlapAD      = errors.New("engine sgoiter: overlap between dst and ad")
	errOverlapMAC     = errors.New("engine sgoiter: overlap between dst and mac")
)

type sgoiterEngine struct{}

// Default returns the sgoiter-emitted monocypher AEAD backend.
func Default() AEAD { return sgoiterEngine{} }

func (sgoiterEngine) LockDst(dstCipher []byte, mac *[16]byte, key, nonce, ad, plain []byte) error {
	if len(dstCipher) != len(plain) {
		return fmt.Errorf("engine sgoiter: dst len %d != plain len %d", len(dstCipher), len(plain))
	}
	if err := checkAEADAlias(dstCipher, plain, ad, mac[:]); err != nil {
		return err
	}
	return sgoi.LockDst(dstCipher, mac[:], key, nonce, ad, plain)
}

func (sgoiterEngine) UnlockDst(dstPlain []byte, key, nonce, ad, cipher, mac []byte) ([]byte, error) {
	if len(dstPlain) < len(cipher) {
		return nil, fmt.Errorf("engine sgoiter: dst len %d < cipher len %d", len(dstPlain), len(cipher))
	}
	pt := dstPlain[:len(cipher)]
	if err := checkAEADAlias(pt, cipher, ad, mac); err != nil {
		return nil, err
	}
	if err := sgoi.UnlockDst(pt, key, nonce, mac, ad, cipher); err != nil {
		return nil, err
	}
	return pt, nil
}

func (sgoiterEngine) LockSubkeyDst(dstCipher []byte, mac *[16]byte, subkey, nonce12, ad, plain []byte) error {
	if len(dstCipher) != len(plain) {
		return fmt.Errorf("engine sgoiter: dst len %d != plain len %d", len(dstCipher), len(plain))
	}
	if err := checkAEADAlias(dstCipher, plain, ad, mac[:]); err != nil {
		return err
	}
	return sgoi.LockSubkeyDst(dstCipher, mac[:], subkey, nonce12, ad, plain)
}

func (sgoiterEngine) UnlockSubkeyDst(dstPlain []byte, subkey, nonce12, ad, cipher, mac []byte) ([]byte, error) {
	if len(dstPlain) < len(cipher) {
		return nil, fmt.Errorf("engine sgoiter: dst len %d < cipher len %d", len(dstPlain), len(cipher))
	}
	pt := dstPlain[:len(cipher)]
	if err := checkAEADAlias(pt, cipher, ad, mac); err != nil {
		return nil, err
	}
	if err := sgoi.UnlockSubkeyDst(pt, subkey, nonce12, mac, ad, cipher); err != nil {
		return nil, err
	}
	return pt, nil
}

func (sgoiterEngine) HChaCha20(out, key, in []byte) {
	// monocypher crypto_chacha20_h
	sgoi.Crypto_chacha20_h(out, key, in)
}
