// Correction : Clés/secrets codés en dur (CWE-798)
// Le secret est chargé depuis une variable d'environnement injectée au runtime,
// jamais versionnée. Une absence de secret bloque explicitement le démarrage.
const express = require('express');
const router = express.Router();

const PAYMENT_API_SECRET = process.env.PAYMENT_API_SECRET;
if (!PAYMENT_API_SECRET) {
  throw new Error('PAYMENT_API_SECRET manquant dans l\'environnement');
}

router.post('/payments/charge', async (req, res) => {
  const { amount, token } = req.body;

  const result = await chargeWithProvider(PAYMENT_API_SECRET, amount, token);
  res.json({ status: result.status });
});

async function chargeWithProvider(secret, amount, token) {
  // Appel simulé au fournisseur de paiement
  return { status: 'ok' };
}

module.exports = router;
