// Corrigé — state et nonce générés et vérifiés, ID token validé intégralement
// par la bibliothèque (signature, iss, aud, exp) (CWE-287).

const express = require('express');
const session = require('express-session');
const { Issuer, generators } = require('openid-client');
const app = express();
app.use(session({ secret: process.env.SESSION_SECRET, resave: false, saveUninitialized: false }));

let client;

app.get('/login', (req, res) => {
  const state = generators.state();
  const nonce = generators.nonce();
  req.session.state = state;
  req.session.nonce = nonce;
  res.redirect(client.authorizationUrl({ scope: 'openid profile', state, nonce }));
});

app.get('/callback', async (req, res) => {
  const params = client.callbackParams(req);
  const tokenSet = await client.callback(redirectUri, params, {
    state: req.session.state,
    nonce: req.session.nonce, // la lib valide iss/aud/exp/signature/nonce
  });
  res.json(tokenSet.claims());
});

module.exports = app;
