---
id: graphql-api
category: api
cwe: CWE-200
owasp: API3:2023-Broken Object Property Level Authorization
severity_default: medium
languages: [js, python, java, csharp, go]
---

# Sécurité des API GraphQL

## Description
Les API GraphQL introduisent des risques spécifiques liés à leur flexibilité : un schéma d'introspection laissé actif en production révèle l'intégralité de la structure interne (types, champs, mutations) à toute personne effectuant une requête d'introspection. Par ailleurs, l'absence de limitation de profondeur ou de complexité des requêtes permet de construire des requêtes fortement imbriquées qui consomment des ressources serveur disproportionnées. Enfin, comme pour REST, l'autorisation doit être vérifiée au niveau de chaque champ résolu et non uniquement au niveau de la requête globale.

## Où ça apparaît typiquement
- Endpoint GraphQL unique (`/graphql`) exposant l'introspection (`__schema`, `__type`) en environnement de production.
- Résolveurs (resolvers) de champs sensibles ne réappliquant pas de contrôle d'autorisation propre, en s'appuyant uniquement sur un contrôle au niveau de la requête parente.
- Requêtes imbriquées récursives (relations circulaires entre types) sans limite de profondeur ni de complexité configurée.
- Messages d'erreur GraphQL détaillés (stack trace, requête SQL) renvoyés tels quels au client.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Configuration du serveur GraphQL sans désactivation explicite de l'introspection en production (`introspection: true` par défaut non modifié).
- Absence de bibliothèque ou de règle de limitation de profondeur/complexité de requête (depth limit, cost analysis) dans la configuration du serveur.
- Résolveurs accédant directement à la base de données sans passer par la même couche d'autorisation que les endpoints REST équivalents.
- Formatteur d'erreur par défaut du framework GraphQL laissé actif, exposant la stack trace complète dans la réponse.

## Remédiation
- Désactiver l'introspection GraphQL en production, ou la restreindre aux environnements internes/authentifiés.
- Mettre en place une analyse de complexité et une limite de profondeur de requête (depth limiting, query cost analysis) pour prévenir le déni de service applicatif.
- Appliquer les contrôles d'autorisation au niveau de chaque résolveur de champ sensible, pas uniquement au niveau de la requête racine.
- Personnaliser le formatteur d'erreurs pour ne jamais renvoyer de détails d'implémentation interne au client.
- Voir `rules/remediation/graphql-api.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js/graphql-api/` (et les répertoires équivalents pour python, java, csharp, go).

## Références
- OWASP API Security Top 10 (recommandations générales applicables à GraphQL)
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
- CWE-770: Allocation of Resources Without Limits or Throttling (pour le risque de déni de service via requêtes profondément imbriquées)
