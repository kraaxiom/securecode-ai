// Correction : Utilisation de MD5 (CWE-328)
// Remplacement de MD5 par Argon2id, un algorithme de hachage de mot de passe
// dédié avec sel automatique et coût mémoire/temps adaptatif.
const express = require('express');
const argon2 = require('argon2');
const router = express.Router();

router.post('/users/register', async (req, res) => {
  const { email, password } = req.body;

  // Argon2id : sel intégré, résistant au bruteforce et aux attaques GPU/ASIC.
  const passwordHash = await argon2.hash(password, { type: argon2.argon2id });

  saveUser(email, passwordHash);
  res.status(201).json({ email });
});

function saveUser(email, passwordHash) {
  // Persistance simulée
}

module.exports = router;
