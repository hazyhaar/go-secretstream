# Export map

**Source of truth (development):** `/devhoros/wal-g/internal/crypto/libsodium/`  
**Exportable module (this repo):** `/data/GITREMOTE/go-secretstream`

| wal-g (`package libsodium`) | go-secretstream (`package secretstream`) |
|-----------------------------|------------------------------------------|
| `secretstream.go` | `secretstream.go` (`chunkSize` → exported `ChunkSize`) |
| `reader.go` / `writer.go` | same (stdlib errors, no `pkg/errors`) |
| `crypter.go` | returns `*Crypter` (no `crypto.Crypter` interface) |
| `keytransform.go` | `KeyTransform` exported |
| `*_test.go` | same coverage |

Do **not** develop only in GITREMOTE. Fix in wal-g, run:

```bash
GOEXPERIMENT=jsonv2 GOWORK=off CGO_ENABLED=0 go test ./internal/crypto/libsodium/
# then re-export (manual or script) into this tree and go test ./...
```
