// Faille : XML External Entity Injection - XXE (CWE-611)
// Le document XML fourni par l'utilisateur est analysé avec la
// configuration par défaut du parseur, qui autorise le traitement des
// DTD et des entités externes. Un attaquant peut définir une entité
// pointant vers un fichier local ou une URL, menant à une divulgation
// de fichiers sensibles ou à du SSRF.
const express = require('express');
const libxmljs = require('libxmljs2');
const router = express.Router();

router.post('/import/xml', express.text({ type: 'application/xml' }), (req, res) => {
  const userSuppliedXml = req.body;

  // Configuration par défaut : DTD et entités externes traitées.
  const doc = libxmljs.parseXml(userSuppliedXml);

  res.json({ root: doc.root().name() });
});

module.exports = router;
