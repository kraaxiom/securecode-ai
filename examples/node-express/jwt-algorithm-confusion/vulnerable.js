// Vulnérable — JWT Algorithm Confusion (CWE-347)
// jwt.verify() est appelé sans restreindre la liste des algorithmes acceptés.
// Un attaquant peut alors forger un token signé en HS256 en réutilisant la
// clé publique RSA de l'application comme secret HMAC, trompant la vérification.

const express = require('express');
const jwt = require('jsonwebtoken');
const fs = require('fs');
const app = express();

const publicKey = fs.readFileSync('public.pem'); // exposée pour la vérification RS256

app.get('/api/profile', (req, res) => {
  const token = req.headers.authorization?.split(' ')[1];

  try {
    // Algorithme non restreint : déduit implicitement du header du token
    const payload = jwt.verify(token, publicKey);
    res.json({ userId: payload.sub, role: payload.role });
  } catch (err) {
    res.status(401).json({ error: 'Token invalide' });
  }
});

module.exports = app;
