// Correction : SQL Injection (CWE-89)
// Utilisation d'une requête préparée avec paramètre lié (placeholder `?`),
// le pilote se charge de l'échappement et de la séparation code/données.
const express = require('express');
const router = express.Router();
const connection = require('../db');

router.get('/users/search', (req, res) => {
  const name = req.query.name;

  // Paramètre lié : la valeur ne peut jamais altérer la structure SQL.
  connection.query('SELECT id, name, email FROM users WHERE name = ?', [name], (err, rows) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de recherche' });
    }
    res.json(rows);
  });
});

module.exports = router;
