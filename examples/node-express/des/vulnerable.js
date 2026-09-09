// Faille : Utilisation de DES / 3DES (CWE-327)
// Le chiffrement des données sensibles utilise l'algorithme DES-EDE3-CBC,
// dont la taille de clé/bloc est trop faible (vulnérable à Sweet32 et au
// bruteforce). Il ne doit plus être utilisé pour protéger des données.
const express = require('express');
const crypto = require('crypto');
const router = express.Router();

const key = Buffer.from('0123456789abcdef01234567', 'utf8'); // 24 octets pour 3DES
const iv = Buffer.from('01234567', 'utf8');

router.post('/wallet/encrypt', (req, res) => {
  const { data } = req.body;

  // Chiffrement DES-EDE3 : algorithme obsolète, cassable en pratique.
  const cipher = crypto.createCipheriv('des-ede3-cbc', key, iv);
  let encrypted = cipher.update(data, 'utf8', 'hex');
  encrypted += cipher.final('hex');

  res.json({ encrypted });
});

module.exports = router;
