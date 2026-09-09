// Faille : Protocole SSLv2 activé (CWE-326)
// Le serveur HTTPS interne autorise explicitement SSLv2, un protocole
// totalement cassé (attaque DROWN, handshake non authentifié) et interdit
// depuis la RFC 6176.
const express = require('express');
const https = require('https');
const fs = require('fs');
const router = express.Router();

function createLegacyServer(app) {
  // secureProtocol force SSLv2 : à ne jamais autoriser.
  return https.createServer(
    {
      key: fs.readFileSync('server-key.pem'),
      cert: fs.readFileSync('server-cert.pem'),
      secureProtocol: 'SSLv2_method',
    },
    app
  );
}

router.get('/tls/status', (req, res) => {
  res.json({ protocol: 'SSLv2_method' });
});

module.exports = router;
module.exports.createLegacyServer = createLegacyServer;
