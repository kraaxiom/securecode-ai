// Correction : Model Extraction (CWE-200)
// Quota strict par clé API, sortie minimisée (pas de logits bruts) et
// détection des patterns d'usage évoquant une tentative d'extraction du modèle.
const express = require('express');
const router = express.Router();
const rateLimit = require('express-rate-limit');
const { model } = require('../ml');
const { usageMonitor } = require('../security');

const inferLimiter = rateLimit({
  keyGenerator: (req) => req.headers['x-api-key'],
  max: 100,
  windowMs: 3600_000,
});

router.post('/v1/infer', inferLimiter, async (req, res) => {
  const apiKey = req.headers['x-api-key'];
  if (await usageMonitor.detectsExtractionPattern(apiKey)) {
    auditLog.record('suspected_model_extraction', apiKey);
    return res.status(429).json({ error: 'Usage anormal détecté' });
  }

  const result = await model.predict(req.body.input, { returnLogits: false });
  res.json({ output: result.text }); // pas de logits/probabilités bruts exposés
});

module.exports = router;
