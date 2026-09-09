---
id: weak-tls
category: crypto
cwe: CWE-326
owasp: A02:2021-Cryptographic Failures
severity_default: medium
languages: []
---

# Configuration TLS faible (suites de chiffrement, paramètres)

## Description
Au-delà de la version de protocole, une configuration TLS peut rester faible via des suites de chiffrement obsolètes (CBC non-AEAD, absence de forward secrecy), des paramètres Diffie-Hellman trop courts, ou une absence de HSTS. Ces faiblesses réduisent la résistance du canal chiffré même si TLS 1.2/1.3 est utilisé.

## Où ça apparaît typiquement
- Configuration de suites de chiffrement serveur incluant des algorithmes CBC, RC4, export-grade ou sans PFS.
- Paramètres Diffie-Hellman personnalisés (`dhparam`) générés avec une taille inférieure à 2048 bits.
- Absence d'en-tête `Strict-Transport-Security` ou politique de renégociation permissive.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Liste de suites de chiffrement (`ssl_ciphers`, `SSLCipherSuite`) contenant `EXPORT`, `NULL`, `DES`, `RC4`, ou des suites non-AEAD sans autre garde-fou.
- Fichier `dhparam.pem` de taille inférieure à 2048 bits.
- Ordre de préférence des suites non forcé côté serveur (`ssl_prefer_server_ciphers off`).

## Remédiation
- Restreindre les suites aux algorithmes AEAD modernes (AES-GCM, ChaCha20-Poly1305) avec forward secrecy (ECDHE).
- Générer des paramètres DH d'au moins 2048 bits, idéalement 3072+.
- Activer HSTS avec une durée adaptée et `includeSubDomains`.
- Voir `rules/remediation/weak-tls.md` pour les diffs par configuration.

## Exemple avant/après
Voir `examples/config/weak-tls/` (à créer selon le même schéma pour chaque configuration listée ci-dessus).

## Références
- Mozilla SSL Configuration Generator (profil Intermediate/Modern)
- OWASP Transport Layer Protection Cheat Sheet
- CWE-326: Inadequate Encryption Strength
