## XML Injection (CWE-91)

Le code vulnérable construit un document XML par concaténation de chaînes incluant le champ `name` fourni par l'utilisateur, sans échapper les caractères spéciaux (`<`, `>`, `&`), permettant à un attaquant d'altérer la structure du document (ajout de nœuds, falsification de données). La correction remplace la concaténation par `xmlbuilder2`, une bibliothèque de sérialisation XML qui échappe automatiquement le contenu des nœuds et attributs.
