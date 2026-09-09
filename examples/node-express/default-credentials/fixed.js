// Corrigé — mot de passe temporaire généré aléatoirement à la création du
// compte, avec obligation de changement avant tout accès (CWE-1392).

const express = require('express');
const crypto = require('crypto');
const argon2 = require('argon2');
const app = express();
app.use(express.json());

async function provisionAdminAccount() {
  const tempPassword = crypto.randomBytes(18).toString('base64url'); // aléatoire, unique
  const hash = await argon2.hash(tempPassword);

  await db.users.create({
    username: 'admin',
    passwordHash: hash,
    role: 'admin',
    mustChangePassword: true,
  });

  // Transmis via un canal sécurisé hors-bande, jamais journalisé en clair.
  return tempPassword;
}

app.post('/admin/login', async (req, res) => {
  const { username, password } = req.body;
  const admin = await db.users.findByUsername(username);

  const valid = admin && (await argon2.verify(admin.passwordHash, password));
  if (!valid) {
    return res.status(401).json({ error: 'Identifiants invalides' });
  }
  if (admin.mustChangePassword) {
    return res.status(403).json({ error: 'Changement de mot de passe requis' });
  }

  res.json({ token: issueToken({ role: admin.role }) });
});

module.exports = app;
