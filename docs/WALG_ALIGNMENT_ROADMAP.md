# Feuille de Route d'Alignement Industriel & Réponses aux Retours WAL-G (#2518)

> **Statut :** Plan d'action technique préliminaire avant publication amont.  
> **Contexte :** Issue [wal-g/wal-g#2518](https://github.com/wal-g/wal-g/issues/2518) — Revue de sécurité et modélisation des menaces par Ilya Shipitsin (`chipitsine`).

---

## 1. Synthèse des Retours d'Ingénierie Amont

L'analyse de risque fournie par le mainteneur de WAL-G identifie des exigences de gouvernance et de robustesse mécanique indispensables pour un moteur de sauvegarde critique :

1. **Absence d'Oracle Obligatoire en CI (`t.Skip`) :** Les tests d'interopérabilité ne doivent jamais s'ignorer silencieusement si l'oracle de référence C (`gcc -O2` + libsodium officiel) est manquant.
2. **Couverture de Fuzzing :** Nécessité d'un banc de fuzzing guidé par la couverture (`testing.F`) pour éprouver les transitions d'état, les troncatures, les corruptions de tags Poly1305 et les cas limites de frontières de blocs.
3. **Clarification du Profil de Streaming :** Formalisation explicite de la convention de découpage par blocs de 8192 octets + 17 octets d'en-tête/tag avec `TAG_FINAL` à l'EOF (profil streaming WAL-G) par rapport aux APIs de messages discrets de libsodium.
4. **Maturité & Cycle de Vie :** Publication d'une version sémantique taguée (`v0.1.0`) avec garanties de provenance et signatures.
5. **Stratégie d'Intégration Amont :** Conservation du moteur CGO par défaut dans WAL-G et positionnement de `go-secretstream` en option de compilation statique pure-Go (`-tags purego_libsodium`).

---

## 2. Plan d'Implémentation Technique

### Phase 1 : Rigueur CI & Interdiction du Contournement (`interop_libsodium_test.go`)
- Définition d'un mode de validation strict : si la variable d'environnement `CI` ou `REQUIRE_LIBSODIUM_ORACLE=1` est active, l'absence du driver de référence ou du jeu de données *golden* provoque un échec dur immédiat (`t.Fatalf`) plutôt qu'un saut passif (`t.Skip`).
- Intégration de la cible `make interop-driver interop-goldens` dans la chaîne de validation hermétique.

### Phase 2 : Harnais de Fuzzing Natif Go (`fuzz_test.go`)
- Implémentation de `FuzzLibsodiumStreamRoundtrip(f *testing.F)` :
  - Génération de flux de longueurs arbitraires (0 octet, 1 octet, 8191 octets, 8192 octets, 8193 octets, multiples de 8192).
  - Validation du cycle complet Chiffrement -> Déchiffrement -> Parité bit-exacte avec l'entrée brute.
- Implémentation de `FuzzLibsodiumCorruptedStream(f *testing.F)` :
  - Injection de mutations aléatoires (altération d'en-tête, modification d'un octet de tag, troncature prématurée, modification de payload).
  - Preuve du comportement *fail-closed* : le déchiffreur doit rejeter systématiquement le flux avec une erreur explicite sans paniquer, sans allouer de mémoire excessive et sans verrouiller l'état en faux positif.

### Phase 3 : Documentation Formelle du Profil Streaming (`WIRE_LIBSODIUM.md`)
- Spécification détaillée du conteneur de flux WAL-G :
  - Taille de chunk nominale : 8192 octets de données brutes.
  - Séquence de trames : `TAG_MESSAGE` (0x00) pour les blocs intermédiaires, `TAG_FINAL` (0x02) pour le bloc terminal.
  - Structure de chaque fragment : `[Header (24 o)]` puis `N × [Ciphertext (len(plain) o) || Tag (17 o)]`.

---

## 3. Positionnement pour la Réponse Amont

Après validation intégrale des phases 1, 2 et 3 :
- Remercier factuellement Ilya Shipitsin pour la rigueur de son analyse de risques.
- Présenter les garanties mécaniques implémentées (CI bloquante sans `t.Skip`, suite de fuzzing, spécification du profil de streaming).
- Proposer l'intégration dans WAL-G sous la forme d'un pilote statique optionnel (`//go:build purego_libsodium`), sans modification du chemin CGO nominal.
