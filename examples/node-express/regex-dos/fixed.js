// Correction : ReDoS - Regular Expression Denial of Service (CWE-1333)
// La regex est réécrite sans quantificateurs imbriqués (complexité linéaire
// garantie), et une limite de longueur est imposée sur l'entrée avant même
// d'appliquer l'expression régulière.
const express = require('express');
const router = express.Router();

// Regex non ambiguë : pas de groupe répété contenant lui-même un quantificateur.
const EMAIL_RE = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

function isValidEmail(input) {
  if (typeof input !== 'string' || input.length > 254) return false;
  return EMAIL_RE.test(input);
}

router.post('/newsletter/subscribe', (req, res) => {
  const email = req.body.email;

  if (!isValidEmail(email)) {
    return res.status(400).json({ error: 'Email invalide' });
  }
  res.json({ status: 'inscrit' });
});

module.exports = router;
