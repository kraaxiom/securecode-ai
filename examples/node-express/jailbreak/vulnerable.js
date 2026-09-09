// Faille : Jailbreak de modèle (CWE-1427)
// Le system prompt est l'unique barrière de sécurité : aucune couche de
// modération indépendante ne filtre l'entrée ou la sortie, et rien ne limite
// le nombre de tours utilisés pour éroder progressivement les restrictions.
const express = require('express');
const router = express.Router();

const SYSTEM_PROMPT = 'Tu es un assistant interne. Respecte les règles de contenu.';

router.post('/chat', async (req, res) => {
  const { message, history } = req.body;
  // entrée utilisateur non fiable insérée directement dans le prompt système
  // (aucun exemple de payload fourni volontairement)
  const messages = [{ role: 'system', content: SYSTEM_PROMPT }, ...history, { role: 'user', content: message }];

  // Aucune modération d'entrée/sortie, aucune limite de tours suspects.
  const response = await llm.chat({ messages });
  res.json({ reply: response });
});

module.exports = router;
