// SPDX-License-Identifier: Apache-2.0

// Modified in the lateos-ai/wal-g fork. Pure-Go crypto_secretstream_xchacha20poly1305.
// Bit-compatible with libsodium secretstream_xchacha20poly1305.c (manual
// ChaCha20-IETF + Poly1305 construction, not AEAD Seal of padded message).

package lsstream

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/binary"
	"fmt"
	"runtime"

	"github.com/hazyhaar/go-secretstream/internal/monocypher55"
	"golang.org/x/crypto/chacha20"
)

const (
	// KeyBytes is crypto_secretstream_xchacha20poly1305_KEYBYTES.
	KeyBytes = 32
	// HeaderBytes is crypto_secretstream_xchacha20poly1305_HEADERBYTES.
	HeaderBytes = 24
	// ABytes is crypto_secretstream_xchacha20poly1305_ABYTES (1 tag + 16 MAC).
	ABytes = 17

	// TagMessage is the default message tag.
	TagMessage = 0x00
	// TagPush marks end of a message set.
	TagPush = 0x01
	// TagRekey triggers rekey.
	TagRekey = 0x02
	// TagFinal marks the last chunk (includes REKEY bit).
	TagFinal = 0x03

	counterBytes = 4
	inonceBytes  = 8

	chunkSize = 8192
)

type streamState struct {
	k      [32]byte
	nonce  [12]byte // counter_le32 || inonce_8
	cipher *chacha20.Cipher
}

func initPush(key []byte, header []byte) (*streamState, error) {
	if len(key) != KeyBytes {
		return nil, fmt.Errorf("secretstream init_push: key length %d want %d", len(key), KeyBytes)
	}
	if len(header) != HeaderBytes {
		return nil, fmt.Errorf("secretstream init_push: header buffer %d want %d", len(header), HeaderBytes)
	}
	if _, err := rand.Read(header); err != nil {
		return nil, fmt.Errorf("secretstream init_push: rand header: %w", err)
	}
	return initFromHeader(key, header)
}

func initPull(key []byte, header []byte) (*streamState, error) {
	if len(key) != KeyBytes {
		return nil, fmt.Errorf("secretstream init_pull: key length %d want %d", len(key), KeyBytes)
	}
	if len(header) != HeaderBytes {
		return nil, fmt.Errorf("secretstream init_pull: header length %d want %d (truncated or corrupt stream)", len(header), HeaderBytes)
	}
	return initFromHeader(key, header)
}

func initFromHeader(key, header []byte) (*streamState, error) {
	sub, err := chacha20.HChaCha20(key[:KeyBytes], header[:16])
	if err != nil {
		return nil, fmt.Errorf("secretstream init: HChaCha20: %w", err)
	}
	st := &streamState{}
	copy(st.k[:], sub)
	st.counterReset()
	copy(st.nonce[counterBytes:], header[16:24])
	if err := st.initCipher(); err != nil {
		return nil, err
	}
	return st, nil
}

func (st *streamState) initCipher() error {
	c, err := chacha20.NewUnauthenticatedCipher(st.k[:], st.nonce[:])
	if err != nil {
		return err
	}
	st.cipher = c
	return nil
}

func (st *streamState) counterReset() {
	for i := 0; i < counterBytes; i++ {
		st.nonce[i] = 0
	}
	st.nonce[0] = 1
}

func (st *streamState) chacha(ic uint32) (*chacha20.Cipher, error) {
	if st.cipher == nil {
		if err := st.initCipher(); err != nil {
			return nil, err
		}
	}
	st.cipher.SetCounter(ic)
	return st.cipher, nil
}

func (st *streamState) advance(mac []byte) (wrapped bool, err error) {
	for i := 0; i < inonceBytes; i++ {
		st.nonce[counterBytes+i] ^= mac[i]
	}
	// sodium_increment on 4-byte counter
	for i := 0; i < counterBytes; i++ {
		st.nonce[i]++
		if st.nonce[i] != 0 {
			break
		}
	}
	if err := st.initCipher(); err != nil {
		return false, err
	}
	wrapped = true
	for i := 0; i < counterBytes; i++ {
		if st.nonce[i] != 0 {
			wrapped = false
			break
		}
	}
	return wrapped, nil
}

