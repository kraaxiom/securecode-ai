---
id: brute-force
category: auth
cwe: CWE-307
owasp: A07:2021-Identification and Authentication Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Brute Force (absence de limitation des tentatives d'authentification)

## Description
Une vulnérabilité de type brute force existe quand un mécanisme d'authentification ne limite pas le nombre de tentatives possibles pour un même compte ou une même adresse IP. Sans verrouillage progressif, limitation de débit ou CAPTCHA, un attaquant peut essayer un grand nombre de mots de passe de façon automatisée jusqu'à en trouver un valide, en particulier sur des comptes à mot de passe faible.

## Où ça apparaît typiquement
- Formulaires de connexion classiques sans compteur d'échecs ni verrouillage.
- API d'authentification (login, réinitialisation de mot de passe, code OTP) sans limitation de débit.
- Endpoints de vérification de code à usage unique (MFA, réinitialisation) avec un espace de valeurs restreint.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Endpoint d'authentification sans middleware de rate limiting ni compteur d'échecs persistant par compte/IP.
- Absence de verrouillage temporaire ou de délai croissant après plusieurs échecs consécutifs.
- Messages d'erreur ou temps de réponse permettant de distinguer un compte existant d'un compte inexistant (facilite l'énumération combinée au brute force).

## Remédiation
- Mettre en place une limitation de débit (rate limiting) par compte et par IP sur tous les endpoints d'authentification.
- Verrouiller temporairement un compte après un nombre défini d'échecs, avec délai progressif.
- Ajouter un CAPTCHA ou un défi équivalent après plusieurs échecs.
- Journaliser et alerter sur les pics de tentatives d'authentification échouées.
- Voir `rules/remediation/brute-force.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/brute-force/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Authentication
- CWE-307: Improper Restriction of Excessive Authentication Attempts
