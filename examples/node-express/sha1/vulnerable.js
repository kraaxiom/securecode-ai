// Faille : Utilisation de SHA-1 (CWE-328)
// Le mot de passe est haché avec SHA-1, une fonction cryptographiquement
// cassée pour la résistance aux collisions (attaque SHAttered). Elle ne
// doit plus être utilisée pour du stockage de mot de passe.
const express = require('express');
const crypto = require('crypto');
const router = express.Router();

router.post('/users/register', (req, res) => {
  const { email, password } = req.body;

  // Hachage SHA-1 sans sel : cassé pour un usage sécuritaire.
  const passwordHash = crypto.createHash('sha1').update(password).digest('hex');

  saveUser(email, passwordHash);
  res.status(201).json({ email });
});

function saveUser(email, passwordHash) {
  // Persistance simulée
}

module.exports = router;
