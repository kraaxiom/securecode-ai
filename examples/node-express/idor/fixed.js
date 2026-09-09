const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// CORRIGÉ — clause de vérification d'appartenance directement dans la
// requête de récupération : seul le document appartenant à l'utilisateur
// authentifié peut être renvoyé, sans post-traitement après coup.
router.get('/api/documents/:id', requireAuth, async (req, res) => {
  const doc = await db.findOne('documents', { _id: req.params.id, ownerId: req.user.id });
  if (!doc) return res.status(404).json({ error: 'Introuvable' });
  res.json(doc);
});

module.exports = router;
