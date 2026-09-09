// Faille : Time-Based Blind SQL Injection (CWE-89)
// L'identifiant est concaténé dans la requête et les erreurs de base de
// données sont masquées (réponse générique), sans timeout d'exécution
// configuré. Un attaquant peut injecter une condition provoquant une
// pause (ex: SLEEP) et inférer des informations en observant le délai
// de réponse, sans qu'aucune donnée ni erreur ne soit exposée.
const express = require('express');
const router = express.Router();
const pool = require('../db'); // pool mysql2

router.get('/orders/:id', (req, res) => {
  const id = req.params.id;

  // Concaténation directe, aucun timeout, erreurs masquées : dangereux.
  pool.query(`SELECT id, status FROM orders WHERE id = ${id}`, (err, rows) => {
    if (err) {
      // Réponse générique : ne divulgue pas l'erreur, mais ne bloque pas non plus l'injection.
      return res.status(500).json({ error: 'Erreur serveur' });
    }
    res.json(rows[0] || null);
  });
});

module.exports = router;
