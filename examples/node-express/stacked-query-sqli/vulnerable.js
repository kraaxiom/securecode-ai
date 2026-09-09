// Faille : Stacked Query SQL Injection (CWE-89)
// La connexion mysql2 est configurée avec `multipleStatements: true` et
// le nom fourni par l'utilisateur est concaténé dans la requête. Un
// attaquant peut ajouter un point-virgule suivi d'une instruction SQL
// complètement distincte (ex: DROP, INSERT), avec un impact plus large
// qu'une injection classique limitée à un seul SELECT.
const express = require('express');
const mysql = require('mysql2');
const router = express.Router();

// Multi-instructions activé sans nécessité fonctionnelle : dangereux.
const connection = mysql.createConnection({
  host: 'localhost',
  user: 'app',
  database: 'shop',
  multipleStatements: true,
});

router.post('/users/:id/name', (req, res) => {
  const { name } = req.body;
  const id = req.params.id;

  // Concaténation directe : permet l'empilement d'une instruction supplémentaire.
  connection.query(`UPDATE users SET name = '${name}' WHERE id = ${id}`, (err) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de mise à jour' });
    }
    res.json({ status: 'ok' });
  });
});

module.exports = router;
