# Memory Poisoning (empoisonnement de la mémoire d'agent)

## Description de la vulnérabilité

De nombreux agents LLM conservent une mémoire persistante entre les sessions (préférences utilisateur, faits appris, résumés de conversations passées) afin de personnaliser leurs réponses. L'empoisonnement de mémoire consiste à faire enregistrer, via une conversation ou un contenu traité, des informations fausses ou des instructions cachées dans cette mémoire persistante. Ces éléments sont ensuite réinjectés automatiquement dans le contexte de sessions futures, créant une compromission durable qui survit à la session initiale.

Dans `vulnerable.go`, `ProcessTurn` extrait des faits de la réponse de l'agent et les écrit directement en mémoire persistante via `store.Append`, sans confirmation utilisateur. `NextSession` réinjecte ensuite cette mémoire telle quelle comme contexte de confiance dans le prompt système.

## CWE réel utilisé

**CWE-349 : Acceptance of Extraneous Untrusted Data With Trust** (tel que référencé dans `knowledge/llm/memory-poisoning.md`).

## Pourquoi c'est dangereux

- L'écriture en mémoire persistante est déclenchée automatiquement par le modèle, sans validation ni confirmation de l'utilisateur concerné.
- Aucun cloisonnement vérifié entre les espaces mémoire de différents utilisateurs ou tenants : risque de contamination croisée.
- Le contenu de mémoire est réinjecté tel quel dans le prompt système des sessions futures sans revalidation, et traité comme une instruction de confiance.
- Aucun mécanisme de revue, d'édition ou de suppression des entrées mémoire par l'utilisateur concerné : la compromission est durable et invisible.

## Comment le correctif fonctionne

Le fichier `fixed.go` introduit :

1. Une confirmation explicite de l'utilisateur (`UserConfirmer.Confirm`) exigée avant toute écriture durable d'un fait en mémoire persistante.
2. Un cloisonnement strict par tenant via une clé scopée (`tenantScopedKey`), sans partage implicite entre contextes utilisateurs.
3. Une revalidation du contenu mémoire (`revalidate`) avant réinjection : seules les entrées au statut `user_confirmed` sont utilisées, présentées comme "non prescriptives" plutôt que comme instruction de confiance.
4. Des fonctions dédiées (`ListUserMemory`, `DeleteUserMemoryEntry`) permettant à l'utilisateur de consulter et supprimer les entrées mémoire associées à son compte.

## Références

- OWASP Top 10 for LLM Applications: LLM04:2025 – Data and Model Poisoning
- CWE-349: Acceptance of Extraneous Untrusted Data With Trust
- `rules/remediation/memory-poisoning.md`
- `knowledge/llm/memory-poisoning.md`
