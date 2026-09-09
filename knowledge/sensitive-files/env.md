---
id: env
category: sensitive-files
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition du fichier .env

## Description
Le fichier `.env` contient généralement les secrets de configuration d'une application : identifiants de base de données, clés API, secrets de session, jetons de services tiers. Lorsqu'il est déployé dans la racine web publique et accessible directement via HTTP, n'importe qui peut le télécharger et obtenir un accès complet aux systèmes qu'il protège. C'est l'une des fuites de configuration les plus fréquentes et les plus critiques observées sur des applications PHP (Laravel, Symfony) et Node.js.

## Où ça apparaît typiquement
- Racine du projet déployée telle quelle sur le serveur web sans exclusion du `.env`.
- Absence de règle de blocage dans la configuration Apache/Nginx pour les fichiers commençant par un point.
- Pipelines de déploiement qui copient l'intégralité du dépôt (y compris `.env`) dans le répertoire servi publiquement.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `.env` présent dans le webroot déployé (ex: `public/`, `www/`, `htdocs/`) au lieu d'être hors racine web.
- Absence de directive serveur bloquant l'accès aux fichiers dotfiles (`.env`, `.git`, etc.).
- `.env` absent du `.gitignore`/`.dockerignore` alors que le projet en dépend.
- Réponse HTTP 200 avec contenu texte brut sur une requête directe vers `/.env`.

## Remédiation
- Stocker le `.env` hors du répertoire servi publiquement par le serveur web.
- Ajouter une règle explicite de blocage des dotfiles dans la configuration Nginx/Apache.
- Utiliser un gestionnaire de secrets (Vault, AWS Secrets Manager, variables d'environnement injectées par la plateforme) plutôt qu'un fichier statique en production.
- Voir `rules/remediation/env.md` pour les diffs par configuration serveur.

## Exemple avant/après
Voir `examples/config/env/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
