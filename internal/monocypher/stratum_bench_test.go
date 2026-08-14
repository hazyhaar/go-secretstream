package monocypher_test

import (
	"bytes"
	"crypto/sha256"
	"testing"
	"unsafe"

	sgoi "github.com/hazyhaar/go-secretstream/internal/monocypher"
)

// Strate benches: AEAD / Blake / X25519 / EdDSA / Argon2 / Elligator.
// Run: go test -bench=BenchmarkStratum -benchmem ./...

func BenchmarkStratum_AEAD_Lock_64(b *testing.B) {
	key, nonce := keyNonce()
	pt := mkPT(64, "i%251")
	ad := []byte("ad")
	b.SetBytes(64)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_, _, _ = sgoi.AEADLock(key, nonce, ad, pt)
	}
}

func BenchmarkStratum_AEAD_Lock_1K(b *testing.B) {
	key, nonce := keyNonce()
	pt := mkPT(1024, "(i*17+3)%251")
	ad := []byte("HEADER")
	b.SetBytes(1024)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_, _, _ = sgoi.AEADLock(key, nonce, ad, pt)
	}
}

func BenchmarkStratum_AEAD_Lock_64K(b *testing.B) {
	key, nonce := keyNonce()
	pt := mkPT(64*1024, "i%251")
	b.SetBytes(64 * 1024)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		_, _, _ = sgoi.AEADLock(key, nonce, nil, pt)
	}
}

func BenchmarkStratum_AEAD_LockDst_64(b *testing.B) {
	key, nonce := keyNonce()
	pt := mkPT(64, "i%251")
	ad := []byte("ad")
	dst := make([]byte, len(pt))
	var mac [16]byte
	b.SetBytes(64)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		if err := sgoi.LockDst(dst, mac[:], key, nonce, ad, pt); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStratum_AEAD_LockDst_1K(b *testing.B) {
	key, nonce := keyNonce()
	pt := mkPT(1024, "(i*17+3)%251")
	ad := []byte("HEADER")
	dst := make([]byte, len(pt))
	var mac [16]byte
	b.SetBytes(1024)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		if err := sgoi.LockDst(dst, mac[:], key, nonce, ad, pt); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStratum_AEAD_LockDst_64K(b *testing.B) {
	key, nonce := keyNonce()
	pt := mkPT(64*1024, "i%251")
	dst := make([]byte, len(pt))
	var mac [16]byte
	b.SetBytes(64 * 1024)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		if err := sgoi.LockDst(dst, mac[:], key, nonce, nil, pt); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStratum_Blake2b_1K(b *testing.B) {
	msg := mkPT(1024, "i%251")
	var h [64]byte
	b.SetBytes(1024)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		sgoi.Crypto_blake2b(h[:], 64, msg, uint64(len(msg)))
	}
}

func BenchmarkStratum_X25519_DH(b *testing.B) {
	sk := bytes.Repeat([]byte{7}, 32)
	var pk, out [32]byte
	sgoi.Crypto_x25519_public_key(pk[:], sk)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		sgoi.Crypto_x25519(out[:], sk, pk[:])
	}
}

func BenchmarkStratum_EdDSA_Sign_64(b *testing.B) {
	seed := bytes.Repeat([]byte{9}, 32)
	var sk [64]byte
	var pk [32]byte
	sgoi.Crypto_eddsa_key_pair(sk[:], pk[:], append([]byte(nil), seed...))
	msg := mkPT(64, "i%251")
	var sig [64]byte
	b.SetBytes(64)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		sgoi.Crypto_eddsa_sign(sig[:], sk[:], msg, uint64(len(msg)))
	}
}

func BenchmarkStratum_EdDSA_Verify_64(b *testing.B) {
	seed := bytes.Repeat([]byte{9}, 32)
	var sk [64]byte
	var pk [32]byte
	sgoi.Crypto_eddsa_key_pair(sk[:], pk[:], append([]byte(nil), seed...))
	msg := mkPT(64, "i%251")
	var sig [64]byte
	sgoi.Crypto_eddsa_sign(sig[:], sk[:], msg, uint64(len(msg)))
	b.SetBytes(64)
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		if sgoi.Crypto_eddsa_check(sig[:], pk[:], msg, uint64(len(msg))) != 0 {
			b.Fatal("verify")
		}
	}
}

func BenchmarkStratum_EdDSA_Scalarbase(b *testing.B) {
	sc := sha256.Sum256([]byte("bench-sc"))
	sc[0] &= 248
	sc[31] &= 127
	sc[31] |= 64
	var pt [32]byte
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		sgoi.Crypto_eddsa_scalarbase(pt[:], sc[:])
	}
}

func BenchmarkStratum_Argon2i_m8_p1(b *testing.B) {
	const blocks = 8
	workArea := make([]uint64, blocks*128)
	work := unsafe.Slice((*byte)(unsafe.Pointer(&workArea[0])), blocks*1024)
	hash := make([]byte, 32)
	cfg := sgoi.Crypto_argon2_config{Algorithm: sgoi.Crypto_argon2_i, Nb_blocks: blocks, Nb_passes: 1, Nb_lanes: 1}
	in := sgoi.Crypto_argon2_inputs{Pass: []byte("password"), Salt: bytes.Repeat([]byte{1}, 16), Pass_size: 8, Salt_size: 16}
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		sgoi.Crypto_argon2(hash, 32, work, cfg, in, sgoi.Crypto_argon2_no_extras)
	}
}

func BenchmarkStratum_Elligator_KeyPair(b *testing.B) {
	var hidden, sk [32]byte
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		seed := sha256.Sum256([]byte("seed-stratum"))
		sgoi.Crypto_elligator_key_pair(hidden[:], sk[:], seed[:])
	}
}

func BenchmarkStratum_Elligator_Map(b *testing.B) {
	h := sha256.Sum256([]byte("map-bench"))
	var curve [32]byte
	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		sgoi.Crypto_elligator_map(curve[:], h[:])
	}
}
