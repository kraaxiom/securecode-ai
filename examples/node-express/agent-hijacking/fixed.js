// Correction : Agent Hijacking (CWE-1427)
// Séparation de la couche de décision (LLM) et de la couche d'exécution :
// les outils à fort impact exigent une confirmation humaine explicite et une
// validation métier indépendante avant exécution, avec journalisation complète.
const express = require('express');
const router = express.Router();
const { sendPayment, deleteAccount, sendEmail } = require('../tools');

const HIGH_IMPACT_TOOLS = new Set(['sendPayment', 'deleteAccount']);

router.post('/agent/run', async (req, res) => {
  const userInput = req.body.input;

  const plan = await llm.plan(userInput, {
    tools: [sendPayment, deleteAccount, sendEmail],
  });

  for (const step of plan.steps) {
    if (HIGH_IMPACT_TOOLS.has(step.tool.name)) {
      const confirmed = await requireHumanConfirmation(step);
      if (!confirmed) {
        auditLog.record('blocked_high_impact_action', step);
        continue;
      }
    }
    validateAgainstBusinessRules(step); // contrôle indépendant du modèle
    await step.tool.execute(step.args);
    auditLog.record('action_executed', step);
  }

  res.json({ status: 'done' });
});

module.exports = router;
