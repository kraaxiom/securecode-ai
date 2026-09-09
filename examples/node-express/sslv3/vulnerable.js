// Faille : Protocole SSLv3 activé (CWE-326)
// Le client HTTP interne force SSLv3 pour une intégration legacy, un
// protocole vulnérable à l'attaque POODLE (padding oracle CBC) interdit
// depuis la RFC 7568.
const express = require('express');
const https = require('https');
const router = express.Router();

router.get('/legacy/sync', (req, res) => {
  // secureProtocol force SSLv3 : vulnérable à POODLE.
  const request = https.request(
    {
      hostname: 'legacy-partner.example.com',
      path: '/sync',
      secureProtocol: 'SSLv3_method',
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
