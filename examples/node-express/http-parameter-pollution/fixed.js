// Correction : HTTP Parameter Pollution (CWE-235)
// La présence d'un paramètre dupliqué est détectée explicitement et la
// requête est rejetée, éliminant toute ambiguïté d'interprétation.
const express = require('express');
const router = express.Router();

router.post('/account/role', (req, res) => {
  const roleRaw = req.query.role;

  // Rejet explicite si le paramètre "role" est dupliqué (tableau).
  if (Array.isArray(roleRaw)) {
    return res.status(400).json({ error: 'Paramètre dupliqué non autorisé' });
  }
  const role = String(roleRaw || '');

  assignRole(role);
  res.json({ status: 'ok', role });
});

function assignRole(role) {
  // Logique métier fictive d'attribution de rôle.
}

module.exports = router;
