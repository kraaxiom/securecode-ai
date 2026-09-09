---
id: des
category: crypto
cwe: CWE-327
owasp: A02:2021-Cryptographic Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Utilisation de DES / 3DES (chiffrement obsolète)

## Description
DES utilise une clé effective de 56 bits, cassable par force brute en quelques heures avec du matériel moderne. Son successeur 3DES (Triple DES) est également déprécié : il est vulnérable à l'attaque Sweet32 sur les blocs de 64 bits et est retiré des standards NIST depuis 2023. Leur présence dans du code indique un chiffrement qui n'offre plus de garantie de confidentialité sérieuse.

## Où ça apparaît typiquement
- Appels explicites à `DES`, `DESede`/`TripleDES`/`3DES` dans les API crypto (Java `Cipher.getInstance("DES...")`, .NET `TripleDESCryptoServiceProvider`, OpenSSL `EVP_des_*`).
- Anciens systèmes de chiffrement de mots de passe, tokens de session ou fichiers legacy jamais migrés.
- Configurations TLS/VPN autorisant encore des suites de chiffrement DES/3DES.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Chaîne littérale `"DES"`, `"3DES"`, `"DESede"`, `"TripleDES"` passée à une factory ou un provider de chiffrement.
- Import de classes historiques (`javax.crypto.Cipher` avec transformation DES, `Crypto++ DES_EDE3`).
- Taille de clé/bloc de 56 ou 64 bits configurée explicitement.

## Remédiation
- Migrer vers AES-256 en mode authentifié (AES-GCM ou ChaCha20-Poly1305).
- Ne jamais réimplémenter un chiffrement custom ; utiliser les primitives certifiées de la bibliothèque standard du langage.
- Planifier une migration des données déjà chiffrées en DES/3DES avec re-chiffrement.
- Voir `rules/remediation/des.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/des/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cryptographic Storage Cheat Sheet
- NIST SP 800-131A Rev. 2 (retrait de TDEA/3DES)
- CWE-327: Use of a Broken or Risky Cryptographic Algorithm
