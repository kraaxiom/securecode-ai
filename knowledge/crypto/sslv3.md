---
id: sslv3
category: crypto
cwe: CWE-326
owasp: A02:2021-Cryptographic Failures
severity_default: critical
languages: []
---

# Protocole SSLv3 activé

## Description
SSLv3 est un protocole obsolète vulnérable à l'attaque POODLE, qui permet à un attaquant en position de man-in-the-middle de déchiffrer des données via un padding oracle sur le chiffrement par bloc en mode CBC. Il est interdit par la RFC 7568 depuis 2015 et ne doit plus jamais être négociable.

## Où ça apparaît typiquement
- Configuration serveur web/mail conservant SSLv3 pour compatibilité avec de très anciens clients.
- Bibliothèques TLS embarquées dans des applications ou appliances non maintenues.
- Fallback de négociation de protocole encore actif côté serveur ou proxy.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Directive de configuration incluant `SSLv3` dans les protocoles autorisés (`SSLProtocol`, `ssl_protocols`).
- Support de la négociation de repli (downgrade) TLS non désactivé (absence de `TLS_FALLBACK_SCSV`).
- Résultat de scan TLS signalant SSLv3 comme négociable.

## Remédiation
- Désactiver explicitement SSLv3 sur tous les points d'entrée TLS.
- N'autoriser que TLS 1.2 et TLS 1.3, avec suites de chiffrement en mode authentifié (AEAD).
- Vérifier que le mécanisme anti-downgrade est actif côté serveur.
- Voir `rules/remediation/sslv3.md` pour les diffs par configuration.

## Exemple avant/après
Voir `examples/config/sslv3/` (à créer selon le même schéma pour chaque configuration listée ci-dessus).

## Références
- RFC 7568: Deprecating Secure Sockets Layer Version 3.0
- OWASP Transport Layer Protection Cheat Sheet
- CWE-326: Inadequate Encryption Strength
