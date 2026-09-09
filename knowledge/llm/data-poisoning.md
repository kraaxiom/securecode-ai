---
id: data-poisoning
category: llm
cwe: CWE-349
owasp: LLM04:2025-Data-and-Model-Poisoning
severity_default: high
languages: [python, js, java, csharp, go]
---

# Data Poisoning (empoisonnement des données d'entraînement)

## Description
L'empoisonnement de données consiste à corrompre le jeu de données utilisé pour entraîner ou affiner (fine-tuning) un modèle, dans le but d'y introduire des biais, des portes dérobées comportementales (backdoor triggers) ou une dégradation ciblée de la qualité. Contrairement au RAG poisoning qui affecte le contenu récupéré au moment de l'inférence, ce risque agit en amont, sur le modèle lui-même, rendant la compromission persistante et difficile à détecter a posteriori.

## Où ça apparaît typiquement
- Pipelines de fine-tuning utilisant des données collectées automatiquement depuis des sources publiques ou semi-publiques.
- Jeux d'entraînement enrichis par des retours utilisateurs (feedback loop) sans filtrage ni modération.
- Données d'entraînement partagées ou téléchargées depuis des dépôts communautaires non vérifiés.
- Processus d'étiquetage (labeling) externalisé sans contrôle qualité ni détection d'anomalies statistiques.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de contrôle de provenance et d'intégrité sur les jeux de données utilisés pour l'entraînement ou le fine-tuning.
- Pipeline d'ingestion de données d'entraînement sans étape de détection d'anomalies (valeurs aberrantes, motifs répétitifs suspects).
- Réintégration de données issues des interactions utilisateurs dans l'entraînement sans validation ni échantillonnage de contrôle.
- Absence de jeu de données de référence (golden set) permettant de détecter une dérive de comportement après un cycle d'entraînement.

## Remédiation
- Valider la provenance et l'intégrité de chaque source de données d'entraînement, avec signature/hash pour les jeux de données figés.
- Mettre en place une détection d'anomalies statistiques sur les données avant leur intégration au pipeline d'entraînement.
- Isoler et auditer les données issues du feedback utilisateur avant réintégration, avec échantillonnage humain.
- Comparer les performances du modèle sur un jeu de référence fixe après chaque cycle d'entraînement pour détecter une dérive suspecte.
- Voir `rules/remediation/data-poisoning.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/data-poisoning/`.

## Références
- OWASP Top 10 for LLM Applications: LLM04:2025 – Data and Model Poisoning
- CWE-349: Acceptance of Extraneous Untrusted Data With Trust
