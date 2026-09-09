---
id: model-extraction
category: llm
cwe: CWE-200
owasp: LLM10:2025-Unbounded-Consumption
severity_default: medium
languages: [python, js, java, csharp, go]
---

# Model Extraction (vol de modèle)

## Description
L'extraction de modèle consiste à interroger massivement et systématiquement une API de modèle exposée publiquement afin d'en reconstituer un équivalent fonctionnel (modèle de substitution / distillation) ou d'en dériver la logique interne et les paramètres, sans y être autorisé. Cette technique permet à un concurrent de s'approprier la propriété intellectuelle investie dans un modèle propriétaire, ou de faciliter la préparation d'autres attaques (recherche de faiblesses par analyse hors ligne du modèle extrait).

## Où ça apparaît typiquement
- APIs de modèle exposées publiquement sans limite de débit (rate limiting) stricte ni quotas par client.
- Réponses de l'API renvoyant des informations riches (logits, probabilités, scores de confiance) au-delà de la sortie textuelle nécessaire.
- Absence de détection de patterns d'usage automatisés (volumes de requêtes systématiques, requêtes générées par script).
- Modèles propriétaires accessibles sans contrat de service limitant explicitement l'usage autorisé.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de limitation de débit ou de quota par clé API/utilisateur sur les endpoints d'inférence.
- Exposition de métadonnées internes du modèle (logits complets, scores bruts) non nécessaires au cas d'usage métier.
- Absence de surveillance des volumes de requêtes et de détection d'usage automatisé anormal (patterns non humains).
- Pas de watermarking ni de mécanisme de traçabilité des sorties du modèle permettant de détecter une réutilisation non autorisée.

## Remédiation
- Mettre en place des quotas et une limitation de débit stricts par client/clé API sur les endpoints d'inférence.
- Ne renvoyer que les informations strictement nécessaires au cas d'usage (éviter d'exposer logits/probabilités bruts sans besoin).
- Surveiller les patterns de requêtes pour détecter une extraction systématique (diversité, volume, régularité anormale).
- Encadrer contractuellement l'usage autorisé de l'API et envisager du watermarking des sorties pour les modèles à forte valeur.
- Voir `rules/remediation/model-extraction.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/model-extraction/`.

## Références
- OWASP Top 10 for LLM Applications: LLM10:2025 – Unbounded Consumption
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
