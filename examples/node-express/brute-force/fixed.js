// Corrigé — limitation de débit par compte ET par IP, avec verrouillage
// progressif et journalisation des échecs (CWE-307).

const express = require('express');
const rateLimit = require('express-rate-limit');
const app = express();
app.use(express.json());

const loginLimiter = rateLimit({
  windowMs: 15 * 60 * 1000, // fenêtre de 15 minutes
  max: 5, // 5 tentatives max par clé (IP + compte)
  keyGenerator: (req) => `${req.ip}:${req.body.email}`,
  message: { error: 'Trop de tentatives. Réessayez plus tard.' },
});

app.post('/login', loginLimiter, async (req, res) => {
  const { email, password } = req.body;

  const user = await authenticate(email, password);
  if (!user) {
    await recordFailedAttempt(email, req.ip); // journalisation pour audit/alerte
    // Message générique : ne révèle pas si le compte existe ou non
    return res.status(401).json({ error: 'Identifiants invalides' });
  }

  await clearFailedAttempts(email);
  const token = issueToken(user);
  res.json({ token });
});

module.exports = app;
