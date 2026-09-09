## Error-based SQL Injection (CWE-89)

Le code vulnérable construit la requête SQL par concaténation et renvoie le message d'erreur brut du driver (`err.message`) au client en cas d'échec, permettant d'extraire des informations via des erreurs SQL provoquées volontairement. La correction utilise une requête préparée avec paramètre lié, valide que l'identifiant est bien un entier, et remplace le message d'erreur détaillé par un message générique, le détail étant journalisé côté serveur uniquement.
