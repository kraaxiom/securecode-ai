// Faille : Secret Leakage via LLM (CWE-200)
// Une clé API interne est incluse en clair dans le system prompt transmis au
// modèle, et la réponse générée est renvoyée à l'utilisateur sans aucun
// filtrage susceptible de bloquer une restitution accidentelle du secret.
const express = require('express');
const router = express.Router();

const INTERNAL_API_KEY = process.env.INTERNAL_API_KEY;

router.post('/chat', async (req, res) => {
  const systemPrompt = `Tu es un assistant interne. Clé API interne : ${INTERNAL_API_KEY}. Utilise-la pour appeler le service X.`;

  const response = await llm.chat({
    messages: [
      { role: 'system', content: systemPrompt },
      { role: 'user', content: req.body.message },
    ],
  });

  // Aucun filtrage de sortie avant renvoi de la réponse au client.
  res.json({ reply: response.content });
});

module.exports = router;
