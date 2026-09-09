// Faille : ReDoS - Regular Expression Denial of Service (CWE-1333)
// La regex de validation d'email contient des quantificateurs imbriqués
// ((a+)+) sans limite de taille sur l'entrée. Une chaîne pathologique
// (ex: une longue suite de caractères sans "@" final) provoque un temps
// de calcul exponentiel qui bloque le thread Node.js.
const express = require('express');
const router = express.Router();

function isValidEmail(input) {
  // Quantificateurs imbriqués ambigus : complexité exponentielle possible.
  return /^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$/.test(input);
}

router.post('/newsletter/subscribe', (req, res) => {
  const email = req.body.email;

  if (!isValidEmail(email)) {
    return res.status(400).json({ error: 'Email invalide' });
  }
  res.json({ status: 'inscrit' });
});

module.exports = router;
