---
id: hardcoded-keys
category: crypto
cwe: CWE-798
owasp: A02:2021-Cryptographic Failures
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Clés cryptographiques codées en dur

## Description
Une clé de chiffrement, un secret HMAC, ou une clé privée intégrée directement dans le code source (ou un fichier de configuration versionné) devient accessible à quiconque a accès au dépôt, à l'historique Git, ou au binaire compilé. Cela annule toute garantie de confidentialité offerte par l'algorithme cryptographique, puisque la clé n'est plus un secret.

## Où ça apparaît typiquement
- Constante `SECRET_KEY`, `ENCRYPTION_KEY`, `API_SECRET` définie littéralement dans le code source.
- Clé privée PEM/PKCS embarquée dans un fichier de configuration commité.
- Valeur par défaut de clé fournie dans un template/boilerplate jamais régénérée en production.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Chaîne de caractères ressemblant à une clé (longueur/format base64, hex, PEM) assignée directement à une variable dans le code.
- Absence de lecture depuis une variable d'environnement, un coffre-fort de secrets, ou un gestionnaire de configuration externe.
- Même valeur de clé présente dans plusieurs environnements (dev/staging/prod) ou dans l'historique Git.

## Remédiation
- Charger les clés exclusivement depuis un gestionnaire de secrets (Vault, AWS/Azure/GCP Secret Manager) ou des variables d'environnement injectées de façon sécurisée.
- Régénérer immédiatement toute clé ayant été exposée dans le code ou l'historique de version.
- Mettre en place une rotation périodique des clés critiques.
- Voir `rules/remediation/hardcoded-keys.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/hardcoded-keys/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Secrets Management Cheat Sheet
- CWE-798: Use of Hard-coded Credentials
