// Correction : Blind SQL Injection (CWE-89)
// Utilisation d'une requête préparée avec paramètre lié : la valeur utilisateur
// n'est plus jamais interprétée comme du code SQL. Le message d'erreur est
// générique pour ne pas ouvrir de canal d'inférence supplémentaire.
const express = require('express');
const mysql = require('mysql2');
const router = express.Router();

const connection = mysql.createConnection({ host: 'localhost', database: 'app' });

router.get('/users/exists', (req, res) => {
  const user = req.query.user;

  // Requête préparée : le paramètre est lié, jamais concaténé.
  connection.query(
    'SELECT 1 FROM users WHERE username = ? AND active = 1',
    [user],
    (err, rows) => {
      if (err) {
        // Message générique pour éviter toute fuite d'information via l'erreur.
        return res.status(500).json({ error: 'Erreur interne' });
      }
      res.json({ exists: rows.length > 0 });
    }
  );
});

module.exports = router;
