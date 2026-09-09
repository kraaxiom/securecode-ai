// Faille : Indirect Prompt Injection (CWE-1427)
// Le contenu d'une page web récupérée est inséré tel quel dans le contexte du
// modèle, sans marquage de provenance ni délimitation, alors que des outils à
// fort impact restent actifs pendant l'analyse de ce contenu non fiable.
const express = require('express');
const router = express.Router();
const { sendEmail, executeCode } = require('../tools');

router.post('/agent/summarize-url', async (req, res) => {
  const { url } = req.body;
  const pageContent = await fetch(url).then((r) => r.text());

  // Contenu externe non fiable concaténé directement, outils actifs.
  const response = await llm.chat({
    messages: [{ role: 'user', content: `Résume ce contenu : ${pageContent}` }],
    tools: [sendEmail, executeCode],
  });

  res.json({ summary: response });
});

module.exports = router;
