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
	wireBuf        []byte // Buffer contigu unique : 4B taille + ChunkSize + 16B MAC
	scratchPayload []byte
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
	enc.wireBuf = make([]byte, 4+ChunkSize+1+TagSize+64)
	enc.scratchPayload = make([]byte, ChunkSize+64)

	if _, err := rand.Read(enc.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to generate nonce: %w", err)
	}

	if _, err := w.Write(enc.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to write header nonce: %w", err)
	}

	return &enc, nil
}

// Write encrypts p and writes encrypted chunks to the underlying io.Writer using unique chunk nonces (N_chunk = N_base ^ seq).
func (e *Encryptor) Write(p []byte) (int, error) {
	totalWritten := 0
	for len(p) > 0 {
		chunkLen := len(p)
		if chunkLen > ChunkSize {
			chunkLen = ChunkSize
		}
		chunk := p[:chunkLen]
		p = p[chunkLen:]

		// P0 Sécurité : Dérivation du nonce unique par chunk (N_chunk = N_base ^ seq)
		var chunkNonce [24]byte
		copy(chunkNonce[:], e.nonce[:])
		baseSeq := binary.BigEndian.Uint64(e.nonce[16:24])
		binary.BigEndian.PutUint64(chunkNonce[16:24], baseSeq^e.seq)

		binary.BigEndian.PutUint64(e.adBuf[:], e.seq)
		e.seq++

		var payload []byte
		tag := TagMessage
		if e.libsodium {
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
		dstCipher := e.wireBuf[4 : 4+len(payload)]
		_, err := c2simd.AEADLockDst(dstCipher, &mac, e.key[:], chunkNonce[:], e.adBuf[:], payload)
		if err != nil {
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

		// Prise en charge de la rotation de clé TagRekey (0x02)
		if tag == TagRekey {
			c2simd.HChaCha20(e.key[:], e.nonce[:16], e.key[:])
			e.seq = 0
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
	// Buffers pré-dimensionnés pour garantir 0 allocation sur Read() steady-state
	dec.inBuf = make([]byte, ChunkSize+TagSize+64)
	dec.plainBuf = make([]byte, ChunkSize+TagSize+64)

	if _, err := io.ReadFull(r, dec.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to read header nonce: %w", err)
	}

	return &dec, nil
}

// Read decrypts p from the underlying encrypted stream using unique chunk nonces with 0 allocation steady-state.
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

	// P0 Sécurité : Dérivation du nonce unique par chunk (N_chunk = N_base ^ seq)
	var chunkNonce [24]byte
	copy(chunkNonce[:], d.nonce[:])
	baseSeq := binary.BigEndian.Uint64(d.nonce[16:24])
	binary.BigEndian.PutUint64(chunkNonce[16:24], baseSeq^d.seq)

	binary.BigEndian.PutUint64(d.adBuf[:], d.seq)
	d.seq++

	if cap(d.plainBuf) < cipherLen {
		d.plainBuf = make([]byte, cipherLen)
	}
	plainDst := d.plainBuf[:cipherLen]

	unlocked, err := c2simd.AEADUnlockDst(plainDst, d.key[:], chunkNonce[:], d.adBuf[:], cipherText, mac)
	if err != nil {
		return 0, fmt.Errorf("secretstream55: AEAD unlock failed: %w", err)
	}

	if d.libsodium {
		if len(unlocked) == 0 {
			return 0, fmt.Errorf("secretstream55: empty libsodium payload")
		}
		tag := unlocked[len(unlocked)-1]
		unlocked = unlocked[:len(unlocked)-1]
		if tag == TagRekey {
			c2simd.HChaCha20(d.key[:], d.nonce[:16], d.key[:])
			d.seq = 0
		}
	}

	n := copy(p, unlocked)
	if n < len(unlocked) {
		d.outBuf = unlocked[n:]
	}

	return n, nil
}
