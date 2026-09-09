// Correction : Server-Side Includes (SSI) Injection (CWE-97)
// Le contenu utilisateur est échappé en HTML puis toute séquence de
// directive SSI résiduelle (`<!--#`) est neutralisée avant écriture.
// Idéalement, ce fichier ne devrait plus être servi via SSI mais rendu
// par un moteur de templates applicatif.
const express = require('express');
const fs = require('fs');
const escapeHtml = require('escape-html');
const router = express.Router();

// Neutralise toute directive SSI résiduelle qui aurait survécu à l'échappement HTML.
function stripSsiDirectives(input) {
  return input.replace(/<!--#/g, '&lt;!--#');
}

router.post('/comments', (req, res) => {
  const comment = stripSsiDirectives(escapeHtml(req.body.comment || ''));

  fs.appendFileSync('public/comments.shtml', `<p>${comment}</p>\n`);
  res.json({ status: 'ajouté' });
});

module.exports = router;
