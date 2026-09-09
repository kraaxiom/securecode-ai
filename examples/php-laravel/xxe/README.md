# XML External Entity Injection (CWE-611)

Le code vulnérable charge un document XML fourni par l'utilisateur avec la configuration par défaut de `DOMDocument`, qui autorise la résolution d'entités externes, exposant l'application à une divulgation de fichiers ou une SSRF. La correction désactive le chargeur d'entités externes via `libxml_set_external_entity_loader(null)` et charge le document avec `LIBXML_NONET` (sans `LIBXML_NOENT`), empêchant toute résolution d'entité externe (CWE-611 : Improper Restriction of XML External Entity Reference).
