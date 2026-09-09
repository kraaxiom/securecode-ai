---
id: gitlab-secrets
category: cicd
cwe: CWE-798
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Fuite de secrets dans GitLab CI

## Description
Une fuite de secrets GitLab CI survient quand des identifiants sont codés en dur dans `.gitlab-ci.yml`, exposés dans les logs de job, ou accessibles par des branches/merge requests non fiables via des variables CI/CD trop permissives. Les variables protégées mal configurées (non limitées aux branches/tags protégés) peuvent être lues par n'importe quel job, y compris ceux déclenchés depuis un fork ou une branche de contribution externe.

## Où ça apparaît typiquement
- Identifiants en clair dans `.gitlab-ci.yml` au lieu de variables CI/CD masquées.
- `echo`/`cat`/`export -p` affichant une variable sensible dans le log du job.
- Variables CI/CD non marquées "Protected" ou "Masked" dans les paramètres du projet, accessibles depuis n'importe quelle branche.
- Runners partagés/self-hosted mal isolés réutilisant le cache ou l'environnement entre projets.
- Artéfacts de job (`artifacts:`) contenant des fichiers de configuration avec secrets.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Motif de secret (clé API, token, chaîne de connexion) en clair dans le YAML du pipeline.
- Variable CI définie dans les paramètres du projet sans coche "Masked" ni "Protected".
- Script de job qui imprime une variable d'environnement sensible ou la redirige vers un fichier inclus dans les artefacts.
- Utilisation de runners partagés pour des jobs manipulant des secrets de production.

## Remédiation
- Déclarer les secrets comme variables CI/CD "Protected" (limitées aux branches/tags protégés) et "Masked".
- Ne jamais journaliser une variable sensible ; utiliser des mécanismes de secrets externes (Vault, GitLab KMS) pour les environnements critiques.
- Restreindre les artefacts et caches pour ne jamais inclure de fichiers contenant des secrets.
- Limiter l'exécution des jobs sensibles aux runners dédiés et isolés.
- Voir `rules/remediation/gitlab-secrets.md`.

## Exemple avant/après
Voir `examples/yaml/gitlab-secrets/`.

## Références
- OWASP Cheat Sheet: CI/CD Security
- GitLab Docs: CI/CD variable security
- CWE-798: Use of Hard-coded Credentials
