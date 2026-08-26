# secretstream55

Streaming encryption in Pure Go (CGO=0).

## Modes

| API | Wire | Crypto |
|-----|------|--------|
| `NewEncryptor` / `NewDecryptor` | **Maison** (BE4 + AEAD) | `engine` : **monocypher55** (défaut — bascule 2026-08-15) ; moteur externe injectable par `NewEncryptorWithEngine` / `NewDecryptorWithEngine` (ex. `c2simd/aeadengine`, sans dépendance de module) |
| `NewLibsodiumEncryptor` / `NewLibsodiumDecryptor` | **libsodium** `crypto_secretstream_xchacha20poly1305` (wal-g) | `internal/lsstream` (ChaCha20-IETF + Poly1305, bit-compat C) |

Libsodium mode **requires** `Close()` on the encryptor to emit `TAG_FINAL`.

## Format de flux v2

Le mode maison émet désormais un flux versionné. L'écrivain produit exclusivement
ce format. Le lecteur reconnaît à la fois les archives v2 et les archives v1
déjà stockées.

L'en-tête v2 occupe trente-six octets : huit octets magiques `SS55-v2` suivis
d'un octet nul, un numéro de version entier non signé de seize bits grand-boutiste
égal à deux, un champ de drapeaux de seize bits qui doit valoir zéro à la lecture,
puis un nonce aléatoire de vingt-quatre octets. La sous-clé reste dérivée par
HChaCha20 à partir de la clé et des seize premiers octets du nonce, comme en v1.

Chaque trame porte une longueur de quatre octets grand-boutiste, un tag d'un
octet, le chiffré et un MAC de seize octets. La longueur compte le tag, le
chiffré et le MAC. Le tag vaut `0x00` pour un message ou `0x03` pour la trame
terminale ; les valeurs `0x01` et `0x02` sont refusées. Le tag entre dans la
donnée associée authentifiée : le modifier sur le fil fait échouer le MAC.

Le nonce IETF de chaque trame est `nonce[16:20]` suivi du compteur de séquence
grand-boutiste, en clair et sans masque XOR. Le compteur part de zéro et
s'incrémente à chaque trame ; un débordement est une erreur collante.

La donnée associée authentifiée est injective :
`SS55-v2\x00 || seq || tag || longueur_de_l_ad_appelant || ad_appelant`.
Le préfixe de longueur sépare les concaténations qui se confondaient en v1.

`Close()` est obligatoire. Il écrit une trame `TagFinal` à chiffré vide
(longueur dix-sept), puis efface les secrets. Un second `Close()` n'émet rien.
`Write` après `Close` reste une erreur. Le lecteur ne rend `io.EOF` qu'après
avoir authentifié cette trame terminale. Une fin de flux avant ce terminal
est une erreur collante de troncature, jamais un `io.EOF`. Toute donnée après
le terminal, une seconde trame terminale, ou un terminal à chiffré non vide,
est une erreur.

Le décodeur hybride lit d'abord huit octets. S'ils valent le magique, le reste
de l'en-tête v2 est consommé. Sinon ces huit octets sont le début du nonce v1
et les seize suivants le complètent. La probabilité qu'un nonce v1 aléatoire
commence par le magique est 2⁻⁶⁴.

Le format v1 n'a ni magique, ni version, ni bloc terminal. Une archive v1 qui
s'arrête sur une frontière de trame se lit comme complète : cette troncature
n'est pas détectable. Une coupure au milieu d'une trame v1 reste une erreur
(`io.ErrUnexpectedEOF` collant). La lecture v1 est conservée bit à bit ; les
archives figées sous `testdata/v1/` en sont la preuve.

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

## Benchmark (standard / c2simd, historique)

See `secretstream_bench_test.go`. Libsodium-wire benches: `internal/lsstream` / wal-g upstream.

## Contributors

- **Hazyhaar** ([@hazyhaar](https://github.com/hazyhaar)) — Architecture, system design & maintainer
- **Gemini** (Google DeepMind) — Adversarial audits, research & verification
- **Grok** (xAI) — Low-level robustness, protocol inspection & fuzzing
- **Claude** (Anthropic) — Go 1.27 SIMD transpiler passes & CUE formal schemas

