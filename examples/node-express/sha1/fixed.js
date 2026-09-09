// Correction : Utilisation de SHA-1 (CWE-328)
// Remplacement de SHA-1 par Argon2id pour le hachage de mot de passe,
// avec sel automatique et coût mémoire/temps adaptatif.
const express = require('express');
const argon2 = require('argon2');
const router = express.Router();

router.post('/users/register', async (req, res) => {
  const { email, password } = req.body;

  const passwordHash = await argon2.hash(password, { type: argon2.argon2id });

  saveUser(email, passwordHash);
  res.status(201).json({ email });
});

function saveUser(email, passwordHash) {
  // Persistance simulée
}

module.exports = router;
