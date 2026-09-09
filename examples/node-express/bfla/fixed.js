const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// CORRIGÉ — vérification de rôle explicite et centralisée (middleware
// réutilisable), appliquée à chaque fonction sensible, y compris celles
// non visibles dans l'interface utilisateur.
function requireRole(role) {
  return (req, res, next) => {
    if (req.user.role !== role) {
      return res.status(403).json({ error: 'Accès refusé' });
    }
    next();
  };
}

router.delete('/users/:id', requireAuth, requireRole('admin'), async (req, res) => {
  await db.deleteOne('users', { _id: req.params.id });
  res.sendStatus(204);
});

// Le même middleware protège la fonction interne de remboursement,
// qui exige elle aussi le rôle admin, indépendamment de sa découvrabilité.
router.post('/internal/refund', requireAuth, requireRole('admin'), async (req, res) => {
  const { orderId, amount } = req.body;
  await db.updateOne('orders', { _id: orderId }, { refunded: true, refundAmount: amount });
  res.json({ ok: true });
});

module.exports = router;
