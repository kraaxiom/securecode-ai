// Correction : Embedding Poisoning (CWE-349)
// Le contenu est modéré avant indexation, une limite de débit par source
// freine les campagnes d'injection automatisées, et les vecteurs présentant
// une densité anormale sont mis en quarantaine plutôt qu'indexés directement.
const express = require('express');
const router = express.Router();
const { v4: uuidv4 } = require('uuid');
const { embeddingModel, vectorStore } = require('../rag');
const { contentModeration, rateLimiter } = require('../security');

const RATE_LIMIT_PER_SOURCE = 50; // vecteurs / heure

router.post('/embeddings/index', async (req, res) => {
  const { content, sourceId } = req.body;

  if (!(await contentModeration.isAllowed(content))) {
    auditLog.record('rejected_content_moderation', sourceId);
    return res.status(400).json({ error: 'Contenu rejeté par la modération' });
  }
  if (rateLimiter.exceeded(sourceId, RATE_LIMIT_PER_SOURCE)) {
    auditLog.record('rejected_rate_limit', sourceId);
    return res.status(429).json({ error: 'Limite de débit atteinte' });
  }

  const vector = await embeddingModel.embed(content);
  if (await vectorStore.isAnomalousDensity(vector)) {
    auditLog.record('flagged_anomalous_embedding', sourceId);
    await vectorStore.quarantine(vector, { source: sourceId });
    return res.status(202).json({ status: 'quarantined' });
  }

  await vectorStore.upsert({ id: uuidv4(), vector, metadata: { source: sourceId } });
  res.json({ status: 'indexed' });
});

module.exports = router;
