---
id: tool-injection
category: llm
cwe: CWE-1427
owasp: LLM06:2025-Excessive-Agency
severity_default: high
languages: [python, js, java, csharp, go, php]
---

# Tool Injection (abus d'appel d'outils)

## Description
L'injection d'outils survient lorsqu'un contenu non fiable parvient à manipuler les paramètres ou la sélection des outils (function calling / tool use) qu'un agent LLM invoque, par exemple en faisant croire au modèle qu'un outil supplémentaire doit être appelé, ou en altérant les arguments transmis à un outil légitime. Cela permet à un attaquant de détourner des appels API, d'exfiltrer des données via un outil de sortie, ou de forcer l'exécution d'opérations non prévues par le développeur de l'agent.

## Où ça apparaît typiquement
- Agents utilisant le function calling où les descriptions d'outils ou leurs paramètres proviennent en partie de données externes.
- Systèmes MCP (Model Context Protocol) ou plugins tiers exposant des outils avec des descriptions non fiables ou non vérifiées.
- Agents combinant plusieurs outils où la sortie d'un outil non fiable alimente les paramètres d'un outil suivant sans validation.
- Absence de schéma strict de validation des arguments d'outils avant exécution.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Paramètres d'appel d'outil transmis directement du texte généré par le modèle vers une exécution sans validation de schéma stricte.
- Description ou métadonnées d'outils chargées dynamiquement depuis une source externe non contrôlée (registre de plugins tiers).
- Sortie d'un outil réinjectée comme entrée d'un autre outil sans étape de nettoyage/validation intermédiaire.
- Absence de liste blanche d'outils autorisés par contexte d'exécution (tous les outils toujours disponibles quel que soit le niveau de confiance du contenu traité).

## Remédiation
- Valider strictement les arguments de chaque appel d'outil contre un schéma typé avant exécution, indépendamment de ce que le modèle a généré.
- Ne charger que des définitions d'outils provenant de sources de confiance vérifiées, avec revue avant intégration.
- Appliquer une liste blanche contextuelle d'outils autorisés selon le niveau de confiance du contenu en cours de traitement.
- Auditer et journaliser chaque appel d'outil avec ses paramètres réels pour permettre la détection d'anomalies après incident.
- Voir `rules/remediation/tool-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/tool-injection/`.

## Références
- OWASP Top 10 for LLM Applications: LLM06:2025 – Excessive Agency
- CWE-1427: Improper Neutralization of Input Used for LLM Prompting
