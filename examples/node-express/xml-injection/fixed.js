// Correction : XML Injection (CWE-91)
// Le document est généré via xmlbuilder2, une bibliothèque de
// sérialisation XML qui échappe automatiquement le contenu des nœuds
// et des attributs, au lieu d'une concaténation de chaînes.
const express = require('express');
const { create } = require('xmlbuilder2');
const router = express.Router();

router.post('/export/user', (req, res) => {
  const name = req.body.name;

  // Sérialisation via une bibliothèque dédiée : échappement automatique.
  const xml = create({ user: { name, role: 'member' } }).end();

  res.type('application/xml').send(xml);
});

module.exports = router;
