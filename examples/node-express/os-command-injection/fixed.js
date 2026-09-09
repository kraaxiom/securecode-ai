// Correction : OS Command Injection (CWE-78)
// Le nom de fichier est validé selon une liste blanche stricte (caractères
// alphanumériques, point, tiret, underscore uniquement), puis execFile est
// utilisé avec des arguments passés en tableau distinct (aucun shell impliqué).
const express = require('express');
const path = require('path');
const { execFile } = require('child_process');
const router = express.Router();

const UPLOAD_DIR = '/uploads';

router.get('/files/convert', (req, res) => {
  const filename = req.query.filename;

  // Liste blanche stricte : rejette tout caractère hors nom de fichier simple.
  if (typeof filename !== 'string' || !/^[a-zA-Z0-9_.-]{1,255}$/.test(filename)) {
    return res.status(400).json({ error: 'Nom de fichier invalide' });
  }

  // Résolution du chemin et vérification qu'il reste dans le dossier autorisé.
  const inputPath = path.join(UPLOAD_DIR, filename);
  if (path.dirname(inputPath) !== UPLOAD_DIR) {
    return res.status(400).json({ error: 'Chemin invalide' });
  }
  const outputPath = `${inputPath}.png`;

  // Arguments passés en tableau distinct : aucun interpréteur shell impliqué.
  execFile('convert', [inputPath, outputPath], (err) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de conversion' });
    }
    res.json({ status: 'ok' });
  });
});

module.exports = router;
