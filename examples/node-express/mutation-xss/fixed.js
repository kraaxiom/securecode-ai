// Corrigé : XSS par mutation (mXSS) — CWE-79
// La sanitisation maison est remplacée par une bibliothèque activement
// maintenue et connue pour couvrir les vecteurs de mutation liés au parsing
// du navigateur (DOMPurify). Les allers-retours entre chaîne et DOM pour du
// contenu non fiable sont réduits au minimum, et une nouvelle passe de
// sanitisation est appliquée à chaque nouvelle sérialisation du contenu.

const express = require('express');
const router = express.Router();
const createDOMPurify = require('dompurify');
const { JSDOM } = require('jsdom');

const DOMPurify = createDOMPurify(new JSDOM('').window);
const notes = {};

router.post('/notes/:id', (req, res) => {
  // Sanitisation avant stockage avec une bibliothèque maintenue
  notes[req.params.id] = DOMPurify.sanitize(req.body.content);
  res.sendStatus(204);
});

router.get('/notes/:id', (req, res) => {
  // Re-sanitisation avant réaffichage, en défense en profondeur
  const clean = DOMPurify.sanitize(notes[req.params.id] || '');
  res.send(`<div class="note-content">${clean}</div>`);
});

module.exports = router;
