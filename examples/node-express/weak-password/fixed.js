// Corrigé — longueur minimale de 12 caractères et vérification contre une
// liste de mots de passe compromis connus, sans règle de composition
// artificielle (conforme NIST 800-63B) (CWE-521).

const express = require('express');
const bcrypt = require('bcrypt');
const { isPasswordPwned } = require('hibp');
const app = express();
app.use(express.json());

app.post('/register', async (req, res) => {
  const { email, password } = req.body;

  if (password.length < 12) {
    return res.status(400).json({ error: 'Le mot de passe doit contenir au moins 12 caractères' });
  }
  if (await isPasswordPwned(password)) {
    return res.status(400).json({ error: 'Ce mot de passe a été trouvé dans une fuite de données connue' });
  }

  const hash = await bcrypt.hash(password, 12);
  await db.users.create({ email, passwordHash: hash });
  res.status(201).json({ ok: true });
});

module.exports = app;
