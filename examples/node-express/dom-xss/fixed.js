// Corrigé : DOM-based XSS — CWE-79
// La valeur de requête n'est plus concaténée directement dans un bloc
// <script> inline. Elle est transmise via un attribut data-* encodé, puis
// lue côté client par du JavaScript qui l'assigne à textContent — jamais à
// un sink HTML dangereux. Aucune donnée contrôlable par l'attaquant n'est
// jamais réinterprétée comme du code exécutable.

const express = require('express');
const router = express.Router();

router.get('/widget', (req, res) => {
  const q = req.query.q || '';
  const escaped = encodeURIComponent(q);
  res.send(`
    <html>
      <body>
        <div id="app" data-query="${escaped}"></div>
        <script>
          const container = document.getElementById('app');
          const searchTerm = decodeURIComponent(container.dataset.query);
          container.textContent = 'Recherche : ' + searchTerm;
        </script>
      </body>
    </html>
  `);
});

module.exports = router;
