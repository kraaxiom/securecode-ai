---
id: race-condition
category: business-logic
cwe: CWE-362
owasp: A04:2021-Insecure-Design
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Race Condition (condition de concurrence métier)

## Description
Une condition de concurrence survient lorsque plusieurs requêtes concurrentes sur une même ressource métier (solde, stock, code promo, tentative de connexion) ne sont pas correctement synchronisées, permettant à un attaquant d'envoyer des requêtes quasi simultanées pour contourner une vérification censée être atomique. L'exemple classique est l'envoi parallèle de plusieurs requêtes de retrait ou d'utilisation de coupon avant que le solde ou le compteur ne soit mis à jour, produisant un résultat incohérent (retrait multiple, dépassement de quota).

## Où ça apparaît typiquement
- Opérations de type "vérifier puis agir" (check-then-act) sur un solde, un stock ou un quota, sans verrouillage atomique.
- Endpoints de paiement, de retrait ou d'application de coupon acceptant des requêtes concurrentes sans idempotence.
- Systèmes de réservation ou d'allocation de ressources limitées (places, tickets) sans contrôle transactionnel.
- Compteurs applicatifs (tentatives de connexion, essais gratuits) mis à jour en mémoire ou via lecture-modification-écriture non atomique.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Séquence lire-puis-écrire sur une donnée partagée sans transaction, verrou ou opération atomique correspondante.
- Absence de contrainte d'unicité ou de verrou pessimiste/optimiste sur les ressources critiques (solde, stock) en base de données.
- Endpoints financiers ou de consommation de quota sans clé d'idempotence permettant de rejeter les requêtes dupliquées.
- Logique métier critique implémentée au niveau applicatif (vérification en mémoire) plutôt que déléguée à une contrainte de la base de données.

## Remédiation
- Utiliser des transactions avec verrouillage approprié (`SELECT ... FOR UPDATE`, verrous optimistes avec version) pour toute opération lire-modifier-écrire critique.
- Exiger une clé d'idempotence sur les endpoints sensibles (paiement, retrait, application de coupon) pour rejeter les doublons.
- Déplacer les contraintes critiques (solde ne pouvant être négatif, quota non dépassable) au niveau de la base de données via des contraintes CHECK ou des opérations atomiques.
- Tester explicitement le comportement sous requêtes concurrentes lors des revues de sécurité des flux financiers.
- Voir `rules/remediation/race-condition.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/race-condition/`.

## Références
- OWASP Top 10: A04:2021 – Insecure Design
- CWE-362: Concurrent Execution using Shared Resource with Improper Synchronization ('Race Condition')
