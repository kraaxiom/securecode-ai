---
id: smtp-injection
category: injections
cwe: CWE-93
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp]
---

# SMTP Injection (Email Header Injection)

## Description
L'injection SMTP, aussi appelée injection d'en-têtes email, survient lorsqu'une entrée utilisateur (ex: champ "nom" ou "sujet" d'un formulaire de contact) est intégrée sans neutralisation dans les en-têtes d'un message envoyé via SMTP. En injectant des caractères CR/LF, un attaquant peut ajouter des en-têtes arbitraires (destinataires en copie, changement de sujet) voire injecter un corps de message complet, transformant l'application en relais de spam ou de phishing.

## Où ça apparaît typiquement
- Formulaires de contact ou de newsletter construisant les en-têtes email (`From`, `Subject`, `To`) à partir de champs utilisateur.
- Fonctions d'envoi de mail bas niveau (`mail()`, `smtplib`) appelées avec des en-têtes concaténés manuellement.
- Fonctionnalités de partage/notification par email où le sujet ou le nom de l'expéditeur provient de l'utilisateur.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation directe d'un champ utilisateur dans une chaîne d'en-tête email sans filtrage des caractères `\r`/`\n`.
- Utilisation de la fonction native `mail()`/`sendmail` avec des en-têtes construits manuellement plutôt qu'une bibliothèque d'email dédiée.
- Absence de validation du format attendu pour les champs sensibles (adresse email au format strict).

## Remédiation
- Utiliser une bibliothèque d'envoi d'email mature qui échappe automatiquement les en-têtes (ex: PHPMailer, Django `send_mail`, Nodemailer).
- Rejeter tout caractère de contrôle (`\r`, `\n`) dans les champs utilisés pour construire des en-têtes.
- Valider strictement le format des adresses email et limiter la longueur des champs libres.
- Voir `rules/remediation/smtp-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/smtp-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Testing Guide: Testing for IMAP/SMTP Injection
- CWE-93: Improper Neutralization of CRLF Sequences ('CRLF Injection')
