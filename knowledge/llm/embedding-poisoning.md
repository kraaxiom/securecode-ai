---
id: embedding-poisoning
category: llm
cwe: CWE-349
owasp: LLM04:2025-Data-and-Model-Poisoning
severity_default: medium
languages: [python, js, java, csharp, go]
---

# Embedding Poisoning

## Description
L'empoisonnement d'embeddings vise directement l'espace vectoriel utilisé par un système de recherche sémantique ou de RAG. Un attaquant fabrique un contenu dont l'embedding est délibérément proche de requêtes légitimes fréquentes, afin qu'il soit systématiquement remonté par le moteur de similarité même s'il n'est pas pertinent sémantiquement, ou qu'il pollue le clustering utilisé pour des décisions automatisées (classification, déduplication, détection d'anomalies).

## Où ça apparaît typiquement
- Bases vectorielles alimentées par du contenu partiellement contrôlé par des utilisateurs externes.
- Systèmes de recherche sémantique ou de recommandation utilisant des embeddings pour classer/prioriser du contenu.
- Pipelines combinant plusieurs sources d'embeddings sans validation d'origine ni de cohérence.
- Fonctionnalités où un utilisateur peut soumettre du contenu directement indexé (avis, description de produit) influençant la similarité perçue par d'autres utilisateurs.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de contrôle sur le contenu soumis par les utilisateurs avant génération et indexation de son embedding.
- Aucune détection d'anomalies de densité ou de clustering inhabituel dans l'espace vectoriel (points anormalement proches de requêtes fréquentes).
- Pas de limite de fréquence ou de volume sur l'ajout de nouveaux vecteurs par une même source/utilisateur.
- Absence de réévaluation périodique de la pertinence des résultats de recherche sémantique en production.

## Remédiation
- Valider et modérer le contenu avant génération d'embedding et indexation, en particulier pour du contenu soumis par des tiers.
- Surveiller la distribution des vecteurs indexés pour détecter des motifs de clustering anormal ou des injections en volume.
- Limiter la fréquence et le volume d'indexation par source pour freiner les campagnes d'empoisonnement automatisées.
- Réévaluer périodiquement la pertinence des résultats retournés pour des requêtes de référence connues.
- Voir `rules/remediation/embedding-poisoning.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/embedding-poisoning/`.

## Références
- OWASP Top 10 for LLM Applications: LLM04:2025 – Data and Model Poisoning
- CWE-349: Acceptance of Extraneous Untrusted Data With Trust
