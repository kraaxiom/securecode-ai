// Correction : Protocole SSLv3 activé (CWE-326)
// Remplacement de SSLv3 par une plage de versions modernes (TLS 1.2
// minimum, TLS 1.3 recommandé), non vulnérable à POODLE.
const express = require('express');
const https = require('https');
const router = express.Router();

router.get('/legacy/sync', (req, res) => {
  const request = https.request(
    {
      hostname: 'legacy-partner.example.com',
      path: '/sync',
      minVersion: 'TLSv1.2',
      maxVersion: 'TLSv1.3',
    },
    (upstream) => {
      let body = '';
      upstream.on('data', (chunk) => (body += chunk));
      upstream.on('end', () => res.send(body));
    }
  );
  request.end();
});

module.exports = router;
