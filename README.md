# go-secretstream

[![Go Reference](https://pkg.go.dev/badge/github.com/hazyhaar/go-secretstream.svg)](https://pkg.go.dev/github.com/hazyhaar/go-secretstream)
[![License: BSD-3-Clause](https://img.shields.io/badge/License-BSD--3--Clause-blue.svg)](LICENSE)
[![Pure Go: CGO Free](https://img.shields.io/badge/Pure_Go-CGO_Free-success.svg)](#architecture)

High-performance, zero-allocation streaming encryption and cryptographic suite in **Pure Go 1.27**, accelerated by native SIMD (`simd/archsimd` AVX2 256-bit).

---

## 🌟 Overview & Architecture

`go-secretstream` delivers high-throughput cryptography without relying on hand-written Plan 9 assembly (`.s`) or external CGO bindings:

1. **`secretstream`** : Zero-allocation streaming encryption interoperable with Libsodium (`crypto_secretstream_xchacha20poly1305`), WAL-G, and PHP `ext-sodium`.
2. **`chacha20poly1305`** : Standard `crypto/cipher.AEAD` implementation (RFC 8439 IETF & XChaCha20-Poly1305) in Pure Go.
3. **`internal/monocypher`** : Full [Monocypher 4.0.2](https://monocypher.org) engine (ChaCha20, Poly1305, Argon2i, Ed25519, X25519, Blake2b, Elligator 2).

---

## 🚀 Performance & Benchmarks (Intel Core i9-14900K)

Benchmarks executed under `GOTOOLCHAIN=go1.27rc3` + `GOEXPERIMENT=simd` :

| Component / Cipher | Throughput / Latency | Heap Allocations | Status / Notes |
| :--- | :--- | :--- | :--- |
| **ChaCha20 AVX2 SIMD (1MB)** | **1 932.90 MB/s** | **0 B/op (0 allocs)** | Pure Go `archsimd.Uint32x8` |
| **AEAD LockDst (64KB)** | **1 024.40 MB/s** | **0 B/op (0 allocs)** | Zero-Allocation API |
| **AEAD LockDst (1KB)** | **801.04 MB/s** | **0 B/op (0 allocs)** | Zero-Allocation API |
| **Blake2b-512 (1KB)** | **978.51 MB/s** | **0 B/op (0 allocs)** | Zero-Allocation API |
| **Argon2i ($m=8, p=1$)** | **16.13 µs / op** | **0 B/op (0 allocs)** | Hardened memory layout |
| **Ed25519 Scalarbase** | **37.54 µs / op** | **0 B/op (0 allocs)** | Edwards25519 arithmetic |
| **X25519 DH** | **64.50 µs / op** | **0 B/op (0 allocs)** | Curve25519 key exchange |
| **Elligator 2 Map** | **5.62 µs / op** | **0 B/op (0 allocs)** | Constant-time mapping |

---

## 🔒 Security & Verification Guarantees

- **13/13 Bit-Exact C Oracles :** Formally verified against Monocypher 4.0.2 and Libsodium C reference test suites (`-O0`).
- **Differential Fuzzing :** Continuous fuzzing harness (`go test -fuzz`) validating over 960,000 randomized executions without failure.
- **Hardware Fallback & Portability :** Dynamic CPU detection via `archsimd.X86.AVX2()` with fallback on non-AVX2 hardware and non-x86 architectures (`arm64`, `riscv64`).
- **Memory Hygiene :** Constant-time arithmetic and automatic buffer erasure via native `clear()` (`DUFFZERO`/`REP STOSQ`).

---

## 📦 Usage Examples

### 1. Libsodium-Compatible Streaming Encryption

```go
package main

import (
	"bytes"
	"io"
	"log"

	"github.com/hazyhaar/go-secretstream"
)

func main() {
	key := make([]byte, 32) // 32-byte secret key

	// Encrypt
	var cipherStream bytes.Buffer
	enc, err := secretstream.NewLibsodiumEncryptor(&cipherStream, key)
	if err != nil {
		log.Fatal(err)
	}
	enc.Write([]byte("Sensitive streaming payload..."))
	enc.Close() // Emits TAG_FINAL

	// Decrypt
	dec, err := secretstream.NewLibsodiumDecryptor(&cipherStream, key)
	if err != nil {
		log.Fatal(err)
	}
	plainText, err := io.ReadAll(dec)
	if err != nil {
		log.Fatal(err)
	}
	println(string(plainText))
}
```

### 2. Standard `cipher.AEAD` (RFC 8439)

```go
package main

import (
	"fmt"
	"log"

	"github.com/hazyhaar/go-secretstream/chacha20poly1305"
)

func main() {
	key := make([]byte, 32)
	nonce := make([]byte, chacha20poly1305.NonceSize) // 12 bytes

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		log.Fatal(err)
	}

	sealed := aead.Seal(nil, nonce, []byte("Hello Go 1.27 SIMD"), []byte("auth_data"))
	opened, err := aead.Open(nil, nonce, sealed, []byte("auth_data"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decrypted: %s\n", opened)
}
```

---

## 📄 License

Licensed under the [BSD-3-Clause License](LICENSE). Compatible with Go standard library contributions.
