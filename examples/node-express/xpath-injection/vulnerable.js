// Faille : XPath Injection (CWE-643)
// Les champs "user" et "pass" sont concaténés directement dans une
// expression XPath utilisée pour l'authentification contre un document
// XML. Un attaquant peut modifier la logique de sélection de nœuds et
// contourner l'authentification.
const express = require('express');
const xpath = require('xpath');
const { DOMParser } = require('@xmldom/xmldom');
const fs = require('fs');
const router = express.Router();

router.post('/auth/login', (req, res) => {
  const { user, pass } = req.body;
  const doc = new DOMParser().parseFromString(fs.readFileSync('data/users.xml', 'utf8'));

  // Concaténation directe dans l'expression XPath : dangereux.
  const expr = `//user[username='${user}' and password='${pass}']`;
  const nodes = xpath.select(expr, doc);

  if (nodes.length === 0) {
    return res.status(401).json({ error: 'Identifiants invalides' });
  }
  res.json({ status: 'authentifié' });
});

module.exports = router;
