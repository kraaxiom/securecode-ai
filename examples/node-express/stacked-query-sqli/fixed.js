// Correction : Stacked Query SQL Injection (CWE-89)
// L'option multi-instructions du driver est désactivée (non requise
// fonctionnellement), et la requête utilise des paramètres liés :
// un point-virgule dans l'entrée ne peut plus démarrer une seconde
// instruction SQL.
const express = require('express');
const mysql = require('mysql2');
const router = express.Router();

// Multi-instructions désactivé par défaut : le driver n'exécute qu'une seule requête par appel.
const connection = mysql.createConnection({
  host: 'localhost',
  user: 'app',
  database: 'shop',
  multipleStatements: false,
});

router.post('/users/:id/name', (req, res) => {
  const { name } = req.body;
  const id = Number(req.params.id);

  if (!Number.isInteger(id) || id < 0) {
    return res.status(400).json({ error: 'Identifiant invalide' });
  }

  connection.query('UPDATE users SET name = ? WHERE id = ?', [name, id], (err) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de mise à jour' });
    }
    res.json({ status: 'ok' });
  });
});

module.exports = router;
