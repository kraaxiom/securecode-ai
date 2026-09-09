---
id: secret-leakage
category: llm
cwe: CWE-200
owasp: LLM02:2025-Sensitive-Information-Disclosure
severity_default: critical
languages: [python, js, java, csharp, go, php]
---

# Secret Leakage via LLM

## Description
Ce pattern couvre les cas où des secrets (clés API, identifiants, jetons internes, mots de passe) se retrouvent exposés à travers l'usage d'un LLM : soit parce qu'ils ont été inclus par erreur dans le prompt système ou dans le contexte transmis au modèle, soit parce que le modèle a été entraîné/fine-tuné sur des données contenant des secrets, soit parce qu'un agent avec accès à des variables d'environnement ou fichiers de configuration les restitue dans ses réponses.

## Où ça apparaît typiquement
- Prompts système construits dynamiquement incluant des clés API ou des identifiants de connexion pour usage interne par l'agent.
- Agents ayant accès au système de fichiers ou aux variables d'environnement, pouvant les inclure dans un résumé ou une réponse.
- Fine-tuning ou logs de conversation contenant des secrets saisis par erreur par des utilisateurs ou développeurs.
- Intégrations tierces (plugins, MCP) transmettant des jetons d'authentification dans le contexte partagé avec le modèle.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Secrets en clair présents dans le code source du prompt système ou dans les templates de prompt versionnés.
- Absence de filtrage de sortie (output scanning) détectant des motifs de secrets avant renvoi de la réponse du modèle.
- Journaux de conversation conservant les prompts et réponses en clair sans masquage des secrets potentiellement présents.
- Agent disposant d'un accès en lecture aux fichiers de configuration ou variables d'environnement sans restriction de scope.

## Remédiation
- Ne jamais inclure de secrets en clair dans un prompt système ; utiliser des références indirectes résolues côté application uniquement.
- Mettre en place un filtrage de sortie détectant les motifs de secrets (clés API, tokens) avant tout renvoi de réponse.
- Restreindre strictement l'accès des agents aux fichiers/variables sensibles, en excluant les répertoires de configuration critiques.
- Masquer ou expurger les secrets potentiels dans les journaux de conversation et les données utilisées pour l'entraînement.
- Voir `rules/remediation/secret-leakage.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/secret-leakage/`.

## Références
- OWASP Top 10 for LLM Applications: LLM02:2025 – Sensitive Information Disclosure
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
