// Package secretstream55 provides ultra-fast streaming encryption and decryption in Pure Go using SIMD256-accelerated XChaCha20-Poly1305.
package secretstream55

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/hazyhaar/c2simd"
)

const (
	HeaderSize = 24 // 24-byte XChaCha20 nonce header
	TagSize    = 16 // 16-byte Poly1305 MAC tag
	ChunkSize  = 64 * 1024

	// Libsodium Tag Constants
	TagMessage byte = 0x00 // Standard chunk tag
	TagPush    byte = 0x01 // Flush tag
	TagRekey   byte = 0x02 // Key rotation tag
	TagFinal   byte = 0x03 // Final stream chunk tag
)

// Encryptor wraps an io.Writer to encrypt outgoing data stream chunks.
type Encryptor struct {
	w              io.Writer
	key            [32]byte
	nonce          [24]byte
	seq            uint64
	libsodium      bool
	buf            []byte
	scratchPayload []byte
	lenBuf         [4]byte
	adBuf          [8]byte
}

// NewEncryptor creates a streaming AEAD encryptor utilizing c2simd SIMD256 acceleration.
func NewEncryptor(w io.Writer, key []byte) (*Encryptor, error) {
	return newEncryptor(w, key, false)
}

// NewLibsodiumEncryptor creates a streaming AEAD encryptor compatible with Libsodium's crypto_secretstream framing.
func NewLibsodiumEncryptor(w io.Writer, key []byte) (*Encryptor, error) {
	return newEncryptor(w, key, true)
}

func newEncryptor(w io.Writer, key []byte, libsodium bool) (*Encryptor, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream55: key must be 32 bytes")
	}

	var enc Encryptor
	copy(enc.key[:], key)
	enc.w = w
	enc.libsodium = libsodium
	enc.buf = make([]byte, ChunkSize+TagSize+16)
	enc.scratchPayload = make([]byte, ChunkSize+1)

	if _, err := rand.Read(enc.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to generate nonce: %w", err)
	}

	if _, err := w.Write(enc.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to write header nonce: %w", err)
	}

	return &enc, nil
}

// Write encrypts p and writes encrypted chunks to the underlying io.Writer without heap allocations per chunk.
func (e *Encryptor) Write(p []byte) (int, error) {
	totalWritten := 0
	for len(p) > 0 {
		chunkLen := len(p)
		if chunkLen > ChunkSize {
			chunkLen = ChunkSize
		}
		chunk := p[:chunkLen]
		p = p[chunkLen:]

		binary.BigEndian.PutUint64(e.adBuf[:], e.seq)
		e.seq++

		var payload []byte
		if e.libsodium {
			tag := TagMessage
			if len(p) == 0 {
				tag = TagFinal
			}
			copy(e.scratchPayload[:chunkLen], chunk)
			e.scratchPayload[chunkLen] = tag
			payload = e.scratchPayload[:chunkLen+1]
		} else {
			payload = chunk
		}

		var mac [16]byte
		cipherText, err := c2simd.AEADLockDst(e.buf[:len(payload)], &mac, e.key[:], e.nonce[:], e.adBuf[:], payload)
		if err != nil {
			return totalWritten, fmt.Errorf("secretstream55: AEAD lock failed: %w", err)
		}

		binary.BigEndian.PutUint32(e.lenBuf[:], uint32(len(cipherText)+16))

		if _, err := e.w.Write(e.lenBuf[:]); err != nil {
			return totalWritten, err
		}
		if _, err := e.w.Write(cipherText); err != nil {
			return totalWritten, err
		}
		if _, err := e.w.Write(mac[:]); err != nil {
			return totalWritten, err
		}

		totalWritten += chunkLen
	}
	return totalWritten, nil
}

// Decryptor wraps an io.Reader to decrypt incoming stream chunks.
type Decryptor struct {
	r         io.Reader
	key       [32]byte
	nonce     [24]byte
	seq       uint64
	outBuf    []byte
	inBuf     []byte
	plainBuf  []byte
	lenBuf    [4]byte
	adBuf     [8]byte
	libsodium bool
}

// NewDecryptor creates a streaming AEAD decryptor for standard secretstream55 streams.
func NewDecryptor(r io.Reader, key []byte) (*Decryptor, error) {
	return newDecryptor(r, key, false)
}

// NewLibsodiumDecryptor creates a streaming AEAD decryptor compatible with Libsodium C crypto_secretstream archives.
func NewLibsodiumDecryptor(r io.Reader, key []byte) (*Decryptor, error) {
	return newDecryptor(r, key, true)
}

func newDecryptor(r io.Reader, key []byte, libsodium bool) (*Decryptor, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream55: key must be 32 bytes")
	}

	var dec Decryptor
	copy(dec.key[:], key)
	dec.r = r
	dec.libsodium = libsodium
	dec.inBuf = make([]byte, ChunkSize+TagSize+16)
	dec.plainBuf = make([]byte, ChunkSize+TagSize+16)

	if _, err := io.ReadFull(r, dec.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to read header nonce: %w", err)
	}

	return &dec, nil
}

// Read decrypts p from the underlying encrypted stream using c2simd SIMD256.
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

	binary.BigEndian.PutUint64(d.adBuf[:], d.seq)
	d.seq++

	if cap(d.plainBuf) < cipherLen {
		d.plainBuf = make([]byte, cipherLen)
	}

	plainText, err := c2simd.AEADUnlockDst(d.plainBuf[:cipherLen], d.key[:], d.nonce[:], d.adBuf[:], cipherText, mac)
	if err != nil {
		return 0, fmt.Errorf("secretstream55: AEAD authentication check failed: %w", err)
	}

	if d.libsodium {
		if len(plainText) < 1 {
			return 0, fmt.Errorf("secretstream55: invalid libsodium payload format")
		}
		plainText = plainText[:len(plainText)-1]
	}

	n := copy(p, plainText)
	if n < len(plainText) {
		d.outBuf = plainText[n:]
	}
	return n, nil
}
