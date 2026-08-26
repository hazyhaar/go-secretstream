// SPDX-License-Identifier: Apache-2.0 OR MIT

// Package secretstream55 provides streaming AEAD encryption in Pure Go.
//
// Two wire modes:
//   - Standard (NewEncryptor): maison framing v2 (en-tête versionné + tag authentifié + TagFinal).
//     NewDecryptor relit le v2 et les archives v1 (décodeur hybride).
//   - Libsodium (NewLibsodiumEncryptor): crypto_secretstream_xchacha20poly1305 wire (wal-g compatible).
//
// AEAD backend for standard mode: monocypher55 (default — bascule 2026-08-15,
// gate à trois preuves : fil croisé bit-identique, aliasing, rejet de forge).
// Un autre moteur (par exemple c2simd/aeadengine) s'injecte par
// NewEncryptorWithEngine / NewDecryptorWithEngine ; ce module n'importe aucun
// moteur optionnel.
package secretstream55

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/hazyhaar/go-secretstream/internal/engine"
	"github.com/hazyhaar/go-secretstream/internal/lsstream"
)

// Encryptor — standard (maison) framing + engine AEAD.
type Encryptor struct {
	w              io.Writer
	nonce          [24]byte
	subkey         [32]byte // HChaCha20(key, nonce[0:16]) — dérivée une fois par stream
	seq            uint64
	wireBuf        []byte
	scratchPayload []byte
	adBuf          [adPrefixV2Len]byte
	adExt          []byte
	eng            engine.AEAD
	stickyErr      error
	closed         bool
}

// Engine est le jeu de méthodes AEAD qu'un moteur externe doit fournir pour
// être injecté dans le flux maison. Toute valeur satisfaisant ce jeu de
// méthodes convient (appariement structurel, aucune dépendance de module).
type Engine = engine.AEAD

// NewEncryptor creates a standard-mode encryptor (maison wire + engine AEAD).
func NewEncryptor(w io.Writer, key []byte) (*Encryptor, error) {
	return newEncryptor(w, key, engine.Default())
}

// NewEncryptorWithEngine est NewEncryptor avec un moteur AEAD fourni par
// l'appelant à la place du moteur par défaut.
func NewEncryptorWithEngine(w io.Writer, key []byte, eng Engine) (*Encryptor, error) {
	if eng == nil {
		return nil, fmt.Errorf("secretstream55: nil engine")
	}
	return newEncryptor(w, key, eng)
}

