// Faille : Agent Hijacking (CWE-1427)
// L'agent exécute directement tout appel d'outil décidé par le LLM, y compris
// les outils à fort impact (paiement, suppression), sans validation métier
// indépendante ni confirmation humaine. Un contenu externe non fiable traité
// par l'agent peut ainsi déclencher des actions irréversibles.
const express = require('express');
const router = express.Router();
const { sendPayment, deleteAccount, sendEmail } = require('../tools');

router.post('/agent/run', async (req, res) => {
  const userInput = req.body.input;

  const plan = await llm.plan(userInput, {
    tools: [sendPayment, deleteAccount, sendEmail],
  });

  for (const step of plan.steps) {
    // Exécution directe, aucune validation ni confirmation, quel que soit l'impact.
    await step.tool.execute(step.args);
  }

  res.json({ status: 'done' });
});

module.exports = router;
