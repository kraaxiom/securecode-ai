# XML External Entity Injection — CWE-611

La vue vulnérable analyse un fichier XML uploadé avec la configuration par défaut de `lxml.etree`, qui ne désactive pas le traitement des DTD ni des entités externes, exposant l'application à une divulgation de fichiers, un SSRF ou un déni de service par expansion d'entités. La correction configure un `XMLParser` durci (`resolve_entities=False`, `no_network=True`, `load_dtd=False`) appliqué à tous les uploads XML. Voir `rules/remediation/xxe.md` pour d'autres langages.
