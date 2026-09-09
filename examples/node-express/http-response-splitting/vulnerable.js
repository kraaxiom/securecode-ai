// Faille : HTTP Response Splitting (CWE-113)
// La cible de redirection issue de l'utilisateur est écrite directement dans
// l'en-tête Location. Une valeur contenant des séquences CR/LF permettrait de
// scinder la réponse HTTP et d'injecter du contenu supplémentaire.
const express = require('express');
const router = express.Router();

router.get('/go', (req, res) => {
  const next = req.query.next;

  // Écriture brute dans l'en-tête sans validation.
  res.setHeader('Location', next);
  res.status(302).end();
});

module.exports = router;
