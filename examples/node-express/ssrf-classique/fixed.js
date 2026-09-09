// Fix : SSRF classique (CWE-918)
// La destination est restreinte à une whitelist explicite de domaines
// métier, le DNS est résolu puis l'IP obtenue est validée pour exclure les
// plages privées/loopback/link-local avant toute requête. Les redirections
// automatiques sont désactivées et un timeout est appliqué.
const express = require('express');
const dns = require('dns').promises;
const { isIP } = require('net');
const ipaddr = require('ipaddr.js');
const router = express.Router();

const ALLOWED_HOSTS = ['media.partenaire.example.com'];

function isPrivateOrReserved(ip) {
  if (!isIP(ip)) return true;
  const addr = ipaddr.parse(ip);
  const range = addr.range();
  return range !== 'unicast'; // exclut private/loopback/linkLocal/reserved/etc.
}

router.post('/preview', async (req, res) => {
  let parsed;
  try {
    parsed = new URL(req.body.url);
  } catch {
    return res.status(400).json({ error: 'URL invalide' });
  }

  if (parsed.protocol !== 'https:' || !ALLOWED_HOSTS.includes(parsed.hostname)) {
    return res.status(400).json({ error: 'URL non autorisée' });
  }

  const { address } = await dns.lookup(parsed.hostname);
  if (isPrivateOrReserved(address)) {
    return res.status(400).json({ error: 'Destination interdite' });
  }

  try {
    const response = await fetch(parsed.toString(), {
      redirect: 'manual',
      signal: AbortSignal.timeout(5000),
    });
    const body = await response.text();
    res.json({ preview: body.slice(0, 500) });
  } catch {
    res.status(502).json({ error: 'Impossible de récupérer la ressource' });
  }
});

module.exports = router;
