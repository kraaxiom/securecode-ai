// Vulnérable : XSS via Markdown — CWE-79
// Le moteur de rendu Markdown est configuré avec le support du HTML brut
// activé (option html: true) sur du contenu directement fourni par
// l'utilisateur (description d'un ticket), puis le HTML produit est renvoyé
// tel quel sans aucune sanitisation. Un utilisateur peut ainsi glisser du
// balisage HTML actif à l'intérieur de son Markdown pour qu'il soit exécuté
// chez les autres lecteurs du ticket.

const express = require('express');
const router = express.Router();
const MarkdownIt = require('markdown-it');

const md = new MarkdownIt({ html: true });

router.post('/tickets/:id/description', (req, res) => {
  const userMarkdown = req.body.description;
  const renderedHtml = md.render(userMarkdown);
  res.send(`<div class="ticket-description">${renderedHtml}</div>`);
});

module.exports = router;
