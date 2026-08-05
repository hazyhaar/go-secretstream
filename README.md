# go-secretstream

**Pure Go streaming AEAD (XChaCha20-Poly1305)**  
Zero CGO • SIMD via `golang.org/x/crypto` • ~1.06 GB/s on 64 KB chunks

```bash
go get github.com/hazyhaar/go-secretstream
```

## API

- `NewEncryptor` / `NewDecryptor` — length-prefixed chunked stream (sequence counter as AD)
- `NewLibsodiumEncryptor` / `NewLibsodiumDecryptor` — trailing libsodium-style tag framing

Cipher: `golang.org/x/crypto/chacha20poly1305.NewX` (official Go, AVX2/NEON).

## Why not Monocypher/ccgo?

An early prototype used Monocypher transpiled via ccgo. Benches (kept under `internal/monocypher` for comparison) showed native `x/crypto` ~6.7× faster for AEAD. Production path is official Go only.

| Engine | 64 KB | Throughput |
| :--- | :--- | :--- |
| Monocypher (ccgo) | ~412 µs | ~159 MB/s |
| `golang.org/x/crypto` | ~62 µs | **~1.06 GB/s** |

## Example

```go
package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"github.com/hazyhaar/go-secretstream"
)

func main() {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	enc, err := secretstream.NewEncryptor(&buf, key)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := enc.Write([]byte("confidential stream")); err != nil {
		log.Fatal(err)
	}
	dec, err := secretstream.NewDecryptor(&buf, key)
	if err != nil {
		log.Fatal(err)
	}
	plain, err := io.ReadAll(dec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", plain)
}
```

## Bench

```bash
CGO_ENABLED=0 go test -bench=. -benchmem ./...
```

## License

See `LICENSE` / `LICENSE-MIT` / `NOTICE` (upstream lineage where applicable).
