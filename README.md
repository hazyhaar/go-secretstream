# go-secretstream

Pure-Go implementation of **libsodium `crypto_secretstream_xchacha20poly1305`**, wire-compatible with the C library and with [WAL-G](https://github.com/wal-g/wal-g) framing (8192-byte plaintext chunks, `TAG_FINAL` on last flush).

- **No CGO**, no libsodium shared library
- Runtime dep: `golang.org/x/crypto` only
- Optional cross-oracle tests via PyNaCl (`SECRETSTREAM_ORACLE` or system `python3` + `nacl`)

## Install

```bash
go get github.com/lateos-ai/go-secretstream@latest
```

## Usage

```go
package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lateos-ai/go-secretstream"
)

func main() {
	key := bytes.Repeat([]byte{0x42}, secretstream.KeyBytes)
	var buf bytes.Buffer

	w := secretstream.NewWriter(&buf, key)
	_, _ = w.Write([]byte("hello"))
	_ = w.Close() // emits TAG_FINAL

	plain, err := io.ReadAll(secretstream.NewReader(bytes.NewReader(buf.Bytes()), key))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(plain))
}
```

High-level crypter (inline key or file, hex/base64/none transforms — WAL-G env parity):

```go
c := secretstream.CrypterFromKey(hexKey, secretstream.KeyTransformHex)
enc, _ := c.Encrypt(dst)
// ...
dec, _ := c.Decrypt(src)
```

## Wire format

| Field | Size |
|-------|------|
| Header | 24 B |
| Chunk | 1 tag byte + ciphertext + 16 B Poly1305 MAC (`ABytes = 17` overhead) |
| Full plaintext chunk | 8192 B (`ChunkSize`) before last |
| Last chunk | `TAG_FINAL` (0x03), any remaining length including 0 |

Core `push`/`pull` match libsodium C bit-for-bit; `Reader`/`Writer` add WAL-G-style framing.

## Test

```bash
go test ./...
go test -count=1 -short ./...   # skip 2 MiB stream
# optional C↔Go cross (PyNaCl):
export SECRETSTREAM_ORACLE=/path/to/python-with-nacl
go test -count=1 -run Cross ./...
```

## License

- Inherited WAL-G crypter surfaces: **Apache-2.0** (`LICENSE`)
- Pure-Go secretstream core & packaging: **MIT** (`LICENSE-MIT`)

See `NOTICE`.

## Origin

Extracted from the Lateos WAL-G fork (`internal/crypto/libsodium`) after replacing the CGO libsodium binding with a pure-Go secretstream. Standalone so it can be reviewed, vendored, or reused outside WAL-G.
