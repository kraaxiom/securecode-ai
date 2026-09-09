---
id: introspection-enabled
category: graphql
cwe: CWE-200
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: [js, python, java, csharp, go, rust, php]
---

# Introspection GraphQL activée en production

## Description
L'introspection est une fonctionnalité native de GraphQL qui permet à un client d'interroger le schéma complet de l'API (types, champs, mutations, arguments, descriptions). Utile en développement, elle devient une fuite d'information en production : un attaquant peut cartographier entièrement la surface d'attaque de l'API (y compris des champs/mutations internes non documentés) sans aucune connaissance préalable, ce qui facilite la découverte d'endpoints sensibles ou de faiblesses d'autorisation.

## Où ça apparaît typiquement
- Serveurs GraphQL déployés avec la configuration par défaut du framework (introspection activée par défaut dans la plupart des librairies).
- Environnements de staging/production n'ayant pas désactivé explicitement l'introspection avant mise en ligne.
- Outils de développement (GraphiQL, GraphQL Playground) laissés accessibles publiquement.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Configuration du serveur GraphQL sans option `introspection: false` (ou équivalent) en environnement de production.
- Interface GraphiQL/Playground accessible sans authentification sur une URL publique.
- Absence de différenciation de configuration entre environnements (dev/staging/prod) dans le code de démarrage du serveur.

## Remédiation
- Désactiver explicitement l'introspection en production dans la configuration du serveur GraphQL.
- Désactiver ou protéger par authentification les interfaces exploratoires (GraphiQL, Playground, Voyager).
- Séparer clairement la configuration par environnement pour éviter qu'un réglage de développement fuite en production.
- Voir `rules/remediation/introspection-enabled.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/introspection-enabled/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP API Security Top 10: API8:2023 Security Misconfiguration
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
