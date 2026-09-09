// Vulnérable : XSS via SVG — CWE-79
// Un fichier SVG uploadé par l'utilisateur comme avatar est stocké et servi
// tel quel, sans aucune sanitisation de son contenu XML, depuis la même
// origine que l'application. Or le format SVG peut contenir des éléments
// actifs (balises et gestionnaires d'événements), qui s'exécuteraient alors
// dans le contexte d'origine de l'application lors de l'affichage.

const express = require('express');
const router = express.Router();
const multer = require('multer');
const fs = require('fs');
const path = require('path');

const upload = multer({ dest: 'uploads/' });

router.post('/avatar', upload.single('avatar'), (req, res) => {
  const dest = path.join('uploads', req.file.filename + '.svg');
  fs.copyFileSync(req.file.path, dest);
  res.sendStatus(204);
});

// Fichiers servis directement, sans sanitisation ni isolation d'origine
router.use('/uploads', express.static('uploads'));

module.exports = router;
