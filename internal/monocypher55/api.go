// SPDX-License-Identifier: Apache-2.0 OR MIT

// Package monocypher55 — Monocypher 4.0.2 via sgoiter (+ hands), CGO=0.
// API AEAD alignée pour secretstream55 (tag aead_sgoiter).
package monocypher55

import (
	"errors"
	"unsafe"
)

var ErrAEADCheckFailed = errors.New("monocypher55: AEAD authentication check failed")

func anyOverlap(x, y []byte) bool {
	return len(x) > 0 && len(y) > 0 &&
		uintptr(unsafe.Pointer(&x[0])) <= uintptr(unsafe.Pointer(&y[len(y)-1])) &&
		uintptr(unsafe.Pointer(&y[0])) <= uintptr(unsafe.Pointer(&x[len(x)-1]))
}

func inexactOverlap(x, y []byte) bool {
	if len(x) == 0 || len(y) == 0 {
		return false
	}
	if &x[0] == &y[0] {
		return len(x) != len(y)
	}
	return anyOverlap(x, y)
}

func aeadPtrs(ad, payload []byte) (adPtr, payloadPtr []byte) {
	if len(ad) > 0 {
		adPtr = ad
	}
	if len(payload) > 0 {
		payloadPtr = payload
	}
	return adPtr, payloadPtr
}

func checkKeyNonce(key, nonce []byte) error {
	if len(key) != 32 {
		return errors.New("monocypher55: key must be 32 bytes")
	}
	if len(nonce) != 24 {
		return errors.New("monocypher55: XChaCha20 nonce must be 24 bytes")
	}
	return nil
}

// LockDst writes XChaCha20-Poly1305 ciphertext into dstCT and the 16-byte tag into mac.
// dstCT must have len >= len(plainText); mac must have len >= 16. Zero heap allocations.
// nil or empty ad / plainText are accepted (C NULL + size 0).
func LockDst(dstCT, mac, key, nonce, ad, plainText []byte) error {
	if err := checkKeyNonce(key, nonce); err != nil {
		return err
	}
	if len(mac) < 16 {
		return errors.New("monocypher55: mac must be at least 16 bytes")
	}
	if len(dstCT) < len(plainText) {
		return errors.New("monocypher55: dstCT shorter than plaintext")
	}
	if inexactOverlap(dstCT[:len(plainText)], plainText) {
		panic("monocypher55: invalid buffer overlap between dstCT and plainText")
	}
	if anyOverlap(dstCT[:len(plainText)], ad) {
		panic("monocypher55: invalid buffer overlap between dstCT and ad")
	}
	adPtr, ptPtr := aeadPtrs(ad, plainText)
	Crypto_aead_lock(dstCT[:len(plainText)], mac[:16], key, nonce, adPtr, uint64(len(ad)), ptPtr, uint64(len(plainText)))
	return nil
}

// UnlockDst verifies then decrypts into dstPT. dstPT must have len >= len(cipherText).
// Zero heap allocations. nil or empty ad / cipherText are accepted.
func UnlockDst(dstPT, key, nonce, mac, ad, cipherText []byte) error {
	if err := checkKeyNonce(key, nonce); err != nil {
		return err
	}
	if len(mac) < 16 {
		return errors.New("monocypher55: mac must be at least 16 bytes")
	}
	if len(dstPT) < len(cipherText) {
		return errors.New("monocypher55: dstPT shorter than ciphertext")
	}
	if inexactOverlap(dstPT[:len(cipherText)], cipherText) {
		panic("monocypher55: invalid buffer overlap between dstPT and cipherText")
	}
	if anyOverlap(dstPT[:len(cipherText)], ad) {
		panic("monocypher55: invalid buffer overlap between dstPT and ad")
	}
	adPtr, ctPtr := aeadPtrs(ad, cipherText)
	res := Crypto_aead_unlock(dstPT[:len(cipherText)], mac[:16], key, nonce, adPtr, uint64(len(ad)), ctPtr, uint64(len(cipherText)))
	if res != 0 {
		return ErrAEADCheckFailed
	}
	return nil
}

// AEADLock XChaCha20-Poly1305 (crypto_aead_lock monocypher). Allocates ct + mac.
func AEADLock(key, nonce, ad, plainText []byte) (cipherText, mac []byte, err error) {
	mac = make([]byte, 16)
	cipherText = make([]byte, len(plainText))
	err = LockDst(cipherText, mac, key, nonce, ad, plainText)
	if err != nil {
		return nil, nil, err
	}
	return cipherText, mac, nil
}

// AEADUnlock verify+decrypt (crypto_aead_unlock). Allocates plaintext.
func AEADUnlock(key, nonce, mac, ad, cipherText []byte) ([]byte, error) {
	plain := make([]byte, len(cipherText))
	if err := UnlockDst(plain, key, nonce, mac, ad, cipherText); err != nil {
		return nil, err
	}
	return plain, nil
}
