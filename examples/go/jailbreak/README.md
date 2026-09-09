# Jailbreak de modèle

## Description de la vulnérabilité

Un jailbreak est une technique visant à contourner les alignements de sécurité et les restrictions de contenu d'un modèle de langage, souvent par des formulations de rôle-jeu, des instructions hypothétiques imbriquées ou des reformulations progressives sur plusieurs tours de conversation. Contrairement à l'injection de prompt qui vise à détourner l'application, le jailbreak vise directement les garde-fous du modèle lui-même.

**Ce dossier ne documente et ne contient aucune technique de contournement fonctionnelle, aucun exemple de payload de jailbreak.** Seule la faiblesse architecturale de l'application est illustrée : dans `vulnerable.go`, la fonction `Chat` ne s'appuie que sur le `systemPrompt` comme unique barrière — aucune couche de modération indépendante ne filtre l'entrée ou la sortie, et rien ne limite le nombre de tours de conversation suspects.

## CWE réel utilisé

**CWE-1427 : Improper Neutralization of Input Used for LLM Prompting** (tel que référencé dans `knowledge/llm/jailbreak.md`).

## Pourquoi c'est dangereux

- Le system prompt est la seule barrière de sécurité : s'il est neutralisé ou contourné, plus aucun contrôle n'existe en aval.
- Aucune couche de modération indépendante du modèle (classifieur de sécurité) ne filtre ni l'entrée utilisateur ni la sortie du modèle.
- Le system prompt expose explicitement ses propres règles de refus, ce qui facilite leur contournement ciblé en l'absence de défense en profondeur.
- Rien ne surveille ni ne limite le nombre de tours de conversation utilisés pour tenter une érosion progressive des restrictions du modèle.

## Comment le correctif fonctionne

Le fichier `fixed.go` introduit :

1. Une couche de modération indépendante du modèle principal (`SafetyClassifier.Flags`), appliquée à la fois à l'entrée utilisateur et à la sortie du modèle — le system prompt n'est plus l'unique contrôle.
2. Un suivi par session (`Session.suspiciousTurnCount`) du nombre de tentatives suspectes, avec blocage (`maxSuspiciousTurns`) au-delà d'un seuil, pour contrer les tentatives d'érosion progressive sur plusieurs tours.
3. Un message de refus générique (`refusalMessage`) renvoyé sans détail exploitable lorsqu'un blocage est déclenché.
4. Une journalisation systématique (`AuditLogger.Record`) des blocages d'entrée, de sortie et de session, pour permettre des campagnes de red teaming et un suivi des régressions après mise à jour du modèle.

## Références

- OWASP Top 10 for LLM Applications: LLM01:2025 – Prompt Injection
- CWE-1427: Improper Neutralization of Input Used for LLM Prompting
- `rules/remediation/jailbreak.md`
- `knowledge/llm/jailbreak.md`
