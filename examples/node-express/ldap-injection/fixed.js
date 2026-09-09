// Correction : LDAP Injection (CWE-90)
// La valeur utilisateur est échappée avec les règles d'échappement de
// filtre LDAP (RFC 4515) avant d'être insérée dans le filtre, et son
// format est validé au préalable.
const express = require('express');
const ldap = require('ldapjs');
const router = express.Router();

// Échappement des caractères spéciaux de filtre LDAP : * ( ) \ NUL
function escapeLdapFilter(value) {
  return String(value).replace(/[\\*()\0]/g, (char) => {
    return '\\' + char.charCodeAt(0).toString(16).padStart(2, '0');
  });
}

router.post('/directory/search', (req, res) => {
  const uid = req.body.uid;

  // Validation stricte du format attendu (identifiant alphanumérique).
  if (typeof uid !== 'string' || !/^[a-zA-Z0-9._-]{1,64}$/.test(uid)) {
    return res.status(400).json({ error: 'Identifiant invalide' });
  }

  const client = ldap.createClient({ url: 'ldap://directory.example.com' });

  // Valeur échappée avant insertion dans le filtre.
  const filter = `(uid=${escapeLdapFilter(uid)})`;

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
