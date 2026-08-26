// SPDX-License-Identifier: Apache-2.0 OR MIT

package monocypher55

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"math/bits"
	"testing"

	"golang.org/x/crypto/poly1305"
)

func TestPoly1305SaturationH2_VsXCrypto(t *testing.T) {
	key := bytes.Repeat([]byte{0xff}, 32)
	msg := bytes.Repeat([]byte{0xff}, 1024)
	if peak := polyBlocksH2Peak(key, msg); peak <= 3 {
		t.Fatalf("précondition: h2>3 attendu sur clé/message 0xff de 1024 octets, peak=%d", peak)
	}

	var got [16]byte
	Crypto_poly1305(got[:], msg, uint64(len(msg)), key)

	var want [16]byte
	var k [32]byte
	copy(k[:], key)
	poly1305.Sum(&want, msg, &k)
	if got != want {
		t.Fatalf("MAC diverge sur l'entrée qui sature h2\ngot  %x\nwant %x", got, want)
	}
}

func TestPoly1305ParityVsXCrypto_10k(t *testing.T) {
	rng := rand.Reader
	var key [32]byte
	for i := 0; i < 10000; i++ {
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
		var got [16]byte
		Crypto_poly1305(got[:], msg, uint64(len(msg)), key[:])
		var want [16]byte
		poly1305.Sum(&want, msg, &key)
		if got != want {
			t.Fatalf("i=%d n=%d MAC diverge\ngot  %x\nwant %x", i, n, got, want)
		}
	}
}

func polyBlocksH2Peak(key, msg []byte) uint64 {
	var ctx Crypto_poly1305_ctx
	Crypto_poly1305_init(&ctx, key)
	r0 := uint64(ctx.R[0]) | (uint64(ctx.R[1]) << 32)
	r1 := uint64(ctx.R[2]) | (uint64(ctx.R[3]) << 32)
	var h0, h1, h2, peak uint64
	for off := 0; off+16 <= len(msg); off += 16 {
		m0 := binary.LittleEndian.Uint64(msg[off : off+8])
		m1 := binary.LittleEndian.Uint64(msg[off+8 : off+16])
		var c uint64
		h0, c = bits.Add64(h0, m0, 0)
		h1, c = bits.Add64(h1, m1, c)
		h2 += c + 1

		hi0_0, lo0_0 := bits.Mul64(h0, r0)
		hi1_0, lo1_0 := bits.Mul64(h1, r0)
		hi0_1, lo0_1 := bits.Mul64(h0, r1)
		hi1_1, lo1_1 := bits.Mul64(h1, r1)
		lo2_0 := h2 * r0
		lo2_1 := h2 * r1
		m1_lo, c1 := bits.Add64(lo1_0, lo0_1, 0)
		m1_hi, _ := bits.Add64(hi1_0, hi0_1, c1)
		m2_lo, c2 := bits.Add64(lo2_0, lo1_1, 0)
		m2_hi, _ := bits.Add64(0, hi1_1, c2)
		t0 := lo0_0
		t1, c := bits.Add64(m1_lo, hi0_0, 0)
		t2, c := bits.Add64(m2_lo, m1_hi, c)
		t3, _ := bits.Add64(lo2_1, m2_hi, c)

		h0 = t0
		h1 = t1
		h2 = t2 & 3
		ccLo := t2 & ^uint64(3)
		ccHi := t3
		h0, c = bits.Add64(h0, ccLo, 0)
		h1, c = bits.Add64(h1, ccHi, c)
		h2 += c
		cLo := (ccLo >> 2) | (ccHi << 62)
		cHi := ccHi >> 2
		h0, c = bits.Add64(h0, cLo, 0)
		h1, c = bits.Add64(h1, cHi, c)
		h2 += c
		if h2 > peak {
			peak = h2
		}
		if h2 > 3 {
			h0, c = bits.Add64(h0, (h2>>2)*5, 0)
			h1, c = bits.Add64(h1, 0, c)
			h2 = (h2 & 3) + c
		}
	}
	return peak
}
