// Faille : Data Poisoning (CWE-349)
// Les données de fine-tuning sont récupérées depuis des sources externes et
// intégrées telles quelles au jeu d'entraînement, sans vérification de
// provenance, d'intégrité ni détection d'anomalies statistiques.
const express = require('express');
const router = express.Router();
const { fetchAndParse } = require('../datasources');

router.post('/training/build-dataset', async (req, res) => {
  const sources = req.body.sources; // liste d'URLs fournies par l'appelant

  const dataset = [];
  for (const url of sources) {
    // Aucune vérification de provenance/intégrité avant intégration.
    dataset.push(...await fetchAndParse(url));
  }

  await trainingQueue.enqueue(dataset);
  res.json({ count: dataset.length });
});

module.exports = router;
