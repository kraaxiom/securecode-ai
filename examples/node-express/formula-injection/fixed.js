// Correction : Formula Injection (CWE-1236)
// Chaque valeur est passée dans une fonction de neutralisation qui préfixe
// d'une apostrophe tout champ commençant par un caractère déclencheur de
// formule avant écriture dans le fichier d'export.
const express = require('express');
const fs = require('fs');
const router = express.Router();

function neutraliserFormule(valeur) {
  if (typeof valeur === 'string' && /^[=+\-@\t\r]/.test(valeur)) {
    return `'${valeur}`;
  }
  return valeur || '';
}

router.post('/tickets/export', (req, res) => {
  const { nom, commentaire } = req.body;

  const csvLine = `${neutraliserFormule(nom)},${neutraliserFormule(commentaire)}\n`;
  fs.appendFileSync('export.csv', csvLine);

  res.json({ status: 'exported' });
});

module.exports = router;
