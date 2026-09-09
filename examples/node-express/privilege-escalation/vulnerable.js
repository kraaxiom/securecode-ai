const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// VULNÉRABLE — Privilege Escalation
// Le rôle est mis à jour comme n'importe quel autre champ de profil : le
// contenu de req.body.role est appliqué tel quel, sans vérifier que
// l'appelant est lui-même autorisé à accorder ce niveau de privilège.
// Un utilisateur standard peut ainsi s'auto-attribuer le rôle "admin".
router.put('/api/users/:id/role', requireAuth, async (req, res) => {
  const user = await db.updateOne('users', { _id: req.params.id }, { role: req.body.role });
  res.json(user);
});

module.exports = router;
