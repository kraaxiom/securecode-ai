// Faille : RAG Poisoning (CWE-349)
// Les documents crawlés depuis des URLs arbitraires sont découpés et indexés
// directement dans la base vectorielle sans vérification de provenance ni
// détection de contenu suspect, avant d'être injectés tels quels dans le prompt.
const express = require('express');
const router = express.Router();
const { crawl, splitIntoChunks, embed, vectorStore } = require('../rag');

router.post('/rag/ingest', async (req, res) => {
  const { urls } = req.body;

  for (const url of urls) {
    const doc = await crawl(url);
    const chunks = splitIntoChunks(doc);
    for (const chunk of chunks) {
      // Pas de vérification de provenance ni de scan de contenu.
      await vectorStore.upsert(await embed(chunk), { text: chunk });
    }
  }

  res.json({ status: 'ingested' });
});

module.exports = router;
