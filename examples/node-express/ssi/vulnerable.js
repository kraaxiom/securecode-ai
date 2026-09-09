// Faille : Server-Side Includes (SSI) Injection (CWE-97)
// Le commentaire utilisateur est écrit tel quel dans un fichier .shtml
// servi par un serveur web avec SSI activé (ex: Apache mod_include).
// Un attaquant peut injecter une directive `<!--#exec cmd="..." -->`
// interprétée côté serveur, menant à une exécution de commande.
const express = require('express');
const fs = require('fs');
const router = express.Router();

router.post('/comments', (req, res) => {
  const comment = req.body.comment;

  // Écriture brute dans un fichier interprété par SSI : dangereux.
  fs.appendFileSync('public/comments.shtml', `<p>${comment}</p>\n`);
  res.json({ status: 'ajouté' });
});

module.exports = router;
