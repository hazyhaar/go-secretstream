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
| Header | 24 B (not authenticated alone — first chunk MAC detects corruption) |
| Chunk | 1 tag byte + ciphertext + 16 B Poly1305 MAC (`ABytes = 17` overhead) |
| Full plaintext chunk | 8192 B (`ChunkSize`) before last |
| Last chunk | `TAG_FINAL` (0x03), any remaining length including 0 |

Core `push`/`pull` match libsodium C bit-for-bit; `Reader`/`Writer` add WAL-G-style framing.

**WAL-G framing rule:** `TAG_FINAL` on a full-sized wire chunk (`ChunkSize+ABytes`) is rejected (`premature end`). Finalize with a short/empty FINAL (what `Writer.Close` does). Foreign streams that end exactly on a full chunk need a trailing empty FINAL.

## Security notes

- No key/material zeroization after use (normal for Go; no `mlock`).
- Prefer `KeyTransformHex` / `KeyTransformBase64` with full 32-byte keys. `KeyTransformNone` is **legacy WAL-G**: truncates >32 bytes silently and zero-pads short keys (25–31) — reduced entropy / prefix collisions.
- After a MAC failure, abandon the `Reader` (state is not advanced on mismatch).
- `Writer.Close` is idempotent; `Write` after `Close` errors.

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
