# Remédiation — Embedding Poisoning

## Principe
Valider et modérer tout contenu avant génération d'embedding et indexation, surveiller la distribution des vecteurs pour détecter un clustering anormal, et limiter la fréquence/le volume d'indexation par source.

## Python
```python
# Avant — vulnérable : indexation directe du contenu utilisateur sans contrôle
def index_content(user_content, source_id):
    vector = embedding_model.embed(user_content)
    vector_store.upsert(id=uuid4(), vector=vector, metadata={"source": source_id})

# Après — sécurisé : modération, limite de débit par source, détection d'anomalie de densité
RATE_LIMIT_PER_SOURCE = 50  # vecteurs / heure

def index_content(user_content, source_id):
    if not content_moderation.is_allowed(user_content):
        audit_log.record("rejected_content_moderation", source_id)
        return
    if rate_limiter.exceeded(source_id, RATE_LIMIT_PER_SOURCE):
        audit_log.record("rejected_rate_limit", source_id)
        return

    vector = embedding_model.embed(user_content)
    if vector_store.is_anomalous_density(vector):  # proximité anormale de requêtes fréquentes
        audit_log.record("flagged_anomalous_embedding", source_id)
        vector_store.quarantine(vector, metadata={"source": source_id})
        return

    vector_store.upsert(id=uuid4(), vector=vector, metadata={"source": source_id})
```

## JavaScript / Node.js
```js
// Avant — vulnérable
async function indexContent(userContent, sourceId) {
  const vector = await embeddingModel.embed(userContent);
  await vectorStore.upsert({ id: uuidv4(), vector, metadata: { source: sourceId } });
}

// Après — sécurisé
const RATE_LIMIT_PER_SOURCE = 50;

async function indexContent(userContent, sourceId) {
  if (!contentModeration.isAllowed(userContent)) { auditLog.record("rejected_content_moderation", sourceId); return; }
  if (rateLimiter.exceeded(sourceId, RATE_LIMIT_PER_SOURCE)) { auditLog.record("rejected_rate_limit", sourceId); return; }

  const vector = await embeddingModel.embed(userContent);
  if (await vectorStore.isAnomalousDensity(vector)) {
    auditLog.record("flagged_anomalous_embedding", sourceId);
    await vectorStore.quarantine(vector, { source: sourceId });
    return;
  }
  await vectorStore.upsert({ id: uuidv4(), vector, metadata: { source: sourceId } });
}
```

## Checklist de vérification post-patch
- [ ] Tout contenu soumis par un utilisateur passe par une modération avant génération d'embedding et indexation.
- [ ] Une limite de fréquence/volume d'indexation par source est appliquée pour freiner les campagnes automatisées.
- [ ] La distribution des vecteurs est surveillée pour détecter un clustering ou une densité anormale.
- [ ] La pertinence des résultats de recherche sémantique sur des requêtes de référence est réévaluée périodiquement.
