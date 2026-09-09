// Vulnérable — Session Prediction (CWE-330)
// L'identifiant de session est dérivé de valeurs prévisibles (ID utilisateur
// et timestamp) via un compteur et un hash faible, plutôt que généré par un
// générateur cryptographiquement sûr.

const express = require('express');
const app = express();
app.use(express.json());

let counter = 0;
function generateSessionId(userId) {
  return `sess_${userId}_${Date.now()}_${counter++}`; // prévisible
}

const sessions = new Map();

app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });

  const sessionId = generateSessionId(user.id);
  sessions.set(sessionId, { userId: user.id });
  res.cookie('sid', sessionId);
  res.json({ ok: true });
});

module.exports = app;
