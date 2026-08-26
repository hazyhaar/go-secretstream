// SPDX-License-Identifier: Apache-2.0

package lsstream

import (
	"crypto/rand"
	"testing"

	"github.com/hazyhaar/go-secretstream/internal/monocypher55"
	"golang.org/x/crypto/poly1305"
)

func TestPolyParityVsXCrypto_2000(t *testing.T) {
	rng := rand.Reader
	var key [32]byte
	for i := 0; i < 2000; i++ {
		if _, err := rng.Read(key[:]); err != nil {
			t.Fatal(err)
		}
		n := i % 4097
		msg := make([]byte, n)
		if n > 0 {
			if _, err := rng.Read(msg); err != nil {
				t.Fatal(err)
			}
		}
		step := (i % 17) + 1
		var ctx monocypher55.Crypto_poly1305_ctx
		monocypher55.Crypto_poly1305_init(&ctx, key[:])
		for off := 0; off < n; {
			end := off + step
			if end > n {
				end = n
			}
			chunk := msg[off:end]
			monocypher55.Crypto_poly1305_update(&ctx, chunk, uint64(len(chunk)))
			off = end
		}
		var got [16]byte
		monocypher55.Crypto_poly1305_final(&ctx, got[:])

		var want [16]byte
		poly1305.Sum(&want, msg, &key)
		if got != want {
			t.Fatalf("i=%d n=%d step=%d MAC diverge\ngot  %x\nwant %x", i, n, step, got, want)
		}
	}
}

func BenchmarkLsstreamStratum8K(b *testing.B) {
	const size = 8192
	key := make([]byte, KeyBytes)
	header := make([]byte, HeaderBytes)
	if _, err := rand.Read(key); err != nil {
		b.Fatal(err)
	}
	if _, err := rand.Read(header); err != nil {
		b.Fatal(err)
	}
	plain := make([]byte, size)
	if _, err := rand.Read(plain); err != nil {
		b.Fatal(err)
	}
	wire := make([]byte, 1+size+16)
	st, err := initFromHeader(key, header)
	if err != nil {
		b.Fatal(err)
	}
	k0, nonce0 := st.k, st.nonce
	b.SetBytes(int64(size))
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		st.k = k0
		st.nonce = nonce0
		if err := st.initCipher(); err != nil {
			b.Fatal(err)
		}
		if _, err := st.pushTo(plain, TagMessage, wire); err != nil {
			b.Fatal(err)
		}
	}
}
