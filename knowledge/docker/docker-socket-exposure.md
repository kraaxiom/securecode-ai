---
id: docker-socket-exposure
category: docker
cwe: CWE-250
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Montage du socket Docker (`/var/run/docker.sock`) dans un conteneur

## Description
Monter le socket Docker de l'hôte à l'intérieur d'un conteneur donne à ce conteneur un contrôle total sur le daemon Docker, et donc sur l'hôte lui-même : il peut créer un nouveau conteneur privilégié montant le système de fichiers racine. C'est une pratique courante dans des outils de CI/CD ou de monitoring, mais elle équivaut à donner un accès root à l'hôte.

## Où ça apparaît typiquement
- Directive `-v /var/run/docker.sock:/var/run/docker.sock` dans une commande `docker run` ou un `docker-compose.yml`.
- Outils "Docker-in-Docker" ou agents CI (Jenkins, GitLab Runner) montant le socket pour builder des images.
- Sidecars de monitoring/logging nécessitant d'inspecter les conteneurs voisins.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence de `/var/run/docker.sock` dans une section `volumes` de `docker-compose.yml` ou un manifeste de déploiement.
- Conteneur ayant à la fois le socket monté et des capacités étendues ou le mode privilégié.
- Absence de proxy d'API Docker filtrant les commandes autorisées (ex: `docker-socket-proxy`).

## Remédiation
- Éviter de monter le socket Docker dans un conteneur applicatif ; privilégier des architectures sans accès direct au daemon.
- Si indispensable, passer par un proxy d'API Docker restreignant les endpoints accessibles en lecture seule.
- Isoler les workloads nécessitant ce montage dans un environnement dédié et surveillé, jamais exposé à du code non fiable.
- Voir `rules/remediation/docker-socket-exposure.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/docker-socket-exposure/` (à créer selon le même schéma).

## Références
- Docker documentation: Protect the Docker daemon socket
- CIS Docker Benchmark
- CWE-250: Execution with Unnecessary Privileges
