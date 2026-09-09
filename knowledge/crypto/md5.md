---
id: md5
category: crypto
cwe: CWE-328
owasp: A02:2021-Cryptographic Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Utilisation de MD5 (fonction de hachage cassée)

## Description
MD5 est une fonction de hachage cryptographique cassée depuis 2004 : des collisions peuvent être générées en quelques secondes, et il n'offre aucune résistance face à des attaques par force brute modernes (GPU/ASIC). Son usage pour du hachage de mot de passe, une signature d'intégrité, ou une preuve cryptographique n'apporte plus aucune garantie de sécurité.

## Où ça apparaît typiquement
- Hachage de mots de passe utilisateurs (`md5($password)`) avant stockage en base.
- Génération de tokens, identifiants de session ou clés d'API à partir de `md5(uniqid())` ou équivalent.
- Vérification d'intégrité de fichiers téléchargés ou de payloads d'API.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel direct à `md5()`, `hashlib.md5`, `MessageDigest.getInstance("MD5")`, `MD5CryptoServiceProvider`.
- MD5 utilisé dans un contexte de stockage de credentials ou de génération de token sensible.
- Absence de sel (salt) ou de facteur de coût associé au hachage.

## Remédiation
- Pour les mots de passe : utiliser Argon2id (ou bcrypt/scrypt à défaut) avec un coût adapté.
- Pour l'intégrité de données non sensibles : utiliser SHA-256 ou SHA-3.
- Pour les tokens/identifiants aléatoires : utiliser un générateur cryptographiquement sûr (CSPRNG), pas un hachage.
- Voir `rules/remediation/md5.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/md5/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Password Storage Cheat Sheet
- NIST SP 800-131A Rev. 2 (dépréciation de MD5)
- CWE-328: Use of Weak Hash
