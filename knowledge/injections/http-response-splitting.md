---
id: http-response-splitting
category: injections
cwe: CWE-113
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# HTTP Response Splitting

## Description
Le fractionnement de réponse HTTP survient lorsqu'une entrée utilisateur non neutralisée contenant des caractères CR/LF (`\r\n`) est intégrée dans une réponse HTTP construite par l'application (en-tête ou redirection). En injectant ces caractères, un attaquant peut terminer la réponse HTTP en cours et injecter le début d'une seconde réponse, permettant notamment l'empoisonnement de cache ou le détournement de contenu affiché aux autres utilisateurs.

## Où ça apparaît typiquement
- Construction manuelle d'en-têtes de réponse ou de redirections à partir de paramètres utilisateur (`Location`, `Set-Cookie`).
- Frameworks anciens ou API bas niveau ne filtrant pas automatiquement les caractères de contrôle dans les valeurs d'en-tête.
- Génération de pages ou de flux de réponse combinant sortie utilisateur et écriture brute du flux HTTP.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Écriture d'en-têtes de réponse via des API bas niveau avec une valeur non filtrée provenant d'une entrée utilisateur.
- Absence de rejet des caractères `\r` et `\n` avant assemblage de la réponse.
- Utilisation de versions de framework connues pour ne pas encoder automatiquement les en-têtes.

## Remédiation
- Utiliser les API modernes du framework qui rejettent ou encodent automatiquement les caractères CR/LF.
- Valider strictement les valeurs utilisées dans les en-têtes et redirections (liste blanche).
- Maintenir les frameworks à jour, les versions récentes bloquant nativement l'injection CR/LF.
- Voir `rules/remediation/http-response-splitting.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/http-response-splitting/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: HTTP Response Splitting Prevention
- CWE-113: Improper Neutralization of CRLF Sequences in HTTP Headers ('HTTP Response Splitting')
