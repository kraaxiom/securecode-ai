// Faille : Tool Injection (CWE-1427)
// Les arguments d'appel d'outil générés par le modèle sont transmis
// directement à l'exécution sans validation de schéma, et tous les outils
// restent disponibles quel que soit le niveau de confiance du contenu traité.
const express = require('express');
const router = express.Router();
const { toolRegistry } = require('../tools');

router.post('/agent/tool-call', async (req, res) => {
  const call = req.body; // { name, args } généré par le modèle

  const tool = toolRegistry[call.name];
  // Aucune validation de schéma, aucune liste blanche contextuelle.
  const result = await tool.execute(call.args);

  res.json({ result });
});

module.exports = router;
