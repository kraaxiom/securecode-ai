const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// Vérifie côté serveur que l'appelant a le droit d'accorder le rôle demandé
// (ex. seul un admin peut attribuer le rôle "admin" ou "manager").
function canGrantRole(actor, requestedRole) {
  const rank = { user: 0, manager: 1, admin: 2 };
  return actor.role === 'admin' || rank[actor.role] > rank[requestedRole];
}

// CORRIGÉ — l'attribution de rôle vérifie explicitement le droit
// d'attribution de l'appelant, puis invalide les sessions actives de la
// cible pour éviter la persistance d'anciens privilèges dans son token.
router.put('/api/users/:id/role', requireAuth, async (req, res) => {
  const requestedRole = req.body.role;
  if (!canGrantRole(req.user, requestedRole)) {
    return res.status(403).json({ error: 'Attribution de rôle non autorisée' });
  }
  const user = await db.updateOne('users', { _id: req.params.id }, { role: requestedRole });
  await db.deleteMany('sessions', { userId: user._id }); // invalide les sessions actives
  res.json(user);
});

module.exports = router;
