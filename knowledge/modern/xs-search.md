---
id: xs-search
category: modern
cwe: CWE-203
owasp: A01:2021-Broken Access Control
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# XS-Search (Cross-Site Search)

## Description
Le XS-Search est une variante des XS-Leaks ciblant spécifiquement les fonctionnalités de recherche ou de filtrage internes à une application authentifiée. Un site attaquant déclenche des requêtes de recherche cross-origin vers l'application ciblée (via des balises `<img>`, `<script>`, ou des requêtes `fetch` avec `no-cors`) et mesure des différences observables (temps de réponse, taille de la ressource chargée) pour déduire, terme par terme, si le résultat de recherche de la victime contient des données spécifiques, révélant progressivement des informations privées sans jamais lire directement la réponse.

## Où ça apparaît typiquement
- Fonctionnalités de recherche interne accessibles en GET dont le temps de réponse ou la taille varie significativement selon le nombre de résultats.
- Endpoints de recherche/filtrage sans protection `SameSite` stricte sur le cookie de session utilisé pour l'authentification.
- Absence d'en-têtes `Cross-Origin-Resource-Policy` empêchant le chargement cross-origin de la ressource par des balises passives.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Endpoint de recherche accessible en GET sans jeton anti-CSRF ni vérification d'origine, chargeable via une balise cross-origin.
- Temps de réponse ou taille de réponse fortement corrélés au nombre de résultats trouvés, sans padding ni normalisation.
- Absence de `Cross-Origin-Resource-Policy: same-origin` sur les endpoints de recherche retournant des données sensibles.

## Remédiation
- Exiger une méthode POST avec vérification d'origine/jeton pour toute recherche portant sur des données sensibles.
- Ajouter `Cross-Origin-Resource-Policy: same-origin` sur les réponses de recherche pour bloquer leur chargement cross-origin passif.
- Normaliser le temps de réponse et éviter les variations de taille directement corrélées au volume de résultats sensibles.
- Voir `rules/remediation/xs-search.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/xs-search/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP: XS-Leaks Wiki (XS-Search Attacks)
- CWE-203: Observable Discrepancy
