// Vulnérable — Brute Force / absence de limitation des tentatives (CWE-307)
// L'endpoint de connexion n'applique aucune limitation de débit ni verrouillage
// de compte. Un attaquant peut essayer un nombre illimité de mots de passe
// contre un même compte sans être ralenti ni bloqué.

const express = require('express');
const app = express();
app.use(express.json());

app.post('/login', async (req, res) => {
  const { email, password } = req.body;

  const user = await authenticate(email, password); // aucune limite d'essais
  if (!user) {
    return res.status(401).json({ error: 'Identifiants invalides' });
  }

  const token = issueToken(user);
  res.json({ token });
});

module.exports = app;
