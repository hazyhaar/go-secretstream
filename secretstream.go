// Pure-Go crypto_secretstream_xchacha20poly1305.
// Bit-compatible with libsodium secretstream_xchacha20poly1305.c (manual
// ChaCha20-IETF + Poly1305 construction, not AEAD Seal of padded message).

package secretstream

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/poly1305"
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

	// ChunkSize is the plaintext framing size used by WAL-G (and this package's
	// Reader/Writer). Core push/pull accept any length.
	ChunkSize = 8192
)

type streamState struct {
	k     [32]byte
	nonce [12]byte // counter_le32 || inonce_8
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
		return nil, err
	}
	st := &streamState{}
	copy(st.k[:], sub)
	st.counterReset()
	copy(st.nonce[counterBytes:], header[16:24])
	return st, nil
}

func (st *streamState) counterReset() {
	for i := 0; i < counterBytes; i++ {
		st.nonce[i] = 0
	}
	st.nonce[0] = 1
}

func (st *streamState) chacha(ic uint32) (*chacha20.Cipher, error) {
	c, err := chacha20.NewUnauthenticatedCipher(st.k[:], st.nonce[:])
	if err != nil {
		return nil, err
	}
	if ic != 0 {
		c.SetCounter(ic)
	}
	return c, nil
}

func (st *streamState) advance(mac []byte) {
	for i := 0; i < inonceBytes; i++ {
		st.nonce[counterBytes+i] ^= mac[i]
	}
	for i := 0; i < counterBytes; i++ {
		st.nonce[i]++
		if st.nonce[i] != 0 {
			break
		}
	}
	zero := true
	for i := 0; i < counterBytes; i++ {
		if st.nonce[i] != 0 {
			zero = false
			break
		}
	}
	if zero {
		st.rekey()
	}
}

func (st *streamState) rekey() {
	msg := make([]byte, 32+inonceBytes)
	copy(msg[:32], st.k[:])
	copy(msg[32:], st.nonce[counterBytes:])
	c, err := st.chacha(0)
	if err != nil {
		return
	}
	out := make([]byte, len(msg))
	c.XORKeyStream(out, msg)
	copy(st.k[:], out[:32])
	copy(st.nonce[counterBytes:], out[32:32+inonceBytes])
	st.counterReset()
}

func sodiumPad(n int) int {
	return (0x10 - n) & 0xf
}

func sodiumPadBlockMlen(mlen int) int {
	return (0x10 - 64 + mlen) & 0xf
}

func (st *streamState) push(m []byte, tag byte) ([]byte, error) {
	c0, err := st.chacha(0)
	if err != nil {
		return nil, err
	}
	var block0 [64]byte
	c0.XORKeyStream(block0[:], block0[:])
	var polyKey [32]byte
	copy(polyKey[:], block0[:32])

	mac := poly1305.New(&polyKey)
	var zeros [16]byte
	_, _ = mac.Write(zeros[:sodiumPad(0)])

	var block [64]byte
	block[0] = tag
	c1, err := st.chacha(1)
	if err != nil {
		return nil, err
	}
	c1.XORKeyStream(block[:], block[:])
	_, _ = mac.Write(block[:])
	out0 := block[0]

	ciph := make([]byte, len(m))
	if len(m) > 0 {
		c2, err := st.chacha(2)
		if err != nil {
			return nil, err
		}
		c2.XORKeyStream(ciph, m)
	}
	_, _ = mac.Write(ciph)
	_, _ = mac.Write(zeros[:sodiumPadBlockMlen(len(m))])

	var slen [8]byte
	binary.LittleEndian.PutUint64(slen[:], 0)
	_, _ = mac.Write(slen[:])
	binary.LittleEndian.PutUint64(slen[:], uint64(64+len(m)))
	_, _ = mac.Write(slen[:])

	var tagOut [16]byte
	mac.Sum(tagOut[:0])

	wire := make([]byte, 1+len(m)+16)
	wire[0] = out0
	copy(wire[1:1+len(m)], ciph)
	copy(wire[1+len(m):], tagOut[:])

	st.advance(tagOut[:])
	if tag&TagRekey != 0 {
		st.rekey()
	}
	return wire, nil
}

func (st *streamState) pull(wire []byte) (m []byte, tag byte, err error) {
	if len(wire) < ABytes {
		return nil, 0, fmt.Errorf("secretstream pull: chunk too short (%d < ABytes=%d)", len(wire), ABytes)
	}
	mlen := len(wire) - ABytes
	storedMAC := wire[1+mlen:]
	ciph := wire[1 : 1+mlen]

	c0, err := st.chacha(0)
	if err != nil {
		return nil, 0, fmt.Errorf("secretstream pull: chacha ic0: %w", err)
	}
	var block0 [64]byte
	c0.XORKeyStream(block0[:], block0[:])
	var polyKey [32]byte
	copy(polyKey[:], block0[:32])

	mac := poly1305.New(&polyKey)
	var zeros [16]byte
	_, _ = mac.Write(zeros[:sodiumPad(0)])

	var block [64]byte
	block[0] = wire[0]
	c1, err := st.chacha(1)
	if err != nil {
		return nil, 0, fmt.Errorf("secretstream pull: chacha ic1: %w", err)
	}
	c1.XORKeyStream(block[:], block[:])
	tag = block[0]
	block[0] = wire[0]
	_, _ = mac.Write(block[:])
	_, _ = mac.Write(ciph)
	_, _ = mac.Write(zeros[:sodiumPadBlockMlen(mlen)])

	var slen [8]byte
	binary.LittleEndian.PutUint64(slen[:], 0)
	_, _ = mac.Write(slen[:])
	binary.LittleEndian.PutUint64(slen[:], uint64(64+mlen))
	_, _ = mac.Write(slen[:])

	var computed [16]byte
	mac.Sum(computed[:0])
	if subtle.ConstantTimeCompare(computed[:], storedMAC) != 1 {
		return nil, 0, fmt.Errorf("secretstream pull: MAC mismatch (wire_len=%d plaintext_len=%d) — wrong key, truncated, or bitrot", len(wire), mlen)
	}

	m = make([]byte, mlen)
	if mlen > 0 {
		c2, err := st.chacha(2)
		if err != nil {
			return nil, 0, fmt.Errorf("secretstream pull: chacha ic2: %w", err)
		}
		c2.XORKeyStream(m, ciph)
	}

	st.advance(storedMAC)
	if tag&TagRekey != 0 {
		st.rekey()
	}
	return m, tag, nil
}
