// Faille : SQL Injection (CWE-89)
// Le paramètre "name" est interpolé directement dans la requête SQL via
// un template literal. Un attaquant peut injecter un caractère `'` ou
// un mot-clé SQL pour altérer la logique de la requête.
const express = require('express');
const router = express.Router();
const connection = require('../db'); // connexion mysql2

router.get('/users/search', (req, res) => {
  const name = req.query.name;

  // Interpolation directe dans la requête SQL : dangereux.
  connection.query(`SELECT id, name, email FROM users WHERE name = '${name}'`, (err, rows) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de recherche' });
    }
    res.json(rows);
  });
});

module.exports = router;
