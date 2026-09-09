---
id: deep-query
category: graphql
cwe: CWE-674
owasp: A04:2021-Insecure Design
severity_default: medium
languages: [js, python, java, csharp, go, rust, php]
---

# Requêtes GraphQL profondément imbriquées

## Description
Le schéma GraphQL permet naturellement d'imbriquer des relations entre types (ex: `user { friends { friends { friends { ... } } } }`). Sans limite de profondeur, un attaquant peut construire une requête récursive dont le coût d'exécution croît de façon exponentielle avec le nombre de niveaux imbriqués, provoquant un déni de service (consommation CPU, mémoire, ou requêtes en cascade vers la base de données) avec une seule requête relativement courte à écrire.

## Où ça apparaît typiquement
- Schémas GraphQL exposant des relations circulaires ou récursives (utilisateurs/amis, catégories/sous-catégories, commentaires imbriqués).
- Serveurs GraphQL sans limite de profondeur de requête configurée (`graphql-depth-limit` absent ou non appliqué).
- APIs publiques ou introspectables où le schéma complet des relations est visible par un attaquant.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de middleware/validation de profondeur maximale dans la configuration du serveur GraphQL.
- Types de schéma avec relations récursives directes sans limite de pagination imposée à chaque niveau.
- Resolvers qui déclenchent une requête base de données par niveau d'imbrication sans dataloader ni cache.

## Remédiation
- Limiter la profondeur maximale des requêtes acceptées (ex: `graphql-depth-limit`, `graphql-armor-max-depth`).
- Combiner avec une analyse de coût de requête (query complexity) pour rejeter les requêtes trop coûteuses avant exécution.
- Imposer une pagination obligatoire sur toute relation potentiellement volumineuse ou récursive.
- Voir `rules/remediation/deep-query.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/deep-query/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP API Security Top 10: API4:2023 Unrestricted Resource Consumption
- CWE-674: Uncontrolled Recursion
