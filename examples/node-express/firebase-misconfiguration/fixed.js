// Correction : Mauvaise configuration Firebase (CWE-284)
// Le serveur vérifie explicitement le token Firebase et l'identité de
// l'appelant avant d'accéder au document, en complément de règles
// Firestore restrictives (propriétaire uniquement) côté client.
const express = require('express');
const admin = require('firebase-admin');
const router = express.Router();

const db = admin.firestore();

router.get('/users/:uid/profile', async (req, res) => {
  const idToken = req.headers.authorization?.replace('Bearer ', '');
  if (!idToken) {
    return res.status(401).json({ error: 'Authentification requise' });
  }

  // Vérification du token et de la propriété de la ressource.
  const decoded = await admin.auth().verifyIdToken(idToken).catch(() => null);
  if (!decoded || decoded.uid !== req.params.uid) {
    return res.status(403).json({ error: 'Accès refusé' });
  }

  const doc = await db.collection('users').doc(req.params.uid).get();
  if (!doc.exists) {
    return res.status(404).json({ error: 'Profil introuvable' });
  }
  res.json(doc.data());
});

module.exports = router;
