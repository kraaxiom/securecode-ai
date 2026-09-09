---
id: backup-zip
category: sensitive-files
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition d'une archive de sauvegarde (.zip/.tar.gz)

## Description
Les archives de sauvegarde complètes d'un site (code source, configuration, parfois base de données) sont souvent générées par des scripts ou des plugins d'administration et nommées de façon prévisible (`backup.zip`, `site-backup.tar.gz`, `nomdedomaine.zip`). Lorsqu'elles restent accessibles dans le webroot après leur création, elles offrent à un attaquant l'intégralité du code source et de la configuration de l'application en un seul téléchargement, sans effort d'énumération.

## Où ça apparaît typiquement
- Plugins de sauvegarde CMS (WordPress, Joomla) écrivant l'archive dans un répertoire accessible publiquement par défaut.
- Scripts d'administration générant une archive temporaire dans le webroot avant transfert externe, jamais nettoyée.
- Noms de fichiers prévisibles basés sur la date ou le nom de domaine, facilitant leur découverte.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichiers d'archive (`.zip`, `.tar.gz`, `.rar`, `.7z`) présents dans le répertoire servi publiquement.
- Nom de fichier suivant un motif prévisible (date, nom de domaine, "backup").
- Absence de nettoyage automatique après génération dans le processus de sauvegarde.

## Remédiation
- Générer les sauvegardes dans un répertoire strictement hors webroot, avec transfert direct vers un stockage externe sécurisé.
- Bloquer au niveau serveur l'accès direct aux extensions d'archive dans les répertoires publics.
- Automatiser la suppression des archives temporaires immédiatement après leur transfert.
- Voir `rules/remediation/backup-zip.md` pour les diffs par configuration serveur.

## Exemple avant/après
Voir `examples/config/backup-zip/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
