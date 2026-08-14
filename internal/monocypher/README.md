# monocypher55 — Pure Go 1.27 SIMD Monocypher 4.0.2

`monocypher55` est une implémentation complète, sans CGO et sans assembleur externe, de la bibliothèque cryptographique [Monocypher 4.0.2](https://monocypher.org) en langage Go, exploitant les intrinsèques vectorielles natives `simd/archsimd` de Go 1.27.

---

## 1. Caractéristiques & Architecture

- **Zéro Assembleur Externe (Plan 9 `.s`) & Zéro CGO :** 100 % Pure Go, portable, typé et analysable par le compilateur.
- **Accélération Matérielle AVX2 (256-bit) :**
  - Noyau ChaCha20 vectoriel 4-voies (`archsimd.Uint32x8`) traitant 256 octets par itération.
  - Déroulage 2-Way ILP et réduction 32-bit pour Poly1305.
- **Garde Runtime CPU & Repli Multi-Architecture :**
  - Détection dynamique AVX2 via `archsimd.X86.AVX2()` (aucun `SIGILL` sur matériel ancien).
  - Repli scalaire automatique sur architectures non-x86 (`arm64`, `riscv64`) et sans `GOEXPERIMENT=simd`.
- **API Zero-Allocation (`LockDst` / `UnlockDst`) :**
  - Chiffrement et déchiffrement AEAD à **0 B/op (0 allocation sur le tas)**.

---

## 2. Performances & Benchmarks (Intel Core i9-14900K)

Mesures standardisées avec `GOTOOLCHAIN=go1.27rc3`, `GOEXPERIMENT=simd` et bancs `for b.Loop()` :

| Algorithme / Scénario | Débit / Latence Pure Go AVX2 | Référence `x/crypto` (Assembleur Plan 9) | Allocations | Gain vs Assembleur |
| :--- | :--- | :--- | :--- | :--- |
| **ChaCha20 AVX2 (64 Ko)** | **913,70 Mo/s** | 625,96 Mo/s | **0 B/op** | **+46,0 %** |
| **Poly1305 64-bit MULQ (64 Ko)** | **2 308,74 Mo/s** (2,31 Go/s) | 2 098,63 Mo/s (2,10 Go/s) | **0 B/op** | **+10,0 %** |
| **AEAD Open Zero-Alloc (64 Ko)** | **1 510,43 Mo/s** (1,51 Go/s) | 2 719,79 Mo/s (2,72 Go/s) | **0 B/op** | 0 allocation heap |
| **AEAD Lock Zero-Alloc (64 Ko)** | **1 081,77 Mo/s** (1,08 Go/s) | 3 365,38 Mo/s (3,36 Go/s) | **0 B/op** | 0 allocation heap |
| **AEAD AD-Heavy (64 Ko AD + 64 B PT)** | **3 463,20 Mo/s** (3,46 Go/s) | Non mesuré | **0 B/op** | Débit crête |
| **Blake2b-512 (1 Ko)** | **397,82 Mo/s** | Non mesuré | **0 B/op** | 0 allocation heap |
| **Ed25519 Scalarbase** | **43,20 µs / op** | Non mesuré | **0 B/op** | 0 allocation heap |
| **X25519 DH** | **88,43 µs / op** | Non mesuré | **0 B/op** | 0 allocation heap |
| **Argon2i ($m=8, p=1$)** | **28,40 µs / op** | Non mesuré | **0 B/op** | 0 allocation heap |

---

## 3. Matrice de Compilation & Build Tags

| Configuration | Build Tags | Moteur Actif |
| :--- | :--- | :--- |
| `amd64` + `GOEXPERIMENT=simd` | `goexperiment.simd && amd64` | SIMD AVX2 `Uint32x8` + garde runtime `archsimd.X86.AVX2()` |
| `amd64` sans `GOEXPERIMENT=simd` | `!goexperiment.simd` | Scalaire pur Go bit-exact |
| `arm64`, `riscv64`, etc. | `!amd64` | Scalaire pur Go bit-exact |

---

## 4. Conformité & Tests de Sécurité

- **Oracles C :** 13/13 suites de tests validées bit-à-bit contre `monocypher.c` 4.0.2 compilé en `-O0`.
- **Fuzzing Différentiel :** Plus de 500 000 exécutions aléatoires via `go test -fuzz`.
- **Project Wycheproof :** Rejet systématique des tags altérés, nonces invalides et points hors courbe.
- **Hygiène Mémoire :** `Crypto_wipe` à effacement réel garanti et zéro état mutable global.
