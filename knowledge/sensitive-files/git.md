---
id: git
category: sensitive-files
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition du répertoire .git

## Description
Un répertoire `.git` accessible publiquement expose l'intégralité de l'historique du dépôt source, y compris des fichiers supprimés, des secrets commités par erreur puis retirés, et le code source complet même si l'application ne le sert pas normalement. Des outils automatisés permettent de reconstruire l'arborescence complète du projet à partir des objets Git exposés, ce qui équivaut souvent à une fuite totale du code source et de son historique.

## Où ça apparaît typiquement
- Déploiement effectué par simple copie/clone du dépôt (`git clone` ou `git pull`) directement dans le webroot.
- Absence de règle serveur bloquant l'accès aux répertoires cachés (`.git`, `.svn`, `.hg`).
- Pipelines CI/CD qui déploient l'arborescence complète du projet sans étape de nettoyage.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `.git/` présent dans le répertoire servi publiquement par le serveur web.
- Réponse HTTP 200 sur une requête directe vers `/.git/HEAD` ou `/.git/config`.
- Absence de directive de blocage des répertoires dotfiles dans la configuration serveur.

## Remédiation
- Déployer uniquement les artefacts buildés, jamais le répertoire `.git`, via un pipeline dédié (build séparé du clone).
- Ajouter une règle explicite bloquant l'accès à `.git/` au niveau du serveur web.
- Auditer l'historique Git à la recherche de secrets commités si une exposition a eu lieu, et les révoquer/rotater.
- Voir `rules/remediation/git.md` pour les diffs par configuration serveur.

## Exemple avant/après
Voir `examples/config/git/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
