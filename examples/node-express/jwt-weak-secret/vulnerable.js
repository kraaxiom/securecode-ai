// Vulnérable — JWT signé avec un secret faible (CWE-326)
// Le secret HMAC utilisé pour signer les tokens est une valeur d'exemple de
// documentation, courte et devinable. Un attaquant peut le retrouver hors
// ligne par force brute/dictionnaire puis forger des tokens valides.

const express = require('express');
const jwt = require('jsonwebtoken');
const app = express();
app.use(express.json());

const JWT_SECRET = 'your-256-bit-secret'; // jamais changé depuis la doc

app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });

  const token = jwt.sign({ sub: user.id, role: user.role }, JWT_SECRET, { algorithm: 'HS256' });
  res.json({ token });
});

module.exports = app;
