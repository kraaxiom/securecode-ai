---
id: prompt-injection
category: llm
cwe: CWE-1427
owasp: LLM01:2025-Prompt-Injection
severity_default: high
languages: [python, js, java, csharp, go, php]
---

# Prompt Injection

## Description
L'injection de prompt survient lorsqu'une entrée fournie directement par l'utilisateur parvient à modifier le comportement prévu d'un modèle de langage en y insérant des instructions concurrentes à celles du system prompt. Contrairement à l'injection indirecte, l'attaquant interagit ici en direct avec l'interface de conversation ou l'API. Le modèle, incapable de distinguer structurellement "instruction du développeur" et "instruction de l'utilisateur", peut être amené à ignorer ses garde-fous, révéler son prompt système ou exécuter des actions non autorisées via les outils qui lui sont connectés.

## Où ça apparaît typiquement
- Chatbots et assistants exposant directement l'entrée utilisateur au modèle sans séparation structurelle des rôles.
- Applications concaténant le prompt système et l'entrée utilisateur dans une seule chaîne de texte brute.
- Agents LLM disposant d'outils (exécution de code, appels API, accès fichiers) accessibles depuis le contexte conversationnel.
- Interfaces permettant à l'utilisateur de définir ou surcharger des "instructions personnalisées".

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Construction du prompt par simple concaténation de chaînes plutôt que via les rôles structurés de l'API (system/user/assistant).
- Absence de séparation claire entre contenu de confiance (instructions développeur) et contenu non fiable (entrée utilisateur).
- Aucune validation ni filtrage de sortie avant que la réponse du modèle ne déclenche une action sensible (appel d'outil, écriture en base).
- Absence de limite de privilèges sur les outils accessibles à l'agent (le modèle peut appeler n'importe quel outil sans contrôle applicatif additionnel).
- Pas de journalisation des prompts système et des sorties utilisées pour des décisions automatisées.

## Remédiation
- Utiliser les API supportant une séparation structurée des rôles (system/developer vs user) plutôt que la concaténation de texte.
- Traiter toute sortie du modèle comme une donnée non fiable avant de l'utiliser pour déclencher une action (validation côté application, pas seulement côté prompt).
- Appliquer le principe du moindre privilège aux outils/API exposés à l'agent, avec confirmation explicite pour les actions sensibles.
- Mettre en place une détection/filtrage des tentatives d'injection connues en entrée et en sortie, sans s'y fier comme unique barrière.
- Voir `rules/remediation/prompt-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/prompt-injection/`.

## Références
- OWASP Top 10 for LLM Applications: LLM01:2025 – Prompt Injection
- CWE-1427: Improper Neutralization of Input Used for LLM Prompting
