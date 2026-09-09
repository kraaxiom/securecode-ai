---
id: backup-files
category: server
cwe: CWE-530
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: []
---

# Fichiers de sauvegarde exposés

## Description
Des fichiers de sauvegarde exposés (`.bak`, `.old`, `.zip`, `~`, copies horodatées) laissés dans le webroot par des éditeurs de texte, des scripts de déploiement ou des sauvegardes manuelles peuvent contenir le code source complet de l'application, y compris des identifiants de base de données ou des clés en dur. Contrairement au fichier original (souvent un `.php` exécuté par le serveur), une copie de sauvegarde est fréquemment servie en texte brut, révélant tout son contenu.

## Où ça apparaît typiquement
- Fichiers laissés par des éditeurs (`file.php~`, `file.php.swp`) après édition directe sur le serveur de production.
- Archives de déploiement (`backup.zip`, `site-2024.tar.gz`) copiées dans le webroot puis oubliées.
- Scripts de sauvegarde automatisés écrivant leur sortie dans un dossier accessible publiquement.
- Copies renommées avant une modification risquée (`config.php.old`, `index_backup.php`).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence de fichiers avec extensions/suffixes typiques de sauvegarde (`.bak`, `.old`, `.orig`, `~`, `.swp`, `.zip`, `.tar.gz`) dans un répertoire servi publiquement.
- Scripts de déploiement/CI écrivant des archives ou copies dans le webroot sans nettoyage post-déploiement.
- Absence de règle serveur bloquant explicitement l'accès à ces extensions.

## Remédiation
- Interdire au niveau du serveur web l'accès aux extensions de sauvegarde courantes (`.bak`, `.old`, `~`, `.zip`, etc.) sur l'ensemble du webroot.
- Ne jamais éditer de fichiers directement en production ; utiliser un pipeline de déploiement qui ne laisse aucun artefact temporaire.
- Stocker les sauvegardes en dehors du webroot, idéalement dans un stockage dédié et chiffré.
- Voir `rules/remediation/backup-files.md`.

## Exemple avant/après
Voir `examples/nginx/backup-files/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Testing Guide: Review Old Backup and Unreferenced Files
- CWE-530: Exposure of Backup File to an Unauthorized Control Sphere
