const express = require('express');
const path = require('path');
const app = express();
const { requireAuth, requireRole } = require('../middleware/auth');

// CORRIGÉ — plus aucun répertoire sensible n'est servi en statique. Chaque
// fichier est délivré via un contrôleur qui vérifie explicitement
// l'autorisation, indépendamment de sa découvrabilité.
function resolveSafePath(baseDir, fileName) {
  const resolved = path.resolve(baseDir, fileName);
  if (!resolved.startsWith(path.resolve(baseDir))) {
    throw new Error('Chemin invalide');
  }
  return resolved;
}

app.get('/backups/:file', requireAuth, requireRole('admin'), (req, res) => {
  const filePath = resolveSafePath('backups', req.params.file);
  res.sendFile(filePath);
});

// La confirmation de paiement vérifie côté serveur que l'étape
// d'autorisation du paiement a réellement été franchie avant d'afficher
// la page finale — impossible de sauter les étapes précédentes.
app.get('/checkout/confirmation', requireAuth, (req, res) => {
  const checkout = req.session.checkout;
  if (!checkout || checkout.status !== 'payment_authorized') {
    return res.redirect('/checkout/start');
  }
  res.render('confirmation', { order: checkout.order });
});

module.exports = app;
