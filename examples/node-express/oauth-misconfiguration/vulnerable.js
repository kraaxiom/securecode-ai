// Vulnérable — Mauvaise configuration OAuth (CWE-287)
// Le paramètre 'state' n'est ni généré ni vérifié au retour du callback, et
// l'ID token reçu du fournisseur n'est pas validé (issuer/audience). Le
// callback est donc exposé à une CSRF et à l'usurpation d'un token forgé.

const express = require('express');
const { Issuer } = require('openid-client');
const app = express();

let client; // initialisé via Issuer.discover(...) au démarrage

app.get('/login', (req, res) => {
  res.redirect(client.authorizationUrl({ scope: 'openid profile' })); // pas de state
});

app.get('/callback', async (req, res) => {
  const params = client.callbackParams(req);
  const tokenSet = await client.callback(redirectUri, params); // pas de state/nonce vérifié
  res.json(tokenSet.claims()); // claims utilisés sans validation stricte
});

module.exports = app;
