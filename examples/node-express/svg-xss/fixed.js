// Corrigé : XSS via SVG — CWE-79
// Le contenu SVG est sanitisé (suppression des éléments actifs et des
// gestionnaires d'événements) avant stockage, à l'aide d'une bibliothèque
// dédiée. Le fichier est ensuite servi en téléchargement forcé plutôt qu'en
// affichage inline, ce qui limite encore l'impact d'un éventuel contournement.

const express = require('express');
const router = express.Router();
const multer = require('multer');
const fs = require('fs');
const path = require('path');
const sanitizeSvg = require('svg-sanitizer'); // bibliothèque dédiée à la sanitisation SVG

const upload = multer({ dest: 'uploads/' });

router.post('/avatar', upload.single('avatar'), (req, res) => {
  const rawSvg = fs.readFileSync(req.file.path, 'utf8');
  const clean = sanitizeSvg(rawSvg);
  const dest = path.join('uploads', req.file.filename + '.svg');
  fs.writeFileSync(dest, clean);
  res.sendStatus(204);
});

// Fichier servi en téléchargement forcé plutôt qu'en affichage inline
router.get('/uploads/:filename', (req, res) => {
  res.setHeader('Content-Type', 'image/svg+xml');
  res.setHeader('Content-Disposition', `attachment; filename="${path.basename(req.params.filename)}"`);
  res.sendFile(path.resolve('uploads', req.params.filename));
});

module.exports = router;
