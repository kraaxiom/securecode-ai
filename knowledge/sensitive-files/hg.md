---
id: hg
category: sensitive-files
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: []
---

# Exposition du répertoire .hg

## Description
Le répertoire `.hg` (Mercurial) contient l'historique complet des changesets d'un dépôt, y compris le code source et son évolution. Comme pour `.git` et `.svn`, un déploiement qui copie directement la working copy Mercurial dans le webroot expose l'historique complet du projet à toute personne connaissant ou devinant l'URL, permettant la reconstruction intégrale du code source et de son historique de modifications.

## Où ça apparaît typiquement
- Déploiement par clone Mercurial direct dans le répertoire servi publiquement.
- Absence de règle serveur bloquant les dossiers `.hg`.
- Anciens projets Mercurial déployés sans étape de build/export séparée du contrôle de version.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence de `.hg/` à la racine du webroot déployé.
- Réponse HTTP 200 sur une requête vers `/.hg/store` ou `/.hg/requires`.
- Absence de directive de blocage des dotfiles/dossiers cachés dans la configuration du serveur web.

## Remédiation
- Déployer via un export/archive (`hg archive`) plutôt qu'un clone direct dans le webroot.
- Bloquer explicitement l'accès aux répertoires `.hg` au niveau du serveur web.
- Séparer strictement le répertoire de contrôle de version du répertoire servi en production.
- Voir `rules/remediation/hg.md` pour les diffs par configuration serveur.

## Exemple avant/après
Voir `examples/config/hg/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
