// Correction : UNION-based SQL Injection (CWE-89)
// L'identifiant est casté en entier et la requête utilise un paramètre
// lié : plus aucune chaîne construite dynamiquement ne peut recevoir
// une clause UNION SELECT.
const express = require('express');
const router = express.Router();
const connection = require('../db');

router.get('/products/:id', (req, res) => {
  const id = Number(req.params.id);

  // Validation de type stricte : rejette toute valeur non numérique.
  if (!Number.isInteger(id) || id < 0) {
    return res.status(400).json({ error: 'Identifiant invalide' });
  }

  connection.query('SELECT id, name, price FROM products WHERE id = ?', [id], (err, rows) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de récupération' });
    }
    res.json(rows);
  });
});

module.exports = router;
