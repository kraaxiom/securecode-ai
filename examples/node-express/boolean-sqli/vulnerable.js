// Faille : Boolean-based SQL Injection (CWE-89)
// Le paramètre "name" est concaténé dans la clause WHERE. Un attaquant peut
// injecter une expression logique pour modifier le nombre de résultats
// retournés et en déduire des informations sur la base de données.
const express = require('express');
const mysql = require('mysql2');
const router = express.Router();

const connection = mysql.createConnection({ host: 'localhost', database: 'shop' });

router.get('/products/search', (req, res) => {
  const name = req.query.name;

  // Comparaison SQL construite par concaténation de chaîne.
  connection.query(
    `SELECT * FROM products WHERE name = '${name}'`,
    (err, rows) => {
      if (err) {
        return res.status(500).json({ error: 'Erreur interne' });
      }
      res.json(rows);
    }
  );
});

module.exports = router;
