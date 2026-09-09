const express = require('express');
const router = express.Router();
const db = require('../db');
const { requireAuth } = require('../middleware/auth');

// VULNÉRABLE — Broken Access Control (catégorie générale)
// Le téléchargement de facture ne repose que sur l'authentification.
// Aucune couche d'autorisation centralisée ne vérifie que l'utilisateur
// a réellement le droit d'accéder à cette ressource précise : le contrôle
// d'accès repose à tort uniquement sur le masquage du bouton côté UI.
router.get('/invoices/:id/download', requireAuth, async (req, res) => {
  const invoice = await db.findOne('invoices', { _id: req.params.id });
  if (!invoice) return res.status(404).end();
  res.download(invoice.filePath);
});

module.exports = router;
