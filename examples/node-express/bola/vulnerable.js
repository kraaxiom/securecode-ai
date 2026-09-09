const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// VULNÉRABLE — Broken Object Level Authorization (BOLA)
// L'objet "order" est récupéré uniquement à partir de l'ID fourni par le
// client dans l'URL, sans jamais vérifier que cet order appartient bien à
// l'utilisateur authentifié. Un utilisateur peut donc lire (ou modifier)
// la commande d'un autre utilisateur en changeant simplement l'ID.
router.get('/api/orders/:id', requireAuth, async (req, res) => {
  const order = await db.findOne('orders', { _id: req.params.id });
  if (!order) return res.status(404).json({ error: 'Introuvable' });
  res.json(order);
});

router.put('/api/orders/:id', requireAuth, async (req, res) => {
  const order = await db.updateOne('orders', { _id: req.params.id }, req.body);
  res.json(order);
});

module.exports = router;
