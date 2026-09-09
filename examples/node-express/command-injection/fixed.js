// Correction : Command Injection (CWE-78)
// Utilisation de execFile avec des arguments passés en tableau (pas de shell),
// et validation stricte que l'entrée est bien une adresse IP avant exécution.
const express = require('express');
const { execFile } = require('child_process');
const net = require('net');
const router = express.Router();

router.get('/network/ping', (req, res) => {
  const host = req.query.host;

  // Validation stricte : seule une IP valide est acceptée.
  if (!net.isIP(host)) {
    return res.status(400).json({ error: 'Hôte invalide' });
  }

  // Arguments passés en tableau distinct : aucun interpréteur shell impliqué.
  execFile('ping', ['-c', '3', host], (err, stdout) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur lors du ping' });
    }
    res.type('text/plain').send(stdout);
  });
});

module.exports = router;
