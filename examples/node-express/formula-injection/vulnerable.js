// Faille : Formula Injection (CWE-1236)
// Le commentaire utilisateur est ajouté tel quel dans un fichier d'export
// CSV. Une valeur commençant par un déclencheur de formule sera exécutée par
// le tableur à l'ouverture du fichier.
const express = require('express');
const fs = require('fs');
const router = express.Router();

router.post('/tickets/export', (req, res) => {
  const { nom, commentaire } = req.body;

  // Aucune neutralisation avant écriture dans le fichier d'export.
  const csvLine = `${nom},${commentaire}\n`;
  fs.appendFileSync('export.csv', csvLine);

  res.json({ status: 'exported' });
});

module.exports = router;
