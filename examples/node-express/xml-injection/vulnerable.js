// Faille : XML Injection (CWE-91)
// Le document XML est construit par concaténation de chaînes incluant
// le nom fourni par l'utilisateur, sans échappement des caractères
// spéciaux (`<`, `>`, `&`). Un attaquant peut altérer la structure du
// document (ajout de nœuds, falsification de données).
const express = require('express');
const router = express.Router();

router.post('/export/user', (req, res) => {
  const name = req.body.name;

  // Concaténation directe, sans échappement : dangereux.
  const xml = `<user><name>${name}</name><role>member</role></user>`;

  res.type('application/xml').send(xml);
});

module.exports = router;
