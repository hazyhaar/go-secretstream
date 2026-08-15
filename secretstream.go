// Package secretstream55 provides streaming AEAD encryption in Pure Go.
//
// Two wire modes:
//   - Standard (NewEncryptor): maison framing (BE4 length + XChaCha20-Poly1305 AEAD via engine).
//   - Libsodium (NewLibsodiumEncryptor): crypto_secretstream_xchacha20poly1305 wire (wal-g compatible).
//
// AEAD backend for standard mode: monocypher55 (default — bascule 2026-08-15,
// gate à trois preuves : fil croisé bit-identique, aliasing, rejet de forge)
// ou c2simd (-tags aead_c2simd).
package secretstream55

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/hazyhaar/go-secretstream/internal/engine"
	"github.com/hazyhaar/go-secretstream/internal/lsstream"
)

const (
	HeaderSize = 24
	TagSize    = 16
	ChunkSize  = 64 * 1024

	TagMessage byte = 0x00
	TagPush    byte = 0x01
	TagRekey   byte = 0x02
	TagFinal   byte = 0x03
)

// Encryptor — standard (maison) framing + engine AEAD.
type Encryptor struct {
	w              io.Writer
	key            [32]byte
	nonce          [24]byte
	subkey         [32]byte // HChaCha20(key, nonce[0:16]) — dérivée une fois par stream
	seq            uint64
	wireBuf        []byte
	scratchPayload []byte
	adBuf          [8]byte
	eng            engine.AEAD
}

// NewEncryptor creates a standard-mode encryptor (maison wire + engine AEAD).
func NewEncryptor(w io.Writer, key []byte) (*Encryptor, error) {
	return newEncryptor(w, key, engine.Default())
}

func newEncryptor(w io.Writer, key []byte, eng engine.AEAD) (*Encryptor, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream55: key must be 32 bytes")
	}
	var enc Encryptor
	copy(enc.key[:], key)
	enc.w = w
	enc.eng = eng
	enc.wireBuf = make([]byte, 4+ChunkSize+1+TagSize+64)
	enc.scratchPayload = make([]byte, ChunkSize+64)
	if _, err := rand.Read(enc.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to generate nonce: %w", err)
	}
	if _, err := w.Write(enc.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to write header nonce: %w", err)
	}
	eng.HChaCha20(enc.subkey[:], enc.key[:], enc.nonce[0:16])
	return &enc, nil
}

// Write encrypts p using unique chunk nonces (N_chunk = N_base ^ seq).
func (e *Encryptor) Write(p []byte) (int, error) {
	totalWritten := 0
	for len(p) > 0 {
		chunkLen := len(p)
		if chunkLen > ChunkSize {
			chunkLen = ChunkSize
		}
		chunk := p[:chunkLen]
		p = p[chunkLen:]

		// Nonce IETF 12 octets : 4 octets zéro || (base ^ seq). La sous-clé
		// HChaCha20 (nonce[0:16] constant par stream) est cachée dans e.subkey.
		var chunkNonce12 [12]byte
		baseSeq := binary.BigEndian.Uint64(e.nonce[16:24])
		binary.BigEndian.PutUint64(chunkNonce12[4:12], baseSeq^e.seq)
		binary.BigEndian.PutUint64(e.adBuf[:], e.seq)
		e.seq++

		payload := chunk
		var mac [16]byte
		dstCipher := e.wireBuf[4 : 4+len(payload)]
		if err := e.eng.LockSubkeyDst(dstCipher, &mac, e.subkey[:], chunkNonce12[:], e.adBuf[:], payload); err != nil {
			return totalWritten, fmt.Errorf("secretstream55: AEAD lock failed: %w", err)
		}
		macOffset := 4 + len(payload)
		copy(e.wireBuf[macOffset:macOffset+16], mac[:])
		totalWireLen := uint32(len(payload) + 16)
		binary.BigEndian.PutUint32(e.wireBuf[0:4], totalWireLen)
		frameLen := 4 + len(payload) + 16
		if _, err := e.w.Write(e.wireBuf[:frameLen]); err != nil {
			return totalWritten, err
		}
		totalWritten += chunkLen
	}
	return totalWritten, nil
}

