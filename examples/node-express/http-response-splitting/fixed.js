// Correction : HTTP Response Splitting (CWE-113)
// La cible de redirection est restreinte à une liste blanche de chemins
// connus, puis passée à res.redirect() qui encode l'en-tête correctement.
const express = require('express');
const router = express.Router();

const ALLOWED_PATHS = new Set(['/dashboard', '/profile', '/account']);

router.get('/go', (req, res) => {
  // Liste blanche stricte : aucune valeur arbitraire n'est acceptée.
  const next = ALLOWED_PATHS.has(req.query.next) ? req.query.next : '/dashboard';
  res.redirect(302, next);
});

module.exports = router;
