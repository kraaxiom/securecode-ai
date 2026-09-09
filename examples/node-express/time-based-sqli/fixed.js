// Correction : Time-Based Blind SQL Injection (CWE-89)
// La requête préparée avec paramètre lié élimine toute la classe de
// vulnérabilité (pas seulement la variante "time-based"). Un timeout
// d'exécution est en complément configuré au niveau du pool de
// connexion pour limiter l'impact d'une éventuelle injection résiduelle.
const express = require('express');
const router = express.Router();
const pool = require('../db'); // pool mysql2 configuré avec un timeout d'exécution

router.get('/orders/:id', (req, res) => {
  const id = Number(req.params.id);

  if (!Number.isInteger(id) || id < 0) {
    return res.status(400).json({ error: 'Identifiant invalide' });
  }

  // Paramètre lié : la valeur ne peut jamais altérer la structure SQL,
  // qu'elle contienne ou non une fonction de pause conditionnelle.
  pool.query('SELECT id, status FROM orders WHERE id = ?', [id], (err, rows) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur serveur' });
    }
    res.json(rows[0] || null);
  });
});

module.exports = router;
