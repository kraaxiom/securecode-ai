const express = require('express');
const { exec } = require('child_process');
const jwt = require('jsonwebtoken');
const router = express.Router();

// Fixture volontairement vulnérable — NE PAS DÉPLOYER.
// Utilisée pour valider le pipeline scan -> detect -> patch -> test du skill SecureCode AI.

// Vulnérabilité : Command Injection (voir knowledge/injections/command-injection.md).
// Le nom de fichier fourni par le client est concaténé directement dans une commande shell.
router.get('/convert', (req, res) => {
  const filename = req.query.file;
  exec(`convert ${filename} ${filename}.png`, (err, stdout) => {
    if (err) return res.status(500).send('Erreur de conversion');
    res.send(stdout);
  });
});

// Vulnérabilité : Broken Access Control (voir knowledge/authorization/broken-access-control.md).
// Aucune vérification que req.user.id correspond à l'ID demandé avant de renvoyer le profil.
router.get('/:id', (req, res) => {
  db.users.findById(req.params.id, (err, user) => {
    if (err) return res.status(500).send('Erreur');
    res.json(user);
  });
});

// Vulnérabilité : Weak JWT secret (voir knowledge/crypto/weak-jwt-secret.md).
// Secret court et codé en dur, trivialement brute-forçable hors ligne.
router.post('/login', (req, res) => {
  const token = jwt.sign({ userId: req.body.userId }, 'secret123');
  res.json({ token });
});

module.exports = router;
