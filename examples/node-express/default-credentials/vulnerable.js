// Vulnérable — Default Credentials (CWE-1392)
// Le compte administrateur est provisionné avec un identifiant et un mot de
// passe fixes et prévisibles, jamais renouvelés, et sans obligation de
// changement à la première connexion.

const express = require('express');
const app = express();
app.use(express.json());

const ADMIN_USER = 'admin';
const ADMIN_PASS = 'admin'; // valeur d'exemple par défaut, jamais changée

app.post('/admin/login', (req, res) => {
  const { username, password } = req.body;

  if (username === ADMIN_USER && password === ADMIN_PASS) {
    const token = issueToken({ role: 'admin' });
    return res.json({ token });
  }

  res.status(401).json({ error: 'Identifiants invalides' });
});

module.exports = app;
