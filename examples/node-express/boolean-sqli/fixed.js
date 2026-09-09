// Correction : Boolean-based SQL Injection (CWE-89)
// La comparaison utilise désormais une requête préparée avec paramètre lié.
// L'entrée utilisateur ne peut plus altérer la structure logique de la requête.
const express = require('express');
const mysql = require('mysql2');
const router = express.Router();

const connection = mysql.createConnection({ host: 'localhost', database: 'shop' });

router.get('/products/search', (req, res) => {
  const name = req.query.name;

  // Requête préparée : la valeur est liée, pas interprétée comme du SQL.
  connection.query(
    'SELECT * FROM products WHERE name = ?',
    [name],
    (err, rows) => {
      if (err) {
        return res.status(500).json({ error: 'Erreur interne' });
      }
      res.json(rows);
    }
  );
});

module.exports = router;
