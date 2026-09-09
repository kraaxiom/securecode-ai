// Corrigé : XSS via Markdown — CWE-79
// Le support du HTML brut est désactivé dans le moteur Markdown (html: false)
// pour tout contenu utilisateur non fiable, et le HTML final généré est
// systématiquement passé dans une bibliothèque de sanitisation dédiée
// (DOMPurify) avant d'être renvoyé, avec une liste blanche stricte de
// schémas d'URL autorisés pour les liens.

const express = require('express');
const router = express.Router();
const MarkdownIt = require('markdown-it');
const createDOMPurify = require('dompurify');
const { JSDOM } = require('jsdom');

const md = new MarkdownIt({ html: false, linkify: true });
const DOMPurify = createDOMPurify(new JSDOM('').window);

router.post('/tickets/:id/description', (req, res) => {
  const userMarkdown = req.body.description;
  const rendered = md.render(userMarkdown);
  const clean = DOMPurify.sanitize(rendered, {
    ALLOWED_URI_REGEXP: /^(?:https?|mailto):/i,
  });
  res.send(`<div class="ticket-description">${clean}</div>`);
});

module.exports = router;
