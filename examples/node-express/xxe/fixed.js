// Correction : XML External Entity Injection - XXE (CWE-611)
// Le traitement des entités et le chargement de DTD externes sont
// explicitement désactivés sur le parseur avant toute analyse d'un
// document XML fourni par l'utilisateur.
const express = require('express');
const libxmljs = require('libxmljs2');
const router = express.Router();

router.post('/import/xml', express.text({ type: 'application/xml' }), (req, res) => {
  const userSuppliedXml = req.body;

  // DTD et entités externes explicitement désactivées : sûr par défaut.
  const doc = libxmljs.parseXml(userSuppliedXml, {
    noent: false,   // ne pas substituer les entités
    dtdload: false, // ne pas charger de DTD externe
    noblanks: true,
  });

  res.json({ root: doc.root().name() });
});

module.exports = router;
