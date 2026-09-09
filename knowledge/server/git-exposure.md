---
id: git-exposure
category: server
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Répertoire .git exposé

## Description
Un répertoire `.git` accessible publiquement via le serveur web permet à un attaquant de reconstituer l'intégralité de l'historique du dépôt (objets, commits, branches) en téléchargeant les fichiers internes de Git. Même si le code actuellement déployé ne contient plus de secrets, l'historique Git peut révéler des identifiants supprimés dans des commits antérieurs, du code source complet, ou des informations sur l'architecture interne.

## Où ça apparaît typiquement
- Déploiement effectué par un simple `git clone`/`git pull` directement dans le webroot, sans exclure `.git`.
- Absence de règle serveur bloquant l'accès aux dossiers cachés (dotfiles/dotdirs).
- Pipelines CI/CD copiant l'ensemble du répertoire de travail (y compris `.git`) vers le serveur de production.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Dossier `.git` présent dans un répertoire servi publiquement par le serveur web.
- Requête sur `/.git/HEAD` ou `/.git/config` retournant un contenu au lieu d'une erreur 403/404.
- Processus de déploiement basé sur `git pull` en production sans séparation entre dépôt et webroot.

## Remédiation
- Ne jamais déployer en clonant directement le dépôt dans le webroot ; utiliser un build/artefact séparé du répertoire `.git`.
- Bloquer explicitement au niveau du serveur web l'accès à tous les dossiers/fichiers commençant par un point.
- Si une exposition a eu lieu, considérer tout secret présent dans l'historique comme compromis et le faire tourner.
- Voir `rules/remediation/git-exposure.md`.

## Exemple avant/après
Voir `examples/nginx/git-exposure/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Testing Guide: Review Webserver Metafiles for Information Leakage
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
