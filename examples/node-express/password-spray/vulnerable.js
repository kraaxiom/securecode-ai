// Vulnérable — Password Spraying (CWE-307)
// La limitation des échecs d'authentification n'est appliquée que par
// compte. Un attaquant testant un même mot de passe courant contre de
// nombreux comptes distincts n'est jamais bloqué, quel que soit le volume.

const express = require('express');
const app = express();
app.use(express.json());

app.post('/login', async (req, res) => {
  const { email, password } = req.body;

  const attempts = await getFailedAttempts(email); // limité par compte uniquement
  if (attempts > 5) {
    return res.status(429).json({ error: 'Trop de tentatives' });
  }

  const user = await authenticate(email, password);
  if (!user) {
    await recordFailedAttempt(email);
    return res.status(401).json({ error: 'Identifiants invalides' });
  }

  res.json({ token: issueToken(user) });
});

module.exports = app;
