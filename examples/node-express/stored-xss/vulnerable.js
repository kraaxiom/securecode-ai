// Vulnérable : XSS stocké — CWE-79
// La biographie d'un profil utilisateur est persistée telle quelle en base
// de données, puis réaffichée à tout visiteur consultant le profil par
// simple concaténation de chaînes, sans aucun encodage de sortie. Comme la
// donnée est stockée, l'exécution touche potentiellement chaque visiteur du
// profil, sans nécessiter d'interaction particulière avec un lien piégé.

const express = require('express');
const router = express.Router();

// Simulation d'une couche d'accès aux données
const User = {
  async findById(id) {
    return { id, bio: fakeDb[id] };
  },
};
const fakeDb = {};

router.get('/profile/:id', async (req, res) => {
  const user = await User.findById(req.params.id);
  res.send(`<div class="bio">${user.bio || ''}</div>`);
});

module.exports = router;
