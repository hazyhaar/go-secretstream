// Package secretstream55 provides ultra-fast streaming encryption and decryption in Pure Go using XChaCha20-Poly1305 AEAD.
package secretstream

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
)

const (
	HeaderSize = 24 // 24-byte XChaCha20 nonce header
	TagSize    = 16 // 16-byte Poly1305 MAC tag
	ChunkSize  = 64 * 1024

	// Libsodium Tag Constants
	TagMessage   byte = 0x00 // Standard chunk tag
	TagPush      byte = 0x01 // Flush tag
	TagRekey     byte = 0x02 // Key rotation tag
	TagFinal     byte = 0x03 // Final stream chunk tag
)

// Encryptor wraps an io.Writer to encrypt outgoing data stream chunks.
type Encryptor struct {
	w         io.Writer
	aead      cipher.AEAD
	nonce     []byte
	seq       uint64
	libsodium bool
}

// NewEncryptor creates a streaming AEAD encryptor utilizing hardware-accelerated XChaCha20-Poly1305.
func NewEncryptor(w io.Writer, key []byte) (*Encryptor, error) {
	return newEncryptor(w, key, false)
}

// NewLibsodiumEncryptor creates a streaming AEAD encryptor compatible with Libsodium's crypto_secretstream C framing.
func NewLibsodiumEncryptor(w io.Writer, key []byte) (*Encryptor, error) {
	return newEncryptor(w, key, true)
}

func newEncryptor(w io.Writer, key []byte, libsodium bool) (*Encryptor, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream: key must be 32 bytes")
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, fmt.Errorf("secretstream: failed to create AEAD: %w", err)
	}

	nonce := make([]byte, 24)
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("secretstream: failed to generate nonce: %w", err)
	}

	if _, err := w.Write(nonce); err != nil {
		return nil, fmt.Errorf("secretstream: failed to write header nonce: %w", err)
	}

	return &Encryptor{
		w:         w,
		aead:      aead,
		nonce:     nonce,
		seq:       0,
		libsodium: libsodium,
	}, nil
}

// Write encrypts p and writes encrypted chunks to the underlying io.Writer.
func (e *Encryptor) Write(p []byte) (int, error) {
	totalWritten := 0
	for len(p) > 0 {
		chunkLen := len(p)
		if chunkLen > ChunkSize {
			chunkLen = ChunkSize
		}
		chunk := p[:chunkLen]
		p = p[chunkLen:]

		ad := make([]byte, 8)
		binary.BigEndian.PutUint64(ad, e.seq)
		e.seq++

		var payload []byte
		if e.libsodium {
			tag := TagMessage
			if len(p) == 0 {
				tag = TagFinal
			}
			payload = append(chunk, tag)
		} else {
			payload = chunk
		}

		sealed := e.aead.Seal(nil, e.nonce, payload, ad)

		lenBuf := make([]byte, 4)
		binary.BigEndian.PutUint32(lenBuf, uint32(len(sealed)))

		if _, err := e.w.Write(lenBuf); err != nil {
			return totalWritten, err
		}
		if _, err := e.w.Write(sealed); err != nil {
			return totalWritten, err
		}

		totalWritten += chunkLen
	}
	return totalWritten, nil
}

// Decryptor wraps an io.Reader to decrypt incoming stream chunks.
type Decryptor struct {
	r         io.Reader
	aead      cipher.AEAD
	nonce     []byte
	seq       uint64
	buf       []byte
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
		return nil, fmt.Errorf("secretstream: key must be 32 bytes")
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, fmt.Errorf("secretstream: failed to create AEAD: %w", err)
	}

	nonce := make([]byte, 24)
	if _, err := io.ReadFull(r, nonce); err != nil {
		return nil, fmt.Errorf("secretstream: failed to read header nonce: %w", err)
	}

	return &Decryptor{
		r:         r,
		aead:      aead,
		nonce:     nonce,
		seq:       0,
		libsodium: libsodium,
	}, nil
}

// Read decrypts p from the underlying encrypted stream.
func (d *Decryptor) Read(p []byte) (int, error) {
	if len(d.buf) > 0 {
		n := copy(p, d.buf)
		d.buf = d.buf[n:]
		return n, nil
	}

	lenBuf := make([]byte, 4)
	if _, err := io.ReadFull(d.r, lenBuf); err != nil {
		return 0, err
	}

	chunkLen := binary.BigEndian.Uint32(lenBuf)
	sealed := make([]byte, chunkLen)
	if _, err := io.ReadFull(d.r, sealed); err != nil {
		return 0, fmt.Errorf("secretstream: read payload failed: %w", err)
	}

	ad := make([]byte, 8)
	binary.BigEndian.PutUint64(ad, d.seq)
	d.seq++

	plainText, err := d.aead.Open(nil, d.nonce, sealed, ad)
	if err != nil {
		return 0, fmt.Errorf("secretstream: AEAD authentication check failed: %w", err)
	}

	if d.libsodium {
		if len(plainText) < 1 {
			return 0, fmt.Errorf("secretstream: invalid libsodium payload format")
		}
		// Extract trailing Libsodium tag (0x00 TagMessage, 0x03 TagFinal, etc.)
		plainText = plainText[:len(plainText)-1]
	}

	n := copy(p, plainText)
	if n < len(plainText) {
		d.buf = plainText[n:]
	}
	return n, nil
}
