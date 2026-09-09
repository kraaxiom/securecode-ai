---
id: weak-jwt-secret
category: crypto
cwe: CWE-1391
owasp: A02:2021-Cryptographic Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Secret JWT faible ou prévisible

## Description
Un JSON Web Token signé avec HMAC (`HS256`) repose sur la confidentialité et l'entropie du secret partagé. Un secret court, prévisible, une valeur par défaut de framework, ou une chaîne présente dans des dictionnaires publics permet à un attaquant de forger des tokens valides par force brute hors ligne, compromettant intégralement l'authentification.

## Où ça apparaît typiquement
- Configuration d'application avec `JWT_SECRET` court, laissé à une valeur d'exemple (`"secret"`, `"changeme"`).
- Génération de secret via un mot de passe humain plutôt qu'une valeur aléatoire à haute entropie.
- Algorithme JWT laissé configurable côté client, permettant un repli vers `alg: none` ou une confusion HS256/RS256.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Secret JWT de longueur insuffisante (moins de 32 octets/256 bits) codé en dur ou en valeur par défaut.
- Bibliothèque JWT configurée sans vérification stricte de l'algorithme attendu.
- Absence de rotation ou d'expiration courte pour les tokens signés en HS256.

## Remédiation
- Générer le secret JWT via un CSPRNG avec au moins 256 bits d'entropie, stocké dans un gestionnaire de secrets.
- Épingler explicitement l'algorithme attendu côté vérification (rejeter `none` et tout algorithme non prévu).
- Privilégier RS256/ES256 (clé asymétrique) lorsque la vérification doit être distribuée à plusieurs services.
- Voir `rules/remediation/weak-jwt-secret.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/weak-jwt-secret/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP JSON Web Token Cheat Sheet for Java
- CWE-1391: Use of Weak Credentials
