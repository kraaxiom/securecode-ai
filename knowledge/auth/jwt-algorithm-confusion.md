---
id: jwt-algorithm-confusion
category: auth
cwe: CWE-347
owasp: A02:2021-Cryptographic Failures
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# JWT Algorithm Confusion

## Description
Cette vulnérabilité survient quand une bibliothèque de vérification de JWT fait confiance à l'algorithme annoncé dans l'en-tête du token (`alg`) plutôt que d'imposer côté serveur l'algorithme attendu. Un attaquant peut alors forger un token signé avec un algorithme différent de celui prévu par le développeur — le cas le plus connu étant de faire passer un token signé en HMAC (`HS256`) en réutilisant la clé publique RSA de l'application comme secret HMAC, trompant ainsi la vérification.

## Où ça apparaît typiquement
- Middlewares d'authentification qui appellent une fonction de vérification JWT générique en lui laissant déduire l'algorithme depuis le token lui-même.
- API exposant sa clé publique RSA/EC (pour la vérification par des tiers) tout en acceptant plusieurs familles d'algorithmes de signature.
- Bibliothèques JWT anciennes ou mal configurées ne permettant pas de restreindre explicitement la liste des algorithmes acceptés.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à la fonction de décodage/vérification JWT sans paramètre explicite restreignant l'algorithme attendu (`algorithms=['RS256']` ou équivalent).
- Code lisant l'algorithme depuis le header du token pour décider dynamiquement de la clé/méthode de vérification à utiliser.
- Clé publique RSA accessible publiquement alors que l'algorithme HMAC reste accepté par le vérificateur.

## Remédiation
- Toujours spécifier explicitement et de façon figée l'algorithme (ou la liste restreinte d'algorithmes) attendu lors de la vérification, sans jamais le déduire du token.
- Utiliser des clés et des vérificateurs distincts par algorithme, sans jamais réutiliser une clé publique comme secret partagé potentiel.
- Mettre à jour vers une bibliothèque JWT maintenue qui impose cette restriction par défaut.
- Voir `rules/remediation/jwt-algorithm-confusion.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/jwt-algorithm-confusion/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: JSON Web Token for Java
- CWE-347: Improper Verification of Cryptographic Signature
