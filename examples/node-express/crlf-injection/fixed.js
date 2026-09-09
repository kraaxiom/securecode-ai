// Correction : CRLF Injection (CWE-93)
// La cible de redirection est validée par une expression régulière stricte
// (chemin relatif, sans \r ni \n) puis passée à res.redirect(), qui encode
// correctement l'en-tête.
const express = require('express');
const router = express.Router();

router.get('/redirect', (req, res) => {
  const next = req.query.next;

  // Liste blanche : uniquement un chemin relatif sans caractère de contrôle.
  const safe = typeof next === 'string' && /^\/[^\r\n]*$/.test(next) ? next : '/';

  res.redirect(302, safe);
});

module.exports = router;
