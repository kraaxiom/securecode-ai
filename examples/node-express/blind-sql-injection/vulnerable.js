// Faille : Blind SQL Injection (CWE-89)
// La valeur "user" est concaténée directement dans la requête SQL. Même si le
// résultat n'est pas affiché tel quel, le booléen "exists" dérivé de la requête
// permet à un attaquant d'inférer des informations en observant les réponses.
const express = require('express');
const mysql = require('mysql2');
const router = express.Router();

const connection = mysql.createConnection({ host: 'localhost', database: 'app' });

router.get('/users/exists', (req, res) => {
  const user = req.query.user;

  // Concaténation directe : aucune requête préparée, aucune validation.
  connection.query(
    `SELECT 1 FROM users WHERE username = '${user}' AND active = 1`,
    (err, rows) => {
      if (err) {
        return res.status(500).json({ error: 'Erreur interne' });
      }
      res.json({ exists: rows.length > 0 });
    }
  );
});

module.exports = router;
