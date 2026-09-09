// Corrigé — état intermédiaire à privilèges limités tant que le second
// facteur n'a pas été validé côté serveur (CWE-287).

const express = require('express');
const app = express();
app.use(express.json());

app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });

  if (user.mfaEnabled) {
    // Token restreint, scope insuffisant pour accéder aux endpoints protégés
    const partialToken = issuePartialToken(user, { scope: 'mfa_pending' });
    return res.json({ mfaRequired: true, partialToken });
  }

  res.json({ token: issueFullToken(user) });
});

app.post('/mfa/verify', requirePartialToken, async (req, res) => {
  const valid = await verifyMfaCode(req.user, req.body.code);
  if (!valid) return res.status(401).json({ error: 'Code invalide' });

  // Token pleinement privilégié uniquement après validation serveur du second facteur
  res.json({ token: issueFullToken(req.user) });
});

module.exports = app;
