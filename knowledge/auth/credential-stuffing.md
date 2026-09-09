---
id: credential-stuffing
category: auth
cwe: CWE-307
owasp: A07:2021-Identification and Authentication Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Credential Stuffing

## Description
Le credential stuffing consiste à tester en masse, sur une application cible, des paires identifiant/mot de passe issues de fuites de données provenant d'autres services. Contrairement au brute force qui essaie de nombreux mots de passe sur un compte donné, le credential stuffing exploite la réutilisation de mots de passe par les utilisateurs sur plusieurs sites, et réussit dès qu'une fraction des couples testés est valide.

## Où ça apparaît typiquement
- Formulaires et API de connexion exposés publiquement, sans détection de trafic automatisé.
- Applications sans authentification multi-facteurs proposée ou imposée.
- Absence de corrélation entre volume de tentatives, diversité des identifiants testés et une même origine (IP, empreinte client).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de limitation de débit globale sur l'endpoint de connexion, en plus de la limitation par compte.
- Absence de détection d'anomalies (volume élevé de tentatives avec des identifiants différents depuis une même source).
- Absence de proposition/imposition de MFA, qui neutraliserait l'usage de couples identifiant/mot de passe volés.

## Remédiation
- Proposer et encourager fortement l'authentification multi-facteurs (MFA).
- Mettre en place une détection comportementale (vélocité des tentatives, diversité des comptes testés depuis une même source) en complément du rate limiting classique.
- Vérifier les mots de passe utilisateurs contre des listes de fuites connues (au moment de l'inscription/changement) pour forcer leur renouvellement.
- Voir `rules/remediation/credential-stuffing.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/credential-stuffing/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Credential Stuffing Prevention
- CWE-307: Improper Restriction of Excessive Authentication Attempts
