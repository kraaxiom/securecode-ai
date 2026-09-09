# Data Poisoning (empoisonnement des données d'entraînement)

## Description de la vulnérabilité

L'empoisonnement de données consiste à corrompre le jeu de données utilisé pour entraîner ou affiner (fine-tuning) un modèle, afin d'y introduire des biais, des portes dérobées comportementales ou une dégradation ciblée de la qualité. Contrairement au RAG poisoning qui affecte le contenu récupéré au moment de l'inférence, ce risque agit en amont, sur le modèle lui-même, rendant la compromission persistante et difficile à détecter a posteriori.

Dans `vulnerable.go`, `BuildTrainingSet` télécharge et agrège les données de toutes les sources fournies sans vérifier leur provenance ni leur intégrité, et `IngestUserFeedback` réintègre directement les retours utilisateurs dans le pipeline d'entraînement sans filtrage.

## CWE réel utilisé

**CWE-349 : Acceptance of Extraneous Untrusted Data With Trust** (tel que référencé dans `knowledge/llm/data-poisoning.md`).

## Pourquoi c'est dangereux

- Aucun contrôle de provenance ni d'intégrité sur les sources de données d'entraînement : n'importe quelle source peut injecter des exemples malveillants.
- Pas de détection d'anomalies statistiques : des motifs répétitifs suspects ou des valeurs aberrantes passent inaperçus.
- Les retours utilisateurs sont réintégrés sans échantillonnage humain, ouvrant la voie à une campagne de contamination progressive (feedback loop empoisonné).
- Aucun jeu de référence (golden set) pour détecter une dérive de comportement après un cycle d'entraînement : la compromission peut rester invisible longtemps.

## Comment le correctif fonctionne

Le fichier `fixed.go` introduit :

1. Une liste explicite `trustedSources` : seules les sources identifiées et autorisées sont ingérées.
2. Une vérification d'intégrité (`verifyIntegrity`) comparant un hash SHA-256 calculé au hash signé attendu (`SignedHash`) avant toute utilisation des données.
3. Un filtrage des valeurs aberrantes (`filterStatisticalOutliers`) appliqué à chaque lot avant intégration au dataset.
4. Un circuit d'échantillonnage et de validation humaine (`sampleForHumanReview`, `humanApproves`) pour les retours utilisateurs avant toute réintégration.
5. Une validation post-entraînement (`ValidateModelAfterTraining`) comparant les performances du modèle à un jeu de référence fixe, avec rejet (`errModelDrift`) si le score chute sous le seuil acceptable.

## Références

- OWASP Top 10 for LLM Applications: LLM04:2025 – Data and Model Poisoning
- CWE-349: Acceptance of Extraneous Untrusted Data With Trust
- `rules/remediation/data-poisoning.md`
- `knowledge/llm/data-poisoning.md`
