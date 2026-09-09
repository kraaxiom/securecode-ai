// Corrigé : XSS stocké — CWE-79
// L'affichage de la biographie passe par un moteur de templates avec
// échappement automatique conservé (EJS <%= %>). La donnée persistée en
// base reste traitée comme non fiable au moment de l'affichage, quel que
// soit le nombre de fois où elle a déjà été relue ou réaffichée.

const express = require('express');
const router = express.Router();

const User = {
  async findById(id) {
    return { id, bio: fakeDb[id] };
  },
};
const fakeDb = {};

router.get('/profile/:id', async (req, res) => {
  const user = await User.findById(req.params.id);
  // template profile.ejs : <div class="bio"><%= bio %></div>
  res.render('profile', { bio: user.bio || '' });
});

module.exports = router;
