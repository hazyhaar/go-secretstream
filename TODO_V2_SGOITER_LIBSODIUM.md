# TODO INTÉGRALE V2 — go-secretstream libsodium-compatible via sgoiter

| | |
|--|--|
| **Horizon** | 2 jours autonomie |
| **Contrat wire** | libsodium `crypto_secretstream_xchacha20poly1305` (cible wal-g) |
| **Moteur crypto V2** | AEAD **sgoiter** (Go pur, CGO=0), régénérable, sans stub sur le graphe |
| **Moteur perf V1** | **c2simd** SIMD — dual, non détruit |
| **HPM55 projet** | `019fd12d-01e6-7d0c-bbec-813b2c7eb575` (`secretstream_go`) |
| **Goals ouverts** | (1) crypter Go pur bit-compatible multi-chunk framing wal-g · (2) intégrer write/read vs libsodium C |
| **Fichier canonique** | **ce document** |
| **Démarrage machine** | `pkg-config libsodium` → **présent** (pc=0) sur l’hôte d’audit 2026-08-13 |

---

## A. Doctrine (non négociable pendant les 2j)

1. **Libsodium-compatible** = interop **octet à octet** avec libsodium **C**  
   (encrypt Go → decrypt C **et** encrypt C → decrypt Go).  
   Un round-trip Go↔Go nommé « Libsodium » **ne compte pas**.
2. **sgoiter** livre **S1–S2** (primitives + AEAD + HChaCha).  
   **S3 framing + S4 API** = Go produit (maintenable, testable).
3. **c2simd** = V1 perf, reste buildable et testé en défaut.  
   **sgoiter** = V2 pureté emit / runtime sous tag `aead_sgoiter`.
4. **Pas de scope creep** : pas EdDSA, X25519, argon2, full monocypher.  
   Graphe = secretstream + AEAD XChaCha20-Poly1305 + rekey.
5. **Gates fail-closed** : skip explicite si outil absent ; goldens versionnés = oracle offline obligatoire une fois posés.
6. **Ne pas ship** un regen monocypher qui casse les MAC (constat 2026-08-13).

---

## B. Architecture en strates

```
S4  API Encryptor/Decryptor (io.Writer/Reader)     ← secretstream.go
S3  Framing stream (header, chunks, tags, rekey)   ← secretstream.go (à coller libsodium C)
S2  AEAD one-shot + HChaCha rekey                  ← c2simd AUJOURD’HUI ; sgoiter CIBLE V2
S1  ChaCha / Poly1305 / verify / wipe              ← sous sgoiter harvest
```

| ID | Couche | État 2026-08-13 | Done V2 |
|----|--------|-----------------|---------|
| S4 | API publique | existe | bascule backend injectable |
| S3 | Framing | self-roundtrip « libsodium » ; **interop C absente** ; **préfixe BE4 probable hors wire C** | wire exact vs C |
| S2 | AEAD | c2simd prod ; monocypher_sgoiter oracle vert (blob figé) | sgoiter runtime + regen stable |
| S1 | primitives | harvest monocypher partiel (8 stubs hors AEAD) | zéro stub sur graphe AEAD/rekey |

---

## C. État de départ (checklist vérité)

| # | Fait | État |
|---|------|------|
| C1 | `go test ./...` secretstream55 (défaut c2simd) | vert (self RT) |
| C2 | `TestLibsodiumFramingRoundtrip` | vert **self only** |
| C3 | Interop libsodium **C** | **ABSENT** |
| C4 | monocypher_sgoiter vs C monocypher AEAD (CT+MAC) | vert blob HEAD |
| C5 | monocypher_sgoiter vs ccgo | vert blob HEAD |
| C6 | Regen monocypher == versionné (`ci_check`) | **ROUGE** |
| C7 | Backend stream sgoiter | **ABSENT** |
| C8 | libsodium-dev sur machine | **OUI** (`pkg-config` ok) |
| C9 | Goldens wire versionnés | **ABSENT** |
| C10 | Driver C push/pull | **ABSENT** |
| C11 | `docs/WIRE_LIBSODIUM.md` | brouillon écarts |
| C12 | DECISION V2 (plus oracle-only plafond) | posée |

---

## D. Jour 1 — Wire + testing de sortie

### D0. Prérequis machine (15 min)

