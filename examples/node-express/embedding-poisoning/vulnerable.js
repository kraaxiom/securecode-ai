// Faille : Embedding Poisoning (CWE-349)
// Le contenu soumis par l'utilisateur est directement transformé en embedding
// et indexé dans la base vectorielle, sans modération ni limite de fréquence,
// ce qui permet de polluer les résultats de recherche sémantique en volume.
const express = require('express');
const router = express.Router();
const { v4: uuidv4 } = require('uuid');
const { embeddingModel, vectorStore } = require('../rag');

router.post('/embeddings/index', async (req, res) => {
  const { content, sourceId } = req.body;

  // Aucune modération, aucune limite de débit par source.
  const vector = await embeddingModel.embed(content);
  await vectorStore.upsert({ id: uuidv4(), vector, metadata: { source: sourceId } });

  res.json({ status: 'indexed' });
});

module.exports = router;
