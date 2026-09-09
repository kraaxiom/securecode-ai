// Faille : Utilisation de RC4 (CWE-327)
// Les notes internes sont chiffrées avec RC4, un chiffrement par flux
// présentant des biais statistiques exploitables, interdit dans TLS
// depuis la RFC 7465. Il ne doit plus être utilisé.
const express = require('express');
const crypto = require('crypto');
const router = express.Router();

const key = Buffer.from(process.env.LEGACY_KEY_HEX || '0011223344556677', 'hex');

router.post('/notes/encrypt', (req, res) => {
  const { content } = req.body;

  // Chiffrement par flux RC4 : biais statistiques connus, cassé.
  const cipher = crypto.createCipheriv('rc4', key, '');
  let encrypted = cipher.update(content, 'utf8', 'hex');
  encrypted += cipher.final('hex');

  res.json({ encrypted });
});

module.exports = router;
