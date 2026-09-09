// Faille : SSRF aveugle (Blind SSRF) (CWE-918)
// Un worker asynchrone déclenche une requête sortante vers une URL fournie
// par l'utilisateur pour valider un webhook, sans jamais renvoyer le
// contenu récupéré au client. L'absence de retour visible ne supprime pas
// le risque : aucune whitelist ni validation de destination n'est appliquée.
const express = require('express');
const router = express.Router();

router.post('/webhooks/register', (req, res) => {
  const callbackUrl = req.body.callback_url;

  // Traitement "fire and forget" : le résultat n'est jamais exposé au client.
  fetch(callbackUrl, { method: 'HEAD' })
    .then(() => console.log('Webhook validé'))
    .catch(() => console.log('Webhook injoignable'));

  res.status(202).json({ status: 'en cours de validation' });
});

module.exports = router;
