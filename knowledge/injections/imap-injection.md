---
id: imap-injection
category: injections
cwe: CWE-93
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp]
---

# IMAP Injection

## Description
L'injection IMAP se produit lorsqu'une entrée utilisateur est insérée sans neutralisation dans une commande IMAP (protocole de gestion de courrier électronique), permettant à un attaquant d'injecter des commandes IMAP supplémentaires. Cela peut conduire à la lecture ou manipulation de messages d'autres utilisateurs, ou à un contournement de l'authentification sur le serveur de messagerie, selon l'implémentation.

## Où ça apparaît typiquement
- Webmails ou intégrations applicatives qui construisent dynamiquement des commandes IMAP (recherche, sélection de dossier) à partir d'entrées utilisateur.
- Formulaires de recherche dans une boîte mail transmettant directement le terme de recherche à la commande `SEARCH`.
- Systèmes d'authentification déléguant à IMAP avec des identifiants insuffisamment échappés.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de variables utilisateur dans une chaîne de commande IMAP sans échappement des littéraux.
- Absence d'utilisation d'une bibliothèque cliente IMAP encodant correctement les arguments (littéraux de chaîne préfixés par leur longueur).
- Entrées utilisateur transmises telles quelles à des fonctions bas niveau de dialogue IMAP.

## Remédiation
- Utiliser une bibliothèque cliente IMAP mature qui encode correctement les arguments en littéraux.
- Valider et limiter le jeu de caractères autorisé dans les termes transmis (recherche, noms de dossiers).
- Ne jamais construire de commandes IMAP par concaténation de chaînes brutes.
- Voir `rules/remediation/imap-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/imap-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Testing Guide: Testing for IMAP/SMTP Injection
- CWE-93: Improper Neutralization of CRLF Sequences ('CRLF Injection')
