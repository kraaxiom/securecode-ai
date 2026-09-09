---
id: password-spray
category: auth
cwe: CWE-307
owasp: A07:2021-Identification and Authentication Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Password Spraying

## Description
Le password spraying consiste à tester un petit nombre de mots de passe très courants (ex. `Welcome1`, `Password123!`) contre un grand nombre de comptes distincts, plutôt que d'essayer de nombreux mots de passe sur un seul compte. Cette approche permet de contourner les mécanismes de verrouillage de compte classiques, qui ne détectent généralement que des échecs répétés sur un même identifiant.

## Où ça apparaît typiquement
- Portails d'authentification d'entreprise exposant une politique de nommage de compte prévisible (ex. prénom.nom).
- Applications sans corrélation des échecs d'authentification au niveau global (tous comptes confondus).
- Services SSO/annuaire exposés publiquement sans détection de trafic distribué anormal.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Limitation des tentatives d'authentification uniquement par compte, sans agrégation par IP/source ou par fenêtre de temps globale.
- Absence de politique de mot de passe empêchant les valeurs triviales couramment ciblées.
- Absence de détection d'anomalie sur un volume d'échecs répartis sur de nombreux comptes distincts en peu de temps.

## Remédiation
- Compléter la limitation par compte par une détection globale (volume d'échecs par IP/source, par fenêtre de temps, tous comptes confondus).
- Imposer une politique de mot de passe robuste interdisant les valeurs les plus courantes.
- Déployer une authentification multi-facteurs, qui neutralise l'essentiel de l'impact d'un password spraying réussi.
- Voir `rules/remediation/password-spray.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/password-spray/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Credential Stuffing Prevention
- CWE-307: Improper Restriction of Excessive Authentication Attempts
