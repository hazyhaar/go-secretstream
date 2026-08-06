# go-secretstream (v2)

Pure Go implementation of high-throughput streaming AEAD encryption using **`c2simd`** vector acceleration on Go 1.27 (`GOEXPERIMENT=simd`). Zero CGO, zero `.s` assembly, 100% Go safety.

## Certified Benchmark Results (Intel Core i9-14900K, Go 1.27 SIMD)

| Layer / Benchmark | Throughput | Allocations | Description |
| :--- | :---: | :---: | :--- |
| **`SecretStream55_SteadyState_WriteOnly_1MB`** | **1,134.93 MB/s (1.13 GB/s)** | **`0 B/op` (0 allocs)** | Single-pass wire buffer steady-state write |
| **`SecretStream55_SteadyState_ReadOnly_1MB`** | **978.00 MB/s (0.98 GB/s)** | **4 allocs/op** | Decryptor steady-state read |
| **`SecretStream55_FullDuplex_1MB`** | **480.78 MB/s** | **7 allocs/op** | Full duplex stream cycle (Encrypt + Decrypt) |

## Features & Cryptographic Discipline

- **Zero Allocation Hot Path:** Single contiguous `wireBuf` coalesces frame header (4B), ciphertext, and Poly1305 MAC tag (16B) into **1 single system `Write` call**.
- **Unique Chunk Nonce Derivation (P0 Security):** $N_{\text{chunk}} = N_{\text{base}} \oplus \text{seq}$ prevents keystream reuse across chunks.
- **Libsodium Framing Mode:** Compatible with Libsodium C `crypto_secretstream` tags (`TagMessage`, `TagPush`, `TagRekey`, `TagFinal`).
- **Dual-Path Build Support:** Seamless fallback to standard Go / `golang.org/x/crypto` when compiled without `GOEXPERIMENT=simd`.

## License

MIT License.
