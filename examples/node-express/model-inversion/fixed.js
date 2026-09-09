// Correction : Model Inversion (CWE-200)
// Limitation stricte du volume de requêtes par clé API et détection des
// patterns de reconstruction progressive, en complément d'un entraînement
// avec confidentialité différentielle en amont (hors périmètre de ce fichier).
const express = require('express');
const router = express.Router();
const rateLimit = require('express-rate-limit');
const { model } = require('../ml');
const { usageMonitor } = require('../security');

const inferLimiter = rateLimit({
  keyGenerator: (req) => req.headers['x-api-key'],
  max: 50,
  windowMs: 3600_000,
});

router.post('/v1/infer', inferLimiter, async (req, res) => {
  const apiKey = req.headers['x-api-key'];
  if (await usageMonitor.detectsReconstructionPattern(apiKey)) {
    auditLog.record('suspected_model_inversion', apiKey);
    return res.status(429).json({ error: 'Usage anormal détecté' });
  }

  const result = await model.predict(req.body.input);
  res.json(result);
});

module.exports = router;
