# XML Injection (CWE-91)

Le code vulnérable construit le document XML par concaténation directe de chaînes incluant les paramètres `name` et `role`, ce qui permet à un attaquant d'injecter des balises ou des caractères spéciaux pour altérer la structure du document et falsifier des champs non prévus. La version corrigée construit le document via l'API DOM (`createElement`/`createTextNode`), qui échappe automatiquement tout contenu textuel inséré dans les nœuds, empêchant toute modification de la structure XML quelle que soit la valeur fournie par l'utilisateur.
