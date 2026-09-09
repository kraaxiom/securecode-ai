// Faille : HTTP Parameter Pollution (CWE-235)
// Si le paramètre "role" est envoyé plusieurs fois (?role=user&role=admin),
// Express le transforme en tableau. Le code l'utilise tel quel sans vérifier
// sa nature, ce qui peut induire un comportement inattendu selon la logique
// métier en aval (WAF laissant passer une valeur, backend en retenant une autre).
const express = require('express');
const router = express.Router();

router.post('/account/role', (req, res) => {
  const role = req.query.role;

  // Aucune vérification que "role" est bien une chaîne unique.
  assignRole(role);
  res.json({ status: 'ok', role });
});

function assignRole(role) {
  // Logique métier fictive d'attribution de rôle.
}

module.exports = router;
