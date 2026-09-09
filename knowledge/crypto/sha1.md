---
id: sha1
category: crypto
cwe: CWE-328
owasp: A02:2021-Cryptographic Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Utilisation de SHA-1 (fonction de hachage affaiblie)

## Description
SHA-1 est considéré comme cryptographiquement cassé depuis la démonstration d'une collision pratique (attaque "SHAttered", 2017). Il reste utilisable pour certains cas non cryptographiques (checksums non adversariaux), mais ne doit plus être employé pour des signatures numériques, des certificats, du hachage de mot de passe ou toute preuve d'intégrité opposable à un attaquant.

## Où ça apparaît typiquement
- Signature de certificats TLS ou de commits/artefacts logiciels.
- Hachage de mots de passe ou de secrets avant stockage.
- Génération de HMAC avec SHA-1 dans des protocoles legacy (certains flux OAuth1, anciens JWT).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à `sha1()`, `hashlib.sha1`, `MessageDigest.getInstance("SHA-1")`, `SHA1CryptoServiceProvider`.
- Configuration de signature de certificat/CSR utilisant `sha1WithRSAEncryption`.
- SHA-1 utilisé comme fonction de dérivation de mot de passe.

## Remédiation
- Remplacer par SHA-256 ou SHA-3 pour le hachage général et les signatures.
- Pour les mots de passe : utiliser Argon2id (ou bcrypt/scrypt).
- Régénérer les certificats signés en SHA-1 avec un algorithme moderne (SHA-256+).
- Voir `rules/remediation/sha1.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/sha1/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cryptographic Storage Cheat Sheet
- NIST SP 800-131A Rev. 2 (retrait de SHA-1)
- CWE-328: Use of Weak Hash
