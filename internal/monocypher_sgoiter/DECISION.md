# Décision produit — monocypher_sgoiter / secretstream V2

**2026-08-13 (final)**

## Rôles

| Composant | Rôle |
|-----------|------|
| `NewLibsodium*` + `internal/lsstream` | Wire **libsodium C** (wal-g) |
| `NewEncryptor*` + `internal/engine` | Framing maison ; **c2simd** défaut / **sgoiter** (`-tags aead_sgoiter`) |
| `monocypher_sgoiter` | AEAD monocypher émis sgoiter — backend maison + parity C |

## Regen AEAD

```bash
cd /devhoros/c2simd && ./sgoiter/scripts/regen_monocypher_dogfood.sh
# ou ./sgoiter/scripts/regen_aead_sgoiter.sh
```

Prérequis : règle `and_ones_u64` ne traite **pas** `0xffffffff` comme identité u64  
(fix 2026-08-13 — sinon MAC poly1305 faux).  
`ci_check` vérifie emit frais == blob versionné.

## Gates

```bash
cd /devhoros/pkg/secretstream55 && export GOWORK=off
make ci && make test-sgoiter
cd /devhoros/c2simd && ./sgoiter/scripts/ci_check.sh
```
