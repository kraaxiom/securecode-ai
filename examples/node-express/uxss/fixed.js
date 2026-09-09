// Corrigé : Universal XSS (UXSS) — CWE-79
// Une Content Security Policy stricte restreint les origines de script et de
// frame autorisées (helmet), le script tiers est chargé avec Subresource
// Integrity (integrity + crossorigin) pour garantir qu'il n'a pas été altéré,
// et l'iframe tierce est isolée avec un attribut sandbox limité au strict
// nécessaire. Une politique Permissions-Policy limite en plus les capacités
// accordées aux composants embarqués.

const express = require('express');
const router = express.Router();
const helmet = require('helmet');

router.use(helmet.contentSecurityPolicy({
  directives: {
    defaultSrc: ["'self'"],
    scriptSrc: ["'self'", 'https://cdn.example.com'],
    frameSrc: ["'self'", 'https://widget-tiers.example.com'],
  },
}));

router.use((req, res, next) => {
  res.setHeader('Permissions-Policy', 'camera=(), microphone=(), geolocation=()');
  next();
});

router.get('/dashboard', (req, res) => {
  res.send(`
    <html>
      <body>
        <div id="app">Tableau de bord</div>
        <script
          src="https://cdn.example.com/widget.js"
          integrity="sha384-EXEMPLE-HASH-A-REMPLACER"
          crossorigin="anonymous">
        </script>
        <iframe
          src="https://widget-tiers.example.com/chat"
          sandbox="allow-scripts allow-same-origin"
          referrerpolicy="no-referrer">
        </iframe>
      </body>
    </html>
  `);
});

module.exports = router;
