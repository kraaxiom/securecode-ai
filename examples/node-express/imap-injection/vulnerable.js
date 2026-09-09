// Faille : IMAP Injection (CWE-93)
// Le terme de recherche fourni par l'utilisateur est concaténé directement
// dans une commande IMAP brute, permettant l'injection de guillemets ou de
// mots-clés de commande supplémentaires.
const express = require('express');
const router = express.Router();

router.get('/mailbox/search', (req, res) => {
  const term = req.query.q;

  // Construction manuelle d'une commande IMAP par concaténation.
  const command = `SEARCH SUBJECT "${term}"`;
  imapClient.exec(command, (err, results) => {
    if (err) return res.status(500).json({ error: 'Erreur de recherche' });
    res.json(results);
  });
});

module.exports = router;