- [ ] **D0.1** Confirmer libsodium :  
  `pkg-config --modversion libsodium` · headers `sodium.h` · `crypto_secretstream_xchacha20poly1305.h`
- [ ] **D0.2** Noter versions dans `docs/WIRE_LIBSODIUM.md` (libsodium X.Y, gcc, go).
- [ ] **D0.3** Créer arbo :  
  `testdata/libsodium_interop/{golden,bin}` · `scripts/gen_goldens.sh` · `testdata/libsodium_interop/driver_secretstream.c`

### D1. Spec wire libsodium C (2 h) — livrable `docs/WIRE_LIBSODIUM.md` complet

- [ ] **D1.1** Constants C (valeurs numériques exactes) :  
  `ABYTES`, `HEADERBYTES`, `KEYBYTES`, `MESSAGEBYTES_MAX`, tags `MESSAGE|PUSH|REKEY|FINAL`.
- [ ] **D1.2** Format header (24 o) : contenu exact après `init_push` / ce que le pull consomme.
- [ ] **D1.3** Format **un chunk** C : ordre des octets (tag ? ct ? mac ?) — **pas de préfixe longueur BE4 en libsodium natif**.
- [ ] **D1.4** AD / état interne : ce qui entre dans le Poly1305 côté C.
- [ ] **D1.5** Dérivation nonce / compteur / rekey côté C (vs Go `N_base⊕seq` + `adBuf=seq`).
- [ ] **D1.6** Tableau d’écarts Go actuel ligne à ligne (`secretstream.go` ~L70–130 encrypt, decrypt symétrique).
- [ ] **D1.7** Décision binaire écrite :  
  - **A (défaut mission)** : patcher S3 Go pour wire C exact  
  - **B** : renommer l’API actuelle « dialecte maison » et créer un **second** mode `NewLibsodiumCEncryptor` wire-exact  
  → **choisir A ou B explicitement dans WIRE_*.md**

### D2. Driver C + goldens (3 h)

- [ ] **D2.1** `driver_secretstream.c` CLI :
  - `init-push --key HEX --out-header FILE` (ou key file 32 o)
  - `push --state … --tag N` : stdin plain → stdout chunk bytes
  - `push-stream --key …` : plain file → wire file complet (header+chunks+final)
  - `pull-stream --key …` : wire file → plain file ; exit 1 si MAC fail
  - `tamper-mac` helper optionnel
- [ ] **D2.2** Build : `cc $(pkg-config --cflags --libs libsodium) -o bin/driver_secretstream driver_secretstream.c`
- [ ] **D2.3** `scripts/gen_goldens.sh` matrices plain :
  - tailles : `0, 1, 15, 16, 63, 64, 65, 1000, 65535, 65536, 65537, 131072, 200000`
  - patterns : zeros, `i%251`, random seed fixe
  - multi-chunk forcé (write en plusieurs push si API le permet)
  - **1 scénario rekey** (tag REKEY puis suite)
  - **1 scénario TagFinal** seul petit message
- [ ] **D2.4** Sortie golden :  
  `golden/<id>.key` (32 o) · `<id>.plain` · `<id>.wire` · `manifest.json`  
  (`id`, `bytes`, `sha256_plain`, `sha256_wire`, `tags`)
- [ ] **D2.5** **Versionner** au moins 8 goldens représentatifs (même sans driver en CI froide).
- [ ] **D2.6** Makefile :  
  `make interop-driver` · `make interop-goldens` · `make interop-test`

### D3. Tests Go de sortie (2–3 h) — `interop_libsodium_test.go`

| ID | Test | Critère pass |
|----|------|----------------|
| T1 | `TestInterop_Golden_CWire_GoDecrypt` | chaque golden wire → plain exact via decryptor Go wire-exact |
| T2 | `TestInterop_GoEncrypt_CDecrypt` | Go encrypt → driver `pull-stream` → plain exact (live libsodium) |
| T3 | `TestInterop_CEncrypt_GoDecrypt` | driver `push-stream` → Go decrypt (live) |
| T4 | `TestInterop_MACTamper_FailClosed` | flip 1 o MAC → erreur Go **et** C |
| T5 | `TestInterop_MultiChunk_Rekey` | golden ou live rekey |
| T6 | `TestSelfRoundtrip_LibsodiumMode` | conserver l’existant (non suffisant seul) |
| T7 | `TestSelfRoundtrip_StandardMode` | mode non-libsodium inchangé (pas de régression) |

