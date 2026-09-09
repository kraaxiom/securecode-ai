# rag-poisoning (CWE-349)

## Description de la vulnérabilité

L'empoisonnement RAG (Retrieval-Augmented Generation) consiste à insérer des documents malveillants ou trompeurs dans la base de connaissances qu'un système consulte pour enrichir ses réponses. Lorsque le mécanisme de récupération sélectionne ce contenu empoisonné, il est injecté dans le contexte du modèle et traité comme une source fiable.

Dans `vulnerable.go`, `ingestDocuments` crawle et indexe n'importe quelle URL fournie sans vérifier que le domaine appartient à une liste de sources de confiance, sans scanner le contenu pour détecter des motifs d'instructions cachées, et sans conserver la moindre métadonnée de provenance sur les chunks indexés. `answerQuery` restitue ensuite un contexte au modèle sans exposer les sources utilisées, rendant impossible tout audit d'une réponse suspecte.

## CWE réel utilisé

**CWE-349 : Acceptance of Extraneous Untrusted Data With Trust**, tel que documenté dans `knowledge/llm/rag-poisoning.md` (OWASP LLM04:2025 – Data and Model Poisoning).

## Pourquoi c'est dangereux

- Un pipeline d'ingestion sans validation de provenance accepte du contenu de n'importe quelle source externe ou semi-publique.
- Un document indexé sans scan de contenu peut contenir des instructions destinées à détourner le comportement de l'agent lorsqu'il est réinjecté dans le contexte du modèle.
- Sans traçabilité de provenance, il est impossible d'auditer a posteriori quelle source a produit une réponse erronée ou malveillante, ni de retirer rapidement un document compromis.

## Comment le correctif fonctionne

Le fichier `fixed.go` applique la remédiation décrite dans `rules/remediation/rag-poisoning.md` :

1. **Allowlist de sources de confiance** (`trustedSources`) : seuls les domaines authentifiés et validés peuvent alimenter la base documentaire ; toute autre source est rejetée et journalisée.
2. **Scan de contenu suspect** (`containsSuspiciousInstructions`) avant indexation, pour détecter des motifs d'instructions impératives cachées dans un document.
3. **Traçabilité de provenance** : chaque chunk indexé conserve `source_url`, `ingested_at` et `trust_level`, exposés jusqu'à la réponse générée (`answerQuery` retourne les sources).
4. **Procédure de retrait rapide** (`revokeDocument`) permettant de désindexer un document identifié comme malveillant.

## Références

- OWASP Top 10 for LLM Applications : LLM04:2025 – Data and Model Poisoning
- CWE-349 : Acceptance of Extraneous Untrusted Data With Trust
- `rules/remediation/rag-poisoning.md`
- `knowledge/llm/rag-poisoning.md`
