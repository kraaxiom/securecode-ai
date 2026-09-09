// Faille : Model Inversion (CWE-200)
// L'API d'inférence, fine-tunée sur des données internes sensibles, ne limite
// pas le volume de requêtes, permettant une reconstruction progressive de
// données mémorisées par interrogation répétée et méthodique du modèle.
const express = require('express');
const router = express.Router();
const { model } = require('../ml');

router.post('/v1/infer', async (req, res) => {
  // Aucune limite de volume ni de structure de requêtes.
  const result = await model.predict(req.body.input);
  res.json(result);
});

module.exports = router;
