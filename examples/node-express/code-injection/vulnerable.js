// Faille : Code Injection (CWE-94)
// La formule fournie par l'utilisateur est passée directement à eval(),
// ce qui permet d'exécuter du code JavaScript arbitraire dans le contexte
// du serveur.
const express = require('express');
const router = express.Router();

router.post('/calc/evaluate', (req, res) => {
  const formula = req.body.formula;

  // Évaluation dynamique d'une chaîne non fiable.
  let result;
  try {
    result = eval(formula);
  } catch (err) {
    return res.status(400).json({ error: 'Formule invalide' });
  }

  res.json({ result });
});

module.exports = router;
