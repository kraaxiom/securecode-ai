---
id: tls1-0
category: crypto
cwe: CWE-326
owasp: A02:2021-Cryptographic Failures
severity_default: high
languages: []
---

# Protocole TLS 1.0 (et 1.1) activé

## Description
TLS 1.0 et 1.1 sont des protocoles obsolètes officiellement dépréciés par la RFC 8996 (2021). Ils ne supportent pas les suites de chiffrement authentifiées modernes, restent exposés à des attaques comme BEAST, et sont exclus des référentiels de conformité actuels (PCI-DSS notamment) depuis 2018.

## Où ça apparaît typiquement
- Configuration serveur web/API acceptant encore TLS 1.0/1.1 pour compatibilité client.
- Intégrations avec des systèmes tiers legacy imposant un ancien protocole.
- SDK ou bibliothèques clientes figeant explicitement une version minimale trop basse.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Directive de configuration incluant `TLSv1` ou `TLSv1.1` dans les protocoles autorisés.
- Paramètre client HTTP fixant explicitement `TLSv1` comme version minimale ou maximale.
- Absence de restriction imposant TLS 1.2 minimum dans la configuration serveur.

## Remédiation
- Configurer un minimum de TLS 1.2, avec support préférentiel de TLS 1.3.
- Retirer les suites de chiffrement non-AEAD associées aux anciennes versions.
- Planifier la migration des clients/intégrations legacy qui dépendent encore de TLS 1.0/1.1.
- Voir `rules/remediation/tls1-0.md` pour les diffs par configuration.

## Exemple avant/après
Voir `examples/config/tls1-0/` (à créer selon le même schéma pour chaque configuration listée ci-dessus).

## Références
- RFC 8996: Deprecating TLS 1.0 and TLS 1.1
- OWASP Transport Layer Protection Cheat Sheet
- CWE-326: Inadequate Encryption Strength