Règles d’exécution :

- [ ] **D3.1** Goldens présents → T1 **obligatoire** sans libsodium runtime.  
- [ ] **D3.2** Driver + libsodium → T2 T3 T4 live.  
- [ ] **D3.3** Sinon → `t.Skip("libsodium_interop_unavailable")` **nommé**, jamais silent pass.  
- [ ] **D3.4** Build tag optionnel `libsodium_c` pour tests live si préférable à la détection PATH.

### D4. Patch framing S3 si décision A/B (reste J1)

- [ ] **D4.1** Implémenter le mode wire-exact (A sur place ou B nouvelle API).  
- [ ] **D4.2** Ne **pas** casser le mode stream actuel sans test de non-régression T7.  
- [ ] **D4.3** Documenter dans README : quel constructeur est **libsodium C wire**, lequel est historique.  
- [ ] **D4.4** T1 vert sur goldens C.

### D5. AEAD unitaire (en parallèle court)

- [ ] **D5.1** Garder verts : `go test ./internal/monocypher_sgoiter/`.  
- [ ] **D5.2** Unifier tailles dans `testdata/aead_sizes.json` (optionnel J1 si temps).  
- [ ] **D5.3** Interdiction : appeler « libsodium-compatible » un test monocypher-only.

### Exit criteria Jour 1

| # | Critère | Preuve |
|---|---------|--------|
| E1 | Wire documenté + décision A/B | `docs/WIRE_LIBSODIUM.md` complet |
| E2 | Driver buildable | `make interop-driver` |
| E3 | ≥8 goldens versionnés | `testdata/libsodium_interop/golden/` |
| E4 | T1 vert (Go decrypt wire C) | `go test -run Interop_Golden` |
| E5 | T2 ou T3 vert (live) | log test |
| E6 | Liste patches S3 restants ou S3 collé | note fin de jour dans ce TODO §H |

---

## E. Jour 2 — Moteur sgoiter + bascule runtime

### E1. Harvest AEAD sans stub (sgoiter) — 3 h

- [ ] **E1.1** Liste fermée des racines (à ajuster au sol) :  
  `crypto_aead_lock`, `crypto_aead_unlock`, `crypto_chacha20_h`, `crypto_chacha20_x`,  
  `crypto_poly1305*`, `crypto_verify16`, `crypto_wipe`, et tout callees **nécessaires**.  
  **Exclure** eddsa / scalarmult / argon du harvest runtime.
- [ ] **E1.2** Flags sgoiter : `-roots …` et/ou script `sgoiter/scripts/regen_aead_sgoiter.sh`.  
- [ ] **E1.3** Test mécanique package :  
  `TestNoPanicStubsOnAEADPath` — reflect ou grep CI : pas de `args ...any` + panic sur symboles AEAD.  
- [ ] **E1.4** Regen → diff 0 vs `monocypher_aead_sgoiter.go` versionné **ou** remplacement versionné **après** parity verte.  
- [ ] **E1.5** `go test ./internal/monocypher_sgoiter/` :  
  parity C monocypher + ccgo + golden JSON + chacha ietf — **tout vert**.  
- [ ] **E1.6** `c2simd/sgoiter/scripts/ci_check.sh` : section monocypher **verte**.

### E2. Interface engine — 2 h

- [ ] **E2.1** Créer `internal/engine/engine.go` :

```go
package engine

type AEAD interface {
    LockDst(dstCipher []byte, mac *[16]byte, key, nonce, ad, plain []byte) error
    UnlockDst(dstPlain []byte, key, nonce, ad, cipher, mac []byte) ([]byte, error)
    HChaCha20(out, key, in []byte)
}
```

- [ ] **E2.2** `engine/c2simd.go` — wrap `c2simd.AEADLockDst` / `UnlockDst` / `HChaCha20`.  
- [ ] **E2.3** `engine/sgoiter.go` — wrap `monocypher_sgoiter` (+ HChaCha exposé si besoin API).  
- [ ] **E2.4** `secretstream.go` : champ `eng engine.AEAD` ; constructeurs internes `newEncryptor(..., eng)`.  
- [ ] **E2.5** Build tags :  
  | Tag | Backend défaut stream |  
  |-----|------------------------|  
  | _(aucun)_ | c2simd (V1) |  
  | `aead_sgoiter` | monocypher_sgoiter (V2) |  
