---
id: price-manipulation
category: business-logic
cwe: CWE-840
owasp: A04:2021-Insecure-Design
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Price Manipulation (manipulation de prix)

## Description
La manipulation de prix survient lorsqu'une application accepte du client un montant, un prix unitaire ou une remise censés être déterminés côté serveur, permettant à un attaquant de modifier ces valeurs avant leur transmission (paramètre de requête, champ caché, appel direct de l'API de paiement). Le serveur, faisant confiance à la valeur reçue plutôt que de la recalculer à partir du catalogue et des règles de tarification, finalise alors la transaction au prix falsifié.

## Où ça apparaît typiquement
- Endpoints de paiement/panier acceptant un prix, un total ou un pourcentage de remise directement depuis la requête client.
- Calcul de prix effectué côté client (JavaScript) puis transmis en confiance au serveur pour finaliser la transaction.
- Système de tarification dynamique (devises, remises négociées) où le taux appliqué provient d'un paramètre modifiable plutôt que d'une source serveur.
- API mobile/partenaire différente de l'API web principale, n'appliquant pas les mêmes vérifications de prix.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Paramètre `price`, `amount`, `discount` ou équivalent accepté depuis la requête client sans recalcul serveur à partir du catalogue.
- Absence de revalidation du prix final au moment de la capture du paiement (le prix affiché en amont n'est pas revérifié).
- Logique de calcul de remise dupliquée entre client et serveur, avec le serveur faisant confiance au résultat déjà calculé.
- Écart possible entre le prix affiché à l'utilisateur et le prix effectivement soumis à l'API de paiement, sans contrôle de cohérence.

## Remédiation
- Toujours recalculer le prix final côté serveur à partir du catalogue et des règles de tarification, sans jamais faire confiance à une valeur envoyée par le client.
- Revérifier le montant au moment de la capture effective du paiement, indépendamment du montant affiché plus tôt dans le parcours.
- Centraliser le moteur de tarification côté serveur et le réutiliser pour tous les clients (web, mobile, partenaires).
- Journaliser les écarts significatifs entre prix catalogue et prix soumis pour détecter les tentatives de manipulation.
- Voir `rules/remediation/price-manipulation.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/price-manipulation/`.

## Références
- OWASP Top 10: A04:2021 – Insecure Design
- CWE-840: Business Logic Errors
