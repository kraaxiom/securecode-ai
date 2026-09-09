---
id: jailbreak
category: llm
cwe: CWE-1427
owasp: LLM01:2025-Prompt-Injection
severity_default: medium
languages: [python, js, java, csharp, go, php]
---

# Jailbreak de modèle

## Description
Un jailbreak est une technique visant à contourner les alignements de sécurité et les restrictions de contenu d'un modèle de langage, souvent par des formulations de rôle-jeu, des instructions hypothétiques imbriquées ou des reformulations progressives, pour obtenir des réponses que le modèle refuserait normalement de produire. Contrairement à l'injection de prompt qui vise à détourner l'application, le jailbreak vise directement les garde-fous du modèle lui-même. Ce document couvre uniquement la détection défensive de patterns applicatifs favorisant ce risque, sans documenter de techniques de contournement fonctionnelles.

## Où ça apparaît typiquement
- Applications exposant un accès conversationnel libre au modèle sans filtrage de contenu additionnel côté application.
- Absence de limite sur la longueur/complexité des instructions système que l'utilisateur peut faire "oublier" ou surcharger.
- Chatbots publics sans supervision ni limitation de contexte multi-tours permettant une érosion progressive des restrictions.
- Produits intégrant un LLM pour générer du contenu destiné à publication directe sans revue.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de couche de modération de contenu indépendante du modèle (classifieur de sécurité, filtre de sortie).
- Aucune limite sur le nombre de tours de conversation utilisés pour "négocier" un changement de comportement du modèle.
- System prompt exposant explicitement ses propres règles de refus (ce qui facilite leur contournement ciblé) sans défense en profondeur.
- Pas de tests de robustesse (red teaming) documentés avant mise en production d'un assistant conversationnel exposé publiquement.

## Remédiation
- Ajouter une couche de modération indépendante du modèle principal (classifieur de sécurité en entrée et en sortie).
- Ne jamais considérer le system prompt comme seule barrière de sécurité ; imposer des contrôles applicatifs en aval des réponses.
- Limiter le contexte conversationnel exploitable et surveiller les schémas de conversation anormaux (tentatives répétées de contournement).
- Effectuer des campagnes de red teaming régulières et documenter les régressions de sécurité du modèle après mise à jour.
- Voir `rules/remediation/jailbreak.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/jailbreak/`.

## Références
- OWASP Top 10 for LLM Applications: LLM01:2025 – Prompt Injection
- CWE-1427: Improper Neutralization of Input Used for LLM Prompting
