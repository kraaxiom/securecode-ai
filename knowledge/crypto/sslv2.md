---
id: sslv2
category: crypto
cwe: CWE-326
owasp: A02:2021-Cryptographic Failures
severity_default: critical
languages: []
---

# Protocole SSLv2 activé

## Description
SSLv2 est un protocole de chiffrement de transport publié en 1995 et officiellement interdit par la RFC 6176 depuis 2011. Il présente des failles structurelles graves (négociation non authentifiée, MAC faible) qui permettent des attaques de type interception ou downgrade, notamment exploitées par l'attaque DROWN contre des serveurs partageant une même clé RSA.

## Où ça apparaît typiquement
- Configuration serveur web/mail (Apache, Nginx, Postfix) autorisant explicitement SSLv2.
- Load balancers ou reverse proxies avec une configuration TLS héritée non durcie.
- Équipements réseau ou appliances legacy dont le firmware n'a jamais été mis à jour.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Directive de configuration incluant `SSLv2` dans la liste des protocoles autorisés (`SSLProtocol`, `ssl_protocols`).
- Absence de directive excluant explicitement SSLv2/SSLv3.
- Résultat de scan de configuration TLS signalant SSLv2 comme protocole négociable.

## Remédiation
- Désactiver explicitement SSLv2 (et SSLv3) dans toute configuration serveur.
- N'autoriser que TLS 1.2 et TLS 1.3.
- Auditer les certificats RSA partagés entre serveurs pour écarter le risque DROWN.
- Voir `rules/remediation/sslv2.md` pour les diffs par configuration.

## Exemple avant/après
Voir `examples/config/sslv2/` (à créer selon le même schéma pour chaque configuration listée ci-dessus).

## Références
- RFC 6176: Prohibiting Secure Sockets Layer (SSL) Version 2.0
- OWASP Transport Layer Protection Cheat Sheet
- CWE-326: Inadequate Encryption Strength
