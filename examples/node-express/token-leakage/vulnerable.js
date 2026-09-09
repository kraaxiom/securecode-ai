// Faille : Token Leakage (CWE-522)
// La clé API du fournisseur LLM est codée en dur et exposée dans un endpoint
// destiné à être appelé côté client, et les en-têtes d'autorisation sont
// journalisés en clair, exposant le jeton à toute personne lisant les logs.
const express = require('express');
const router = express.Router();
const { AnthropicClient } = require('../llm-client');

const client = new AnthropicClient({ apiKey: 'sk-ant-api03-XXXXXXXXXXXXXXXX' });

router.post('/api/assistant-config', (req, res) => {
  // La clé API est renvoyée au client pour qu'il appelle le fournisseur directement.
  console.log('Requête entrante', req.headers); // journalise l'en-tête Authorization en clair
  res.json({ apiKey: client.apiKey, model: 'claude-3' });
});

module.exports = router;
