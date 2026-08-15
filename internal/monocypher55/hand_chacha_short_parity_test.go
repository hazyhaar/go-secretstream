package monocypher55

import (
	"bytes"
	"crypto/sha256"
	"testing"
)

// détPseudo remplit dst de façon déterministe à partir d'une graine textuelle.
func detPseudo(dst []byte, seed string) {
	h := sha256.Sum256([]byte(seed))
	for i := range dst {
		if i%32 == 0 && i > 0 {
			h = sha256.Sum256(h[:])
		}
		dst[i] = h[i%32]
	}
}

// TestHChaCha20SIMDParity vérifie la parité bit-exacte entre la voie SIMD 128-bit
// et le corps scalaire émis, sur des clés et entrées variées.
func TestHChaCha20SIMDParity(t *testing.T) {
	if !hasAVX2() {
		t.Skip("AVX2 indisponible : la voie SIMD n'est pas exercée")
	}
	var key [32]byte
	var in [16]byte
	var outSIMD, outScal [32]byte
	for i := 0; i < 200; i++ {
		detPseudo(key[:], "hchacha-key-"+string(rune(i)))
		detPseudo(in[:], "hchacha-in-"+string(rune(i*7)))
		hchacha20_simd128(outSIMD[:], key[:], in[:])
		crypto_chacha20_h_scalar(outScal[:], key[:], in[:])
		if !bytes.Equal(outSIMD[:], outScal[:]) {
			t.Fatalf("divergence HChaCha20 SIMD vs scalaire, itération %d\nsimd=%x\nscal=%x", i, outSIMD, outScal)
		}
	}
}

// TestDeriv2SIMDParity vérifie que la dérivation 2 blocs en une passe 256-bit
// produit exactement le keystream du chemin scalaire, compteurs variés inclus
// (bords de wrap 32-bit du mot bas du compteur).
func TestDeriv2SIMDParity(t *testing.T) {
	if !hasAVX2() {
		t.Skip("AVX2 indisponible : la voie SIMD n'est pas exercée")
	}
	var key [32]byte
	var nonce [8]byte
	ctrs := []uint64{0, 1, 2, 41, 0xFFFFFFFE, 0xFFFFFFFF, 1 << 32, (1 << 32) + 3, 0xFFFFFFFFFFFFFFFE, 12345678901234567}
	var simdBuf, scalBuf [128]byte
	for i := 0; i < 30; i++ {
		detPseudo(key[:], "deriv2-key-"+string(rune(i)))
		detPseudo(nonce[:], "deriv2-nonce-"+string(rune(i*13)))
		for _, ctr := range ctrs {
			chacha20_deriv2_simd256(simdBuf[:], key[:], nonce[:], ctr)
			Crypto_chacha20_djb(scalBuf[:], nil, 128, key[:], nonce[:], ctr)
			if !bytes.Equal(simdBuf[:], scalBuf[:]) {
				t.Fatalf("divergence deriv2 SIMD vs scalaire, key#%d ctr=%d\nsimd=%x\nscal=%x", i, ctr, simdBuf, scalBuf)
			}
		}
	}
}

// refFastPathWrite rejoue le fast-path AEAD court avec les seules primitives
// scalaires (comportement d'avant vectorisation) : oracle de parité complet.
func refFastPathWrite(ctx *Crypto_aead_ctx, cipher, mac, ad, plain []byte) {
	var buf [128]byte
	derivSize := uint64(64)
	if len(plain) > 0 {
		derivSize = 128
	}
	Crypto_chacha20_djb(buf[:], nil, derivSize, ctx.Key[:], ctx.Nonce[:], ctx.Counter)
	for i := range plain {
		cipher[i] = plain[i] ^ buf[64+i]
	}
	Lock_auth(mac, buf[:32], ad, uint64(len(ad)), cipher, uint64(len(plain)))
	copy(ctx.Key[:], buf[32:64])
	if len(plain) > 0 {
		ctx.Counter += 2
	} else {
		ctx.Counter += 1
	}
}

// TestAEADShortSIMDvsScalarAllSizes couvre TOUTES les tailles 0..64 avec nonces
// et compteurs variés : le fast-path vectorisé (write puis read) doit reproduire
// bit à bit l'oracle scalaire, y compris l'état du contexte après opération.
func TestAEADShortSIMDvsScalarAllSizes(t *testing.T) {
	var key [32]byte
	var nonce [8]byte
	ctrs := []uint64{0, 1, 7, 0xFFFFFFFF, 1 << 32}
	for size := 0; size <= 64; size++ {
		for ci, ctr0 := range ctrs {
			detPseudo(key[:], "aead-key-"+string(rune(size))+string(rune(ci)))
			detPseudo(nonce[:], "aead-nonce-"+string(rune(size*3+ci)))
			plain := make([]byte, size)
			ad := make([]byte, (size*5+ci)%40)
			detPseudo(plain, "aead-pt")
			detPseudo(ad, "aead-ad")

			mkCtx := func() Crypto_aead_ctx {
				var c Crypto_aead_ctx
				copy(c.Key[:], key[:])
				copy(c.Nonce[:], nonce[:])
				c.Counter = ctr0
				return c
			}

			// Voie testée (SIMD si dispo)
			ctxA := mkCtx()
			cipherA := make([]byte, size)
			var macA [16]byte
			Crypto_aead_write(&ctxA, cipherA, macA[:], ad, uint64(len(ad)), plain, uint64(size))

			// Oracle scalaire
			ctxB := mkCtx()
			cipherB := make([]byte, size)
			var macB [16]byte
			refFastPathWrite(&ctxB, cipherB, macB[:], ad, plain)

			if !bytes.Equal(cipherA, cipherB) || !bytes.Equal(macA[:], macB[:]) {
				t.Fatalf("divergence write size=%d ctr=%d\nct simd=%x\nct scal=%x\nmac simd=%x mac scal=%x",
					size, ctr0, cipherA, cipherB, macA, macB)
			}
			if ctxA.Key != ctxB.Key || ctxA.Counter != ctxB.Counter {
				t.Fatalf("divergence d'état ctx size=%d ctr=%d (rekey/compteur)", size, ctr0)
			}

			// Read : déchiffrement authentifié depuis un contexte frais
			ctxC := mkCtx()
			plainC := make([]byte, size)
			if diff := Crypto_aead_read(&ctxC, plainC, macA[:], ad, uint64(len(ad)), cipherA, uint64(size)); diff != 0 {
				t.Fatalf("read: tag rejeté à tort size=%d ctr=%d", size, ctr0)
			}
			if !bytes.Equal(plainC, plain) {
				t.Fatalf("read: plaintext divergent size=%d ctr=%d", size, ctr0)
			}
			if ctxC.Key != ctxB.Key || ctxC.Counter != ctxB.Counter {
				t.Fatalf("read: divergence d'état ctx size=%d ctr=%d", size, ctr0)
			}
		}
	}
}
