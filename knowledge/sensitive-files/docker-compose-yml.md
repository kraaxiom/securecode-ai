---
id: docker-compose-yml
category: sensitive-files
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: []
---

# Exposition du fichier docker-compose.yml

## Description
`docker-compose.yml` décrit l'architecture complète des services d'une application et contient fréquemment, en clair, des identifiants de base de données, des mots de passe de services (Redis, RabbitMQ, MinIO) et des variables d'environnement sensibles définies directement dans le fichier plutôt que référencées depuis un `.env` externe. Sa divulgation révèle à la fois la topologie interne des services et des secrets exploitables immédiatement.

## Où ça apparaît typiquement
- Fichier `docker-compose.yml` copié par erreur dans le webroot ou une archive de build exposée.
- Valeurs de secrets écrites en dur dans le fichier (`POSTGRES_PASSWORD: motdepasse123`) au lieu d'une référence à `.env`.
- Dépôt public contenant le fichier avec des identifiants de production réutilisés depuis le développement.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Valeurs de mots de passe/secrets écrites en clair directement dans `docker-compose.yml` plutôt que via `env_file` ou `${VAR}`.
- Fichier présent dans un répertoire servi publiquement par le serveur web.
- Absence de `.gitignore` pour les variantes locales (`docker-compose.override.yml`) contenant des secrets de développement.

## Remédiation
- Externaliser tous les secrets vers un fichier `.env` non versionné, référencé via `${VAR}` dans `docker-compose.yml`.
- S'assurer que `docker-compose.yml` n'est jamais copié dans le webroot ni inclus dans une image publiée.
- En production, préférer un gestionnaire de secrets (Docker Secrets, Vault) aux variables d'environnement en clair.
- Voir `rules/remediation/docker-compose-yml.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/docker-compose-yml/` (fichier avant/après externalisation des secrets).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
