// SPDX-License-Identifier: Apache-2.0 OR MIT

// Package monocypher_sgoiter re-exports monocypher55 (compat path for engine tag aead_sgoiter).
package monocypher_sgoiter

import (
	"errors"
	"unsafe"

	mono "github.com/hazyhaar/go-secretstream/internal/monocypher55"
)

var ErrAEADCheckFailed = mono.ErrAEADCheckFailed

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

func AEADLock(key, nonce, ad, plain []byte) (cipher, mac []byte, err error) {
	return mono.AEADLock(key, nonce, ad, plain)
}

func AEADUnlock(key, nonce, mac, ad, cipher []byte) (plain []byte, err error) {
	return mono.AEADUnlock(key, nonce, mac, ad, cipher)
}

func LockDst(dstCT, mac, key, nonce, ad, plain []byte) error {
	return mono.LockDst(dstCT, mac, key, nonce, ad, plain)
}

func UnlockDst(dstPT, key, nonce, mac, ad, cipher []byte) error {
	return mono.UnlockDst(dstPT, key, nonce, mac, ad, cipher)
}

func Crypto_chacha20_h(out, key, in []byte) {
	mono.Crypto_chacha20_h(out, key, in)
}

// LockSubkeyDst seals with a pre-derived HChaCha20 subkey and a 12-byte IETF
// nonce (monocypher crypto_aead_init_ietf + crypto_aead_write). With nonce12 =
// 0x00000000 || nonce24[16:24] and subkey = crypto_chacha20_h(key, nonce24[0:16]),
// the output is byte-identical to LockDst(key, nonce24).
func LockSubkeyDst(dstCT, mac, subkey, nonce12, ad, plain []byte) error {
	if len(subkey) != 32 {
		return errors.New("monocypher_sgoiter: subkey must be 32 bytes")
	}
	if len(nonce12) != 12 {
		return errors.New("monocypher_sgoiter: IETF nonce must be 12 bytes")
	}
	if len(mac) < 16 {
		return errors.New("monocypher_sgoiter: mac must be at least 16 bytes")
	}
	if len(dstCT) < len(plain) {
		return errors.New("monocypher_sgoiter: dstCT shorter than plaintext")
	}
	if inexactOverlap(dstCT[:len(plain)], plain) {
		panic("monocypher_sgoiter: invalid buffer overlap between dstCT and plain")
	}
	if anyOverlap(dstCT[:len(plain)], ad) {
		panic("monocypher_sgoiter: invalid buffer overlap between dstCT and ad")
	}
	var ctx mono.Crypto_aead_ctx
	mono.Crypto_aead_init_ietf(&ctx, subkey, nonce12)
	adPtr, ptPtr := aeadPtrs(ad, plain)
	mono.Crypto_aead_write(&ctx, dstCT[:len(plain)], mac[:16], adPtr, uint64(len(ad)), ptPtr, uint64(len(plain)))
	return nil
}

// UnlockSubkeyDst verifies then decrypts with a pre-derived subkey (fail-closed).
func UnlockSubkeyDst(dstPT, subkey, nonce12, mac, ad, cipher []byte) error {
	if len(subkey) != 32 {
		return errors.New("monocypher_sgoiter: subkey must be 32 bytes")
	}
	if len(nonce12) != 12 {
		return errors.New("monocypher_sgoiter: IETF nonce must be 12 bytes")
	}
	if len(mac) < 16 {
		return errors.New("monocypher_sgoiter: mac must be at least 16 bytes")
	}
	if len(dstPT) < len(cipher) {
		return errors.New("monocypher_sgoiter: dstPT shorter than ciphertext")
	}
	if inexactOverlap(dstPT[:len(cipher)], cipher) {
		panic("monocypher_sgoiter: invalid buffer overlap between dstPT and cipher")
	}
	if anyOverlap(dstPT[:len(cipher)], ad) {
		panic("monocypher_sgoiter: invalid buffer overlap between dstPT and ad")
	}
	var ctx mono.Crypto_aead_ctx
	mono.Crypto_aead_init_ietf(&ctx, subkey, nonce12)
	adPtr, ctPtr := aeadPtrs(ad, cipher)
	if mono.Crypto_aead_read(&ctx, dstPT[:len(cipher)], mac[:16], adPtr, uint64(len(ad)), ctPtr, uint64(len(cipher))) != 0 {
		return ErrAEADCheckFailed
	}
	return nil
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
