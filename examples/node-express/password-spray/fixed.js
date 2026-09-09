// Corrigé — agrégation globale des échecs par IP/source, en complément de la
// limite par compte, pour détecter un spray distribué (CWE-307).

const express = require('express');
const app = express();
app.use(express.json());

app.post('/login', async (req, res) => {
  const { email, password } = req.body;

  const [accountAttempts, ipAttempts] = await Promise.all([
    getFailedAttempts(email),
    getFailedAttemptsByIp(req.ip),
  ]);

  if (accountAttempts > 5 || ipAttempts > 50) {
    if (ipAttempts > 50) await flagSuspiciousSource(req.ip, 'password_spray');
    return res.status(429).json({ error: 'Trop de tentatives.' });
  }

  const user = await authenticate(email, password);
  if (!user) {
    await recordFailedAttempt(email, req.ip);
    return res.status(401).json({ error: 'Identifiants invalides' });
  }

  res.json({ token: issueToken(user) });
});

module.exports = app;
