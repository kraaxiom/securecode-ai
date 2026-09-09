---
id: batch-attack
category: graphql
cwe: CWE-770
owasp: A04:2021-Insecure Design
severity_default: medium
languages: [js, python, java, csharp, go, rust, php]
---

# Attaque par batching GraphQL

## Description
De nombreux clients et serveurs GraphQL supportent l'envoi de plusieurs opérations dans un seul tableau JSON (batching). Sans limite sur la taille de ce tableau, un attaquant peut envoyer des centaines d'opérations indépendantes (par exemple des tentatives de connexion avec des couples identifiant/mot de passe différents) dans une seule requête HTTP, contournant ainsi les protections de rate-limiting classiques et amplifiant fortement la charge sur le serveur et la base de données.

## Où ça apparaît typiquement
- Endpoints GraphQL acceptant un corps de requête sous forme de tableau (`[{query...}, {query...}]`) sans limite de taille.
- Serveurs Apollo/Yoga/graphql-http avec le batching activé par défaut sans configuration de plafond.
- Rate limiting comptant les requêtes HTTP au lieu des opérations GraphQL individuelles qu'elles contiennent.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Configuration serveur GraphQL avec batching activé sans option `maxBatchSize`/équivalent définie.
- Absence de validation de la taille du tableau de requêtes avant traitement.
- Logs montrant un rate limiter qui compte 1 requête HTTP alors que N opérations métier ont été exécutées.

## Remédiation
- Désactiver le batching si non nécessaire, ou plafonner strictement le nombre d'opérations par requête batchée.
- Faire compter le rate limiter au niveau du nombre d'opérations GraphQL exécutées, pas des requêtes HTTP.
- Appliquer une limitation dédiée sur les mutations sensibles indépendamment du batching.
- Voir `rules/remediation/batch-attack.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/batch-attack/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP API Security Top 10: API4:2023 Unrestricted Resource Consumption
- CWE-770: Allocation of Resources Without Limits or Throttling
