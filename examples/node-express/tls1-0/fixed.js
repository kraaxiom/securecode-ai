// Correction : Protocole TLS 1.0 activé (CWE-326)
// L'agent HTTPS accepte désormais uniquement TLS 1.2 et TLS 1.3, conformes
// aux exigences PCI-DSS et à la RFC 8996 (dépréciation de TLS 1.0/1.1).
const express = require('express');
const https = require('https');
const router = express.Router();

const modernAgent = new https.Agent({ minVersion: 'TLSv1.2', maxVersion: 'TLSv1.3' });

router.get('/providers/rates', (req, res) => {
  https.get('https://rates.example.com/latest', { agent: modernAgent }, (upstream) => {
    let body = '';
    upstream.on('data', (chunk) => (body += chunk));
    upstream.on('end', () => res.type('application/json').send(body));
  });
});

module.exports = router;
