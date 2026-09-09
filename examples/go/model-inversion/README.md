# model-inversion (CWE-200)

## Description de la vulnérabilité

L'inversion de modèle (« model inversion ») est une technique par laquelle un attaquant, en interrogeant de façon répétée et méthodique un modèle exposé via une API, parvient à reconstruire progressivement des informations sensibles présentes dans ses données d'entraînement (données personnelles, documents propriétaires, secrets internes), même sans y avoir un accès direct.

Dans `vulnerable.go`, le gestionnaire `inferHandler` expose le modèle sans aucune limite de volume de requêtes, sans surveillance des motifs d'usage, et sans qu'aucun test de mémorisation n'ait été réalisé avant la mise en production. Un client peut ainsi interroger le modèle un nombre illimité de fois avec des variations minimes de la même entrée pour extraire progressivement des données mémorisées.

## CWE réel utilisé

**CWE-200 : Exposure of Sensitive Information to an Unauthorized Actor**, tel que documenté dans `knowledge/llm/model-inversion.md` (OWASP LLM02:2025 – Sensitive Information Disclosure).

## Pourquoi c'est dangereux

- Un modèle fine-tuné sur des données internes (dossiers clients, données médicales, code propriétaire) peut mémoriser et restituer partiellement des exemples d'entraînement.
- Sans limite de requêtes ni surveillance, un attaquant dispose d'un budget d'interrogation illimité pour affiner sa reconstruction.
- L'absence de test de mémorisation avant déploiement signifie que le niveau de risque n'est jamais mesuré ni maîtrisé.

## Comment le correctif fonctionne

Le fichier `fixed.go` introduit :

1. **Limitation de requêtes (`rate limiting`)** par clé API via `usageMonitor.allow`, bornant le volume d'interrogations possibles par fenêtre de temps.
2. **Détection de motifs de reconstruction** via `detectsReconstructionPattern`, qui déclenche un blocage anticipé en cas d'usage suspect.
3. **Journalisation d'audit** (`auditLog`) de tout événement de dépassement de quota ou d'usage anormal, pour permettre l'investigation.
4. En amont (pipeline d'entraînement, non représenté dans ce handler HTTP), l'application de techniques de confidentialité différentielle (DP-SGD) et de tests de mémorisation avant mise en production, conformément à `rules/remediation/model-inversion.md`.

## Références

- OWASP Top 10 for LLM Applications : LLM02:2025 – Sensitive Information Disclosure
- CWE-200 : Exposure of Sensitive Information to an Unauthorized Actor
- `rules/remediation/model-inversion.md`
- `knowledge/llm/model-inversion.md`
