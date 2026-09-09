# Model Extraction (vol de modèle)

## Description de la vulnérabilité

L'extraction de modèle consiste à interroger massivement et systématiquement une API de modèle exposée publiquement afin d'en reconstituer un équivalent fonctionnel (modèle de substitution / distillation) ou d'en dériver la logique interne, sans y être autorisé. Cette technique permet à un concurrent de s'approprier la propriété intellectuelle investie dans un modèle propriétaire, ou de faciliter la préparation d'autres attaques.

Dans `vulnerable.go`, `InferHandler` expose l'endpoint d'inférence sans aucune limite de débit ni quota par client, et renvoie systématiquement les logits bruts du modèle en plus du texte généré.

## CWE réel utilisé

**CWE-200 : Exposure of Sensitive Information to an Unauthorized Actor** (tel que référencé dans `knowledge/llm/model-extraction.md`).

## Pourquoi c'est dangereux

- Aucune limitation de débit ou de quota par clé API/utilisateur sur l'endpoint d'inférence : un script peut interroger l'API sans restriction.
- Exposition de métadonnées internes du modèle (logits complets) non nécessaires au cas d'usage métier, ce qui facilite grandement la reconstruction d'un modèle de substitution par distillation.
- Aucune surveillance des volumes de requêtes ni détection d'usage automatisé anormal (patterns non humains).
- Pas de mécanisme de traçabilité des sorties permettant de détecter une réutilisation non autorisée du modèle extrait.

## Comment le correctif fonctionne

Le fichier `fixed.go` introduit :

1. Un quota strict par clé API (`RateLimiter.Allow`, `requestsPerHourLimit = 100`) appliqué avant tout traitement de la requête.
2. La suppression de l'exposition des logits bruts : `predict` est appelé avec `returnLogits=false` et seule la sortie textuelle (`Text`) est renvoyée au client.
3. Une détection de patterns d'usage anormal (`UsageMonitor.DetectsExtractionPattern`) qui bloque les clients dont le volume ou la régularité des requêtes évoque une tentative d'extraction systématique.
4. Une journalisation systématique (`AuditLogger.Record`) des dépassements de quota et des détections suspectes, pour appuyer un encadrement contractuel de l'usage de l'API.

## Références

- OWASP Top 10 for LLM Applications: LLM10:2025 – Unbounded Consumption
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
- `rules/remediation/model-extraction.md`
- `knowledge/llm/model-extraction.md`
