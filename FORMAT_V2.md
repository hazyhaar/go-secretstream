# Format de flux maison v2

Le descripteur déclaratif `spec/format_v2.cue` gouverne ce format. Les constantes Go, la présente spécification et les vecteurs de test en sont des produits générés. Le code manuscrit ne redéfinit pas ces valeurs.

L'ordre des octets du fil est grand-boutiste.

## En-tête

L'en-tête v2 occupe 36 octets. L'en-tête v1 (nonce seul, également taille libsodium) occupe 24 octets.

Le magique tient sur 8 octets à l'offset 0 (hexadécimal 535335352d763200, ASCII « SS55-v2 » suivi d'un octet nul). Le numéro de version occupe 2 octets à l'offset 8 et vaut 2. Le champ de drapeaux occupe 2 octets à l'offset 10 et doit valoir 0. Le nonce aléatoire occupe 24 octets à l'offset 12.

La sous-clé se dérive par HChaCha20 à partir de la clé et des seize premiers octets du nonce.

## Trame

Chaque trame commence par une longueur de 4 octets grand-boutiste, qui compte le tag, le chiffré et le MAC. Le tag tient sur 1 octet. Le MAC tient sur 16 octets. La charge utile minimale vaut 17 octets (tag et MAC, chiffré vide). La charge utile maximale vaut 65553 octets (tag, fragment d'au plus 65536 octets, MAC).

Le tag de message vaut 0x00. Le tag terminal vaut 0x03. Les valeurs 0x01 (push) et 0x02 (rekey) sont refusées. Le tag entre dans la donnée associée authentifiée : le modifier sur le fil fait échouer le MAC.

Le nonce IETF de chaque trame (12 octets) est formé des 4 octets du nonce d'en-tête à l'offset 16, suivis du compteur de séquence grand-boutiste de 8 octets, en clair et sans masque XOR. Le compteur part de zéro et s'incrémente à chaque trame. Un débordement est une erreur collante.

## Donnée associée

La donnée associée authentifiée est le magique (8 octets), le compteur de séquence (8 octets), le tag (1 octet), la longueur de l'ad d'appelant (4 octets grand-boutiste), puis l'ad d'appelant. Le préfixe sans ad d'appelant tient sur 21 octets. Le préfixe de longueur sépare les concaténations qui se confondaient en v1. L'ad d'appelant n'est pas transmise sur le fil : le lecteur la fournit.

## Clôture et lecture

Close écrit une trame terminale à chiffré vide (longueur 17), puis efface les secrets. Un second Close n'émet rien. Write après Close reste une erreur. Le lecteur ne rend io.EOF qu'après avoir authentifié cette trame terminale. Une fin de flux avant ce terminal est une erreur collante de troncature, jamais un io.EOF. Toute donnée après le terminal, une seconde trame terminale, ou un terminal à chiffré non vide, est une erreur.

## Décodeur hybride

Le décodeur lit d'abord 8 octets. Si les huit premiers octets valent le magique, le reste de l'en-tête v2 est consommé. Sinon ces huit octets sont le début du nonce v1 et les seize suivants le complètent. Le format v1 n'a ni magique, ni version, ni bloc terminal. Une archive v1 qui s'arrête sur une frontière de trame se lit comme complète. La probabilité qu'un nonce v1 aléatoire commence par le magique est 2^-64.
