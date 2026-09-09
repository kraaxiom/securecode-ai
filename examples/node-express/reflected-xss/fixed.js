// Corrigé : XSS réfléchi — CWE-79
// Le terme de recherche transite désormais par un moteur de templates (EJS)
// dont l'échappement automatique est conservé (<%= %>), et n'est jamais
// désactivé pour cette donnée issue de la requête. Toute balise HTML saisie
// par l'utilisateur est ainsi affichée comme texte inerte plutôt qu'exécutée.

const express = require('express');
const router = express.Router();

router.get('/search', (req, res) => {
  const term = req.query.q || '';
  // template results.ejs :
  // <h1>Résultats pour : <%= term %></h1>
  // <p>Aucun résultat trouvé.</p>
  res.render('results', { term });
});

module.exports = router;
