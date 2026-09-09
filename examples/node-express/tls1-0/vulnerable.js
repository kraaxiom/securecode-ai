// Faille : Protocole TLS 1.0 activé (CWE-326)
// L'agent HTTPS utilisé pour appeler un fournisseur tiers est figé sur
// TLSv1_method, un protocole obsolète vulnérable à BEAST et exclu de la
// conformité PCI-DSS depuis 2018.
const express = require('express');
const https = require('https');
const router = express.Router();

// Agent forçant TLS 1.0 : protocole déprécié (RFC 8996).
const legacyAgent = new https.Agent({ secureProtocol: 'TLSv1_method' });

router.get('/providers/rates', (req, res) => {
  https.get('https://rates.example.com/latest', { agent: legacyAgent }, (upstream) => {
    let body = '';
    upstream.on('data', (chunk) => (body += chunk));
    upstream.on('end', () => res.type('application/json').send(body));
  });
});

module.exports = router;
