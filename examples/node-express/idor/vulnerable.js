const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// VULNÉRABLE — Insecure Direct Object Reference (IDOR)
// Le document est récupéré uniquement à partir de l'ID transmis par le
// client dans l'URL. Aucune clause de filtrage sur le propriétaire n'est
// appliquée : la vérification d'autorisation se limite à "l'utilisateur
// est connecté", pas à "l'utilisateur possède ce document précis".
router.get('/api/documents/:id', requireAuth, async (req, res) => {
  const doc = await db.findOne('documents', { _id: req.params.id });
  if (!doc) return res.status(404).json({ error: 'Introuvable' });
  res.json(doc);
});

module.exports = router;
