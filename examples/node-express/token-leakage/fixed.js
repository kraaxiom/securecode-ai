// Correction : Token Leakage (CWE-522)
// La clé API du fournisseur LLM reste exclusivement côté serveur, résolue
// depuis un gestionnaire de secrets, les appels sont proxifiés derrière une
// authentification, et les en-têtes sensibles sont masqués dans les journaux.
const express = require('express');
const router = express.Router();
const { AnthropicClient } = require('../llm-client');
const { requireAuth, redactHeaders } = require('../security');

const client = new AnthropicClient({ apiKey: process.env.LLM_API_KEY }); // via secret manager

router.use(redactHeaders(['authorization', 'x-api-key']));

router.post('/api/assistant', requireAuth, async (req, res) => {
  const response = await client.messages.create({
    model: 'claude-3',
    messages: [{ role: 'user', content: req.body.question }],
  });
  res.json(response); // aucun secret transmis au client
});

module.exports = router;
