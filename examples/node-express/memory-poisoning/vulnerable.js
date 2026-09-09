// Faille : Memory Poisoning (CWE-349)
// Les "faits" extraits automatiquement d'une conversation sont écrits en
// mémoire persistante sans confirmation de l'utilisateur, puis réinjectés tels
// quels comme contexte de confiance dans les sessions futures.
const express = require('express');
const router = express.Router();
const { memoryStore } = require('../agent');

router.post('/agent/turn', async (req, res) => {
  const { userId, message, agentResponse } = req.body;

  const facts = await llm.extractFacts(agentResponse);
  // Écriture automatique, aucune confirmation ni cloisonnement vérifié.
  await memoryStore.append(userId, facts);

  res.json({ status: 'ok' });
});

router.get('/agent/session/:userId', async (req, res) => {
  const memory = await memoryStore.get(req.params.userId);
  const response = await llm.chat({
    messages: [{ role: 'system', content: `Contexte connu : ${JSON.stringify(memory)}` }],
  });
  res.json({ reply: response });
});

module.exports = router;
