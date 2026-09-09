// Faille : Secret JWT faible ou codé en dur (CWE-1391)
// Le secret utilisé pour signer les tokens JWT est une chaîne courte et
// codée en dur, ce qui permet un bruteforce hors ligne pour forger des
// tokens valides et usurper n'importe quel utilisateur.
const express = require('express');
const jwt = require('jsonwebtoken');
const router = express.Router();

// Secret court et prévisible, présent en clair dans le code source.
const JWT_SECRET = 'mysecret';

router.post('/auth/login', (req, res) => {
  const { userId } = req.body;

  const token = jwt.sign({ sub: userId }, JWT_SECRET);
  res.json({ token });
});

module.exports = router;