// Decryptor — standard (maison) framing.
type Decryptor struct {
	r        io.Reader
	key      [32]byte
	nonce    [24]byte
	subkey   [32]byte // HChaCha20(key, nonce[0:16]) — dérivée une fois par stream
	seq      uint64
	outBuf   []byte
	inBuf    []byte
	plainBuf []byte
	lenBuf   [4]byte
	adBuf    [8]byte
	eng      engine.AEAD
}

// NewDecryptor creates a standard-mode decryptor.
func NewDecryptor(r io.Reader, key []byte) (*Decryptor, error) {
	return newDecryptor(r, key, engine.Default())
}

func newDecryptor(r io.Reader, key []byte, eng engine.AEAD) (*Decryptor, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream55: key must be 32 bytes")
	}
	var dec Decryptor
	copy(dec.key[:], key)
	dec.r = r
	dec.eng = eng
	dec.inBuf = make([]byte, ChunkSize+TagSize+64)
	dec.plainBuf = make([]byte, ChunkSize+TagSize+64)
	if _, err := io.ReadFull(r, dec.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to read header nonce: %w", err)
	}
	eng.HChaCha20(dec.subkey[:], dec.key[:], dec.nonce[0:16])
	return &dec, nil
}

// Read decrypts from the maison stream.
func (d *Decryptor) Read(p []byte) (int, error) {
	if len(d.outBuf) > 0 {
		n := copy(p, d.outBuf)
		d.outBuf = d.outBuf[n:]
		return n, nil
	}
	if _, err := io.ReadFull(d.r, d.lenBuf[:]); err != nil {
		return 0, err
	}
	chunkLen := binary.BigEndian.Uint32(d.lenBuf[:])
	if chunkLen < 16 {
		return 0, fmt.Errorf("secretstream55: payload length too short")
	}
	totalPayloadLen := int(chunkLen)
	if cap(d.inBuf) < totalPayloadLen {
		d.inBuf = make([]byte, totalPayloadLen)
	}
	payloadBuf := d.inBuf[:totalPayloadLen]
	if _, err := io.ReadFull(d.r, payloadBuf); err != nil {
		return 0, fmt.Errorf("secretstream55: read payload failed: %w", err)
	}
	cipherLen := totalPayloadLen - 16
	cipherText := payloadBuf[:cipherLen]
	mac := payloadBuf[cipherLen:]

	var chunkNonce12 [12]byte
	baseSeq := binary.BigEndian.Uint64(d.nonce[16:24])
	binary.BigEndian.PutUint64(chunkNonce12[4:12], baseSeq^d.seq)
	binary.BigEndian.PutUint64(d.adBuf[:], d.seq)
	d.seq++

	if cap(d.plainBuf) < cipherLen {
		d.plainBuf = make([]byte, cipherLen)
	}
	plainDst := d.plainBuf[:cipherLen]
	unlocked, err := d.eng.UnlockSubkeyDst(plainDst, d.subkey[:], chunkNonce12[:], d.adBuf[:], cipherText, mac)
	if err != nil {
		return 0, fmt.Errorf("secretstream55: AEAD unlock failed: %w", err)
	}
	n := copy(p, unlocked)
	if n < len(unlocked) {
		d.outBuf = unlocked[n:]
	}
	return n, nil
}

// --- Libsodium wire (crypto_secretstream_xchacha20poly1305, wal-g framing) ---

// NewLibsodiumEncryptor returns a WriteCloser using true libsodium secretstream wire.
// Caller must Close() to emit TAG_FINAL.
func NewLibsodiumEncryptor(w io.Writer, key []byte) (io.WriteCloser, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream55: key must be 32 bytes")
	}
	return lsstream.NewWriter(w, key), nil
}

// NewLibsodiumDecryptor returns a Reader for true libsodium secretstream wire.
func NewLibsodiumDecryptor(r io.Reader, key []byte) (io.Reader, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream55: key must be 32 bytes")
	}
	return lsstream.NewReader(r, key), nil
}
