// Corrigé — algorithme de vérification explicitement figé côté serveur,
// exclusion de toute confusion HS256/RS256 (CWE-347).

const express = require('express');
const jwt = require('jsonwebtoken');
const fs = require('fs');
const app = express();

const publicKey = fs.readFileSync('public.pem');

app.get('/api/profile', (req, res) => {
  const token = req.headers.authorization?.split(' ')[1];

  try {
    // Liste fermée d'algorithmes : exclut explicitement HS256 et toute autre famille
    const payload = jwt.verify(token, publicKey, { algorithms: ['RS256'] });
    res.json({ userId: payload.sub, role: payload.role });
  } catch (err) {
    res.status(401).json({ error: 'Token invalide' });
  }
});

module.exports = app;
