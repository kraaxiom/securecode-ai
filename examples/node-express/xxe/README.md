## XML External Entity Injection - XXE (CWE-611)

Le code vulnérable analyse un document XML fourni par l'utilisateur avec la configuration par défaut de `libxmljs2`, qui autorise le traitement des DTD et des entités externes. Un attaquant peut définir une entité pointant vers un fichier local ou une URL interne, menant à une divulgation de fichiers sensibles, du SSRF, voire un déni de service par expansion d'entités. La correction désactive explicitement le chargement de DTD (`dtdload: false`) et la substitution d'entités (`noent: false`) sur le parseur.
