// Package monocypher_sgoiter re-exports internal monocypher (compat path for engine tag aead_sgoiter).
package monocypher_sgoiter

import mono "github.com/hazyhaar/go-secretstream/internal/monocypher"

var ErrAEADCheckFailed = mono.ErrAEADCheckFailed

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
