// Faille : CRLF Injection (CWE-93)
// La valeur "next" est écrite directement dans l'en-tête Location sans
// vérifier l'absence de retour chariot/saut de ligne, ce qui permet d'injecter
// des en-têtes supplémentaires ou de scinder la réponse HTTP.
const express = require('express');
const router = express.Router();

router.get('/redirect', (req, res) => {
  const next = req.query.next;

  // Écriture brute dans l'en-tête, sans validation ni passage par une API sûre.
  res.setHeader('Location', next);
  res.status(302).end();
});

module.exports = router;
