// Vulnérable : Universal XSS (UXSS) — CWE-79
// L'application charge un script tiers depuis un CDN externe sans vérifier
// son intégrité (pas d'attribut integrity/crossorigin), et embarque un
// widget tiers dans une iframe sans attribut sandbox restrictif, sans aucune
// Content Security Policy pour limiter les origines de script autorisées.
// La surface d'exposition à une faille du composant tiers ou du navigateur
// n'est donc pas contenue par l'application.

const express = require('express');
const router = express.Router();

router.get('/dashboard', (req, res) => {
  res.send(`
    <html>
      <body>
        <div id="app">Tableau de bord</div>
        <script src="https://cdn.example.com/widget.js"></script>
        <iframe src="https://widget-tiers.example.com/chat"></iframe>
      </body>
    </html>
  `);
});

module.exports = router;
