// Vulnérable — Credential Stuffing (CWE-307)
// Aucune corrélation n'est faite entre les tentatives échouées sur des
// comptes différents provenant d'une même origine. Un attaquant peut donc
// tester en masse des couples identifiant/mot de passe issus de fuites
// externes sans être détecté ni bloqué.

const express = require('express');
const app = express();
app.use(express.json());

app.post('/login', async (req, res) => {
  const { email, password } = req.body;

  const user = await authenticate(email, password); // pas de détection de vélocité
  if (!user) {
    return res.status(401).json({ error: 'Identifiants invalides' });
  }

  const token = issueToken(user);
  res.json({ token });
});

module.exports = app;
