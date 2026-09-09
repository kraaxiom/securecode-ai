// Corrigé — secret chargé depuis un gestionnaire de secrets, avec entropie
// minimale vérifiée au démarrage (256 bits pour HS256) (CWE-326).

const express = require('express');
const jwt = require('jsonwebtoken');
const app = express();
app.use(express.json());

const JWT_SECRET = process.env.JWT_SECRET; // généré une fois via crypto.randomBytes(32), stocké dans un vault
if (!JWT_SECRET || Buffer.byteLength(JWT_SECRET, 'utf8') < 32) {
  throw new Error('JWT_SECRET manquant ou insuffisant (256 bits minimum requis).');
}

app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });

  const token = jwt.sign({ sub: user.id, role: user.role }, JWT_SECRET, { algorithm: 'HS256' });
  res.json({ token });
});

module.exports = app;
