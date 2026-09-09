---
id: ds-store
category: sensitive-files
cwe: CWE-200
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: []
---

# Exposition du fichier .DS_Store

## Description
`.DS_Store` est un fichier généré automatiquement par macOS Finder dans chaque répertoire parcouru, contenant des métadonnées d'affichage mais aussi, souvent, la liste des noms de fichiers et sous-répertoires présents à cet emplacement au moment de sa création. Exposé publiquement, il permet à un attaquant d'énumérer le contenu d'un répertoire sans que le listing de répertoire (`Directory Listing`) ne soit activé, révélant potentiellement des fichiers non liés depuis l'application (scripts de debug, sauvegardes oubliées).

## Où ça apparaît typiquement
- Développeurs macOS déployant leur répertoire de travail directement sur le serveur, `.DS_Store` inclus.
- Absence de règle serveur bloquant les fichiers `.DS_Store`.
- Fichier commité par erreur dans le dépôt puis inclus dans l'artefact de déploiement.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichiers `.DS_Store` présents à un ou plusieurs niveaux du webroot déployé.
- Absence de `.gitignore` global excluant `.DS_Store` pour tous les contributeurs macOS.
- Absence de directive serveur bloquant l'accès aux dotfiles.

## Remédiation
- Ajouter `.DS_Store` à un `.gitignore` global et à l'exclusion de build/déploiement.
- Bloquer explicitement l'accès aux dotfiles au niveau du serveur web.
- Auditer périodiquement le webroot déployé pour détecter des fichiers de métadonnées système oubliés.
- Voir `rules/remediation/ds-store.md` pour les diffs de configuration serveur.

## Exemple avant/après
Voir `examples/config/ds-store/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
