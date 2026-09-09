---
id: privileged-container
category: docker
cwe: CWE-250
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Conteneur en mode privilégié

## Description
Le mode privilégié (`--privileged`) désactive quasiment toutes les protections d'isolation de Docker : le conteneur obtient l'accès à l'ensemble des périphériques de l'hôte, peut charger des modules kernel, et contourne les restrictions de capabilities, seccomp et AppArmor. Un attaquant compromettant un tel conteneur peut généralement s'échapper vers l'hôte.

## Où ça apparaît typiquement
- Option `--privileged` sur une commande `docker run` ou `privileged: true` dans `docker-compose.yml`.
- Conteneurs nécessitant un accès matériel direct (ex: outils de virtualisation imbriquée) configurés sans restriction fine.
- Images tierces demandant le mode privilégié par simplicité plutôt que par nécessité réelle.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence de `privileged: true` ou `--privileged` dans un fichier de composition ou une commande de lancement.
- Absence de liste explicite de `cap_add` restreinte (le mode privilégié rend cette liste inutile car tout est déjà accordé).
- Montage combiné avec des volumes sensibles (`/`, `/dev`, `/proc`) renforçant le risque d'évasion.

## Remédiation
- Ne jamais utiliser `--privileged` sauf nécessité technique strictement justifiée et isolée.
- Accorder uniquement les capabilities Linux réellement nécessaires via `cap_add`, en partant d'une base `cap_drop: ALL`.
- Activer les profils seccomp et AppArmor/SELinux par défaut plutôt que de les désactiver.
- Voir `rules/remediation/privileged-container.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/privileged-container/` (à créer selon le même schéma).

## Références
- Docker documentation: Runtime privilege and Linux capabilities
- CIS Docker Benchmark
- CWE-250: Execution with Unnecessary Privileges
