// Correction : Prompt Injection (CWE-1427)
// Séparation structurée des rôles system/user via l'API du modèle plutôt que
// la concaténation de chaîne, validation de schéma sur chaque appel d'outil,
// et moindre privilège avec confirmation humaine pour les actions sensibles.
const express = require('express');
const router = express.Router();
const { sendEmail } = require('../tools');

const SYSTEM_PROMPT = 'Tu es un assistant support. Réponds aux demandes des utilisateurs.';
const HIGH_IMPACT_TOOLS = new Set(['deleteFile']);

router.post('/chat', async (req, res) => {
  const userInput = req.body.message;

  // entrée utilisateur non fiable transmise via le rôle "user" structuré
  // (aucun exemple de payload fourni volontairement)
  const response = await llm.chat({
    messages: [
      { role: 'system', content: SYSTEM_PROMPT },
      { role: 'user', content: userInput },
    ],
    tools: [sendEmail], // outils à faible impact uniquement, par défaut
  });

  for (const call of response.toolCalls) {
    validateAgainstSchema(call.tool.name, call.args); // validation applicative indépendante
    if (HIGH_IMPACT_TOOLS.has(call.tool.name)) {
      await requireHumanConfirmation(call);
    }
    await call.tool.execute(call.args);
  }

  auditLog.record('chat_turn', userInput, response);
  res.json({ reply: response });
});

module.exports = router;