func (st *streamState) rekey() error {
	var msg [32 + inonceBytes]byte
	copy(msg[:32], st.k[:])
	copy(msg[32:], st.nonce[counterBytes:])
	c, err := st.chacha(0)
	if err != nil {
		memzero(msg[:])
		return fmt.Errorf("secretstream rekey: chacha: %w", err)
	}
	var out [32 + inonceBytes]byte
	c.XORKeyStream(out[:], msg[:])
	copy(st.k[:], out[:32])
	copy(st.nonce[counterBytes:], out[32:32+inonceBytes])
	st.counterReset()
	if err := st.initCipher(); err != nil {
		memzero(msg[:])
		memzero(out[:])
		return err
	}
	memzero(msg[:])
	memzero(out[:])
	return nil
}

// memzero is best-effort scrubbing of sensitive bytes.
// Uses runtime.KeepAlive to prevent compiler Dead Store Elimination (DSE).
func memzero(b []byte) {
	for i := range b {
		b[i] = 0
	}
	runtime.KeepAlive(b)
}

func (st *streamState) wipe() {
	memzero(st.k[:])
	memzero(st.nonce[:])
}

// sodiumPad matches libsodium `(0x10 - n) & 0xf` (two's complement).
func sodiumPad(n int) int {
	return (0x10 - n) & 0xf
}

// sodiumPadBlockMlen matches `(0x10 - sizeof(block) + mlen) & 0xf` with block=64.
func sodiumPadBlockMlen(mlen int) int {
	return (0x10 - 64 + mlen) & 0xf
}

func (st *streamState) push(m []byte, tag byte) ([]byte, error) {
	wire := make([]byte, 1+len(m)+16)
	n, err := st.pushTo(m, tag, wire)
	if err != nil {
		return nil, err
	}
	return wire[:n], nil
}

var staticZeros [64]byte

func (st *streamState) pushTo(m []byte, tag byte, wire []byte) (int, error) {
	wireLen := 1 + len(m) + 16
	if len(wire) < wireLen {
		return 0, fmt.Errorf("secretstream pushTo: wire buffer too small (%d < %d)", len(wire), wireLen)
	}

	// Single chacha cipher pass starting at counter 0
	c, err := st.chacha(0)
	if err != nil {
		return 0, err
	}
	var polyBlock [64]byte
	c.XORKeyStream(polyBlock[:], staticZeros[:64])
	var polyKey [32]byte
	copy(polyKey[:], polyBlock[:32])
	memzero(polyBlock[:])

	var polyCtx monocypher55.Crypto_poly1305_ctx
	monocypher55.Crypto_poly1305_init(&polyCtx, polyKey[:])
	if pad := sodiumPad(0); pad > 0 {
		var zeros [16]byte
		monocypher55.Crypto_poly1305_update(&polyCtx, zeros[:pad], uint64(pad))
	}

	// block = tag||zeros, xor with chacha (counter moves to 1); auth full 64-byte block
	var block [64]byte
	block[0] = tag
	c.XORKeyStream(block[:], block[:])
	monocypher55.Crypto_poly1305_update(&polyCtx, block[:], uint64(len(block)))
	out0 := block[0]

	// encrypt message with chacha (counter moves to 2..) directly into wire[1:1+len(m)]
	ciph := wire[1 : 1+len(m)]
	if len(m) > 0 {
		c.XORKeyStream(ciph, m)
		monocypher55.Crypto_poly1305_update(&polyCtx, ciph, uint64(len(ciph)))
	}
	if pad := sodiumPadBlockMlen(len(m)); pad > 0 {
		var zeros [16]byte
		monocypher55.Crypto_poly1305_update(&polyCtx, zeros[:pad], uint64(pad))
	}

	var slen [16]byte
	binary.LittleEndian.PutUint64(slen[0:8], 0)
	binary.LittleEndian.PutUint64(slen[8:16], uint64(64+len(m)))
	monocypher55.Crypto_poly1305_update(&polyCtx, slen[:], uint64(len(slen)))

	var tagOut [16]byte
	monocypher55.Crypto_poly1305_final(&polyCtx, tagOut[:])
	clear(polyCtx.C[:])
	clear(polyCtx.R[:])
	clear(polyCtx.Pad[:])
	clear(polyCtx.H[:])
	polyCtx.C_idx = 0
	runtime.KeepAlive(&polyCtx)
	memzero(polyKey[:])
	memzero(block[:])

	wire[0] = out0
	copy(wire[1+len(m):], tagOut[:])

	wrapped, err := st.advance(tagOut[:])
	if err != nil {
		return 0, err
	}
	if tag&TagRekey != 0 || wrapped {
		if err := st.rekey(); err != nil {
			return 0, err
		}
	}
	return wireLen, nil
}

