# Error-based SQL Injection

La version vulnérable concatène `order_id` dans la requête SQL et renvoie `str(e)` au client en cas d'exception, ce qui correspond à CWE-89 : un attaquant peut provoquer des erreurs de syntaxe ciblées pour extraire des données via les messages d'erreur du driver. La version corrigée utilise une requête paramétrée avec un identifiant validé en entier, et journalise le détail de l'erreur côté serveur tout en renvoyant un message générique au client.
