// Corrigé — vérification de signature active, liste fermée d'algorithmes
// excluant structurellement 'none' (CWE-347).

const express = require('express');
const jwt = require('jsonwebtoken');
const app = express();

const JWT_SECRET = process.env.JWT_SECRET;

app.get('/api/admin', (req, res) => {
  const token = req.headers.authorization?.split(' ')[1];

  try {
    // jwt.verify() vérifie la signature ; 'none' ne peut pas figurer dans la liste
    const payload = jwt.verify(token, JWT_SECRET, { algorithms: ['HS256'] });
    if (payload.role !== 'admin') {
      return res.status(403).json({ error: 'Accès refusé' });
    }
    res.json({ secret: 'données administrateur' });
  } catch (err) {
    res.status(401).json({ error: 'Token invalide' });
  }
});

module.exports = app;
