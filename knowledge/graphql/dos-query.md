---
id: dos-query
category: graphql
cwe: CWE-400
owasp: A04:2021-Insecure Design
severity_default: high
languages: [js, python, java, csharp, go, rust, php]
---

# Déni de service par requête GraphQL coûteuse

## Description
Une requête GraphQL peut combiner de larges listes, une pagination excessive, des alias multiples et des champs coûteux en une seule opération, sans que sa taille en octets ne reflète le coût réel d'exécution côté serveur. Sans contrôle de complexité (query cost), une requête syntaxiquement valide et de taille modeste peut déclencher un traitement massif en base de données ou en mémoire, entraînant un déni de service.

## Où ça apparaît typiquement
- Champs de liste (`items(first: N)`) sans plafond sur la valeur de `N` demandée par le client.
- APIs GraphQL sans analyse de complexité/coût de requête avant exécution.
- Resolvers effectuant des jointures ou calculs coûteux déclenchés directement par les arguments d'une requête.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Arguments de pagination (`first`, `limit`, `take`) sans valeur maximale imposée côté serveur.
- Absence de bibliothèque d'analyse de coût de requête (ex: `graphql-cost-analysis`, `graphql-query-complexity`) dans la configuration.
- Aucun timeout d'exécution ni limite de ressources configuré sur le serveur GraphQL.

## Remédiation
- Plafonner strictement toute valeur de pagination fournie par le client.
- Mettre en place une analyse de complexité de requête avec un budget de coût maximal par requête.
- Ajouter des timeouts d'exécution et une limite de ressources par requête au niveau du serveur.
- Voir `rules/remediation/dos-query.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/dos-query/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP API Security Top 10: API4:2023 Unrestricted Resource Consumption
- CWE-400: Uncontrolled Resource Consumption
