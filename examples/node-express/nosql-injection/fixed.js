// Correction : NoSQL Injection (CWE-943)
// Chaque champ est validé par type (chaîne stricte) avant d'être utilisé
// dans le filtre de requête, rejetant tout objet/opérateur MongoDB
// (`$ne`, `$gt`, `$where`, `$regex`) envoyé à la place d'une valeur simple.
const express = require('express');
const router = express.Router();
const User = require('../models/User');

router.post('/auth/login', async (req, res) => {
  const { username, password } = req.body;

  // Validation stricte de type : rejette tout objet/opérateur injecté.
  if (typeof username !== 'string' || typeof password !== 'string') {
    return res.status(400).json({ error: 'Format invalide' });
  }

  const user = await User.findOne({ username, password });

  if (!user) {
    return res.status(401).json({ error: 'Identifiants invalides' });
  }
  res.json({ id: user._id, username: user.username });
});

module.exports = router;
