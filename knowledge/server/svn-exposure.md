---
id: svn-exposure
category: server
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: []
---

# Répertoire .svn exposé

## Description
Un répertoire `.svn` accessible publiquement via le serveur web expose les métadonnées internes de gestion de version Subversion, y compris potentiellement des copies en clair de fichiers sources dans `.svn/pristine` ou `.svn/text-base` selon la version. Comme pour l'exposition `.git`, cela permet de reconstituer tout ou partie du code source de l'application et parfois des révisions antérieures contenant des secrets déjà retirés du code actuel.

## Où ça apparaît typiquement
- Déploiement effectué via un checkout/export Subversion directement dans le webroot sans nettoyer les métadonnées `.svn`.
- Absence de règle serveur bloquant l'accès aux dossiers cachés.
- Anciennes applications maintenues avec un workflow de déploiement basé sur `svn update` en production.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Dossier `.svn` présent dans un répertoire servi publiquement par le serveur web.
- Requête sur `/.svn/entries` ou `/.svn/wc.db` retournant un contenu au lieu d'une erreur 403/404.
- Déploiement basé sur `svn update`/`svn checkout` directement dans le webroot.

## Remédiation
- Utiliser `svn export` (qui ne conserve pas les métadonnées `.svn`) plutôt qu'un checkout direct pour le déploiement.
- Bloquer explicitement au niveau du serveur web l'accès aux dossiers/fichiers commençant par un point.
- Si une exposition a eu lieu, faire tourner tout secret présent dans les révisions du dépôt.
- Voir `rules/remediation/svn-exposure.md`.

## Exemple avant/après
Voir `examples/nginx/svn-exposure/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Testing Guide: Review Webserver Metafiles for Information Leakage
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
