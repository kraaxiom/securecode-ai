const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// CORRIGÉ — couche d'autorisation centralisée et réutilisable, appliquée
// selon le principe "deny by default" : l'accès est refusé sauf si la
// fonction de vérification retourne explicitement vrai.
function authorize(loadResource, checkFn) {
  return async (req, res, next) => {
    const resource = await loadResource(req);
    if (!resource || !checkFn(req.user, resource)) {
      return res.status(403).json({ error: 'Accès refusé' });
    }
    req.resource = resource;
    next();
  };
}

router.get(
  '/invoices/:id/download',
  requireAuth,
  authorize(
    (req) => db.findOne('invoices', { _id: req.params.id }),
    (user, invoice) => invoice.ownerId === user.id || user.role === 'admin'
  ),
  (req, res) => res.download(req.resource.filePath)
);

module.exports = router;