func (st *streamState) pull(wire []byte) (m []byte, tag byte, err error) {
	if len(wire) < ABytes {
		return nil, 0, fmt.Errorf("secretstream pull: chunk too short (%d < ABytes=%d)", len(wire), ABytes)
	}
	mlen := len(wire) - ABytes
	m = make([]byte, mlen)
	n, tag, err := st.pullTo(wire, m)
	if err != nil {
		return nil, 0, err
	}
	return m[:n], tag, nil
}

func (st *streamState) pullTo(wire []byte, dst []byte) (n int, tag byte, err error) {
	if len(wire) < ABytes {
		return 0, 0, fmt.Errorf("secretstream pull: chunk too short (%d < ABytes=%d)", len(wire), ABytes)
	}
	mlen := len(wire) - ABytes
	if len(dst) < mlen {
		return 0, 0, fmt.Errorf("secretstream pullTo: dst buffer too small (%d < %d)", len(dst), mlen)
	}
	storedMAC := wire[1+mlen:]
	ciph := wire[1 : 1+mlen]

	// Single chacha cipher pass starting at counter 0
	c, err := st.chacha(0)
	if err != nil {
		return 0, 0, fmt.Errorf("secretstream pull: chacha ic0: %w", err)
	}
	var polyBlock [64]byte
	c.XORKeyStream(polyBlock[:], staticZeros[:64])
	var polyKey [32]byte
	copy(polyKey[:], polyBlock[:32])
	memzero(polyBlock[:])

	var polyCtx monocypher55.Crypto_poly1305_ctx
	monocypher55.Crypto_poly1305_init(&polyCtx, polyKey[:])
	if pad := sodiumPad(0); pad > 0 {
		var zeros [16]byte
		monocypher55.Crypto_poly1305_update(&polyCtx, zeros[:pad], uint64(pad))
	}

	var block [64]byte
	block[0] = wire[0]
	c.XORKeyStream(block[:], block[:])
	tag = block[0]
	block[0] = wire[0] // auth uses ciphertext byte, not decrypted tag — libsodium C order, do not change
	monocypher55.Crypto_poly1305_update(&polyCtx, block[:], uint64(len(block)))
	if mlen > 0 {
		monocypher55.Crypto_poly1305_update(&polyCtx, ciph, uint64(len(ciph)))
	}
	if pad := sodiumPadBlockMlen(mlen); pad > 0 {
		var zeros [16]byte
		monocypher55.Crypto_poly1305_update(&polyCtx, zeros[:pad], uint64(pad))
	}

	var slen [16]byte
	binary.LittleEndian.PutUint64(slen[0:8], 0)
	binary.LittleEndian.PutUint64(slen[8:16], uint64(64+mlen))
	monocypher55.Crypto_poly1305_update(&polyCtx, slen[:], uint64(len(slen)))

	var computed [16]byte
	monocypher55.Crypto_poly1305_final(&polyCtx, computed[:])
	clear(polyCtx.C[:])
	clear(polyCtx.R[:])
	clear(polyCtx.Pad[:])
	clear(polyCtx.H[:])
	polyCtx.C_idx = 0
	runtime.KeepAlive(&polyCtx)
	memzero(polyKey[:])
	memzero(block[:])
	if subtle.ConstantTimeCompare(computed[:], storedMAC) != 1 {
		return 0, 0, fmt.Errorf("secretstream pull: MAC mismatch (wire_len=%d plaintext_len=%d) — wrong key, truncated, or bitrot", len(wire), mlen)
	}

	if mlen > 0 {
		c.XORKeyStream(dst[:mlen], ciph)
	}

	// Advance only after MAC success so a failed pull leaves state coherent
	// for diagnostics; Reader always abandons on error.
	wrapped, err := st.advance(storedMAC)
	if err != nil {
		return 0, 0, err
	}
	if tag&TagRekey != 0 || wrapped {
		if err := st.rekey(); err != nil {
			return 0, 0, err
		}
	}
	return mlen, tag, nil
}
