---
id: http-parameter-pollution
category: injections
cwe: CWE-235
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# HTTP Parameter Pollution

## Description
La pollution de paramètres HTTP (HPP) consiste à envoyer plusieurs paramètres portant le même nom dans une requête HTTP. Selon la technologie serveur ou le composant qui traite la requête, seul le premier, le dernier, ou une concaténation des paramètres peut être retenu, ce qui crée des divergences d'interprétation exploitables pour contourner des contrôles de validation, des filtres WAF ou de la logique métier. Le risque vient du décalage d'interprétation entre plusieurs couches (proxy, framework, backend).

## Où ça apparaît typiquement
- Applications derrière un reverse proxy ou WAF dont le comportement de parsing diffère du serveur applicatif.
- Formulaires ou API construisant des requêtes en dupliquant des noms de paramètres (query string, form-data).
- Logique métier ou de validation qui ne considère qu'une seule valeur par nom de paramètre alors que l'appel HTTP sous-jacent en transporte plusieurs.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Récupération de paramètres via des API qui retournent silencieusement une seule valeur (`request.getParameter`) sans vérifier l'existence de doublons.
- Absence de normalisation/validation du nombre d'occurrences d'un paramètre attendu comme unique.
- Divergence documentée ou observable entre les couches d'infrastructure (proxy vs backend) sur le traitement des paramètres dupliqués.

## Remédiation
- Définir explicitement le comportement attendu (rejet si paramètre dupliqué alors qu'il doit être unique).
- Harmoniser le comportement de parsing entre toutes les couches de l'infrastructure (proxy, WAF, backend).
- Valider strictement le nombre d'occurrences de chaque paramètre côté serveur applicatif.
- Voir `rules/remediation/http-parameter-pollution.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/nodejs-express/http-parameter-pollution/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Testing Guide: Testing for HTTP Parameter Pollution
- CWE-235: Improper Handling of Extra Parameters
