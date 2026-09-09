// Correction : XPath Injection (CWE-643)
// Les valeurs utilisateur sont transformées en littéraux XPath sûrs
// (fonction dédiée gérant les apostrophes via concat()) avant insertion
// dans l'expression. Idéalement, XPath ne devrait pas servir de
// mécanisme d'authentification : préférer une base avec hachage de mot
// de passe.
const express = require('express');
const xpath = require('xpath');
const { DOMParser } = require('@xmldom/xmldom');
const fs = require('fs');
const router = express.Router();

// Construit un littéral XPath sûr même si la valeur contient des apostrophes.
function xpathLiteral(value) {
  if (!value.includes("'")) return `'${value}'`;
  return "concat('" + value.split("'").join("', \"'\", '") + "')";
}

router.post('/auth/login', (req, res) => {
  const { user, pass } = req.body;

  if (typeof user !== 'string' || typeof pass !== 'string') {
    return res.status(400).json({ error: 'Format invalide' });
  }

  const doc = new DOMParser().parseFromString(fs.readFileSync('data/users.xml', 'utf8'));

  // Valeurs échappées en littéraux XPath : ne peuvent plus altérer la logique de sélection.
  const expr = `//user[username=${xpathLiteral(user)} and password=${xpathLiteral(pass)}]`;
  const nodes = xpath.select(expr, doc);

  if (nodes.length === 0) {
    return res.status(401).json({ error: 'Identifiants invalides' });
  }
  res.json({ status: 'authentifié' });
});

module.exports = router;
