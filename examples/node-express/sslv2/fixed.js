// Correction : Protocole SSLv2 activé (CWE-326)
// Remplacement du protocole figé SSLv2 par une plage de versions modernes
// (TLS 1.2 minimum, TLS 1.3 recommandé).
const express = require('express');
const https = require('https');
const fs = require('fs');
const router = express.Router();

function createServer(app) {
  return https.createServer(
    {
      key: fs.readFileSync('server-key.pem'),
      cert: fs.readFileSync('server-cert.pem'),
      minVersion: 'TLSv1.2',
      maxVersion: 'TLSv1.3',
    },
    app
  );
}

router.get('/tls/status', (req, res) => {
  res.json({ minVersion: 'TLSv1.2', maxVersion: 'TLSv1.3' });
});

module.exports = router;
module.exports.createServer = createServer;
