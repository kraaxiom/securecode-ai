# Indirect Prompt Injection (injection de prompt indirecte)

## Description de la vulnérabilité

L'injection de prompt indirecte se produit lorsqu'un contenu malveillant est déposé dans une source externe que le modèle va consulter plus tard (page web, document, e-mail, résultat de recherche), plutôt que d'être saisi directement par l'utilisateur. Lorsque le LLM ingère ce contenu, les instructions cachées qu'il contient sont interprétées comme faisant partie du contexte de confiance, permettant à un tiers de détourner le comportement de l'agent sans jamais interagir directement avec lui.

Dans `vulnerable.go`, `SummarizeURL` récupère le contenu d'une page web et l'insère tel quel dans le message utilisateur envoyé au LLM, sans aucun marquage de provenance ni délimitation, alors que des outils à fort impact (`send_email`, `execute_code`) restent actifs pendant cette analyse.

## CWE réel utilisé

**CWE-1427 : Improper Neutralization of Input Used for LLM Prompting** (tel que référencé dans `knowledge/llm/indirect-prompt-injection.md`).

## Pourquoi c'est dangereux

- Le contenu externe est injecté directement dans le contexte du modèle sans marquage de provenance ni délimitation : le modèle ne peut pas distinguer "texte à résumer" de "instruction à exécuter".
- Des outils actifs (exécution de code, envoi d'e-mail) restent disponibles pendant le traitement de contenu externe non fiable, ce qui permet à des instructions cachées dans ce contenu de déclencher des actions réelles.
- Aucune règle ne limite les actions qu'un agent peut déclencher suite à l'analyse d'un document externe.
- L'attaque ne nécessite aucune interaction directe entre l'attaquant et l'agent : il suffit de déposer le contenu piégé là où l'agent ira le lire.

## Comment le correctif fonctionne

Le fichier `fixed.go` introduit :

1. Un message système explicite (`untrustedSystemPrompt`) rappelant au modèle que le contenu à suivre provient d'une source externe non fiable et ne doit jamais être traité comme une instruction.
2. Une délimitation stricte du contenu externe par des balises (`<untrusted_external_content>...</untrusted_external_content>`), séparant clairement "contenu à analyser" et "instruction à exécuter".
3. La désactivation de tout outil (`Tools: []string{}`) pendant la phase de lecture/analyse de contenu externe non fiable, appliquant le principe du moindre privilège contextuel.
4. Une revalidation de la sortie du modèle (`validateOutput`) côté application avant toute utilisation ultérieure, y compris après un simple résumé.

## Références

- OWASP Top 10 for LLM Applications: LLM01:2025 – Prompt Injection
- CWE-1427: Improper Neutralization of Input Used for LLM Prompting
- `rules/remediation/indirect-prompt-injection.md`
- `knowledge/llm/indirect-prompt-injection.md`
