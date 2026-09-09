---
id: weak-encryption
category: database
cwe: CWE-326
owasp: A02:2021-Cryptographic-Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Weak Encryption (chiffrement faible des données)

## Description
Ce pattern concerne les données sensibles stockées en base de données avec un chiffrement absent, obsolète (algorithmes dépréciés) ou mal implémenté (clé statique codée en dur, mode opératoire non sécurisé, absence de sel pour le hachage de mots de passe). Une base compromise ou une fuite de données révèle alors directement les informations sensibles en clair ou facilement récupérables, alors qu'un chiffrement correctement implémenté aurait limité l'impact.

## Où ça apparaît typiquement
- Mots de passe stockés avec un hachage non adapté (MD5, SHA-1 sans sel) au lieu d'un algorithme dédié (argon2id, bcrypt).
- Données sensibles (numéros de carte, données personnelles) stockées en clair alors qu'un chiffrement au niveau colonne serait requis.
- Clé de chiffrement codée en dur dans le code source ou stockée dans la même base que les données qu'elle protège.
- Utilisation d'algorithmes ou de modes de chiffrement obsolètes (DES, ECB) pour protéger des données sensibles.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fonction de hachage de mot de passe non adaptée à cet usage (hachage rapide générique) au lieu d'un algorithme dédié avec facteur de coût.
- Colonnes contenant des données sensibles sans mécanisme de chiffrement au niveau applicatif ou base de données.
- Clé de chiffrement présente en clair dans le code source, un fichier de configuration versionné, ou la même base que les données chiffrées.
- Mode opératoire de chiffrement non authentifié (ECB, CBC sans MAC) utilisé pour des données sensibles.

## Remédiation
- Hacher les mots de passe avec un algorithme dédié à facteur de coût réglable (argon2id recommandé), jamais un hachage générique rapide.
- Chiffrer les données sensibles au niveau colonne ou applicatif avec un algorithme moderne authentifié (AES-GCM).
- Gérer les clés de chiffrement via un service dédié (KMS/HSM), séparé du stockage des données chiffrées, avec rotation régulière.
- Auditer les algorithmes cryptographiques utilisés et migrer proactivement ceux jugés obsolètes ou dépréciés.
- Voir `rules/remediation/weak-encryption.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/weak-encryption/`.

## Références
- OWASP Top 10: A02:2021 – Cryptographic Failures
- CWE-326: Inadequate Encryption Strength
