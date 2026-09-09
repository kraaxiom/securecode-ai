// Vulnérable — Politique de mot de passe faible (CWE-521)
// Seule une longueur minimale insuffisante est contrôlée, sans vérification
// contre les mots de passe déjà compromis. N'importe quel mot de passe
// trivial est accepté dès lors qu'il fait 6 caractères.

const express = require('express');
const bcrypt = require('bcrypt');
const app = express();
app.use(express.json());

app.post('/register', async (req, res) => {
  const { email, password } = req.body;

  if (password.length < 6) {
    return res.status(400).json({ error: 'Mot de passe trop court' });
  }

  const hash = await bcrypt.hash(password, 10);
  await db.users.create({ email, passwordHash: hash });
  res.status(201).json({ ok: true });
});

module.exports = app;
