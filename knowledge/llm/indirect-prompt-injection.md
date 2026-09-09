---
id: indirect-prompt-injection
category: llm
cwe: CWE-1427
owasp: LLM01:2025-Prompt-Injection
severity_default: high
languages: [python, js, java, csharp, go, php]
---

# Indirect Prompt Injection

## Description
L'injection de prompt indirecte se produit lorsqu'un contenu malveillant est déposé dans une source externe que le modèle va consulter plus tard (page web, document, e-mail, résultat de recherche, fichier joint), plutôt que d'être saisi directement par l'utilisateur. Lorsque le LLM ingère ce contenu (résumé de page, RAG, agent naviguant sur le web), les instructions cachées qu'il contient sont interprétées comme faisant partie du contexte de confiance, permettant à un tiers de détourner le comportement de l'agent sans jamais interagir directement avec lui.

## Où ça apparaît typiquement
- Agents capables de naviguer sur le web ou de récupérer des URLs fournies par l'utilisateur ou trouvées dynamiquement.
- Systèmes RAG ingérant des documents provenant de sources externes ou partiellement contrôlées par des tiers (PDF, pages wiki, tickets support).
- Résumé automatique d'e-mails, de commentaires ou de contenus générés par des utilisateurs non fiables.
- Agents lisant des métadonnées de fichiers (nom, description, contenu EXIF/commentaires) transmises telles quelles au modèle.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Contenu récupéré depuis une source externe injecté directement dans le contexte du modèle sans marquage de provenance ni délimitation.
- Absence de distinction entre "contenu à résumer/analyser" et "instructions à exécuter" dans le pipeline de traitement.
- Agent disposant d'outils actifs (exécution, appels API) alors qu'il traite simultanément du contenu externe non fiable.
- Pas de règle limitant les actions qu'un agent peut déclencher suite à l'analyse d'un document externe.

## Remédiation
- Marquer explicitement le contenu externe comme non fiable dans le prompt (délimiteurs, balises de provenance) et instruire le modèle à ne jamais exécuter d'instructions qui en proviennent.
- Séparer strictement les phases "lecture/analyse de contenu externe" et "exécution d'action", avec validation applicative entre les deux.
- Limiter les outils disponibles pendant le traitement de contenu non fiable (principe du moindre privilège contextuel).
- Mettre en place une revue humaine pour les actions à fort impact déclenchées suite à l'analyse d'un contenu externe.
- Voir `rules/remediation/indirect-prompt-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/indirect-prompt-injection/`.

## Références
- OWASP Top 10 for LLM Applications: LLM01:2025 – Prompt Injection
- CWE-1427: Improper Neutralization of Input Used for LLM Prompting
