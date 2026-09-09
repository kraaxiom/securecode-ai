// Vulnérable : DOM-based XSS — CWE-79
// Le serveur injecte une valeur de requête directement dans un bloc <script>
// inline via concaténation de chaînes. Cette valeur, entièrement contrôlée
// par l'attaquant via le paramètre d'URL, se retrouve interprétée comme du
// code JavaScript actif côté client au lieu d'être traitée comme une simple
// donnée textuelle.

const express = require('express');
const router = express.Router();

router.get('/widget', (req, res) => {
  const q = req.query.q || '';
  res.send(`
    <html>
      <body>
        <div id="app"></div>
        <script>
          const searchTerm = "${q}";
          document.getElementById('app').textContent = 'Recherche : ' + searchTerm;
        </script>
      </body>
    </html>
  `);
});

module.exports = router;
