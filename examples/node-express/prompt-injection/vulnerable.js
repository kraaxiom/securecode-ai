// Faille : Prompt Injection (CWE-1427)
// Le prompt système et l'entrée utilisateur sont concaténés en une seule
// chaîne de texte brute, sans séparation structurelle des rôles, et la sortie
// du modèle déclenche directement des appels d'outils sans validation.
const express = require('express');
const router = express.Router();
const { deleteFile, sendEmail } = require('../tools');

const SYSTEM_PROMPT = 'Tu es un assistant support. Réponds aux demandes des utilisateurs.';

router.post('/chat', async (req, res) => {
  const userInput = req.body.message;
  // entrée utilisateur non fiable insérée directement dans le prompt système
  // (aucun exemple de payload fourni volontairement)
  const prompt = SYSTEM_PROMPT + '\n' + userInput;

  const response = await llm.complete(prompt, { tools: [deleteFile, sendEmail] });
  for (const call of response.toolCalls) {
    await call.tool.execute(call.args); // exécution directe sans validation
  }

  res.json({ reply: response });
});

module.exports = router;
