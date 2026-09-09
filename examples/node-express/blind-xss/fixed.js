// Corrigé : XSS aveugle (Blind XSS) — CWE-79
// Toute donnée provenant de l'extérieur du périmètre de confiance (ici le
// User-Agent et le message d'un visiteur non authentifié) est traitée comme
// non fiable, même quand elle n'est affichée que dans une interface interne.
// Le rendu passe par un moteur de templates avec échappement automatique
// (EJS <%= %>, jamais <%- %> sur donnée externe), ce qui neutralise tout
// contenu HTML injecté avant qu'il n'atteigne le navigateur de l'admin.

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

// Vue interne (admin) : rendu via template EJS avec échappement automatique
// template.ejs :
// <% logEntries.forEach(function(entry) { %>
//   <div class="log-entry"><%= entry.userAgent %> — <%= entry.message %></div>
// <% }); %>
router.get('/admin/logs', (req, res) => {
  res.render('logs', { logEntries });
});

module.exports = router;
