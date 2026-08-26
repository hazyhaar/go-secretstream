# go-secretstream

[![Go Reference](https://pkg.go.dev/badge/github.com/hazyhaar/go-secretstream.svg)](https://pkg.go.dev/github.com/hazyhaar/go-secretstream)
[![License](https://img.shields.io/badge/License-Apache_2.0_|_MIT-blue.svg)](LICENSE)
[![Pure Go](https://img.shields.io/badge/Pure_Go-CGO%3D0-brightgreen.svg)]()
[![Zero Allocation](https://img.shields.io/badge/Memory-0_Allocs%2Fop-orange.svg)]()

High-performance, memory-safe streaming authenticated encryption (**XChaCha20-Poly1305**) in Pure Go (`CGO=0`).

`go-secretstream` delivers chunked, constant-overhead authenticated encryption for unbounded streams (`io.Reader` / `io.Writer`), featuring hardware-accelerated SIMD kernels, strict anti-truncation defenses, domain-separated AAD authentication, and bit-exact interoperability with C `libsodium`.

---

## Key Highlights

* **Pure Go & CGO=0** : Completely standalone, cross-platform, zero CGO dependencies.
* **SIMD Hardware Acceleration** : Powered by `sgoiter`-transpiled AVX2 SIMD fused kernels (`c2fused`), achieving **>70% of hand-tuned Plan 9 assembly speed** and up to **36× faster than standard C-to-Go transpilation (`ccgo`)**.
* **Zero-Allocation Hot Paths** : `0 allocs/op` and `0 B/op` during stream encryption and decryption.
* **Format V2 Framing** : Versioned header, injective domain-separated authenticated data, and mandatory `TagFinal` terminal frame for robust stream truncation defense.
* **Hybrid Decoder** : Seamless, transparent backward compatibility with legacy v1 stream archives.
* **Libsodium Wire Interoperability** : First-class support for `crypto_secretstream_xchacha20poly1305` wire format, bit-exact and interoperable with C Libsodium and WAL-G.
* **Hardened Memory Safety** : Rigorous in-place anti-aliasing guards, sequence counter overflow detection, and compiler-resistant secret zeroing on `Close()`.

---

## Performance Benchmarks

Measured on **Intel® Core™ i9-14900K** (Ubuntu Linux 6.8, Go 1.27 with `GOEXPERIMENT=simd`):

### AEAD Seal Throughput (XChaCha20-Poly1305)

| Payload Size | `ccgo` (Scalar C-to-Go) | `go-secretstream` (Pure Go AVX2) | `x/crypto` (Plan 9 Assembly) | Speedup vs `ccgo` | % Assembly Speed |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **64 B** | 3,200 ns (20.0 MB/s) | **228.6 ns (280.0 MB/s)** | 170.8 ns (374.8 MB/s) | **14.0×** | **74.7 %** |
| **1 KB** | 17,289 ns (59.2 MB/s) | **809.6 ns (1,264.8 MB/s)** | 492.7 ns (2,078.4 MB/s) | **21.4×** | **60.8 %** |
| **64 KB** | 1,080 µs (60.7 MB/s) | **29.4 µs (2,223.8 MB/s)** | 20.8 µs (3,144.5 MB/s) | **36.7×** | **70.7 %** |
| **1 MB** | 17.2 ms (60.9 MB/s) | **469.2 µs (2,234.8 MB/s)** | 332.0 µs (3,157.9 MB/s) | **36.7×** | **70.8 %** |

### Memory & Allocation Profile

```
BenchmarkCompare_AEAD_ZeroAlloc/Seal/purego_avx2/1KB-32   809.6 ns/op   1264.81 MB/s   0 B/op   0 allocs/op
BenchmarkCompare_AEAD_ZeroAlloc/Open/purego_avx2/1KB-32   828.1 ns/op   1236.55 MB/s   0 B/op   0 allocs/op
```

* **Heap Allocations** : **0 allocs/op** on both encryption and decryption streaming paths.
* **Bounds-Check Elimination** : 80% reduction in dead branching paths (`panicBounds` reduced from 640 to 126 calls), minimizing instruction cache and frontend decoder pressure.

---

## Installation

```bash
go get github.com/hazyhaar/go-secretstream
```

---

## Operational Wire Modes

| Mode | Constructor | Wire Protocol | Target Use Case |
| :--- | :--- | :--- | :--- |
| **Format V2 (Standard)** | `NewEncryptor` / `NewDecryptor` | **SS55-v2** Chunked Framing (BE4 Length + Tag + Ciphertext + MAC) | Modern Go streaming pipelines, anti-truncation defenses, cloud storage archives. |
| **Libsodium Interop** | `NewLibsodiumEncryptor` / `NewLibsodiumDecryptor` | `crypto_secretstream_xchacha20poly1305` | Interoperability with C `libsodium` toolchains, WAL-G PostgreSQL backups. |

---

## Quickstart

### 1. Standard V2 Streaming Encryption (`io.Writer`)

```go
package main

import (
	"crypto/rand"
	"os"

	"github.com/hazyhaar/go-secretstream"
)

func main() {
	// 32-byte secret key
	key := make([]byte, secretstream55.KeyBytes)
	rand.Read(key)

	destFile, err := os.Create("encrypted.stream")
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	// Initialize V2 streaming encryptor
	enc, err := secretstream55.NewEncryptor(destFile, key)
	if err != nil {
		panic(err)
	}

	// Write arbitrary stream chunks
	enc.Write([]byte("Hello, secure streaming world!"))
	enc.Write([]byte(" Chunk 2 payload..."))

	// CRITICAL: Close() emits the authenticated TagFinal frame and wipes keys
	if err := enc.Close(); err != nil {
		panic(err)
	}
}
```

### 2. Standard V2 Streaming Decryption (`io.Reader`)

```go
package main

import (
	"io"
	"os"

	"github.com/hazyhaar/go-secretstream"
)

func main() {
	key := []byte("...32-byte secret key here...")

	sourceFile, err := os.Open("encrypted.stream")
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	// Decryptor transparently handles both V2 streams and legacy V1 archives
	dec, err := secretstream55.NewDecryptor(sourceFile, key)
	if err != nil {
		panic(err)
	}

	plaintext, err := io.ReadAll(dec)
	if err != nil {
		// Truncation or authentication errors will be caught here
		panic(err)
	}

	println(string(plaintext))
}
```

### 3. Libsodium Interoperability Mode

```go
// Direct compatibility with libsodium crypto_secretstream_xchacha20poly1305
enc, err := secretstream55.NewLibsodiumEncryptor(destWriter, key)
if err != nil {
    return err
}
_, err = enc.Write(payload)
err = enc.Close() // Emits Libsodium TAG_FINAL

// Read with Libsodium compatibility
dec, err := secretstream55.NewLibsodiumDecryptor(sourceReader, key)
```

---

## Wire Format V2 Specification

The V2 format resolves truncation ambiguities and injective domain separation.

### Header Structure (36 Bytes)

```
+------------------+---------------+----------------+----------------------+
|  Magic (8 Bytes) | Ver (2 Bytes) | Flags (2 Bytes)|   Nonce (24 Bytes)   |
|   "SS55-v2\0"    |    0x0002     |     0x0000     |  Random XChaCha20 IV |
+------------------+---------------+----------------+----------------------+
```

1. **Magic (`8 bytes`)** : `0x53 0x53 0x35 0x35 0x2D 0x76 0x32 0x00` (`SS55-v2\0`).
2. **Version (`2 bytes, BE`)** : `0x0002`.
3. **Flags (`2 bytes, BE`)** : Reserved (`0x0000`).
4. **Nonce (`24 bytes`)** : 192-bit cryptographic random initialization vector.

### Frame Structure

Each frame within the stream is encapsulated as:

```
+--------------------+--------------+---------------------+-------------------+
| Length (4 Bytes BE)| Tag (1 Byte) | Ciphertext (N Bytes)|  MAC (16 Bytes)   |
+--------------------+--------------+---------------------+-------------------+
```

* **Length (`uint32_be`)** : Total frame payload size: $1 + N + 16$ bytes (Tag + Ciphertext + MAC).
* **Tag (`1 byte`)** :
  * `0x00` : Standard message frame (`TagMessage`).
  * `0x03` : Terminal final frame (`TagFinal`). Must contain an empty ciphertext ($N = 0$, length = $17$).
  * Any other tag value is rejected with a fatal authentication error.
* **Ciphertext (`N bytes`)** : Payload encrypted via XChaCha20 with per-frame subkey derivation.
* **MAC (`16 bytes`)** : Poly1305 authenticator.

### Injective Domain-Separated AAD

To prevent cross-stream manipulation and parameter ambiguity, the Associated Authenticated Data (AAD) for each chunk is injectively constructed as:

$$\text{AAD} = \text{Magic (8B)} \parallel \text{Seq (4B BE)} \parallel \text{Tag (1B)} \parallel \text{Len}(\text{UserAD}) \text{ (4B BE)} \parallel \text{UserAD}$$

### Anti-Truncation & Close Semantics

* An encryptor **must** call `Close()`. This emits a zero-length `TagFinal` frame and immediately zeroes internal state.
* The decryptor **only returns `io.EOF`** after successfully verifying this authenticated `TagFinal` frame.
* Any stream termination prior to the `TagFinal` frame returns a sticky `io.ErrUnexpectedEOF` (preventing silent truncation attacks).

---

## Testing & Verification

Run the comprehensive test suite, including bit-exact cross-language oracles:

```bash
# Standard test suite
go test -v ./...

# Hardware SIMD test suite (Go 1.27+)
GOEXPERIMENT=simd go test -count=1 ./...

# Libsodium C cross-verification (requires libsodium-dev)
make interop-test
```

---

## Formal Schema Verification

The V2 wire format is formally verified against closed CUE specifications located in [`spec/format_v2.cue`](spec/format_v2.cue):

```bash
cue vet ./spec/...
```

---

## Contributors & Acknowledgements

* **Hazyhaar** ([@hazyhaar](https://github.com/hazyhaar)) — System architecture, core design & project maintainer.
* **Gemini** (Google DeepMind) — Adversarial verification, security audits, and testing harnesses.
* **Grok** (xAI) — Low-level protocol robustness, fuzzing, and boundary condition auditing.
* **Claude** (Anthropic) — Go 1.27 SIMD transpilation passes and formal CUE specifications.

---

## License

This project is licensed under dual **Apache-2.0** OR **MIT** terms at your option. See [LICENSE](LICENSE) and [LICENSE-MIT](LICENSE-MIT) for details.
