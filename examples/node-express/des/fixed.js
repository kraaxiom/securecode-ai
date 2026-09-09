// Correction : Utilisation de DES / 3DES (CWE-327)
// Remplacement de DES-EDE3-CBC par AES-256-GCM (chiffrement authentifié),
// avec un IV/nonce unique généré aléatoirement à chaque chiffrement.
const express = require('express');
const crypto = require('crypto');
const router = express.Router();

const key = crypto.randomBytes(32); // clé AES-256, à charger depuis un secret manager en production

router.post('/wallet/encrypt', (req, res) => {
  const { data } = req.body;

  // Nonce unique par chiffrement : jamais réutilisé.
  const iv = crypto.randomBytes(12);
  const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);
  let encrypted = cipher.update(data, 'utf8', 'hex');
  encrypted += cipher.final('hex');
  const authTag = cipher.getAuthTag();

  res.json({
    encrypted,
    iv: iv.toString('hex'),
    authTag: authTag.toString('hex'),
  });
});

module.exports = router;
