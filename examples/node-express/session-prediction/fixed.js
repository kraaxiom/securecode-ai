// Corrigé — identifiant de session généré via un CSPRNG, entropie de 256
// bits, régénéré à chaque connexion, cookie sécurisé (CWE-330).

const express = require('express');
const session = require('express-session');
const crypto = require('crypto');
const app = express();
app.use(express.json());

app.use(session({
  genid: () => crypto.randomBytes(32).toString('hex'), // 256 bits d'entropie
  secret: process.env.SESSION_SECRET,
  resave: false,
  saveUninitialized: false,
  cookie: { httpOnly: true, secure: true, sameSite: 'strict' },
}));

app.post('/login', async (req, res) => {
  const user = await authenticate(req.body.email, req.body.password);
  if (!user) return res.status(401).json({ error: 'Identifiants invalides' });

  // Régénération de l'ID de session à chaque changement de niveau de privilège
  req.session.regenerate(() => {
    req.session.userId = user.id;
    res.json({ ok: true });
  });
});

module.exports = app;
