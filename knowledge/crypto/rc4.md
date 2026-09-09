---
id: rc4
category: crypto
cwe: CWE-327
owasp: A02:2021-Cryptographic Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Utilisation de RC4 (chiffrement par flux cassé)

## Description
RC4 est un chiffrement par flux présentant des biais statistiques connus dans son flux de sortie, exploitables pour récupérer du texte en clair (attaques sur WEP, sur les cookies TLS via RC4, etc.). Il est interdit dans TLS depuis la RFC 7465 (2015) et ne doit plus être utilisé dans aucun protocole ou stockage.

## Où ça apparaît typiquement
- Suites de chiffrement TLS/SSL héritées encore activées côté serveur.
- Chiffrement de fichiers ou de flux propriétaires dans du code legacy.
- Bibliothèques ou SDK tiers embarquant RC4 par défaut pour la compatibilité.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Chaîne `"RC4"` ou `"ARC4"` passée à une API de chiffrement (`Cipher.getInstance("RC4")`, `crypto.createCipher('rc4', ...)`).
- Suite de chiffrement contenant `RC4` dans une configuration serveur TLS (Apache, Nginx, IIS).
- Dépendance à une bibliothèque crypto obsolète exposant RC4 comme algorithme par défaut.

## Remédiation
- Désactiver toutes les suites RC4 côté serveur et client.
- Migrer vers AES-GCM ou ChaCha20-Poly1305 pour tout chiffrement symétrique.
- Mettre à jour les bibliothèques tierces qui imposent RC4 par compatibilité descendante.
- Voir `rules/remediation/rc4.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/rc4/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- RFC 7465: Prohibiting RC4 Cipher Suites
- OWASP Transport Layer Protection Cheat Sheet
- CWE-327: Use of a Broken or Risky Cryptographic Algorithm
