---
id: wp-config-php
category: sensitive-files
cwe: CWE-540
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition du fichier wp-config.php

## Description
`wp-config.php` est le fichier central de configuration de WordPress : il contient les identifiants de connexion à la base de données, les clés de sécurité (`AUTH_KEY`, `SECRET_KEY`, etc.) et parfois des clés d'API tierces. Sa divulgation, généralement via une copie de sauvegarde mal nommée ou une erreur de configuration serveur, donne un accès direct à la base de données du site et permet de forger des cookies d'authentification valides grâce aux clés de sécurité exposées.

## Où ça apparaît typiquement
- Copies de sauvegarde (`wp-config.php.bak`, `wp-config.old`, `wp-config.php~`) laissées à la racine après une migration ou un plugin de sauvegarde.
- Éditeurs de code laissant des fichiers temporaires accessibles à côté de `wp-config.php`.
- Serveur mal configuré ne traitant plus l'extension `.php` comme exécutable dans certains contextes (sous-domaines, reverse proxy).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence de variantes de nommage (`wp-config.php.bak`, `wp-config-backup.php`, `wp-config.php.save`) dans le webroot.
- Absence de règle serveur bloquant les extensions de sauvegarde courantes autour de `wp-config.php`.
- Plugins de sauvegarde WordPress générant des archives accessibles publiquement contenant `wp-config.php`.

## Remédiation
- Interdire au niveau serveur toute extension dérivée non exécutable pour `wp-config.php` (`*.bak`, `*~`, `*.save`, `*.old`).
- Déplacer `wp-config.php` un niveau au-dessus du webroot lorsque l'hébergement le permet (WordPress le supporte nativement).
- Régénérer immédiatement les clés de sécurité et changer les identifiants de base de données en cas d'exposition suspectée.
- Voir `rules/remediation/wp-config-php.md` pour les diffs par configuration serveur.

## Exemple avant/après
Voir `examples/config/wp-config-php/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-540: Inclusion of Sensitive Information in Source Code
