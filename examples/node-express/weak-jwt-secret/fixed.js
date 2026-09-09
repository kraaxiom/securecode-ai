// Correction : Secret JWT faible ou codé en dur (CWE-1391)
// Le secret provient d'une variable d'environnement à haute entropie
// (>= 256 bits), l'algorithme est épinglé explicitement, et une expiration
// courte limite la fenêtre d'exploitation d'un token compromis.
const express = require('express');
const jwt = require('jsonwebtoken');
const router = express.Router();

const JWT_SECRET = process.env.JWT_SECRET; // généré via crypto.randomBytes(32), stocké en secret manager
if (!JWT_SECRET) {
  throw new Error('JWT_SECRET manquant dans l\'environnement');
}

router.post('/auth/login', (req, res) => {
  const { userId } = req.body;

  const token = jwt.sign({ sub: userId }, JWT_SECRET, {
    algorithm: 'HS256',
    expiresIn: '15m',
  });
  res.json({ token });
});

module.exports = router;
