// Correction : Configuration TLS faible (CWE-326)
// La vérification du certificat est réactivée et une version minimale de
// TLS moderne est imposée, y compris en environnement de test.
const express = require('express');
const https = require('https');
const router = express.Router();

const secureAgent = new https.Agent({ rejectUnauthorized: true, minVersion: 'TLSv1.2' });

router.get('/partners/status', (req, res) => {
  https.get('https://partner-api.example.com/status', { agent: secureAgent }, (upstream) => {
    let body = '';
    upstream.on('data', (chunk) => (body += chunk));
    upstream.on('end', () => res.type('application/json').send(body));
  });
});

module.exports = router;
