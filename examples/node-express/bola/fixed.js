const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// CORRIGÉ — le filtrage d'appartenance (ownerId) est appliqué directement
// dans la requête de données, pas en post-traitement après coup : la
// commande d'un autre utilisateur ne peut jamais être renvoyée ni modifiée.
router.get('/api/orders/:id', requireAuth, async (req, res) => {
  const order = await db.findOne('orders', { _id: req.params.id, ownerId: req.user.id });
  if (!order) return res.status(404).json({ error: 'Introuvable' });
  res.json(order);
});

router.put('/api/orders/:id', requireAuth, async (req, res) => {
  const order = await db.updateOne(
    'orders',
    { _id: req.params.id, ownerId: req.user.id },
    req.body
  );
  if (!order) return res.status(404).json({ error: 'Introuvable' });
  res.json(order);
});

module.exports = router;
