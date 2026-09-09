const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// VULNÉRABLE — Mass Assignment
// Le corps entier de la requête est passé tel quel à la mise à jour du
// modèle utilisateur. Un attaquant peut donc injecter des champs non
// prévus par le formulaire d'origine, comme "role" ou "isAdmin", et se
// les faire appliquer directement.
router.put('/api/users/:id', requireAuth, async (req, res) => {
  const user = await db.updateOne('users', { _id: req.params.id }, req.body);
  res.json(user);
});

module.exports = router;
