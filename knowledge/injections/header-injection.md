---
id: header-injection
category: injections
cwe: CWE-113
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# Header Injection

## Description
L'injection d'en-têtes HTTP survient lorsqu'une valeur contrôlée par l'utilisateur est insérée dans un en-tête de réponse HTTP sans neutralisation des caractères de fin de ligne (CR, LF). Un attaquant peut alors injecter des en-têtes supplémentaires ou terminer prématurément la section d'en-têtes pour influencer la réponse. C'est un cas particulier de CRLF injection appliqué spécifiquement à la construction d'en-têtes applicatifs.

## Où ça apparaît typiquement
- Définition manuelle d'en-têtes de réponse (`Set-Header`, `header()`, `response.setHeader()`) avec une valeur issue d'un paramètre utilisateur.
- Redirections construites dynamiquement (`Location:`) à partir d'une entrée non validée.
- En-têtes personnalisés reflétant des données utilisateur (ex: nom de fichier dans `Content-Disposition`).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à une fonction de définition d'en-tête avec une chaîne concaténée contenant une variable non filtrée.
- Absence de suppression/rejet des caractères `\r` et `\n` avant écriture dans un en-tête.
- Utilisation d'API bas niveau d'écriture de réponse plutôt que des API du framework qui encodent automatiquement.

## Remédiation
- Utiliser les API du framework qui gèrent nativement l'échappement des en-têtes plutôt que d'écrire des chaînes brutes.
- Rejeter ou supprimer tout caractère de contrôle (`\r`, `\n`) dans les valeurs destinées à un en-tête.
- Valider les valeurs selon une liste blanche stricte (ex: format attendu d'un nom de fichier).
- Voir `rules/remediation/header-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/header-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: HTTP Response Splitting Prevention / Input Validation
- CWE-113: Improper Neutralization of CRLF Sequences in HTTP Headers
