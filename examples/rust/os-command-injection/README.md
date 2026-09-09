# OS Command Injection (CWE-78)

Le code vulnérable construisait une commande shell par concaténation d'une entrée utilisateur (`host`), permettant l'injection de méta-caractères shell pour exécuter des commandes arbitraires. La correction remplace l'appel `sh -c` par un appel direct de binaire avec arguments passés séparément (aucune interprétation shell), et ajoute une validation stricte par liste blanche du format de l'hôte avant exécution. Cela élimine la classe de vulnérabilité CWE-78 en supprimant tout point d'interprétation shell des données utilisateur.
