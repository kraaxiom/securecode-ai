---
id: root-container
category: docker
cwe: CWE-250
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: []
---

# Conteneur exécuté en tant que root

## Description
Par défaut, un conteneur Docker exécute son processus principal en tant qu'utilisateur `root` (UID 0) s'il n'est pas configuré autrement. Bien que le namespace utilisateur limite partiellement l'impact, un processus root dans le conteneur reste root sur le système de fichiers monté et augmente significativement la surface d'une évasion de conteneur réussie.

## Où ça apparaît typiquement
- `Dockerfile` sans instruction `USER` définissant un utilisateur non privilégié.
- Image de base exécutée telle quelle sans durcissement (beaucoup d'images officielles démarrent en root par défaut).
- `docker-compose.yml` ou manifeste de déploiement sans `user:` explicite.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence d'instruction `USER` dans le `Dockerfile`, ou `USER root` explicite.
- Absence de paramètre `user:`/`securityContext.runAsNonRoot` dans la configuration de déploiement.
- Processus applicatif ne nécessitant aucun privilège root pour fonctionner (serveur web, API) mais lancé sans changement d'utilisateur.

## Remédiation
- Créer un utilisateur non privilégié dans le `Dockerfile` et l'activer via `USER`.
- Définir `runAsNonRoot: true` dans les contextes de sécurité d'orchestrateur (Kubernetes, Compose).
- Vérifier que les permissions de fichiers dans l'image sont compatibles avec l'utilisateur non root choisi.
- Voir `rules/remediation/root-container.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/root-container/` (à créer selon le même schéma).

## Références
- Docker documentation: Dockerfile best practices (USER)
- CIS Docker Benchmark
- CWE-250: Execution with Unnecessary Privileges
