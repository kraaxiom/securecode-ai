---
id: agent-hijacking
category: llm
cwe: CWE-1427
owasp: LLM01:2025-Prompt-Injection
severity_default: critical
languages: [python, js, java, csharp, go, php]
---

# Agent Hijacking

## Description
Le détournement d'agent (agent hijacking) désigne le scénario où un attaquant, via une injection de prompt directe ou indirecte, parvient à faire exécuter par un agent LLM autonome des actions non prévues via les outils, API ou permissions dont il dispose : envoi de données, transactions, modification de fichiers, appels réseau. La gravité dépasse celle d'une simple mauvaise réponse textuelle car l'agent dispose d'une capacité d'action réelle dans des systèmes externes. Le risque augmente avec le niveau d'autonomie et de privilèges accordés à l'agent.

## Où ça apparaît typiquement
- Agents autonomes multi-étapes (planification + exécution) connectés à des outils sensibles (envoi d'e-mails, paiements, exécution de commandes).
- Agents combinant lecture de contenu non fiable (web, documents) et capacité d'action dans la même session.
- Orchestrateurs multi-agents où un agent compromis peut transmettre des instructions à d'autres agents en aval.
- Absence de séparation entre l'agent qui "décide" et le composant qui "exécute" une action sensible.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Outils à fort impact (exécution système, paiement, suppression de données) invocables directement par le modèle sans validation applicative intermédiaire.
- Agent conservant les mêmes permissions élevées pendant toute la session, y compris lors du traitement de contenu externe non fiable.
- Absence de journalisation détaillée des décisions et actions de l'agent (traçabilité insuffisante en cas d'incident).
- Aucune confirmation humaine (human-in-the-loop) requise avant l'exécution d'actions irréversibles ou à fort impact.

## Remédiation
- Exiger une confirmation humaine explicite pour toute action irréversible ou à fort impact financier/opérationnel.
- Cloisonner les permissions de l'agent selon le contexte (moins de privilèges lors du traitement de contenu externe non fiable).
- Séparer la couche de décision (LLM) de la couche d'exécution (application) avec validation métier indépendante du modèle.
- Journaliser exhaustivement les décisions, appels d'outils et résultats pour permettre l'audit et la détection d'anomalies.
- Voir `rules/remediation/agent-hijacking.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/agent-hijacking/`.

## Références
- OWASP Top 10 for LLM Applications: LLM01:2025 – Prompt Injection
- CWE-1427: Improper Neutralization of Input Used for LLM Prompting
