// Correction : Jailbreak de modèle (CWE-1427)
// Ajout d'une couche de modération indépendante du modèle principal en entrée
// et en sortie, et surveillance du nombre de tours suspects par session pour
// détecter une tentative d'érosion progressive des restrictions.
const express = require('express');
const router = express.Router();

const SYSTEM_PROMPT = 'Tu es un assistant interne. Respecte les règles de contenu.';
const REFUSAL_MESSAGE = { reply: "Je ne peux pas traiter cette demande." };
const MAX_SUSPICIOUS_TURNS = 3;

router.post('/chat', async (req, res) => {
  const { message, history, session } = req.body;

  if (await safetyClassifier.flags(message)) {
    auditLog.record('blocked_input_flagged', session.id);
    return res.json(REFUSAL_MESSAGE);
  }
  if (session.suspiciousTurnCount() > MAX_SUSPICIOUS_TURNS) {
    auditLog.record('session_flagged_repeated_attempts', session.id);
    return res.json(REFUSAL_MESSAGE);
  }

  // entrée utilisateur non fiable insérée dans le tour de conversation
  // (aucun exemple de payload fourni volontairement)
  const messages = [{ role: 'system', content: SYSTEM_PROMPT }, ...history, { role: 'user', content: message }];
  const response = await llm.chat({ messages });

  if (await safetyClassifier.flags(response.content)) {
    auditLog.record('blocked_output_flagged', session.id);
    return res.json(REFUSAL_MESSAGE);
  }

  res.json({ reply: response });
});

module.exports = router;
