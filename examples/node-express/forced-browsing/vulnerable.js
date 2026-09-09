const express = require('express');
const app = express();

// VULNÉRABLE — Forced Browsing
// Le répertoire de sauvegardes est servi statiquement sans aucun contrôle
// d'accès. Ces fichiers ne sont liés nulle part dans l'interface, mais
// restent accessibles à quiconque devine ou découvre le chemin exact
// (sécurité par obscurité, pas par contrôle d'accès réel).
app.use('/backups', express.static('backups/'));

// Étape finale d'un flux de paiement accessible directement, sans
// vérifier côté serveur que les étapes précédentes ont été complétées.
app.get('/checkout/confirmation', (req, res) => {
  res.render('confirmation', { order: req.session.lastOrder });
});

module.exports = app;
