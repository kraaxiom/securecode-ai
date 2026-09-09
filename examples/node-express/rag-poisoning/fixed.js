// Correction : RAG Poisoning (CWE-349)
// Seules les sources en allowlist sont ingérées, chaque document est scanné
// pour détecter des instructions cachées avant indexation, et la provenance
// est conservée jusqu'à la réponse générée pour permettre l'audit.
const express = require('express');
const router = express.Router();
const { crawl, splitIntoChunks, embed, vectorStore } = require('../rag');
const { contentScanner, extractDomain } = require('../security');

const TRUSTED_SOURCES = new Set(['docs.internal.acme.com', 'verified-partner-wiki.acme.com']);

router.post('/rag/ingest', async (req, res) => {
  const { urls } = req.body;

  for (const url of urls) {
    const domain = extractDomain(url);
    if (!TRUSTED_SOURCES.has(domain)) {
      auditLog.record('rejected_untrusted_source', url);
      continue;
    }

    const doc = await crawl(url);
    if (await contentScanner.containsSuspiciousInstructions(doc)) {
      auditLog.record('flagged_suspicious_document', url);
      continue;
    }

    const chunks = splitIntoChunks(doc);
    for (const chunk of chunks) {
      await vectorStore.upsert(await embed(chunk), {
        text: chunk,
        sourceUrl: url,
        ingestedAt: Date.now(),
        trustLevel: 'verified',
      });
    }
  }

  res.json({ status: 'ingested' });
});

module.exports = router;
