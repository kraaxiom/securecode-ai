// Faille : Configuration TLS faible (CWE-326)
// L'appel sortant désactive la vérification du certificat du serveur
// distant, ce qui rend le canal vulnérable à une interception (MITM)
// même si TLS est nominalement utilisé.
const express = require('express');
const https = require('https');
const router = express.Router();

// rejectUnauthorized: false désactive toute validation de certificat.
const insecureAgent = new https.Agent({ rejectUnauthorized: false });

router.get('/partners/status', (req, res) => {
  https.get('https://partner-api.example.com/status', { agent: insecureAgent }, (upstream) => {
    let body = '';
    upstream.on('data', (chunk) => (body += chunk));
    upstream.on('end', () => res.type('application/json').send(body));
  });
});

module.exports = router;
