# Remédiation — XML External Entity (XXE) Injection

## Principe
Désactiver explicitement le traitement des DTD et des entités externes sur tout parseur XML recevant des entrées non fiables. Utiliser des bibliothèques ou des modes de parsing "sûrs par défaut" lorsqu'ils sont disponibles.

## PHP (libxml / DOMDocument)
```php
// Avant — vulnérable
$doc = new DOMDocument();
$doc->loadXML($userSuppliedXml);

// Après — sécurisé
$doc = new DOMDocument();
libxml_set_external_entity_loader(null); // désactive le chargement d'entités externes
$doc->loadXML($userSuppliedXml, LIBXML_NONET | LIBXML_NOENT === false ? LIBXML_NONET : LIBXML_NONET);
// Note : ne jamais activer LIBXML_NOENT avec une entrée non fiable ; conserver DTD/ENTITY désactivés.
```

## JavaScript / Node.js (libxmljs2)
```js
// Avant — vulnérable
const doc = libxmljs.parseXml(userSuppliedXml);

// Après — sécurisé
const doc = libxmljs.parseXml(userSuppliedXml, {
  noent: false,   // ne pas substituer les entités
  dtdload: false, // ne pas charger de DTD externe
  noblanks: true,
});
```

## Python (lxml.etree)
```python
# Avant — vulnérable
from lxml import etree
tree = etree.parse(user_supplied_file)

# Après — sécurisé
from lxml import etree
parser = etree.XMLParser(resolve_entities=False, no_network=True, dtd_validation=False, load_dtd=False)
tree = etree.parse(user_supplied_file, parser=parser)
```

## Java (DocumentBuilderFactory)
```java
// Avant — vulnérable
DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
Document doc = dbf.newDocumentBuilder().parse(userSuppliedStream);

// Après — sécurisé
DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
dbf.setFeature("http://apache.org/xml/features/disallow-doctype-decl", true);
dbf.setFeature("http://xml.org/sax/features/external-general-entities", false);
dbf.setFeature("http://xml.org/sax/features/external-parameter-entities", false);
dbf.setXIncludeAware(false);
dbf.setExpandEntityReferences(false);
Document doc = dbf.newDocumentBuilder().parse(userSuppliedStream);
```

## Checklist de vérification post-patch
- [ ] Le traitement des déclarations DOCTYPE/DTD est explicitement désactivé sur le parseur.
- [ ] La résolution des entités externes (générales et paramétriques) est désactivée.
- [ ] Les formats basés sur XML uploadés par l'utilisateur (SVG, DOCX, XLSX) passent par ce même parseur durci.
- [ ] Un test confirme qu'un document XML avec une déclaration d'entité externe est rejeté ou traité sans résolution.
- [ ] La configuration durcie est appliquée à tous les points d'entrée XML de l'application, pas seulement au fichier corrigé.
