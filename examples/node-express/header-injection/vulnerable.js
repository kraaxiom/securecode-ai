// Faille : Header Injection (CWE-113)
// Le nom de fichier fourni par l'utilisateur est inséré tel quel dans l'en-tête
// Content-Disposition, sans suppression des caractères de contrôle ni liste
// blanche de caractères autorisés.
const express = require('express');
const router = express.Router();

router.get('/download', (req, res) => {
  const filename = req.query.filename;

  // Concaténation directe dans l'en-tête, sans validation.
  res.setHeader('Content-Disposition', `attachment; filename=${filename}`);
  res.send('contenu du fichier');
});

module.exports = router;
