---
id: workflow-bypass
category: business-logic
cwe: CWE-841
owasp: A04:2021-Insecure-Design
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Workflow Bypass (contournement de séquence métier)

## Description
Ce pattern regroupe les failles permettant à un utilisateur de sauter, réordonner ou rejouer les étapes obligatoires d'un processus métier multi-étapes (validation d'identité avant transaction, paiement avant expédition, approbation avant publication) en appelant directement une étape avancée de l'API sans avoir complété les étapes précédentes. La cause est l'absence d'un état serveur faisant autorité sur la progression réelle de l'utilisateur dans le processus.

## Où ça apparaît typiquement
- Parcours de commande en plusieurs étapes (panier → livraison → paiement → confirmation) où chaque étape est une route indépendante sans vérification de la précédente.
- Processus d'onboarding ou de KYC où l'étape finale d'activation est accessible directement sans preuve que les vérifications intermédiaires ont abouti.
- Workflows d'approbation (validation manager, revue de contenu) où le statut final peut être atteint sans passage par les étapes de validation.
- APIs REST exposant chaque étape comme un endpoint indépendant sans machine à états serveur reliant les appels entre eux.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Étapes d'un processus multi-étapes exposées comme endpoints indépendants sans vérification de l'état d'avancement réel côté serveur.
- Statut de progression du workflow stocké ou vérifié côté client (paramètre d'URL, état local) plutôt que dans une source serveur faisant autorité.
- Absence de machine à états définissant les transitions valides et rejetant les transitions non autorisées.
- Endpoint d'étape finale (confirmation, activation, publication) sans vérification que les prérequis métier ont été satisfaits.

## Remédiation
- Modéliser le processus comme une machine à états côté serveur, chaque transition étant validée contre l'état courant réel de la ressource.
- Ne jamais déterminer la progression d'un workflow à partir de données fournies par le client (paramètre, état local).
- Vérifier explicitement, à chaque étape sensible, que tous les prérequis métier des étapes précédentes sont satisfaits en base de données.
- Auditer et journaliser les transitions d'état pour détecter les tentatives de saut d'étape ou de rejeu.
- Voir `rules/remediation/workflow-bypass.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/workflow-bypass/`.

## Références
- OWASP Top 10: A04:2021 – Insecure Design
- CWE-841: Improper Enforcement of Behavioral Workflow
