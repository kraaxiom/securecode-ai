---
id: config-php
category: sensitive-files
cwe: CWE-540
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: []
---

# Exposition du fichier config.php

## Description
De nombreuses applications PHP historiques stockent leurs identifiants de base de données et secrets applicatifs directement dans un fichier `config.php`. Si ce fichier est accessible sans être interprété par le moteur PHP (mauvaise configuration serveur, extension modifiée, sauvegarde `.php.bak`) ou si une erreur de configuration permet son téléchargement en texte brut, ses secrets sont directement exposés à toute personne accédant à l'URL.

## Où ça apparaît typiquement
- Copies de sauvegarde du fichier (`config.php.bak`, `config.php~`, `config.old.php`) laissées dans le webroot.
- Serveur web mal configuré ne traitant plus les fichiers `.php` comme du code exécutable après un changement d'extension.
- Fichier `config.php` inclus dans une archive de sauvegarde exposée publiquement.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichiers avec extension double ou modifiée (`config.php.bak`, `config.php.old`, `config.php~`) présents dans le webroot.
- Éditeurs laissant des fichiers temporaires (`.swp`, `.tmp`) à côté de `config.php` accessibles publiquement.
- Absence de règle serveur bloquant explicitement les extensions de sauvegarde courantes.

## Remédiation
- Ne jamais laisser de copies de sauvegarde de fichiers de configuration dans le webroot.
- Bloquer au niveau serveur toute extension non exécutable dérivée d'un fichier sensible (`*.bak`, `*~`, `*.old`, `*.swp`).
- Déplacer les secrets vers des variables d'environnement ou un gestionnaire de secrets plutôt qu'un fichier de configuration statique.
- Voir `rules/remediation/config-php.md` pour les diffs par configuration serveur.

## Exemple avant/après
Voir `examples/config/config-php/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-540: Inclusion of Sensitive Information in Source Code
