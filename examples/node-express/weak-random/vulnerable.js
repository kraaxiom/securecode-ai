// Faille : Générateur aléatoire non cryptographique (CWE-338)
// Le token de réinitialisation de mot de passe est généré avec Math.random(),
// un PRNG statistique prévisible, inadapté à la génération de secrets.
const express = require('express');
const router = express.Router();

router.post('/auth/password-reset', (req, res) => {
  const { email } = req.body;

  // Math.random() n'est pas cryptographiquement sûr : séquence prévisible.
  const resetToken = Math.random().toString(36).substring(2);

  storeResetToken(email, resetToken);
  res.json({ status: 'reset_email_sent' });
});

function storeResetToken(email, token) {
  // Persistance simulée
}

module.exports = router;
