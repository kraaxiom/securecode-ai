// Correction : Tool Injection (CWE-1427)
// Chaque appel d'outil est validé contre un schéma typé strict, une liste
// blanche contextuelle limite les outils disponibles selon le niveau de
// confiance du contenu, et chaque exécution est journalisée pour l'audit.
const express = require('express');
const router = express.Router();
const { toolRegistry, toolSchemas, allowedToolsByTrust } = require('../tools');

router.post('/agent/tool-call', async (req, res) => {
  const { call, contextTrustLevel } = req.body;

  if (!allowedToolsByTrust[contextTrustLevel].has(call.name)) {
    auditLog.record('blocked_tool_not_in_allowlist', call.name, contextTrustLevel);
    return res.status(403).json({ error: 'Outil non autorisé dans ce contexte' });
  }

  const validatedArgs = toolSchemas[call.name].validate(call.args); // rejette tout argument hors schéma
  const result = await toolRegistry[call.name].execute(validatedArgs);
  auditLog.record('tool_executed', call.name, validatedArgs);

  res.json({ result });
});

module.exports = router;
