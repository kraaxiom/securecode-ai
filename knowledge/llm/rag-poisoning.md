---
id: rag-poisoning
category: llm
cwe: CWE-349
owasp: LLM04:2025-Data-and-Model-Poisoning
severity_default: high
languages: [python, js, java, csharp, go]
---

# RAG Poisoning (empoisonnement de la base documentaire)

## Description
L'empoisonnement RAG (Retrieval-Augmented Generation) consiste à insérer des documents malveillants ou trompeurs dans la base de connaissances qu'un système consulte pour enrichir ses réponses. Lorsque le mécanisme de récupération sélectionne ce contenu empoisonné, il est injecté dans le contexte du modèle et traité comme une source fiable, ce qui peut produire des réponses erronées, biaisées ou contenant des instructions cachées destinées à détourner l'agent (chevauchant alors l'injection indirecte de prompt).

## Où ça apparaît typiquement
- Bases documentaires alimentées par des contributions externes ou semi-publiques (wiki collaboratif, tickets support, avis clients).
- Pipelines d'ingestion automatisés sans validation de la provenance ni de l'intégrité des documents indexés.
- Systèmes permettant le crawl de sites web tiers pour construire l'index vectoriel sans liste de sources de confiance.
- Absence de processus de mise à jour/suppression contrôlée des documents obsolètes ou signalés comme suspects.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Pipeline d'ingestion sans étape de validation de provenance, de format ou de contenu avant indexation vectorielle.
- Absence de contrôle d'accès en écriture sur les sources alimentant la base documentaire (tout utilisateur peut y contribuer).
- Aucun mécanisme de détection de contenu anormal (instructions impératives, formatage suspect) dans les documents indexés.
- Pas de traçabilité de la provenance d'un document jusqu'à la réponse générée (impossible d'auditer une réponse suspecte).

## Remédiation
- Restreindre et authentifier les sources autorisées à alimenter la base documentaire, avec revue avant indexation pour les sources externes.
- Scanner les documents entrants pour détecter des motifs d'instructions cachées avant indexation.
- Conserver la traçabilité de la provenance de chaque chunk récupéré jusqu'à la réponse générée pour permettre l'audit.
- Mettre en place une procédure de retrait rapide d'un document identifié comme malveillant, avec réindexation.
- Voir `rules/remediation/rag-poisoning.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/rag-poisoning/`.

## Références
- OWASP Top 10 for LLM Applications: LLM04:2025 – Data and Model Poisoning
- CWE-349: Acceptance of Extraneous Untrusted Data With Trust
