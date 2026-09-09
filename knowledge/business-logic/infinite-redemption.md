---
id: infinite-redemption
category: business-logic
cwe: CWE-841
owasp: A04:2021-Insecure-Design
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Infinite Redemption (rachat/échange illimité)

## Description
Ce pattern décrit les failles permettant de réclamer ou d'échanger un avantage à usage unique (récompense de parrainage, cadeau de bienvenue, cashback, ticket de tombola) un nombre de fois non prévu par la règle métier, généralement parce que l'état "déjà réclamé" n'est pas correctement persistant, vérifié de façon atomique, ou lié sans ambiguïté à l'identité de l'utilisateur.

## Où ça apparaît typiquement
- Programmes de parrainage où le statut "récompense réclamée" est stocké côté client ou vérifié de façon non atomique.
- Flux de suppression/recréation de compte permettant de retoucher un avantage réservé aux nouveaux comptes.
- Actions déclenchant une récompense (première commande, inscription) sans verrou empêchant leur déclenchement répété en parallèle.
- Systèmes de points/cashback recalculant le solde à partir d'un historique manipulable plutôt que d'un registre immuable.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- État "avantage déjà réclamé" non stocké de façon atomique et vérifiable côté serveur avant chaque octroi.
- Vérification d'éligibilité à une récompense basée sur des critères facilement recréables (nouvelle adresse e-mail, nouveau compte) sans détection de duplication (device fingerprint, autres signaux).
- Endpoint de réclamation de récompense sans verrou transactionnel, exploitable par des requêtes concurrentes.
- Calcul du solde de points/cashback dérivé d'un historique modifiable plutôt que d'un registre d'événements immuable et audité.

## Remédiation
- Marquer l'octroi d'un avantage de façon atomique et transactionnelle côté serveur avant toute confirmation à l'utilisateur.
- Vérifier l'éligibilité via des signaux multiples et résistants à la duplication triviale (pas seulement l'e-mail ou l'IP).
- Protéger les endpoints de réclamation de récompense contre les requêtes concurrentes via verrouillage ou contrainte d'unicité.
- Maintenir un registre d'événements immuable pour tout octroi d'avantage, permettant l'audit et la détection de doublons.
- Voir `rules/remediation/infinite-redemption.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/infinite-redemption/`.

## Références
- OWASP Top 10: A04:2021 – Insecure Design
- CWE-841: Improper Enforcement of Behavioral Workflow
