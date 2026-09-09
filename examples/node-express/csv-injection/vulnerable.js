// Faille : CSV Injection (CWE-1236)
// Les champs utilisateur "name" et "comment" sont écrits tels quels dans le
// fichier CSV exporté. Si une valeur commence par "=", "+", "-" ou "@", le
// tableur qui ouvrira le fichier peut l'interpréter comme une formule active.
const express = require('express');
const router = express.Router();

router.get('/export/customers', (req, res) => {
  const customers = [
    { name: req.query.name, comment: req.query.comment },
  ];

  // Génération manuelle du CSV, sans neutralisation des cellules.
  const csv = customers.map(c => `${c.name},${c.comment}`).join('\n');

  res.set('Content-Type', 'text/csv');
  res.send(csv);
});

module.exports = router;
