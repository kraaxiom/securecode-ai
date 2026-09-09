// Faille : DNS Rebinding, contournement de filtre SSRF (CWE-918)
// L'hôte est validé une première fois (résolution DNS + vérification d'IP
// publique), puis la requête HTTP effective est déléguée au client `fetch`,
// qui résout le DNS une seconde fois. Entre les deux résolutions, l'attaquant
// peut faire pointer le domaine vers une IP interne (TOCTOU).
const express = require('express');
const dns = require('dns').promises;
const router = express.Router();

router.post('/fetch-avatar', async (req, res) => {
  const avatarUrl = req.body.url;
  const { hostname, protocol } = new URL(avatarUrl);

  // Validation ponctuelle : la résolution utilisée ici n'est pas celle
  // réellement utilisée par fetch() au moment de la connexion TCP.
  const { address } = await dns.lookup(hostname);
  if (protocol !== 'https:' || address.startsWith('10.') || address.startsWith('127.')) {
    return res.status(400).json({ error: 'Destination interdite' });
  }

  // Deuxième résolution DNS ici, potentiellement différente (rebinding).
  const response = await fetch(avatarUrl);
  const buffer = Buffer.from(await response.arrayBuffer());
  res.set('Content-Type', response.headers.get('content-type') || 'application/octet-stream');
  res.send(buffer);
});

module.exports = router;
