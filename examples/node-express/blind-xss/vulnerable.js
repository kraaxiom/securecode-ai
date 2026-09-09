// Vulnérable : XSS aveugle (Blind XSS) — CWE-79
// Le User-Agent envoyé par un visiteur non authentifié est journalisé puis
// réaffiché tel quel dans un dashboard interne réservé aux administrateurs.
// Aucune donnée externe n'est de confiance, y compris dans un outil interne :
// l'absence d'échappement ici permettrait à un contenu HTML injecté par un
// visiteur de s'exécuter dans le navigateur de l'administrateur qui consulte
// les logs, bien plus tard et sans lien direct avec la requête d'origine.

const express = require('express');
const router = express.Router();

const logEntries = [];

router.post('/contact', (req, res) => {
  logEntries.push({
    userAgent: req.headers['user-agent'],
    message: req.body.message,
    date: new Date(),
  });
  res.sendStatus(204);
});

// Vue interne (admin) : affichage brut, sans encodage contextuel
router.get('/admin/logs', (req, res) => {
  const rows = logEntries
    .map((entry) => `<div class="log-entry">${entry.userAgent} — ${entry.message}</div>`)
    .join('');
  res.send(`<html><body>${rows}</body></html>`);
});

module.exports = router;
