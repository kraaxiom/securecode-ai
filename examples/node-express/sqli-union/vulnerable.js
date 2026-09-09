// Faille : UNION-based SQL Injection (CWE-89)
// Le paramètre "id" est concaténé directement dans la requête. Un
// attaquant peut ajouter une clause `UNION SELECT` pour extraire des
// données d'autres tables (ex: mots de passe stockés dans "users").
const express = require('express');
const router = express.Router();
const connection = require('../db');

router.get('/products/:id', (req, res) => {
  const id = req.params.id;

  // Concaténation directe : permet l'ajout d'une clause UNION SELECT.
  connection.query(`SELECT id, name, price FROM products WHERE id = ${id}`, (err, rows) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de récupération' });
    }
    res.json(rows);
  });
});

module.exports = router;
