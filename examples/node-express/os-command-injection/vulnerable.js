// Faille : OS Command Injection (CWE-78)
// Le nom de fichier fourni par l'utilisateur est concaténé dans une
// commande shell exécutée via exec(). Un attaquant peut injecter des
// méta-caractères shell (`;`, `|`, `&&`, backticks) pour exécuter des
// commandes arbitraires sur le serveur.
const express = require('express');
const { exec } = require('child_process');
const router = express.Router();

router.get('/files/convert', (req, res) => {
  const filename = req.query.filename;

  // Concaténation dans une commande shell : dangereux.
  exec(`convert /uploads/${filename} /uploads/${filename}.png`, (err, stdout) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de conversion' });
    }
    res.json({ status: 'ok' });
  });
});

module.exports = router;
