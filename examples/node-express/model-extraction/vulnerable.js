// Faille : Model Extraction (CWE-200)
// L'endpoint d'inférence renvoie les logits complets sans aucune limitation de
// débit par client, permettant une interrogation systématique massive pour
// reconstituer un modèle équivalent (distillation non autorisée).
const express = require('express');
const router = express.Router();
const { model } = require('../ml');

router.post('/v1/infer', async (req, res) => {
  const result = await model.predict(req.body.input, { returnLogits: true });
  // Aucune limite de débit, exposition des logits bruts.
  res.json({ output: result.text, logits: result.logits });
});

module.exports = router;
