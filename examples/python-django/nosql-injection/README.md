# NoSQL Injection

La version vulnérable transmet `username`/`password` du corps JSON directement au filtre `find_one()` sans validation de type, ce qui correspond à CWE-943 : un attaquant peut envoyer un objet opérateur MongoDB (ex: `{"$ne": null}`) à la place d'une chaîne pour contourner l'authentification. La version corrigée vérifie explicitement que ces deux champs sont des chaînes (`isinstance(..., str)`) avant de les utiliser dans la requête, rejetant tout objet ou opérateur injecté.
