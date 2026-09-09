---
id: excessive-data-exposure
category: api
cwe: CWE-213
owasp: API3:2023-Broken Object Property Level Authorization
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# Excessive Data Exposure

## Description
L'excessive data exposure se produit lorsqu'une API retourne dans sa réponse un objet complet (modèle interne, entité base de données) au lieu de sélectionner explicitement les champs destinés au client, en comptant sur le frontend pour filtrer l'affichage. Un attaquant inspectant directement la réponse brute de l'API peut alors récupérer des champs sensibles jamais affichés dans l'interface : mots de passe hachés, jetons internes, données personnelles d'autres utilisateurs, informations de facturation. Ce problème découle généralement d'une sérialisation trop permissive côté serveur.

## Où ça apparaît typiquement
- Sérialisation directe d'une entité ORM/modèle de base de données en JSON sans DTO de sortie dédié.
- Endpoints de recherche ou de listing renvoyant l'objet complet plutôt qu'une projection restreinte.
- APIs internes réutilisées pour un usage public sans retrait des champs techniques ou sensibles.
- Réponses d'erreur ou de debug incluant des informations internes (stack traces, structure de la base).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Sérialisation d'objet complet sans liste blanche de champs de sortie (`return json_encode($user)`, `res.json(userModel)` sans mapping vers un DTO).
- Absence d'annotation d'exclusion (`@JsonIgnore`, `hidden` en Eloquent, `exclude` en serializer DRF) sur des champs sensibles du modèle (mot de passe, token, données internes).
- Endpoints qui documentent un contrat de sortie minimal dans le frontend mais dont la réponse API brute contient davantage de champs.
- Réutilisation du même modèle de sérialisation entre contexte admin (données complètes) et contexte utilisateur final.

## Remédiation
- Définir explicitement un schéma de sortie (DTO, serializer, projection) par endpoint, en liste blanche, plutôt que de sérialiser l'objet interne tel quel.
- Exclure systématiquement les champs sensibles (secrets, hachages, données internes) du modèle de sérialisation, quel que soit le contexte d'appel.
- Différencier les schémas de sortie selon le rôle de l'appelant (utilisateur standard vs administrateur).
- Auditer régulièrement les réponses API brutes (hors frontend) pour vérifier l'absence de sur-exposition.
- Voir `rules/remediation/excessive-data-exposure.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/excessive-data-exposure/` (et les répertoires équivalents pour js, python, java, csharp, go).

## Références
- OWASP API Security Top 10 2023: API3-Broken Object Property Level Authorization (anciennement Excessive Data Exposure)
- CWE-213: Exposure of Sensitive Information Due to Incompatible Policies
