---
id: memory-poisoning
category: llm
cwe: CWE-349
owasp: LLM04:2025-Data-and-Model-Poisoning
severity_default: high
languages: [python, js, java, csharp, go]
---

# Memory Poisoning (empoisonnement de la mémoire d'agent)

## Description
De nombreux agents LLM conservent une mémoire persistante entre les sessions (préférences utilisateur, faits appris, résumés de conversations passées) afin de personnaliser leurs réponses. L'empoisonnement de mémoire consiste à faire enregistrer, via une conversation ou un contenu traité, des informations fausses ou des instructions cachées dans cette mémoire persistante. Ces éléments seront ensuite réinjectés automatiquement dans le contexte de sessions futures, créant une compromission durable qui survit à la session initiale.

## Où ça apparaît typiquement
- Assistants avec mémoire long terme enregistrant automatiquement des "faits" extraits de la conversation.
- Systèmes multi-utilisateurs où la mémoire d'un agent peut être influencée indirectement par du contenu partagé (documents, tickets communs).
- Agents résumant automatiquement chaque session et réinjectant ce résumé comme contexte de confiance à la session suivante.
- Absence de distinction entre mémoire validée par l'utilisateur et mémoire déduite automatiquement par le modèle.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Écriture en mémoire persistante déclenchée automatiquement par le modèle sans validation ni confirmation de l'utilisateur.
- Absence de séparation entre les espaces mémoire de différents utilisateurs ou tenants (risque de contamination croisée).
- Contenu de mémoire réinjecté tel quel dans le prompt système des sessions futures sans revalidation.
- Aucun mécanisme de revue, d'édition ou de suppression des entrées mémoire par l'utilisateur concerné.

## Remédiation
- Exiger une confirmation explicite de l'utilisateur avant l'écriture durable d'une information en mémoire persistante.
- Cloisonner strictement la mémoire par utilisateur/tenant, sans partage implicite entre contextes.
- Traiter le contenu de la mémoire réinjecté comme une donnée à revalider, pas comme une instruction de confiance absolue.
- Offrir à l'utilisateur un moyen de consulter, corriger et supprimer les entrées mémoire associées à son compte.
- Voir `rules/remediation/memory-poisoning.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/memory-poisoning/`.

## Références
- OWASP Top 10 for LLM Applications: LLM04:2025 – Data and Model Poisoning
- CWE-349: Acceptance of Extraneous Untrusted Data With Trust
