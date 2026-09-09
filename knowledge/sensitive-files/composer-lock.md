---
id: composer-lock
category: sensitive-files
cwe: CWE-200
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: []
---

# Exposition du fichier composer.lock

## Description
`composer.lock` liste précisément les versions de toutes les dépendances PHP installées dans le projet. Sans être un secret en soi, sa divulgation permet à un attaquant d'identifier rapidement les versions exactes de librairies utilisées et de croiser cette liste avec des bases de vulnérabilités connues (CVE), facilitant le ciblage de failles déjà publiques dans des dépendances obsolètes.

## Où ça apparaît typiquement
- Fichier `composer.lock` déployé à la racine du webroot avec le reste du projet.
- Absence de règle serveur bloquant l'accès direct aux fichiers de manifeste de dépendances.
- Pipelines de déploiement copiant l'intégralité du dépôt sans distinction entre code exécutable et métadonnées de build.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `composer.lock` accessible directement via une requête HTTP sur le webroot déployé.
- Absence de règle serveur bloquant les fichiers `.lock`/manifestes de gestion de dépendances.
- Dépendances listées avec des versions anciennes non mises à jour depuis longtemps (signal de risque additionnel).

## Remédiation
- Bloquer l'accès direct aux fichiers de manifeste de dépendances (`composer.lock`, `composer.json`) au niveau du serveur web.
- Maintenir les dépendances à jour et surveiller les avis de sécurité via `composer audit` ou un outil équivalent en CI.
- Ne déployer que les fichiers strictement nécessaires à l'exécution, hors métadonnées de build.
- Voir `rules/remediation/composer-lock.md` pour les diffs de configuration serveur.

## Exemple avant/après
Voir `examples/config/composer-lock/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
