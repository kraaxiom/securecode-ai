// Faille : Error-based SQL Injection (CWE-89)
// La requête est construite par concaténation et, en cas d'erreur SQL, le
// message brut du driver est renvoyé au client. Un attaquant peut provoquer
// des erreurs contrôlées pour extraire des informations depuis le message.
const express = require('express');
const mysql = require('mysql2');
const router = express.Router();

const connection = mysql.createConnection({ host: 'localhost', database: 'shop' });

router.get('/orders/:id', (req, res) => {
  const id = req.params.id;

  connection.query(`SELECT * FROM orders WHERE id = ${id}`, (err, rows) => {
    if (err) {
      // Fuite du message d'erreur natif du driver SQL au client.
      return res.status(500).json({ error: err.message });
    }
    res.json(rows);
  });
});

module.exports = router;
