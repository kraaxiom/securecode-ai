# Remédiation — XML Injection

## Principe
Ne jamais construire un document XML par concaténation de chaînes incluant des données utilisateur. Utiliser une bibliothèque de sérialisation XML qui échappe automatiquement le contenu des nœuds et des attributs.

## PHP (DOMDocument)
```php
// Avant — vulnérable
$xml = "<user><name>" . $_POST['name'] . "</name></user>";

// Après — sécurisé
$doc = new DOMDocument('1.0', 'UTF-8');
$user = $doc->createElement('user');
$name = $doc->createElement('name');
$name->appendChild($doc->createTextNode($_POST['name'])); // échappement automatique
$user->appendChild($name);
$doc->appendChild($user);
```

## JavaScript / Node.js (xmlbuilder2)
```js
// Avant — vulnérable
const xml = `<user><name>${req.body.name}</name></user>`;

// Après — sécurisé
const { create } = require('xmlbuilder2');
const xml = create({ user: { name: req.body.name } }).end(); // échappement automatique
```

## Python (lxml.etree)
```python
# Avant — vulnérable
xml = f"<user><name>{name}</name></user>"

# Après — sécurisé
from lxml import etree
user = etree.Element("user")
name_el = etree.SubElement(user, "name")
name_el.text = name  # échappement automatique par lxml
xml = etree.tostring(user)
```

## Checklist de vérification post-patch
- [ ] Le document XML n'est plus construit par concaténation de chaînes incluant une variable utilisateur.
- [ ] Une bibliothèque de sérialisation XML est utilisée pour tout point d'insertion de données dynamiques.
- [ ] Un test confirme qu'une valeur contenant `<`, `>` ou `&` ne modifie plus la structure du document généré.
- [ ] La structure du document produit est validée par un schéma (XSD) lorsque c'est applicable.
