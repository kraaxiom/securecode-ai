// Correction : Utilisation de RC4 (CWE-327)
// Remplacement de RC4 par AES-256-GCM (chiffrement authentifié), avec un
// nonce unique généré aléatoirement à chaque chiffrement.
const express = require('express');
const crypto = require('crypto');
const router = express.Router();

const key = Buffer.from(process.env.NOTES_ENCRYPTION_KEY_HEX, 'hex'); // 32 octets, AES-256

router.post('/notes/encrypt', (req, res) => {
  const { content } = req.body;

  const iv = crypto.randomBytes(12);
  const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);
  let encrypted = cipher.update(content, 'utf8', 'hex');
  encrypted += cipher.final('hex');
  const authTag = cipher.getAuthTag();

  res.json({
    encrypted,
    iv: iv.toString('hex'),
    authTag: authTag.toString('hex'),
  });
});

module.exports = router;
