// Corrigé — détection de vélocité globale par source + MFA proposée/imposée
// pour neutraliser l'usage de couples identifiant/mot de passe volés (CWE-307).

const express = require('express');
const app = express();
app.use(express.json());

const THRESHOLD = 30; // volume anormal de tentatives, tous comptes confondus

app.post('/login', async (req, res) => {
  const { email, password } = req.body;

  const velocity = await getFailedAttemptVelocity(req.ip); // toutes tentatives depuis cette IP
  if (velocity > THRESHOLD) {
    await flagSuspiciousSource(req.ip); // alerte vers la supervision sécurité
    return res.status(429).json({ error: 'Trafic anormal détecté, réessayez plus tard.' });
  }

  const user = await authenticate(email, password);
  if (!user) {
    await recordFailedAttempt(email, req.ip);
    return res.status(401).json({ error: 'Identifiants invalides' });
  }

  if (user.mfaEnabled) {
    return res.json({ mfaRequired: true, challengeId: await createMfaChallenge(user) });
  }

  res.json({ token: issueToken(user) });
});

module.exports = app;
