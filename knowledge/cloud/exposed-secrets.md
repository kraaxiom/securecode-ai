---
id: exposed-secrets
category: cloud
cwe: CWE-798
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Secrets exposés dans le code ou la configuration cloud

## Description
Des secrets exposés désignent des identifiants sensibles (clés API, tokens, chaînes de connexion, clés privées) codés en dur dans le code source, committés dans un dépôt Git, ou stockés en clair dans des fichiers de configuration/variables d'environnement accessibles publiquement (metadata cloud, fichiers `.env` servis par le serveur web). Une fois exposé, un secret doit être considéré compromis même après suppression, car il peut rester dans l'historique Git ou avoir déjà été indexé/scrapé.

## Où ça apparaît typiquement
- Clé API/token codé en dur dans un fichier source, committé puis poussé sur un dépôt public ou privé partagé.
- Fichier `.env` ou `config.php` accessible directement via le serveur web (mauvaise configuration du webroot).
- Secrets injectés en clair dans des metadata d'instance cloud (user-data) sans chiffrement.
- Historique Git contenant un secret supprimé dans un commit ultérieur mais toujours présent dans l'historique.
- Secrets partagés par messagerie/tickets et recopiés dans le code au lieu d'un gestionnaire de secrets.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Motifs reconnaissables de secrets (préfixes de tokens de fournisseurs cloud, blocs `-----BEGIN ... PRIVATE KEY-----`, chaînes de connexion avec identifiants) dans le code source ou les fichiers de configuration versionnés.
- Fichier `.env`/`.git` présent dans le webroot et potentiellement servi par le serveur.
- Absence de `.gitignore` couvrant les fichiers de configuration sensibles.
- Variables d'environnement en clair dans des logs applicatifs ou des messages d'erreur.

## Remédiation
- Utiliser un gestionnaire de secrets dédié (Vault, AWS Secrets Manager, Azure Key Vault, GCP Secret Manager) plutôt que des fichiers en clair.
- Ajouter les fichiers de configuration sensibles au `.gitignore` avant tout premier commit.
- Faire tourner (révoquer et régénérer) immédiatement tout secret ayant fuité, même après suppression du code.
- Mettre en place un scan automatique de secrets en pre-commit et en CI (ex: détection de motifs de type gitleaks/trufflehog).
- Voir `rules/remediation/exposed-secrets.md`.

## Exemple avant/après
Voir `examples/php/exposed-secrets/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Cheat Sheet: Secrets Management
- CWE-798: Use of Hard-coded Credentials
