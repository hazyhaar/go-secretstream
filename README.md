# SecretStream55 (`pkg/secretstream55`)

> **High-Performance Pure Go Streaming AEAD Encryption (XChaCha20-Poly1305)**  
> *Zero CGO • Direct `c2simd` SIMD256 Integration • Libsodium C Framing Support*

---

## 1. Overview & Architecture

`secretstream55` is a streaming AEAD encryption and decryption package for Go applications. It wraps chunked streaming encryption (`NewEncryptor`, `NewDecryptor`) around standard `io.Writer` and `io.Reader` interfaces.

* **AEAD Core:** Powered directly by `c2simd.AEADLockDst` and `c2simd.AEADUnlockDst` (**SIMD256 Pure Go**, `CGO_ENABLED=0`).
* **Framing Protocol:** Compatible with Libsodium C `crypto_secretstream_xchacha20poly1305` tag constants (`TagMessage`, `TagPush`, `TagRekey`, `TagFinal`).
* **Performance:** **1 094 Mo/s** single-pass engine throughput, **494 Mo/s** full-duplex stream roundtrip (encrypt + decrypt + I/O).
* **Memory Guarantees:** **0 volatile allocation per chunk** on the warm stream path (`Write` / `Read`).

---

## 2. Benchmark Verification & Performance Matrix (Intel Core i9-14900K)

| Benchmark Target | Engine / Implementation | Throughput (1 MB) | Allocations per Op | Architectural Model |
| :--- | :--- | :--- | :--- | :--- |
| **`c2simd` Engine Lock** | `c2simd.AEADLockDst` | **1 224 Mo/s** | **0 B/op (0 alloc)** | Fused SIMD256 L1 Cache Pass |
| **`secretstream55` Duplex** | `Encryptor` + `Decryptor` | **494 Mo/s** | 24 allocs (Total Stream Setup) | Stream Chunk Framing + AEAD |
| **Monocypher C Transpiled** | `monocypher.c` (`ccgo`) | 141 Mo/s | 17 allocs/op | Transpiled Scalar C (`libc.TLS`) |

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
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	var cipherBuf bytes.Buffer
	enc, err := secretstream55.NewEncryptor(&cipherBuf, key)
	if err != nil {
		log.Fatal(err)
	}
	enc.Write([]byte("Highly confidential stream payload"))

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
