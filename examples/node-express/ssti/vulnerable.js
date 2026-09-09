// Faille : Server-Side Template Injection - SSTI (CWE-1336)
// Le texte du template lui-même est construit par concaténation avec
// l'entrée utilisateur avant compilation par Handlebars, au lieu de
// passer cette valeur comme simple variable de contexte. Le moteur
// interprète la syntaxe injectée comme du code de template légitime.
const express = require('express');
const Handlebars = require('handlebars');
const router = express.Router();

router.get('/welcome', (req, res) => {
  const name = req.query.name;

  // La structure du template est elle-même construite avec l'entrée utilisateur.
  const source = `<p>Bonjour ${name}, bienvenue !</p>`;
  const template = Handlebars.compile(source);

  res.send(template({}));
});

module.exports = router;
