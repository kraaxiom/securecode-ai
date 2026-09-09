# XML Injection — CWE-91

Le code vulnérable construit le document XML par f-string en insérant `name` et `role` sans échappement, ce qui permet à un attaquant contenant `<`, `>` ou `&` d'altérer la structure du document. La correction utilise `lxml.etree` pour construire l'arbre XML, qui échappe automatiquement le texte des nœuds lors de la sérialisation. Voir `rules/remediation/xml-injection.md` pour d'autres langages.
