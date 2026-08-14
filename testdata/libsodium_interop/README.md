# Goldens + driver interop libsodium C

## Contenu attendu (après J1)

```
libsodium_interop/
  README.md                 ← ce fichier
  driver_secretstream.c     ← push/pull CLI
  golden/
    manifest.json
    *.plain / *.wire / *.key
  gen_goldens.sh
```

## Génération

```bash
# prérequis : libsodium + pkg-config
pkg-config --exists libsodium || exit 1
./gen_goldens.sh
```

## Consommation tests Go

- Goldens versionnés → tests offline sans libsodium runtime.  
- Driver présent → tests live `Go→C` et `C→Go`.

Voir `TODO_V2_SGOITER_LIBSODIUM.md`.