- [ ] **E2.6** Tests table-driven optionnels pour exercer les deux engines sans tag si injection possible en test.

### E3. Prouver le stream sgoiter — 2 h

- [ ] **E3.1** `go test -count=1 ./...` (défaut c2simd) — **aucune régression**.  
- [ ] **E3.2** `go test -count=1 -tags aead_sgoiter ./...` — round-trip stream + T6/T7.  
- [ ] **E3.3** T1–T5 avec tag `aead_sgoiter` (decrypt/encrypt Go sgoiter ↔ C).  
- [ ] **E3.4** Bench informatif non bloquant :  
  `go test -tags aead_sgoiter -bench=SecretStream55_SteadyState -benchtime=1s`  
  noter MB/s vs défaut dans retex (pas un gate).

### E4. Housekeep + HPM55 — 1 h

- [ ] **E4.1** README secretstream55 : section « Backends » + « Interop libsodium » + commandes test.  
- [ ] **E4.2** `DECISION.md` monocypher aligné (déjà V2 — mettre à jour si bascule faite).  
- [ ] **E4.3** `sgoiter/SPEC.md` § secretstream pointe ce TODO et l’état gates.  
- [ ] **E4.4** Retex HPM55 projet secretstream + sgoiter.  
- [ ] **E4.5** `goal-rendu` **uniquement** si E-exit 1–4 verts (sinon laisser goals ouverts).  
- [ ] **E4.6** Cocher §H journal de bord fin de mission.

### Exit criteria Jour 2

| # | Critère | Preuve |
|---|---------|--------|
| F1 | Regen AEAD stable | `ci_check` monocypher OK |
| F2 | Engine sgoiter branché | tag `aead_sgoiter` compile + tests |
| F3 | Interop C avec backend sgoiter | T1+T2 ou T3 verts sous tag |
| F4 | V1 c2simd intact | `go test` défaut vert |
| F5 | Doctrine / README à jour | fichiers + retex |

---

## F. Definition of Done — mission 2 jours (intégrale)

La mission est **DONE** ssi **tous** les points suivants sont vrais :

1. **Wire** libsodium C documenté ; mode Go **wire-exact** existe et est le seul appelé « libsodium-compatible ».  
2. **Goldens** C versionnés + **T1** vert (Go decrypt wire C).  
3. **Au moins une** direction live **T2 ou T3** verte sur la machine (libsodium présente).  
4. **MAC tamper** fail-closed (T4).  
5. **AEAD sgoiter** : parity C + regen = versionné (`ci_check`).  
6. **Stream** sur `aead_sgoiter` : round-trip + interop T1 (et live si possible).  
7. **Défaut c2simd** : tests package verts, pas de régression perf de build.  
8. **Hors scope** respecté (pas de full monocypher).  
9. **Retex** HPM55 écrit ; goals rendus seulement si 1–7 OK.

**NON-DONE** si : seul self-roundtrip Go ; ou regen casse MAC ; ou README ment sur « compatible » sans T1.

---

## G. Checklist testing de sortie (à coller en fin de chaque demi-journée)

```bash
export GOWORK=off

# --- secretstream55 ---
cd /devhoros/pkg/secretstream55
go test -count=1 ./...
go test -count=1 ./internal/monocypher_sgoiter/
go test -count=1 -run Interop ./...                    # goldens / interop
go test -count=1 -tags aead_sgoiter ./...              # dès E2 prêt
go test -count=1 -tags libsodium_c -run Interop ./...  # si tag live

# --- driver ---
make -C /devhoros/pkg/secretstream55 interop-driver || true
make -C /devhoros/pkg/secretstream55 interop-test || true

# --- sgoiter ---
cd /devhoros/c2simd
./sgoiter/scripts/ci_check.sh
./bin/tribench -root /devhoros/c2simd -sgoiter ./bin/sgoiter -skip-ccgo -skip-bench
```

Journal minimal à append en §H : date, commande, exit code, 1 ligne constat.

---

## H. Journal de bord (remplir en autonomie)

