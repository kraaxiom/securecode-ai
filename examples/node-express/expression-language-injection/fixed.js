// Correction : Expression Language Injection (CWE-917)
// Le template est une chaîne statique définie par le développeur ; la donnée
// utilisateur est transmise en tant que variable de rendu et échappée
// automatiquement par le moteur EJS (<%= %>).
const express = require('express');
const ejs = require('ejs');
const router = express.Router();

router.get('/greet', (req, res) => {
  const nom = req.query.nom;

  // Template statique : aucune donnée utilisateur dans la source évaluée.
  const tplSource = '<p>Bonjour <%= nom %></p>';
  const html = ejs.render(tplSource, { nom });

  res.send(html);
});

module.exports = router;
