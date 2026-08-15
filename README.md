# secretstream55

Streaming encryption in Pure Go (CGO=0).

## Modes

| API | Wire | Crypto |
|-----|------|--------|
| `NewEncryptor` / `NewDecryptor` | **Maison** (BE4 + AEAD) | `engine` : **monocypher55** (défaut — bascule 2026-08-15) ou **c2simd** (`-tags aead_c2simd`) |
| `NewLibsodiumEncryptor` / `NewLibsodiumDecryptor` | **libsodium** `crypto_secretstream_xchacha20poly1305` (wal-g) | `internal/lsstream` (ChaCha20-IETF + Poly1305, bit-compat C) |

Libsodium mode **requires** `Close()` on the encryptor to emit `TAG_FINAL`.

## Tests

```bash
export GOWORK=off
make test              # unit + monocypher_sgoiter + lsstream
make interop-goldens   # needs libsodium-dev
make interop-test      # C↔Go interop
make test-sgoiter      # standard mode with sgoiter AEAD backend
```

## Docs

- Wire libsodium : [`docs/WIRE_LIBSODIUM.md`](docs/WIRE_LIBSODIUM.md)
- Plan V2 : [`TODO_V2_SGOITER_LIBSODIUM.md`](TODO_V2_SGOITER_LIBSODIUM.md)

## Benchmark (standard / c2simd, historique)

See `secretstream_bench_test.go`. Libsodium-wire benches: `internal/lsstream` / wal-g upstream.
