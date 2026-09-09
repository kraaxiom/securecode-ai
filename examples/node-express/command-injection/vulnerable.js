// Faille : Command Injection (CWE-78)
// Le paramètre "host" est concaténé dans une commande shell exécutée via exec().
// Un attaquant peut injecter des métacaractères shell pour exécuter des
// commandes arbitraires sur le serveur.
const express = require('express');
const { exec } = require('child_process');
const router = express.Router();

router.get('/network/ping', (req, res) => {
  const host = req.query.host;

  // Concaténation dans une commande shell : dangereux.
  exec(`ping -c 3 ${host}`, (err, stdout) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur lors du ping' });
    }
    res.type('text/plain').send(stdout);
  });
});

module.exports = router;