| Quand | Fait | Exit / preuve | Blocage |
|-------|------|---------------|---------|
| 2026-08-13 final | **and_ones_u64 fix** + regen monocypher mécanique + ci_check OK + goal-rendu G1 G2 | ci=0 mono parity=0 | — |
| 2026-08-13 session | libsodium-dev installé ; lsstream porté ; wire A ; goldens 14 ; interop T1–T4 verts ; engine c2simd/sgoiter ; tests défaut+aead_sgoiter verts | make interop-test OK ; go test ./... OK | regen monocypher full encore non auto (parity) — blob HEAD conservé |
| J1 matin | | | |
| J1 midi | | | |
| J1 soir | | | |
| J2 matin | | | |
| J2 midi | | | |
| J2 soir | | | |


---

## I. Hors scope (refus mécanique)

- [ ] Porter EdDSA / X25519 / argon2 / blake full monocypher  
- [ ] Exiger sgoiter plus rapide que c2simd  
- [ ] Transpiler le framing S3 depuis du C  
- [ ] Publier module public avant T1 vert  
- [ ] `goal-rendu` sans interop C  
- [ ] Commit d’un monocypher_aead regen qui échoue parity  
- [ ] `git add -A` / secrets / binaires driver non reproductibles sans script  

---

## J. Fichiers touchés (carte)

| Chemin | Rôle |
|--------|------|
| `TODO_V2_SGOITER_LIBSODIUM.md` | **cette todo intégrale** |
| `docs/WIRE_LIBSODIUM.md` | spec + écarts wire |
| `testdata/libsodium_interop/**` | driver, goldens, README |
| `scripts/gen_goldens.sh` | génération goldens |
| `Makefile` | interop-driver / goldens / test |
| `interop_libsodium_test.go` | tests de sortie |
| `secretstream.go` | S3/S4 + injection engine |
| `internal/engine/*.go` | backends c2simd / sgoiter |
| `internal/monocypher_sgoiter/**` | AEAD emit + tests parity |
| `README.md` | backends + interop |
| `c2simd/sgoiter/scripts/regen_*.sh` | regen AEAD |
| `c2simd/sgoiter/SPEC.md` | pointeur mission |

---

## K. Planning horaire indicatif

| Bloc | Contenu | Exit partiel |
|------|---------|--------------|
| **J1 0–0,5 h** | D0 libsodium + arbo | headers OK |
| **J1 0,5–2,5 h** | D1 WIRE complet + décision A/B | md rempli |
| **J1 2,5–5,5 h** | D2 driver + goldens versionnés | make goldens |
| **J1 5,5–7,5 h** | D3 T1–T4 + D4 patch S3 si besoin | T1 vert |
| **J1 7,5–8 h** | §H + retex soir | journal |
| **J2 0–3 h** | E1 harvest + regen + ci_check | F1 |
| **J2 3–5 h** | E2 engine + tags | compile dual |
| **J2 5–7 h** | E3 interop sgoiter + stream | F2 F3 |
| **J2 7–8 h** | E4 docs goals retex §H | F4 F5 · DoD |

---

## L. Première commande (réveil agent)

```bash
pkg-config --modversion libsodium
cd /devhoros/pkg/secretstream55
# 1) compléter docs/WIRE_LIBSODIUM.md (D1)
# 2) écrire testdata/libsodium_interop/driver_secretstream.c (D2)
# 3) ne pas toucher monocypher regen avant fin J1 wire
```

---

## M. Liens croisés

- HPM55 secretstream_go : `019fd12d-…`  
- sgoiter (outil) : `019fec65-…` · domicile `/devhoros/c2simd/sgoiter`  
- Pointeur mince sgoiter : `c2simd/sgoiter/TODO_SECRETSTREAM.md` → **ici**  
- **Monocypher API complète via sgoiter** (blake/x25519/eddsa/argon — pas le wire libsodium) :  
  [`/devhoros/c2simd/sgoiter/TODO_MONOCYPHER_FULL.md`](/devhoros/c2simd/sgoiter/TODO_MONOCYPHER_FULL.md)  
- Probe perf c2simd/sgoiter kernels : hors chemin critique secretstream 2j (bonus seulement)

---

*Todo intégrale posée 2026-08-13. Toute session autonome commence en cochant C.* puis D0.* sans improvisation de scope.*
