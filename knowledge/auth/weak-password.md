---
id: weak-password
category: auth
cwe: CWE-521
owasp: A07:2021-Identification and Authentication Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Politique de mot de passe faible

## Description
Une politique de mot de passe faible survient quand une application n'impose pas d'exigences suffisantes sur les mots de passe choisis par les utilisateurs (longueur minimale insuffisante, absence de vérification contre les mots de passe les plus courants ou déjà compromis) ou impose au contraire des règles contre-productives (rotation forcée fréquente, complexité artificielle) qui poussent les utilisateurs vers des choix prévisibles. Ces mots de passe faibles facilitent le brute force et le credential stuffing.

## Où ça apparaît typiquement
- Formulaires d'inscription/changement de mot de passe sans contrôle de longueur minimale raisonnable.
- Absence de vérification contre une liste de mots de passe compromis connus.
- Politiques imposant une rotation périodique obligatoire sans justification (favorise les variations prévisibles).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Validation de mot de passe reposant uniquement sur une longueur minimale faible (ex. inférieure à 8-10 caractères) sans autre contrôle.
- Absence d'appel à un service/liste de vérification de mots de passe compromis lors de l'inscription ou du changement.
- Règles de complexité imposant des contraintes de caractères spéciaux/majuscules sans contrôle de longueur, contraires aux recommandations actuelles (NIST 800-63B).

## Remédiation
- Exiger une longueur minimale d'au moins 12 caractères, sans imposer de règles de composition artificielles.
- Vérifier les mots de passe choisis contre une liste de mots de passe compromis connus (ex. via un service de type "have I been pwned" côté serveur).
- Ne pas imposer de rotation périodique sans motif (fuite avérée), conformément aux recommandations modernes.
- Voir `rules/remediation/weak-password.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-django/weak-password/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Authentication
- CWE-521: Weak Password Requirements
