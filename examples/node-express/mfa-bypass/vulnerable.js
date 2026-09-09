// Vulnérable — MFA Bypass (CWE-287)
// Un token pleinement privilégié est émis dès la vérification du mot de
// passe, avant toute validation du second facteur. La vérification MFA
// n'est donc qu'une étape facultative qui n'empêche rien.

const express = require('express');
const app = express();
app.use(express.json());

app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });

  // Token complet émis immédiatement, MFA activée ou non !
  const token = issueFullToken(user);
  res.json({ token, mfaEnabled: user.mfaEnabled });
});

module.exports = app;
