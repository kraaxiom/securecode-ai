// Vulnérable : XSS par mutation (mXSS) — CWE-79
// Le contenu d'un éditeur riche est "sanitisé" par un aller-retour naïf via
// une div temporaire (string -> DOM -> string), puis stocké et réaffiché.
// Ce type de sanitisation maison ne tient pas compte des quirks de parsing
// du navigateur : le HTML peut être réécrit ("muté") lors de sa réinsertion
// finale dans le DOM, faisant réapparaître une structure active après coup.

const express = require('express');
const router = express.Router();

const notes = {};

function sanitizeNaive(html) {
  // Aller-retour string -> DOM -> string sans bibliothèque dédiée
  const { JSDOM } = require('jsdom');
  const dom = new JSDOM('<div></div>');
  const div = dom.window.document.querySelector('div');
  div.innerHTML = html;
  return div.innerHTML;
}

router.post('/notes/:id', (req, res) => {
  notes[req.params.id] = sanitizeNaive(req.body.content);
  res.sendStatus(204);
});

router.get('/notes/:id', (req, res) => {
  res.send(`<div class="note-content">${notes[req.params.id] || ''}</div>`);
});

module.exports = router;
