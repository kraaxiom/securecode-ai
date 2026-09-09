// Faille : Utilisation de MD5 (CWE-328)
// Les mots de passe utilisateurs sont hachés avec MD5, une fonction cassée
// depuis 2004, sans sel ni coût adaptatif. Un attaquant disposant d'une
// fuite de la base peut retrouver les mots de passe en quelques secondes.
const express = require('express');
const crypto = require('crypto');
const router = express.Router();

router.post('/users/register', (req, res) => {
  const { email, password } = req.body;

  // Hachage MD5 sans sel : cassable par rainbow table ou bruteforce GPU.
  const passwordHash = crypto.createHash('md5').update(password).digest('hex');

  saveUser(email, passwordHash);
  res.status(201).json({ email });
});

function saveUser(email, passwordHash) {
  // Persistance simulée
}

module.exports = router;
