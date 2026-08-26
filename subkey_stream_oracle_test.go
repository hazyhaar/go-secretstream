// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/engine"
)

// oldPathStream reconstruit le flux maison v2 à partir des primitives,
// indépendamment de Encryptor : en-tête versionné, nonce de trame
// nonce[16:20]||seq_be64, AD v2, trame avec tag, bloc TagFinal.
func oldPathStream(t *testing.T, eng engine.AEAD, key []byte, nonce [24]byte, payload []byte) []byte {
	t.Helper()
	var out bytes.Buffer
	var hdr [HeaderSize]byte
	writeHeaderV2(hdr[:], &nonce)
	out.Write(hdr[:])

	var subkey [32]byte
	eng.HChaCha20(subkey[:], key, nonce[0:16])
	seq := uint64(0)
	seal := func(chunk []byte, tag byte) {
		t.Helper()
		n12 := frameNonceV2(&nonce, seq)
		var prefix [adPrefixV2Len]byte
		var ext []byte
		ad := bindChunkADv2(&prefix, &ext, seq, tag, nil)
		cipher := make([]byte, len(chunk))
		var mac [16]byte
		if err := eng.LockSubkeyDst(cipher, &mac, subkey[:], n12[:], ad, chunk); err != nil {
			t.Fatalf("old path LockSubkeyDst: %v", err)
		}
		var lenBuf [4]byte
		binary.BigEndian.PutUint32(lenBuf[:], uint32(1+len(chunk)+16))
		out.Write(lenBuf[:])
		out.Write([]byte{tag})
		out.Write(cipher)
		out.Write(mac[:])
		seq++
	}
	for len(payload) > 0 {
		chunkLen := len(payload)
		if chunkLen > ChunkSize {
			chunkLen = ChunkSize
		}
		seal(payload[:chunkLen], TagMessage)
		payload = payload[chunkLen:]
	}
	seal(nil, TagFinal)
	return out.Bytes()
}

// TestSubkeyStream_ByteExact_VsOldPath — oracle n°1 du chantier sous-clé :
// le flux produit par l'Encryptor (sous-clé cachée, nonce IETF 12 octets) doit
// être identique OCTET PAR OCTET au flux de l'ancien chemin, à clé et nonce
// fixés, pour les tailles 0, 1, 64Ki et 1Mi.
func TestSubkeyStream_ByteExact_VsOldPath(t *testing.T) {
	eng := engine.Default()
	key := make([]byte, 32)
	for i := range key {
		key[i] = byte(0x51 ^ i*13)
	}

	for _, n := range []int{0, 1, 64 * 1024, 1 << 20} {
		payload := make([]byte, n)
		for i := range payload {
			payload[i] = byte(i*89 + 7)
		}

		var newStream bytes.Buffer
		enc, err := NewEncryptor(&newStream, key)
		if err != nil {
			t.Fatalf("n=%d: NewEncryptor: %v", n, err)
		}
		if _, err := enc.Write(payload); err != nil {
			t.Fatalf("n=%d: Write: %v", n, err)
		}
		if err := enc.Close(); err != nil {
			t.Fatalf("n=%d: Close: %v", n, err)
		}

		// Nonce fixé par lecture de l'en-tête v2 réellement émis (octets 12..36).
		var nonce [24]byte
		copy(nonce[:], newStream.Bytes()[12:12+24])

		oldStream := oldPathStream(t, eng, key, nonce, payload)

		if !bytes.Equal(oldStream, newStream.Bytes()) {
			for i := range oldStream {
				if i >= newStream.Len() || oldStream[i] != newStream.Bytes()[i] {
					t.Fatalf("n=%d: divergence at byte %d (old len %d, new len %d)",
						n, i, len(oldStream), newStream.Len())
				}
			}
			t.Fatalf("n=%d: streams differ in length: old %d, new %d", n, len(oldStream), newStream.Len())
		}

		// Le Decryptor (chemin sous-clé) doit rendre le payload d'origine.
		dec, err := NewDecryptor(bytes.NewReader(newStream.Bytes()), key)
		if err != nil {
			t.Fatalf("n=%d: NewDecryptor: %v", n, err)
		}
		got := make([]byte, 0, n)
		buf := make([]byte, 32*1024)
		for {
			m, err := dec.Read(buf)
			got = append(got, buf[:m]...)
			if err != nil {
				break
			}
			if len(got) >= n && n >= 0 {
				if len(got) == n {
					break
				}
			}
		}
		if !bytes.Equal(got, payload) {
			t.Fatalf("n=%d: decrypt roundtrip mismatch (got %d bytes)", n, len(got))
		}
	}
}
