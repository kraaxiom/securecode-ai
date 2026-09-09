---
id: alias-abuse
category: graphql
cwe: CWE-770
owasp: A04:2021-Insecure Design
severity_default: medium
languages: [js, python, java, csharp, go, rust, php]
---

# Abus d'alias GraphQL

## Description
GraphQL permet de renommer (aliaser) un même champ plusieurs fois dans une seule requête, ce qui est utile légitimement mais peut être détourné pour multiplier des centaines d'appels au même resolver en une seule requête HTTP. Cet abus contourne les limitations de rate-limiting basées sur le nombre de requêtes HTTP, car une seule requête réseau déclenche en réalité des milliers d'exécutions de resolver côté serveur, ce qui peut être utilisé pour du bruteforce (ex: tester des centaines de mots de passe en un seul appel) ou pour épuiser les ressources.

## Où ça apparaît typiquement
- API GraphQL exposant une mutation de login/vérification sans limite sur le nombre d'alias par requête.
- Endpoints GraphQL publics sans limite de complexité ni de nombre de champs par requête.
- Rate limiting appliqué uniquement au niveau HTTP (par requête) et non au niveau du nombre d'opérations GraphQL exécutées.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de limite de nombre d'alias/champs dans la configuration du serveur GraphQL (`graphql-armor`, `graphql-depth-limit`, etc. non configurés).
- Rate limiter positionné uniquement sur la couche HTTP/middleware, sans awareness du contenu de la requête GraphQL.
- Resolvers sensibles (authentification, réinitialisation de mot de passe) accessibles sans limitation dédiée du nombre d'appels par requête.

## Remédiation
- Configurer une limite stricte sur le nombre d'alias/champs autorisés par requête (ex: `graphql-armor-max-aliases`).
- Appliquer un rate limiting basé sur le coût réel de la requête (query cost analysis) plutôt que sur le nombre de requêtes HTTP.
- Ajouter une limitation dédiée sur les resolvers sensibles (login, OTP) indépendamment du transport.
- Voir `rules/remediation/alias-abuse.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/alias-abuse/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP API Security Top 10: API4:2023 Unrestricted Resource Consumption
- CWE-770: Allocation of Resources Without Limits or Throttling
