// Faille : NoSQL Injection (CWE-943)
// Le corps de la requête HTTP est transmis tel quel comme filtre de
// requête MongoDB. Un attaquant peut envoyer un objet contenant des
// opérateurs (ex: $ne, $gt) là où une valeur scalaire est attendue,
// ce qui modifie la logique de la requête plutôt que sa syntaxe.
const express = require('express');
const router = express.Router();
const User = require('../models/User'); // modèle Mongoose

router.post('/auth/login', async (req, res) => {
  const { username, password } = req.body;

  // Transmission brute du body en filtre de requête : dangereux.
  const user = await User.findOne({ username, password });

  if (!user) {
    return res.status(401).json({ error: 'Identifiants invalides' });
  }
  res.json({ id: user._id, username: user.username });
});

module.exports = router;
