// Faille : Mauvaise configuration Firebase (CWE-284)
// L'application suppose que les règles Firestore sont en mode test
// (allow read, write: if true), donc elle lit/écrit directement les
// documents utilisateur sans vérifier l'identité appelante côté serveur.
const express = require('express');
const admin = require('firebase-admin');
const router = express.Router();

const db = admin.firestore();

router.get('/users/:uid/profile', async (req, res) => {
  // Aucune vérification que l'appelant est bien l'utilisateur "uid"
  // ou un administrateur : repose entièrement sur des règles Firestore
  // permissives ("if true") côté client.
  const doc = await db.collection('users').doc(req.params.uid).get();

  if (!doc.exists) {
    return res.status(404).json({ error: 'Profil introuvable' });
  }
  res.json(doc.data());
});

module.exports = router;
