// Correction : Générateur aléatoire non cryptographique (CWE-338)
// Remplacement de Math.random() par crypto.randomBytes(), un CSPRNG
// offrant une entropie suffisante pour un token de sécurité.
const express = require('express');
const crypto = require('crypto');
const router = express.Router();

router.post('/auth/password-reset', (req, res) => {
  const { email } = req.body;

  // CSPRNG : au moins 128 bits d'entropie, imprévisible.
  const resetToken = crypto.randomBytes(32).toString('hex');

  storeResetToken(email, resetToken);
  res.json({ status: 'reset_email_sent' });
});

function storeResetToken(email, token) {
  // Persistance simulée
}

module.exports = router;
