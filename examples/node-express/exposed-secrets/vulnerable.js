// Faille : Secrets exposés dans le code (CWE-798)
// La clé API Stripe et les identifiants de connexion à la base de données
// sont codés en dur dans le code source, et la clé est même journalisée,
// ce qui l'expose dans les logs et dans l'historique du dépôt.
const express = require('express');
const Stripe = require('stripe');
const mysql = require('mysql2/promise');
const router = express.Router();

// Secret en dur : visible par quiconque a accès au code ou au dépôt Git.
const stripe = new Stripe('sk_live_EXAMPLE_NOT_A_REAL_KEY');

router.post('/checkout', async (req, res) => {
  console.log('Utilisation de la clé Stripe:', 'sk_live_EXAMPLE_NOT_A_REAL_KEY');

  const db = await mysql.createConnection({
    host: 'db.example.com',
    user: 'admin',
    password: 'Sup3rS3cret!2024', // mot de passe en dur
  });

  const charge = await stripe.charges.create({ amount: req.body.amount, currency: 'xof' });
  await db.end();
  res.json({ id: charge.id });
});

module.exports = router;
