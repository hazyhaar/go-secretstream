package monocypher

import (
	"errors"
	"unsafe"

	"modernc.org/libc"
)

var (
	ErrAEADCheckFailed = errors.New("monocypher: AEAD authentication check failed")
)

// AEADLock encrypts plainText and produces cipherText and mac using XChaCha20-Poly1305 in Pure Go.
func AEADLock(key []byte, nonce []byte, ad []byte, plainText []byte) (cipherText []byte, mac []byte, err error) {
	if len(key) != 32 {
		return nil, nil, errors.New("monocypher: key must be 32 bytes")
	}
	if len(nonce) != 24 {
		return nil, nil, errors.New("monocypher: XChaCha20 nonce must be 24 bytes")
	}

	tls := libc.NewTLS()
	defer tls.Close()

	cKey, _ := libc.CString(string(key))
	defer libc.Xfree(tls, cKey)

	cNonce, _ := libc.CString(string(nonce))
	defer libc.Xfree(tls, cNonce)

	var cAD uintptr
	if len(ad) > 0 {
		cAD, _ = libc.CString(string(ad))
		defer libc.Xfree(tls, cAD)
	}

	var cPlain uintptr
	if len(plainText) > 0 {
		cPlain, _ = libc.CString(string(plainText))
		defer libc.Xfree(tls, cPlain)
	}

	mac = make([]byte, 16)
	cMAC, _ := libc.CString(string(mac))
	defer libc.Xfree(tls, cMAC)

	cipherText = make([]byte, len(plainText))
	var cCipher uintptr
	if len(cipherText) > 0 {
		cCipher, _ = libc.CString(string(cipherText))
		defer libc.Xfree(tls, cCipher)
	}

	crypto_aead_lock(
		tls,
		cCipher,
		cMAC,
		cKey,
		cNonce,
		cAD,
		size_t(len(ad)),
		cPlain,
		size_t(len(plainText)),
	)

	if len(cipherText) > 0 {
		copy(cipherText, unsafe.Slice((*byte)(unsafe.Pointer(cCipher)), len(cipherText)))
	}
	copy(mac, unsafe.Slice((*byte)(unsafe.Pointer(cMAC)), 16))

	return cipherText, mac, nil
}

// AEADUnlock decrypts cipherText and verifies mac using XChaCha20-Poly1305 in Pure Go.
func AEADUnlock(key []byte, nonce []byte, mac []byte, ad []byte, cipherText []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("monocypher: key must be 32 bytes")
	}
	if len(nonce) != 24 {
		return nil, errors.New("monocypher: XChaCha20 nonce must be 24 bytes")
	}
	if len(mac) != 16 {
		return nil, errors.New("monocypher: mac must be 16 bytes")
	}

	tls := libc.NewTLS()
	defer tls.Close()

	cKey, _ := libc.CString(string(key))
	defer libc.Xfree(tls, cKey)

	cNonce, _ := libc.CString(string(nonce))
	defer libc.Xfree(tls, cNonce)

	cMAC, _ := libc.CString(string(mac))
	defer libc.Xfree(tls, cMAC)

	var cAD uintptr
	if len(ad) > 0 {
		cAD, _ = libc.CString(string(ad))
		defer libc.Xfree(tls, cAD)
	}

	var cCipher uintptr
	if len(cipherText) > 0 {
		cCipher, _ = libc.CString(string(cipherText))
		defer libc.Xfree(tls, cCipher)
	}

	plainText := make([]byte, len(cipherText))
	var cPlain uintptr
	if len(plainText) > 0 {
		cPlain, _ = libc.CString(string(plainText))
		defer libc.Xfree(tls, cPlain)
	}

	res := crypto_aead_unlock(
		tls,
		cPlain,
		cMAC,
		cKey,
		cNonce,
		cAD,
		size_t(len(ad)),
		cCipher,
		size_t(len(cipherText)),
	)

	if res != 0 {
		return nil, ErrAEADCheckFailed
	}

	if len(plainText) > 0 {
		copy(plainText, unsafe.Slice((*byte)(unsafe.Pointer(cPlain)), len(plainText)))
	}

	return plainText, nil
}
