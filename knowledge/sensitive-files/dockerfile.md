---
id: dockerfile
category: sensitive-files
cwe: CWE-540
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: []
---

# Exposition et mauvaises pratiques du Dockerfile

## Description
Un `Dockerfile` exposé publiquement révèle la structure interne de build de l'application, ses dépendances et parfois des secrets injectés directement via des instructions `ENV`, `ARG` ou `COPY` de fichiers sensibles (clés, `.env`). Même sans exposition directe du fichier, des secrets passés via `ARG`/`ENV` restent visibles dans l'historique des layers de l'image Docker, accessible à quiconque peut inspecter ou télécharger l'image.

## Où ça apparaît typiquement
- Instructions `ENV`/`ARG` utilisées pour passer des secrets de build (clés API, mots de passe) au lieu de secrets de build dédiés.
- `COPY . .` sans `.dockerignore` adéquat, embarquant `.env`, `.git`, ou des clés privées dans l'image.
- `Dockerfile` lui-même accessible via le webroot suite à une copie non filtrée du dépôt.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Utilisation de `ENV SECRET=...` ou `ARG SECRET=...` pour des valeurs sensibles au lieu de `--mount=type=secret` (BuildKit).
- Absence de fichier `.dockerignore` ou `.dockerignore` incomplet (n'excluant pas `.env`, `.git`, `*.pem`).
- Image publiée sur un registre sans scan de contenu ni vérification des layers pour secrets embarqués.

## Remédiation
- Utiliser les secrets de build BuildKit (`--mount=type=secret`) plutôt que `ARG`/`ENV` pour toute donnée sensible.
- Maintenir un `.dockerignore` complet excluant `.env`, `.git`, les clés privées et les fichiers de sauvegarde.
- Scanner les images publiées à la recherche de secrets embarqués avant publication sur un registre.
- Voir `rules/remediation/dockerfile.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/dockerfile/` (Dockerfile avant/après utilisation de secrets de build).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-540: Inclusion of Sensitive Information in Source Code
