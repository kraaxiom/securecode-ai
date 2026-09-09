// Vulnérable — JWT alg:none (CWE-347)
// jwt.decode() ne vérifie AUCUNE signature. Un attaquant peut donc fabriquer
// un token avec un header alg:none et une signature vide, en contrôlant
// intégralement les claims (y compris le rôle), sans connaître aucun secret.

const express = require('express');
const jwt = require('jsonwebtoken');
const app = express();

app.get('/api/admin', (req, res) => {
  const token = req.headers.authorization?.split(' ')[1];

  const payload = jwt.decode(token); // AUCUNE vérification cryptographique !
  if (!payload || payload.role !== 'admin') {
    return res.status(403).json({ error: 'Accès refusé' });
  }

  res.json({ secret: 'données administrateur' });
});

module.exports = app;
