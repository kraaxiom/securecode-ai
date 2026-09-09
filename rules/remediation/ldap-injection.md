# Remédiation — LDAP Injection

## Principe
Ne jamais construire un filtre de recherche LDAP ou un DN par concaténation directe d'une entrée utilisateur. Utiliser les fonctions d'échappement dédiées fournies par la bibliothèque LDAP du langage (échappement des caractères spéciaux `* ( ) \ NUL`), et valider le format attendu de l'entrée avant construction du filtre.

## PHP (ldap_escape)
```php
// Avant — vulnérable
$uid = $_POST['uid'];
$filter = "(uid=$uid)";
$result = ldap_search($conn, $baseDn, $filter);

// Après — sécurisé
$uid = ldap_escape($_POST['uid'], '', LDAP_ESCAPE_FILTER);
$filter = "(uid=$uid)";
$result = ldap_search($conn, $baseDn, $filter);
```

## JavaScript / Node.js (ldapjs)
```js
// Avant — vulnérable
const uid = req.body.uid;
const filter = `(uid=${uid})`;
client.search(baseDn, { filter }, callback);

// Après — sécurisé
const { escape } = require('ldapjs-filter-escape'); // ou implémentation dédiée
const uid = escape(req.body.uid);
const filter = `(uid=${uid})`;
client.search(baseDn, { filter }, callback);
```

## Python (python-ldap)
```python
# Avant — vulnérable
uid = request.form['uid']
filter_str = f"(uid={uid})"
conn.search_s(base_dn, ldap.SCOPE_SUBTREE, filter_str)

# Après — sécurisé
from ldap.filter import escape_filter_chars

uid = escape_filter_chars(request.form['uid'])
filter_str = f"(uid={uid})"
conn.search_s(base_dn, ldap.SCOPE_SUBTREE, filter_str)
```

## Java (Spring LDAP)
```java
// Avant — vulnérable
String uid = request.getParameter("uid");
String filter = "(uid=" + uid + ")";
ctx.search(baseDn, filter, controls);

// Après — sécurisé
import org.springframework.ldap.filter.EqualsFilter;

Filter filter = new EqualsFilter("uid", uid); // échappement automatique
ldapTemplate.search(baseDn, filter.encode(), controls);
```

## Checklist de vérification post-patch
- [ ] Aucune concaténation directe d'entrée utilisateur dans un filtre LDAP ou un DN.
- [ ] La fonction d'échappement dédiée à la bibliothèque LDAP utilisée est bien appliquée à chaque valeur insérée dans le filtre.
- [ ] Un test confirme qu'une entrée contenant `*`, `(`, `)` ou `\` ne modifie pas la logique du filtre.
- [ ] Le compte de service utilisé pour la connexion LDAP dispose du minimum de privilèges nécessaires (lecture seule si possible).
- [ ] Le format attendu de l'entrée (ex: identifiant alphanumérique) est validé avant construction du filtre.
