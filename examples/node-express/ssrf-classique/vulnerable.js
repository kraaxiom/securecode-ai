// Faille : SSRF classique (CWE-918)
// L'URL fournie par l'utilisateur est utilisée telle quelle pour effectuer
// une requête sortante depuis le serveur, sans whitelist ni validation de
// la destination (hôte, IP). Un attaquant peut ainsi forcer le serveur à
// interroger des ressources internes normalement inaccessibles depuis
// l'extérieur (services d'administration, réseau privé, etc.).
const express = require('express');
const router = express.Router();

router.post('/preview', async (req, res) => {
  const targetUrl = req.body.url;

  try {
    // Aucune whitelist, aucune résolution/validation d'IP : SSRF classique.
    const response = await fetch(targetUrl);
    const contentType = response.headers.get('content-type') || '';
    const body = await response.text();
    res.json({ contentType, preview: body.slice(0, 500) });
  } catch (err) {
    res.status(502).json({ error: 'Impossible de récupérer la ressource' });
  }
});

module.exports = router;
