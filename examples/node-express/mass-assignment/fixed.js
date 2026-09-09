const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// CORRIGÉ — liste blanche explicite des champs modifiables. Les champs
// sensibles ("role", "isAdmin", etc.) ne sont jamais extraits du corps de
// requête et ne peuvent donc jamais transiter jusqu'au modèle.
router.put('/api/users/:id', requireAuth, async (req, res) => {
  const { name, email } = req.body; // extraction explicite, 'role'/'isAdmin' ignorés
  const user = await db.updateOne('users', { _id: req.params.id }, { name, email });
  res.json(user);
});

module.exports = router;
