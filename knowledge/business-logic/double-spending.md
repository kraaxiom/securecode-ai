---
id: double-spending
category: business-logic
cwe: CWE-362
owasp: A04:2021-Insecure-Design
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Double Spending (double dépense)

## Description
La double dépense est une variante spécifique de condition de concurrence appliquée aux flux financiers : elle permet à un attaquant d'utiliser un même solde, crédit, jeton de paiement ou avoir plusieurs fois en soumettant des requêtes concurrentes ou en rejouant une transaction avant que le système n'ait pu marquer la ressource comme consommée. Le résultat est un débit ou une consommation qui n'est jamais reflété correctement dans le solde réel, créant une perte financière directe.

## Où ça apparaît typiquement
- Systèmes de paiement ou de portefeuille interne où le solde est vérifié puis débité en deux opérations distinctes non atomiques.
- Traitement de webhooks de paiement (confirmation asynchrone) pouvant être reçus ou traités plusieurs fois pour un même événement.
- Systèmes de points de fidélité ou de crédits internes convertibles, sans verrouillage transactionnel lors de l'échange.
- Rejeu possible d'une requête de transaction (absence de déduplication) suite à une erreur réseau ou un retry client.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Vérification de solde suffisant et débit effectué dans deux requêtes/étapes séparées non enveloppées dans une transaction atomique.
- Traitement de webhook/callback de paiement sans vérification d'idempotence sur l'identifiant unique de l'événement.
- Absence de verrou ou de contrainte d'unicité empêchant le traitement multiple d'une même transaction rejouée.
- Solde ou crédit stocké et mis à jour sans contrainte empêchant une valeur négative ou incohérente en base de données.

## Remédiation
- Encapsuler toute opération de vérification-puis-débit dans une transaction atomique avec verrouillage approprié.
- Exiger un identifiant unique d'idempotence sur chaque transaction et rejeter tout traitement dupliqué d'un même événement de paiement.
- Ajouter des contraintes de base de données empêchant un solde de devenir négatif ou incohérent.
- Réconcilier régulièrement les soldes calculés avec un registre de transactions immuable pour détecter les anomalies.
- Voir `rules/remediation/double-spending.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/double-spending/`.

## Références
- OWASP Top 10: A04:2021 – Insecure Design
- CWE-362: Concurrent Execution using Shared Resource with Improper Synchronization ('Race Condition')
