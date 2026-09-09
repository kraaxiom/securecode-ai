// Correction : Server-Side Template Injection - SSTI (CWE-1336)
// Le template est une chaîne statique contrôlée par le développeur ;
// l'entrée utilisateur est passée uniquement comme variable de contexte,
// jamais comme structure de template.
const express = require('express');
const Handlebars = require('handlebars');
const router = express.Router();

// Template statique, compilé une seule fois, sans jamais intégrer d'entrée utilisateur.
const template = Handlebars.compile('<p>Bonjour {{name}}, bienvenue !</p>');

router.get('/welcome', (req, res) => {
  const name = req.query.name;

  // La valeur utilisateur transite uniquement par le contexte de variables.
  res.send(template({ name }));
});

module.exports = router;
