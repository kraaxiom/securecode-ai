---
id: business-rule-bypass
category: business-logic
cwe: CWE-840
owasp: A04:2021-Insecure-Design
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Business Rule Bypass (contournement de règle métier)

## Description
Ce pattern couvre les cas où une règle métier essentielle (limite d'âge, plafond de montant, séquence obligatoire d'étapes, restriction géographique) n'est appliquée que côté client ou dans une seule couche de l'application, permettant à un attaquant de la contourner en appelant directement l'API sous-jacente ou en manipulant les paramètres de la requête. La règle métier existe dans l'interface mais n'est jamais revalidée côté serveur au moment de l'exécution effective.

## Où ça apparaît typiquement
- Validation de règles métier (âge minimum, montant maximum, quantité limite) uniquement implémentée en JavaScript côté client.
- Champs cachés ou paramètres de formulaire (prix, statut, rôle) transmis au serveur sans revalidation, en confiance dans le client.
- APIs multiples pour un même flux (web, mobile, partenaire) où seule une des interfaces applique correctement la règle métier.
- Étapes d'un processus (validation KYC avant transaction) contrôlées uniquement par l'ordre de navigation dans l'UI, pas par un état serveur.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Règle métier critique présente uniquement dans le code frontend (validation de formulaire) sans équivalent serveur.
- Endpoint API acceptant un paramètre qui devrait être dérivé/calculé côté serveur (prix, rôle, statut) directement depuis la requête client.
- Absence de machine à états serveur pour les processus multi-étapes, l'ordre étant uniquement garanti par la navigation UI.
- Incohérence de règles métier appliquées entre différents clients consommant la même API (web vs mobile vs API publique).

## Remédiation
- Revalider systématiquement toute règle métier côté serveur, indépendamment de ce que l'interface utilisateur a déjà vérifié.
- Ne jamais faire confiance à des valeurs sensibles (prix, rôle, statut) transmises par le client ; les recalculer ou les vérifier côté serveur.
- Modéliser les processus multi-étapes avec une machine à états tenue côté serveur, rejetant toute transition non autorisée.
- Centraliser la logique métier critique dans une couche partagée par tous les clients consommateurs de l'API.
- Voir `rules/remediation/business-rule-bypass.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/business-rule-bypass/`.

## Références
- OWASP Top 10: A04:2021 – Insecure Design
- CWE-840: Business Logic Errors
