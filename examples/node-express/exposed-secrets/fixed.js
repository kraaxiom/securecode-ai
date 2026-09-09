// Correction : Secrets exposés dans le code (CWE-798)
// Les secrets sont récupérés à l'exécution depuis un gestionnaire de
// secrets dédié (AWS Secrets Manager), jamais codés en dur ni journalisés.
const express = require('express');
const Stripe = require('stripe');
const mysql = require('mysql2/promise');
const { SecretsManagerClient, GetSecretValueCommand } = require('@aws-sdk/client-secrets-manager');
const router = express.Router();

const secretsClient = new SecretsManagerClient({ region: 'eu-west-1' });

async function getSecret(secretId) {
  const { SecretString } = await secretsClient.send(new GetSecretValueCommand({ SecretId: secretId }));
  return JSON.parse(SecretString);
}

router.post('/checkout', async (req, res) => {
  const { stripeSecretKey } = await getSecret('prod/stripe/api-key');
  const { host, user, password } = await getSecret('prod/db/credentials');

  // Aucun secret en clair dans le code ni dans les logs.
  const stripe = new Stripe(stripeSecretKey);
  const db = await mysql.createConnection({ host, user, password });

  const charge = await stripe.charges.create({ amount: req.body.amount, currency: 'xof' });
  await db.end();
  res.json({ id: charge.id });
});

module.exports = router;
