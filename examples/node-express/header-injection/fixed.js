// Correction : Header Injection (CWE-113)
// Les caractères de contrôle sont supprimés puis le nom de fichier est
// validé par une liste blanche stricte de caractères avant d'être inséré
// dans l'en-tête.
const express = require('express');
const path = require('path');
const router = express.Router();

router.get('/download', (req, res) => {
  let filename = String(req.query.filename || '').replace(/[\r\n]/g, '');
  filename = path.basename(filename);

  // Liste blanche stricte de caractères autorisés dans le nom de fichier.
  if (!/^[\w.-]+$/.test(filename)) {
    return res.status(400).send('Nom de fichier invalide');
  }

  res.setHeader('Content-Disposition', `attachment; filename="${filename}"`);
  res.send('contenu du fichier');
});

module.exports = router;
