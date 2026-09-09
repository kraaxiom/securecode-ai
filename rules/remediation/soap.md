# Remédiation — Sécurité des API SOAP

## Principe
Désactiver systématiquement la résolution des entités externes et le traitement de DTD dans tout parseur XML traitant des messages SOAP entrants ; ajouter WS-Security pour les échanges sensibles.

## PHP
```php
// Avant — vulnérable
$doc = new DOMDocument();
$doc->loadXML($soapRequest); // entités externes résolues par défaut sur anciennes libxml

// Après — sécurisé
libxml_disable_entity_loader(true); // pour PHP < 8.0
$doc = new DOMDocument();
$doc->loadXML($soapRequest, LIBXML_NONET | LIBXML_NOENT ^ LIBXML_NOENT); // ne pas activer NOENT
// Utiliser explicitement des options sûres :
$doc->resolveExternals = false;
$doc->substituteEntities = false;
```

```php
// SoapServer — désactiver le cache WSDL et forcer les options sûres
$server = new SoapServer($wsdl, [
    'cache_wsdl' => WSDL_CACHE_NONE,
    'features' => SOAP_SINGLE_ELEMENT_ARRAYS,
]);
```

## Java
```java
// Avant — vulnérable
DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
DocumentBuilder db = dbf.newDocumentBuilder();
Document doc = db.parse(soapInputStream);

// Après — sécurisé
DocumentBuilderFactory dbf = DocumentBuilderFactory.newInstance();
dbf.setFeature("http://apache.org/xml/features/disallow-doctype-decl", true);
dbf.setFeature("http://xml.org/sax/features/external-general-entities", false);
dbf.setFeature("http://xml.org/sax/features/external-parameter-entities", false);
dbf.setXIncludeAware(false);
dbf.setExpandEntityReferences(false);
DocumentBuilder db = dbf.newDocumentBuilder();
Document doc = db.parse(soapInputStream);
```

## C# (.NET)
```csharp
// Avant — vulnérable
XmlDocument doc = new XmlDocument();
doc.Load(soapStream); // XmlResolver actif par défaut sur anciennes versions .NET

// Après — sécurisé
XmlDocument doc = new XmlDocument();
doc.XmlResolver = null; // empêche la résolution d'entités externes
doc.Load(soapStream);
```

## Python (lxml)
```python
# Avant — vulnérable
from lxml import etree
tree = etree.parse(soap_stream)

# Après — sécurisé
from lxml import etree
parser = etree.XMLParser(resolve_entities=False, no_network=True, dtd_validation=False)
tree = etree.parse(soap_stream, parser)
```

## Checklist de vérification post-patch
- [ ] Tout parseur XML utilisé pour les messages SOAP désactive explicitement la résolution des entités externes et le DTD.
- [ ] Les bibliothèques XML/SOAP sont à jour (pas de version connue pour un comportement XXE par défaut).
- [ ] WS-Security (signature/chiffrement) est en place pour les échanges contenant des données sensibles.
- [ ] Les messages SOAP/WSDL entrants sont validés contre leur schéma XSD avant traitement métier.
- [ ] Un test confirme qu'une entité externe déclarée dans une requête SOAP n'est pas résolue.
