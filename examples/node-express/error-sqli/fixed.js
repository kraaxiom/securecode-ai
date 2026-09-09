// Correction : Error-based SQL Injection (CWE-89)
// Requête préparée avec paramètre lié et castage explicite en nombre. En cas
// d'erreur, seul un message générique est renvoyé au client ; le détail est
// journalisé côté serveur.
const express = require('express');
const mysql = require('mysql2');
const router = express.Router();

const connection = mysql.createConnection({ host: 'localhost', database: 'shop' });

router.get('/orders/:id', (req, res) => {
  const id = Number(req.params.id);
  if (!Number.isInteger(id)) {
    return res.status(400).json({ error: 'Identifiant invalide' });
  }

  connection.query('SELECT * FROM orders WHERE id = ?', [id], (err, rows) => {
    if (err) {
      // Le détail de l'erreur est journalisé, jamais renvoyé au client.
      console.error(err);
      return res.status(500).json({ error: 'Une erreur est survenue.' });
    }
    res.json(rows);
  });
});

module.exports = router;
