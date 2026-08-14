# Wire libsodium secretstream — spec opposable (2026-08-13)

**libsodium :** 1.0.18  
**Réf. pure Go :** `internal/lsstream` (port algorithmique wal-g fork, Apache-2.0)  
**Décision framing :** **A** — mode `NewLibsodium*` = wire C exact (wal-g chunking 8192).  
Mode `NewEncryptor*` = framing **maison** (préfixe BE4 + AEAD engine), non libsodium-wire.

## Constants (libsodium 1.0.18)

| Symbole | Valeur |
|---------|--------|
| KEYBYTES | 32 |
| HEADERBYTES | 24 |
| ABYTES | 17 (= 1 tag chiffré + 16 MAC) |
| TAG_MESSAGE | 0x00 |
| TAG_PUSH | 0x01 |
| TAG_REKEY | 0x02 |
| TAG_FINAL | 0x03 (= PUSH\|REKEY) |
| Chunk plaintext (wal-g / ce package) | **8192** |

## Format stream

```
wire = header[24] || chunk0 || chunk1 || … || chunkN
chunk = push(plaintext, tag)   # len = mlen + 17
```

- **Pas** de préfixe longueur uint32 BE (contrairement au mode maison).
- Chunks non-finaux : exactement 8192 o plaintext → 8209 o wire.
- Dernier chunk : TAG_FINAL, longueur variable (y compris 0 o plain + 17 o wire).
- Libsodium C peut mettre TAG_FINAL sur un chunk plein 8192 (sans empty FINAL).
- Go writer (lsstream) : MESSAGE sur chunks pleins, puis FINAL (souvent vide) au `Close()` — les deux formes sont acceptées en lecture.

## Crypto (pas un simple AEAD Seal)

Construction manuelle ChaCha20-IETF + Poly1305 (voir `lsstream/secretstream.go`) :
1. `k = HChaCha20(key, header[0:16])`, `inonce = header[16:24]`, counter init 1
2. Poly key from chacha block 0
3. Auth block tag||zeros (64 o), encrypt message, MAC 16 o
4. Advance inonce ^= mac[0:8], increment counter ; rekey if TAG_REKEY bit

## Mode maison (NewEncryptor) — distinct

```
header = nonce[24]
frame  = be32(len(ct||mac)) || ct || mac[16]
N_chunk = nonce with last 8 o ^= seq
AD = be64(seq)
ct,mac = AEAD_XChaCha20Poly1305(key, N_chunk, AD, plain)  # engine c2simd|sgoiter
```

## Preuves interop

| Test | Sens |
|------|------|
| `TestInterop_Golden_CWire_GoDecrypt` | wire C → Go |
| `TestInterop_GoEncrypt_CDecrypt` | wire Go → C |
| `TestInterop_CEncrypt_GoDecrypt` | live C push → Go |
| `TestInterop_MACTamper_FailClosed` | flip MAC → fail Go+C |

Goldens : `testdata/libsodium_interop/golden/`  
Driver : `make interop-driver` · `make interop-goldens`
