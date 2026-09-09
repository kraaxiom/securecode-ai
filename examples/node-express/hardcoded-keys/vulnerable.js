// Faille : Clés/secrets codés en dur (CWE-798)
// La clé secrète du service de paiement est écrite en dur dans le code source.
// Toute personne ayant accès au dépôt (ou à l'historique git) peut la lire
// et l'utiliser pour usurper des appels à l'API tierce.
const express = require('express');
const router = express.Router();

// Secret codé en dur : visible dans le code source et l'historique git.
const PAYMENT_API_SECRET = 'sk_live_EXAMPLE_NOT_A_REAL_KEY';

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
