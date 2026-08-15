package secretstream55

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/engine"
)

// oldPathStream reconstruit le flux maison par le chemin HISTORIQUE : un appel
// XChaCha20-Poly1305 complet (nonce 24 octets, HChaCha20 recalculée) par chunk.
func oldPathStream(t *testing.T, eng engine.AEAD, key []byte, nonce [24]byte, payload []byte) []byte {
	t.Helper()
	var out bytes.Buffer
	out.Write(nonce[:])
	seq := uint64(0)
	baseSeq := binary.BigEndian.Uint64(nonce[16:24])
	for len(payload) > 0 {
		chunkLen := len(payload)
		if chunkLen > ChunkSize {
			chunkLen = ChunkSize
		}
		chunk := payload[:chunkLen]
		payload = payload[chunkLen:]

		var chunkNonce [24]byte
		copy(chunkNonce[:], nonce[:])
		binary.BigEndian.PutUint64(chunkNonce[16:24], baseSeq^seq)
		var ad [8]byte
		binary.BigEndian.PutUint64(ad[:], seq)
		seq++

		cipher := make([]byte, chunkLen)
		var mac [16]byte
		if err := eng.LockDst(cipher, &mac, key, chunkNonce[:], ad[:], chunk); err != nil {
			t.Fatalf("old path LockDst: %v", err)
		}
		var lenBuf [4]byte
		binary.BigEndian.PutUint32(lenBuf[:], uint32(chunkLen+16))
		out.Write(lenBuf[:])
		out.Write(cipher)
		out.Write(mac[:])
	}
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

		// Nonce fixé par lecture de l'en-tête réellement émis : les deux
		// chemins travaillent sous exactement la même clé et le même nonce.
		var nonce [24]byte
		copy(nonce[:], newStream.Bytes()[:HeaderSize])

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
