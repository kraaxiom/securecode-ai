// Faille : LDAP Injection (CWE-90)
// Le paramètre "uid" est concaténé directement dans le filtre de recherche
// LDAP. Un attaquant peut injecter des méta-caractères de filtre
// (`* ( ) \` ou NUL) pour altérer la logique du filtre.
const express = require('express');
const ldap = require('ldapjs');
const router = express.Router();

router.post('/directory/search', (req, res) => {
  const uid = req.body.uid;
  const client = ldap.createClient({ url: 'ldap://directory.example.com' });

  // Concaténation directe dans le filtre : dangereux.
  const filter = `(uid=${uid})`;

  client.search('ou=users,dc=example,dc=com', { filter }, (err, search) => {
    if (err) {
      return res.status(500).json({ error: 'Erreur de recherche' });
    }
    const entries = [];
    search.on('searchEntry', (entry) => entries.push(entry.pojo));
    search.on('end', () => res.json(entries));
  });
});

module.exports = router;
