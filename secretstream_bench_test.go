// Modified in the lateos-ai/wal-g fork.
// Benchmark comparatif multi-moteurs (Notre lib optimisée vs OpenZiti vs Go x/crypto AEAD)
// avec charges utiles standard (16MB WAL) et charges de VOLUME EXTRÊME (1GB).

package secretstream_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	local "github.com/hazyhaar/go-secretstream"
	openziti "github.com/openziti/secretstream"
	"golang.org/x/crypto/chacha20poly1305"
)

var benchPayloadSizes = []struct {
	name string
	size int
}{
	{"8KB_Chunk", 8192},
	{"16MB_PG_WAL", 16 * 1024 * 1024},          // Fichier WAL standard PostgreSQL
	{"1GB_Extreme_Volume", 1024 * 1024 * 1024}, // Charge EXTRÊME (1 Go continu)
}

func generatePayload(size int) []byte {
	b := make([]byte, size)
	_, _ = rand.Read(b)
	return b
}

// -----------------------------------------------------------------------------
// 1. NOTRE LIB OPTIMISÉE - MODE DIRECT EN MÉMOIRE (Apples to Apples avec OpenZiti)
// -----------------------------------------------------------------------------

func BenchmarkEncrypt_NotreLib_Direct(b *testing.B) {
	key := make([]byte, local.KeyBytes)
	_, _ = rand.Read(key)

	for _, bm := range benchPayloadSizes {
		if testing.Short() && bm.size > 16*1024*1024 {
			continue
		}
		payload := generatePayload(bm.size)
		header := make([]byte, local.HeaderBytes)
		wireBuf := make([]byte, 8192+local.ABytes)

		b.Run(bm.name, func(b *testing.B) {
			b.SetBytes(int64(bm.size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				st, err := local.TestInitPush(key, header)
				if err != nil {
					b.Fatal(err)
				}
				for off := 0; off < len(payload); off += 8192 {
					end := off + 8192
					if end > len(payload) {
						end = len(payload)
					}
					tag := byte(0)
					if end == len(payload) {
						tag = byte(local.TagFinal)
					}
					_, _ = st.TestPushTo(payload[off:end], tag, wireBuf)
				}
			}
		})
	}
}

// -----------------------------------------------------------------------------
// 2. NOTRE LIB OPTIMISÉE - MODE WRITER STREAMING (Avec IO Buffer Zero-Copy)
// -----------------------------------------------------------------------------

func BenchmarkEncrypt_NotreLib_Writer(b *testing.B) {
	key := make([]byte, local.KeyBytes)
	_, _ = rand.Read(key)

	for _, bm := range benchPayloadSizes {
		if testing.Short() && bm.size > 16*1024*1024 {
			continue
		}
		payload := generatePayload(bm.size)

		b.Run(bm.name, func(b *testing.B) {
			b.SetBytes(int64(bm.size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				var out bytes.Buffer
				out.Grow(bm.size + 1024)
				w := local.NewWriter(&out, key)
				_, _ = w.Write(payload)
				_ = w.Close()
			}
		})
	}
}

// -----------------------------------------------------------------------------
// 3. OPENZITI SECRETSTREAM (Référence Pure Go)
// -----------------------------------------------------------------------------

func BenchmarkEncrypt_OpenZiti(b *testing.B) {
	key := make([]byte, 32)
	_, _ = rand.Read(key)

	for _, bm := range benchPayloadSizes {
		if testing.Short() && bm.size > 16*1024*1024 {
			continue
		}
		payload := generatePayload(bm.size)

		b.Run(bm.name, func(b *testing.B) {
			b.SetBytes(int64(bm.size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				ez, _, err := openziti.NewEncryptor(key)
				if err != nil {
					b.Fatal(err)
				}
				for off := 0; off < len(payload); off += 8192 {
					end := off + 8192
					if end > len(payload) {
						end = len(payload)
					}
					tag := byte(openziti.TagMessage)
					if end == len(payload) {
						tag = byte(openziti.TagFinal)
					}
					_, _ = ez.Push(payload[off:end], tag)
				}
			}
		})
	}
}

// -----------------------------------------------------------------------------
// 4. GO STANDARD XCHACHA20-POLY1305 (x/crypto/chacha20poly1305)
// -----------------------------------------------------------------------------

func BenchmarkEncrypt_GoStandardAEAD(b *testing.B) {
	key := make([]byte, chacha20poly1305.KeySize)
	_, _ = rand.Read(key)
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		b.Fatal(err)
	}

	for _, bm := range benchPayloadSizes {
		if testing.Short() && bm.size > 16*1024*1024 {
			continue
		}
		payload := generatePayload(bm.size)
		nonce := make([]byte, aead.NonceSize())
		_, _ = rand.Read(nonce)

		b.Run(bm.name, func(b *testing.B) {
			b.SetBytes(int64(bm.size))
			b.ReportAllocs()
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				dst := make([]byte, 0, len(payload)+aead.Overhead())
				_ = aead.Seal(dst, nonce, payload, nil)
			}
		})
	}
}
