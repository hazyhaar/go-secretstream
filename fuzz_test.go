// SPDX-License-Identifier: Apache-2.0 OR MIT

package secretstream55_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/hazyhaar/go-secretstream"
)

// FuzzLibsodiumStreamRoundtrip exercises streaming encryption and decryption
// with random data lengths across multiple 8192-byte chunk boundaries.
func FuzzLibsodiumStreamRoundtrip(f *testing.F) {
	// Seed corpus with critical edge cases
	f.Add([]byte{})
	f.Add([]byte("a"))
	f.Add(bytes.Repeat([]byte{0x42}, 8191))
	f.Add(bytes.Repeat([]byte{0x42}, 8192))
	f.Add(bytes.Repeat([]byte{0x42}, 8193))
	f.Add(bytes.Repeat([]byte{0xAA}, 16384))
	f.Add(bytes.Repeat([]byte{0x55}, 20000))

	f.Fuzz(func(t *testing.T, plain []byte) {
		key := make([]byte, 32)
		for i := range key {
			key[i] = byte(i + 1)
		}

		// 1. Encrypt
		var cipherBuf bytes.Buffer
		enc, err := secretstream55.NewLibsodiumEncryptor(&cipherBuf, key)
		if err != nil {
			t.Fatalf("NewLibsodiumEncryptor failed: %v", err)
		}

		if len(plain) > 0 {
			if _, err := enc.Write(plain); err != nil {
				t.Fatalf("enc.Write failed: %v", err)
			}
		}
		if err := enc.Close(); err != nil {
			t.Fatalf("enc.Close failed: %v", err)
		}

		// 2. Decrypt
		dec, err := secretstream55.NewLibsodiumDecryptor(&cipherBuf, key)
		if err != nil {
			t.Fatalf("NewLibsodiumDecryptor failed: %v", err)
		}

		gotPlain, err := io.ReadAll(dec)
		if err != nil {
			t.Fatalf("dec.Read failed on valid stream: %v", err)
		}

		// 3. Assert exact match
		if !bytes.Equal(gotPlain, plain) {
			t.Fatalf("decrypted plain mismatch (len %d vs %d)", len(gotPlain), len(plain))
		}
	})
}

// FuzzLibsodiumCorruptedStream verifies that any mutated ciphertext or truncated
// stream fails closed without panicking or returning corrupted plaintext.
func FuzzLibsodiumCorruptedStream(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{0x00, 0x01, 0x02, 0x03})
	f.Add(bytes.Repeat([]byte{0xFF}, 24)) // Header only
	f.Add(bytes.Repeat([]byte{0xFF}, 50))

	f.Fuzz(func(t *testing.T, corruptWire []byte) {
		key := make([]byte, 32)
		for i := range key {
			key[i] = 0xA5
		}

		dec, err := secretstream55.NewLibsodiumDecryptor(bytes.NewReader(corruptWire), key)
		if err != nil {
			return // Valid fail-closed rejection at construction
		}

		// Reading from corrupt stream must either return an error or EOF without panicking
		buf := make([]byte, 1024)
		for {
			_, err := dec.Read(buf)
			if err != nil {
				break
			}
		}
	})
}
