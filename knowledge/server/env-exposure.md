---
id: env-exposure
category: server
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Fichier .env exposé

## Description
Un fichier `.env` (ou équivalent de configuration d'environnement) accessible directement via une requête HTTP expose l'intégralité des secrets de l'application : identifiants de base de données, clés API, secrets de session, clés de chiffrement. Cette exposition provient généralement d'un fichier `.env` placé dans le webroot sans règle de blocage serveur, ou d'un déploiement qui copie l'ensemble du dépôt (y compris les fichiers non destinés au public) dans le répertoire servi.

## Où ça apparaît typiquement
- Fichier `.env` situé à la racine du webroot et accessible via `https://site/.env`.
- Absence de règle serveur (Apache/Nginx) bloquant l'accès aux fichiers commençant par un point.
- Processus de déploiement copiant tout le dépôt Git (y compris `.env`) directement dans le webroot sans étape de filtrage.
- Frameworks où le fichier `.env` n'est pas placé au-dessus du webroot par convention.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichier `.env` (ou `.env.local`, `.env.production`) présent dans un répertoire servi publiquement par le serveur web.
- Absence de directive serveur refusant l'accès aux fichiers dotfiles (`.env`, `.git`, `.htaccess`).
- Configuration de déploiement ne distinguant pas les fichiers applicatifs des fichiers de configuration sensibles.

## Remédiation
- Placer le fichier `.env` en dehors du webroot chaque fois que possible.
- Bloquer explicitement l'accès aux fichiers dotfiles au niveau du serveur web (règle `location ~ /\.` en Nginx, `<FilesMatch "^\.">` en Apache).
- Faire tourner immédiatement tout secret présent dans un `.env` ayant été exposé, même après correction de la configuration.
- Automatiser un contrôle post-déploiement vérifiant que `.env` n'est pas accessible publiquement.
- Voir `rules/remediation/env-exposure.md`.

## Exemple avant/après
Voir `examples/nginx/env-exposure/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Cheat Sheet: Secrets Management
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
