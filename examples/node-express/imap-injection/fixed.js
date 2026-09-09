// Correction : IMAP Injection (CWE-93)
// Utilisation de l'API structurée du client IMAP (imapflow), qui encode
// correctement les littéraux, sans jamais construire de commande texte brute.
const express = require('express');
const router = express.Router();

router.get('/mailbox/search', async (req, res) => {
  // Longueur limitée et pas de construction de commande brute.
  const term = String(req.query.q ?? '').slice(0, 200);

  try {
    await client.mailboxOpen('INBOX');
    // L'API structurée encode elle-même le littéral de recherche.
    const results = await client.search({ subject: term });
    res.json(results);
  } catch (err) {
    res.status(500).json({ error: 'Erreur de recherche' });
  }
});

module.exports = router;
