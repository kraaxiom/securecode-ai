# Error-based SQL Injection (CWE-89)

La version vulnérable concatène l'identifiant dans la requête SQL et renvoie le message d'erreur natif du driver au client, ce qui permet à un attaquant d'extraire des informations via des erreurs provoquées. La correction utilise une requête préparée avec paramètre lié, type l'identifiant en `i32` dès l'extraction de la route, journalise le détail de l'erreur côté serveur (`log::error!`) et ne renvoie qu'un message générique au client, conformément à CWE-89.
