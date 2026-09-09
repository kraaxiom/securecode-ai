const express = require('express');
const router = express.Router();
const db = require('../db'); // objet d'accès aux données simulé (db.find, db.deleteOne, etc.)
const { requireAuth } = require('../middleware/auth');

// VULNÉRABLE — Broken Function Level Authorization (BFLA)
// Cette route de suppression d'utilisateur ne vérifie que l'authentification
// (requireAuth), jamais le rôle de l'appelant. N'importe quel utilisateur
// connecté peut donc appeler cette fonction d'administration, même si le
// bouton correspondant est masqué côté interface pour les non-admins.
router.delete('/users/:id', requireAuth, async (req, res) => {
  await db.deleteOne('users', { _id: req.params.id });
  res.sendStatus(204);
});

// Même défaut sur une fonctionnalité de gestion de facturation interne,
// non liée dans l'UI standard mais accessible sans contrôle de rôle.
router.post('/internal/refund', requireAuth, async (req, res) => {
  const { orderId, amount } = req.body;
  await db.updateOne('orders', { _id: orderId }, { refunded: true, refundAmount: amount });
  res.json({ ok: true });
});

module.exports = router;
