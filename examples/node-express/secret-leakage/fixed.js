// Correction : Secret Leakage via LLM (CWE-200)
// Le secret n'apparaît jamais dans le prompt : il est résolu côté application
// uniquement lors de l'appel d'outil, et la réponse du modèle est filtrée pour
// détecter tout motif de secret avant d'être renvoyée au client.
const express = require('express');
const router = express.Router();
const { resolveSecret, secretScanner } = require('../security');

const SYSTEM_PROMPT =
  "Tu es un assistant interne. Pour appeler le service X, utilise l'outil callServiceX " +
  '(le secret est géré par l\'application, jamais exposé dans ce contexte).';

router.post('/chat', async (req, res) => {
  const response = await llm.chat({
    messages: [
      { role: 'system', content: SYSTEM_PROMPT },
      { role: 'user', content: req.body.message },
    ],
  });

  if (secretScanner.containsSecret(response.content)) {
    auditLog.record('blocked_secret_in_output', response.id);
    return res.json({ reply: '[réponse expurgée : secret détecté]' });
  }

  res.json({ reply: response.content });
});

async function callServiceX(args) {
  return httpClient.post(SERVICE_X_URL, args, {
    headers: { Authorization: `Bearer ${resolveSecret('service_x')}` },
  });
}

module.exports = router;
