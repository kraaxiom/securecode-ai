// Correction : Data Poisoning (CWE-349)
// Seules les sources de confiance déclarées sont acceptées, leur intégrité est
// vérifiée par hash signé, et les données subissent une détection d'anomalies
// statistiques avant d'être intégrées au pipeline de fine-tuning.
const express = require('express');
const router = express.Router();
const { fetchAndParse } = require('../datasources');

const TRUSTED_SOURCES = new Set(['internal-labeled-set', 'verified-partner-feed']);

router.post('/training/build-dataset', async (req, res) => {
  const sources = req.body.sources; // objets { id, url, signedHash }

  const dataset = [];
  for (const source of sources) {
    if (!TRUSTED_SOURCES.has(source.id)) {
      auditLog.record('rejected_untrusted_source', source.id);
      continue;
    }
    verifyIntegrity(source, source.signedHash);
    let batch = await fetchAndParse(source.url);
    batch = filterStatisticalOutliers(batch); // détection d'anomalies
    dataset.push(...batch);
  }

  await trainingQueue.enqueue(dataset);
  res.json({ count: dataset.length });
});

module.exports = router;
