# SecretStream55 (`pkg/secretstream55`)

**High-Performance Pure Go Streaming AEAD Encryption (XChaCha20-Poly1305)**  
*Zero CGO • Hardware SIMD Vector Acceleration (AVX2/NEON) • 1.06 GB/s Encryption Throughput*

---

## 1. Overview & Architecture

`secretstream55` is a production-grade streaming AEAD encryption/decryption package for Go applications. It provides chunked streaming encryption (`NewEncryptor`, `NewDecryptor`) wrapped over standard `io.Writer` and `io.Reader` interfaces.

* **AEAD Cipher:** **XChaCha20-Poly1305** (24-byte Nonce, 32-byte Key, 16-byte Poly1305 MAC tag).
* **Hardware Acceleration:** Native Go SIMD assembly (`golang.org/x/crypto/chacha20poly1305`).
* **Performance:** **1.06 GB/s** encryption throughput on standard 64 KB stream chunks.
* **Security Guarantees:** Strict chunk sequence counter (AD) prevents chunk reordering, dropping, or replay attacks.

---

## 2. Engineering Migration Note & Benchmark Rationale

During initial development, `secretstream55` was prototyped by transpiling the C library **Monocypher** (`monocypher.c`) into Pure Go via `modernc.org/ccgo/v4`. 

While transpiled C code proved to be a massive performance victory for complex graphics/PDF engines (such as `stb_truetype` or `stb_image` in `pdfast55`), a comparative benchmark revealed that native Go cryptographic primitives (`golang.org/x/crypto`) significantly outperform transpiled C for AEAD routines:

| Encryption Engine | Execution Time (64 KB) | Throughput | Allocations per Op | Architectural Model |
| :--- | :--- | :--- | :--- | :--- |
| **Monocypher Transpiled (C via ModernC)** | 411 687 ns | 159.19 MB/s | 17 allocs/op | Scalar C translation (`libc.TLS`, `unsafe`) |
| **Native Go SIMD (`golang.org/x/crypto`)** | **61 748 ns** | **1,061.34 MB/s (1.06 GB/s)** | **2 allocs/op** | **Hardware SIMD Assembly (AVX2/NEON)** |

**Decision:** `secretstream55` was migrated to native `golang.org/x/crypto/chacha20poly1305` (`NewX`), yielding a **6.7× throughput increase** while eliminating all third-party C/libc wrappers.

---

## 3. Usage Example

```go
package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"code.hazyhaar.fr/devhoros/pkg/secretstream55"
)

func main() {
	// 32-byte secret key
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	// Encrypt stream
	var cipherBuf bytes.Buffer
	enc, err := secretstream55.NewEncryptor(&cipherBuf, key)
	if err != nil {
		log.Fatal(err)
	}
	enc.Write([]byte("Highly confidential stream payload"))

	// Decrypt stream
	dec, err := secretstream55.NewDecryptor(&cipherBuf, key)
	if err != nil {
		log.Fatal(err)
	}

	plainText, err := io.ReadAll(dec)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decrypted payload: %s\n", string(plainText))
}
```

---

## 4. Benchmark Verification

To re-run the benchmark suite locally:

```bash
GOWORK=off CGO_ENABLED=0 go test -bench=. -benchmem .
```