func newEncryptor(w io.Writer, key []byte, eng engine.AEAD) (*Encryptor, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream55: key must be 32 bytes")
	}
	var enc Encryptor
	enc.w = w
	enc.eng = eng
	enc.wireBuf = make([]byte, 4+ChunkSize+1+TagSize+64)
	enc.scratchPayload = make([]byte, ChunkSize+64)
	if _, err := rand.Read(enc.nonce[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to generate nonce: %w", err)
	}
	var hdr [HeaderSizeV2]byte
	writeHeaderV2(hdr[:], &enc.nonce)
	if _, err := w.Write(hdr[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to write v2 header: %w", err)
	}
	eng.HChaCha20(enc.subkey[:], key, enc.nonce[0:16])
	return &enc, nil
}

// Close écrit une trame TagFinal à chiffré vide, puis verrouille l'écrivain
// et efface la sous-clé ainsi que les tampons possédés. Un second Close est
// un no-op : un seul bloc terminal est émis. La clé fournie par l'appelant
// n'est pas effacée : elle reste sa responsabilité.
func (e *Encryptor) Close() error {
	if e.closed {
		return nil
	}
	var closeErr error
	if e.stickyErr == nil {
		closeErr = e.writeFrame(nil, nil, TagFinal)
	}
	e.closed = true
	clear(e.subkey[:])
	clear(e.scratchPayload)
	clear(e.wireBuf)
	clear(e.adBuf[:])
	clear(e.adExt)
	return closeErr
}

// Write chiffre p. Équivalent à WriteWithAD(p, nil).
func (e *Encryptor) Write(p []byte) (int, error) {
	return e.WriteWithAD(p, nil)
}

// WriteWithAD chiffre p en liant chaque fragment à ad (format v2).
//
//	AD = "SS55-v2\x00" || seq_be64 || tag || len(ad_appelant)_be32 || ad_appelant
//
// Le numéro de séquence et le tag sont authentifiés. La donnée d'appelant
// n'est pas transmise sur le fil : le lecteur la fournit. Les fragments
// restent bornés à ChunkSize.
func (e *Encryptor) WriteWithAD(p, ad []byte) (int, error) {
	if e.closed {
		return 0, fmt.Errorf("secretstream55: write to closed encryptor")
	}
	if e.stickyErr != nil {
		return 0, e.stickyErr
	}
	totalWritten := 0
	for len(p) > 0 {
		chunkLen := len(p)
		if chunkLen > ChunkSize {
			chunkLen = ChunkSize
		}
		chunk := p[:chunkLen]
		p = p[chunkLen:]
		if err := e.writeFrame(chunk, ad, TagMessage); err != nil {
			return totalWritten, err
		}
		totalWritten += chunkLen
	}
	return totalWritten, nil
}

func (e *Encryptor) writeFrame(payload, ad []byte, tag byte) error {
	if e.seq == ^uint64(0) {
		e.stickyErr = fmt.Errorf("secretstream55: sequence number overflow")
		return e.stickyErr
	}
	chunkNonce12 := frameNonceV2(&e.nonce, e.seq)
	chunkAD := bindChunkADv2(&e.adBuf, &e.adExt, e.seq, tag, ad)

	var mac [16]byte
	e.wireBuf[4] = tag
	dstCipher := e.wireBuf[5 : 5+len(payload)]
	if err := e.eng.LockSubkeyDst(dstCipher, &mac, e.subkey[:], chunkNonce12[:], chunkAD, payload); err != nil {
		e.stickyErr = fmt.Errorf("secretstream55: AEAD lock failed: %w", err)
		return e.stickyErr
	}
	macOffset := 5 + len(payload)
	copy(e.wireBuf[macOffset:macOffset+16], mac[:])
	totalWireLen := uint32(1 + len(payload) + 16)
	binary.BigEndian.PutUint32(e.wireBuf[0:4], totalWireLen)
	frameLen := 4 + int(totalWireLen)
	n, err := e.w.Write(e.wireBuf[:frameLen])
	if err != nil {
		e.stickyErr = err
		return err
	}
	if n != frameLen {
		e.stickyErr = fmt.Errorf("secretstream55: short write (%d/%d bytes)", n, frameLen)
		return e.stickyErr
	}
	e.seq++
	return nil
}

// Decryptor — standard (maison) framing.
type Decryptor struct {
	r         io.Reader
	nonce     [24]byte
	subkey    [32]byte // HChaCha20(key, nonce[0:16]) — dérivée une fois par stream
	seq       uint64
	outBuf    []byte
	inBuf     []byte
	plainBuf  []byte
	lenBuf    [4]byte
	adBuf     [adPrefixV2Len]byte
	adExt     []byte
	eng       engine.AEAD
	stickyErr error
	closed    bool
	format    uint8
	finalSeen bool
}

// NewDecryptor creates a standard-mode decryptor.
func NewDecryptor(r io.Reader, key []byte) (*Decryptor, error) {
	return newDecryptor(r, key, engine.Default())
}

// NewDecryptorWithEngine est NewDecryptor avec un moteur AEAD fourni par
// l'appelant ; il doit être le même moteur (ou un moteur bit-identique) que
// celui qui a produit le flux.
func NewDecryptorWithEngine(r io.Reader, key []byte, eng Engine) (*Decryptor, error) {
	if eng == nil {
		return nil, fmt.Errorf("secretstream55: nil engine")
	}
	return newDecryptor(r, key, eng)
}

func newDecryptor(r io.Reader, key []byte, eng engine.AEAD) (*Decryptor, error) {
	if len(key) != 32 {
		return nil, fmt.Errorf("secretstream55: key must be 32 bytes")
	}
	var dec Decryptor
	dec.r = r
	dec.eng = eng
	dec.inBuf = make([]byte, ChunkSize+TagSize+64)
	dec.plainBuf = make([]byte, ChunkSize+TagSize+64)
	var head8 [8]byte
	if _, err := io.ReadFull(r, head8[:]); err != nil {
		return nil, fmt.Errorf("secretstream55: failed to read header: %w", err)
	}
	if isMagicV2(head8[:]) {
		var rest [28]byte
		if _, err := io.ReadFull(r, rest[:]); err != nil {
			return nil, fmt.Errorf("secretstream55: failed to read v2 header: %w", err)
		}
		ver := binary.BigEndian.Uint16(rest[0:2])
		flags := binary.BigEndian.Uint16(rest[2:4])
		if ver != VersionV2 {
			return nil, fmt.Errorf("secretstream55: unsupported version %d", ver)
		}
		if flags != FlagsV2 {
			return nil, fmt.Errorf("secretstream55: reserved flags must be 0 (got %d)", flags)
		}
		copy(dec.nonce[:], rest[4:28])
		dec.format = 2
	} else {
		copy(dec.nonce[:8], head8[:])
		if _, err := io.ReadFull(r, dec.nonce[8:HeaderSizeV1]); err != nil {
			return nil, fmt.Errorf("secretstream55: failed to read header nonce: %w", err)
		}
		dec.format = 1
	}
	eng.HChaCha20(dec.subkey[:], key, dec.nonce[0:16])
	return &dec, nil
}

// Close verrouille le déchiffreur et efface la sous-clé ainsi que les tampons possédés.
// La clé fournie par l'appelant n'est pas effacée : elle reste sa responsabilité.
func (d *Decryptor) Close() error {
	if d.closed {
		return nil
	}
	d.closed = true
	clear(d.subkey[:])
	clear(d.inBuf)
	clear(d.plainBuf)
	clear(d.adBuf[:])
	clear(d.adExt)
	clear(d.outBuf)
	d.outBuf = nil
	return nil
}

// Read déchiffre depuis le flux maison. Équivalent à ReadWithAD(p, nil).
func (d *Decryptor) Read(p []byte) (int, error) {
	return d.ReadWithAD(p, nil)
}

// ReadWithAD déchiffre un fragment en fournissant la même donnée associée
// que l'écrivain. La donnée n'est pas lue sur le fil. Un reliquat déjà
// déchiffré (outBuf) est servi sans réappliquer ad : l'AEAD a déjà été
// vérifiée pour ce fragment.
func (d *Decryptor) ReadWithAD(p, ad []byte) (int, error) {
	if d.closed {
		return 0, fmt.Errorf("secretstream55: read from closed decryptor")
	}
	if d.stickyErr != nil {
		return 0, d.stickyErr
	}
	if len(d.outBuf) > 0 {
		n := copy(p, d.outBuf)
		d.outBuf = d.outBuf[n:]
		return n, nil
	}
	if d.finalSeen {
		return 0, io.EOF
	}
	if d.format == 2 {
		return d.readV2(p, ad)
	}
	return d.readV1(p, ad)
}

func (d *Decryptor) readV1(p, ad []byte) (int, error) {
	if _, err := io.ReadFull(d.r, d.lenBuf[:]); err != nil {
		d.stickyErr = err
		return 0, err
	}
	chunkLen := binary.BigEndian.Uint32(d.lenBuf[:])
	maxAllowed := uint32(ChunkSize + TagSize + 64)
	if chunkLen < 16 || chunkLen > maxAllowed {
		d.stickyErr = fmt.Errorf("secretstream55: invalid payload length %d (must be between 16 and %d)", chunkLen, maxAllowed)
		return 0, d.stickyErr
	}
	totalPayloadLen := int(chunkLen)
	if cap(d.inBuf) < totalPayloadLen {
		d.inBuf = make([]byte, totalPayloadLen)
	}
	payloadBuf := d.inBuf[:totalPayloadLen]
	if _, err := io.ReadFull(d.r, payloadBuf); err != nil {
		d.stickyErr = fmt.Errorf("secretstream55: read payload failed: %w", err)
		return 0, d.stickyErr
	}
	cipherLen := totalPayloadLen - 16
	cipherText := payloadBuf[:cipherLen]
	mac := payloadBuf[cipherLen:]

	chunkNonce12 := frameNonceV1(&d.nonce, d.seq)
	chunkAD := bindChunkAD(&d.adBuf, &d.adExt, d.seq, ad)
	return d.unlockInto(p, chunkNonce12[:], chunkAD, cipherText, mac)
}

func (d *Decryptor) readV2(p, ad []byte) (int, error) {
	if _, err := io.ReadFull(d.r, d.lenBuf[:]); err != nil {
		d.stickyErr = fmt.Errorf("secretstream55: flux tronqué")
		return 0, d.stickyErr
	}
	chunkLen := binary.BigEndian.Uint32(d.lenBuf[:])
	if chunkLen < MinFramePayloadV2 || chunkLen > MaxFramePayloadV2 {
		d.stickyErr = fmt.Errorf("secretstream55: invalid payload length %d (must be between %d and %d)", chunkLen, MinFramePayloadV2, MaxFramePayloadV2)
		return 0, d.stickyErr
	}
	totalPayloadLen := int(chunkLen)
	if cap(d.inBuf) < totalPayloadLen {
		d.inBuf = make([]byte, totalPayloadLen)
	}
	payloadBuf := d.inBuf[:totalPayloadLen]
	if _, err := io.ReadFull(d.r, payloadBuf); err != nil {
		d.stickyErr = fmt.Errorf("secretstream55: flux tronqué")
		return 0, d.stickyErr
	}
	tag := payloadBuf[0]
	if tag != TagMessage && tag != TagFinal {
		d.stickyErr = fmt.Errorf("secretstream55: tag v2 refusé 0x%02x", tag)
		return 0, d.stickyErr
	}
	cipherLen := totalPayloadLen - 1 - TagSize
	cipherText := payloadBuf[1 : 1+cipherLen]
	mac := payloadBuf[1+cipherLen:]

	chunkNonce12 := frameNonceV2(&d.nonce, d.seq)
	chunkAD := bindChunkADv2(&d.adBuf, &d.adExt, d.seq, tag, ad)
	if tag == TagFinal {
		if cap(d.plainBuf) < cipherLen {
			d.plainBuf = make([]byte, cipherLen)
		}
		unlocked, err := d.eng.UnlockSubkeyDst(d.plainBuf[:cipherLen], d.subkey[:], chunkNonce12[:], chunkAD, cipherText, mac)
		if err != nil {
			d.stickyErr = fmt.Errorf("secretstream55: AEAD unlock failed: %w", err)
			return 0, d.stickyErr
		}
		if cipherLen != 0 {
			clear(unlocked)
			d.stickyErr = fmt.Errorf("secretstream55: TagFinal avec chiffré non vide")
			return 0, d.stickyErr
		}
		d.seq++
		d.finalSeen = true
		var extra [1]byte
		n, _ := d.r.Read(extra[:])
		if n > 0 {
			d.stickyErr = fmt.Errorf("secretstream55: données après TagFinal")
			return 0, d.stickyErr
		}
		return 0, io.EOF
	}
	return d.unlockInto(p, chunkNonce12[:], chunkAD, cipherText, mac)
}

func (d *Decryptor) unlockInto(p, nonce12, chunkAD, cipherText, mac []byte) (int, error) {
	cipherLen := len(cipherText)
	if len(p) >= cipherLen {
		unlocked, err := d.eng.UnlockSubkeyDst(p[:cipherLen], d.subkey[:], nonce12, chunkAD, cipherText, mac)
		if err != nil {
			d.stickyErr = fmt.Errorf("secretstream55: AEAD unlock failed: %w", err)
			return 0, d.stickyErr
		}
		d.seq++
		return len(unlocked), nil
	}

	if cap(d.plainBuf) < cipherLen {
		d.plainBuf = make([]byte, cipherLen)
	}
	plainDst := d.plainBuf[:cipherLen]
	unlocked, err := d.eng.UnlockSubkeyDst(plainDst, d.subkey[:], nonce12, chunkAD, cipherText, mac)
	if err != nil {
		d.stickyErr = fmt.Errorf("secretstream55: AEAD unlock failed: %w", err)
		return 0, d.stickyErr
	}
	d.seq++
	n := copy(p, unlocked)
	if n < len(unlocked) {
		d.outBuf = unlocked[n:]
	}
	return n, nil
}

// bindChunkAD assemble AD v1 = seq_be64 || ad_appelant.
// Une ad vide renvoie les huit premiers octets de seqBuf, sans allocation.
func bindChunkAD(seqBuf *[adPrefixV2Len]byte, ext *[]byte, seq uint64, ad []byte) []byte {
	binary.BigEndian.PutUint64(seqBuf[:8], seq)
	if len(ad) == 0 {
		return seqBuf[:8]
	}
	need := 8 + len(ad)
	if cap(*ext) < need {
		*ext = make([]byte, need)
	}
	buf := (*ext)[:need]
	copy(buf[:8], seqBuf[:8])
	copy(buf[8:], ad)
	*ext = buf
	return buf
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
