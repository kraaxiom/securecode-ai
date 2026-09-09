// Correction : Code Injection (CWE-94)
// Suppression totale de eval(). Les opérations autorisées sont déclarées dans
// une liste blanche de fonctions ; seul le nom de l'opération est fourni par
// le client, jamais du code exécutable.
const express = require('express');
const router = express.Router();

const ALLOWED_OPERATIONS = {
  add: (a, b) => a + b,
  sub: (a, b) => a - b,
  mul: (a, b) => a * b,
  div: (a, b) => (b !== 0 ? a / b : NaN),
};

router.post('/calc/evaluate', (req, res) => {
  const { op, a, b } = req.body;

  // Liste blanche stricte : seules les opérations déclarées sont exécutables.
  if (!Object.prototype.hasOwnProperty.call(ALLOWED_OPERATIONS, op)) {
    return res.status(400).json({ error: 'Opération non autorisée' });
  }
  if (typeof a !== 'number' || typeof b !== 'number') {
    return res.status(400).json({ error: 'Opérandes invalides' });
  }

  const result = ALLOWED_OPERATIONS[op](a, b);
  res.json({ result });
});

module.exports = router;
