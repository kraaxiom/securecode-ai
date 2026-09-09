# Embedding Poisoning (empoisonnement d'embeddings)

## Description de la vulnérabilité

L'empoisonnement d'embeddings vise directement l'espace vectoriel utilisé par un système de recherche sémantique ou de RAG. Un attaquant fabrique un contenu dont l'embedding est délibérément proche de requêtes légitimes fréquentes, afin qu'il soit systématiquement remonté par le moteur de similarité même s'il n'est pas pertinent, ou qu'il pollue le clustering utilisé pour des décisions automatisées.

Dans `vulnerable.go`, `IndexContent` transforme n'importe quel contenu utilisateur en embedding et l'indexe immédiatement dans la base vectorielle, sans modération, sans limite de fréquence par source, et sans détection de clustering anormal.

## CWE réel utilisé

**CWE-349 : Acceptance of Extraneous Untrusted Data With Trust** (tel que référencé dans `knowledge/llm/embedding-poisoning.md`).

## Pourquoi c'est dangereux

- Aucun contrôle du contenu soumis avant génération et indexation de son embedding : du contenu malveillant peut entrer directement dans l'espace vectoriel.
- Aucune détection d'anomalies de densité ou de clustering inhabituel : un attaquant peut positionner ses vecteurs à proximité de requêtes fréquentes sans être détecté.
- Pas de limite de fréquence ou de volume par source : une campagne automatisée peut injecter un grand nombre de vecteurs empoisonnés rapidement.
- Aucune réévaluation périodique de la pertinence des résultats de recherche sémantique en production, ce qui laisse la dérive s'installer durablement.

## Comment le correctif fonctionne

Le fichier `fixed.go` introduit :

1. Une étape de modération (`ContentModeration.IsAllowed`) appliquée avant toute génération d'embedding.
2. Une limite de débit par source (`RateLimiter.Exceeded`, `rateLimitPerSource = 50`) qui freine les campagnes d'indexation automatisées.
3. Une détection de densité anormale (`VectorStore.IsAnomalousDensity`) après génération de l'embedding : les vecteurs suspects sont mis en quarantaine (`Quarantine`) plutôt qu'indexés directement.
4. Une journalisation systématique (`AuditLogger.Record`) de chaque rejet ou anomalie détectée.

## Références

- OWASP Top 10 for LLM Applications: LLM04:2025 – Data and Model Poisoning
- CWE-349: Acceptance of Extraneous Untrusted Data With Trust
- `rules/remediation/embedding-poisoning.md`
- `knowledge/llm/embedding-poisoning.md`
