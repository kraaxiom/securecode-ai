// Vulnérable : XSS réfléchi — CWE-79
// Le terme de recherche fourni dans la query string est directement
// concaténé dans la réponse HTML générée par le serveur, sans passer par un
// moteur de templates à échappement automatique ni par aucun encodage
// contextuel. La donnée n'est jamais persistée : elle n'est active que dans
// la réponse immédiate à la requête contenant le paramètre.

const express = require('express');
const router = express.Router();

router.get('/search', (req, res) => {
  const term = req.query.q || '';
  res.send(`
    <html>
      <body>
        <h1>Résultats pour : ${term}</h1>
        <p>Aucun résultat trouvé.</p>
      </body>
    </html>
  `);
});

module.exports = router;
