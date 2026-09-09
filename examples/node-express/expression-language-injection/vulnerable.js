// Faille : Expression Language Injection (CWE-917)
// La source du template EJS est construite par concaténation d'une entrée
// utilisateur, ce qui permet d'injecter des balises EJS (<% %>) exécutant du
// code JavaScript arbitraire côté serveur lors du rendu.
const express = require('express');
const ejs = require('ejs');
const router = express.Router();

router.get('/greet', (req, res) => {
  const nom = req.query.nom;

  // Concaténation dans la source du template avant compilation/rendu.
  const tplSource = `<p>Bonjour ${nom}</p>`;
  const html = ejs.render(tplSource);

  res.send(html);
});

module.exports = router;
