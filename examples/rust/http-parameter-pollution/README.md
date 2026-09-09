# HTTP Parameter Pollution (CWE-235)

Le code vulnérable utilise l'extracteur `web::Query` qui ne retient silencieusement qu'une valeur pour le paramètre `role`, sans détecter une éventuelle duplication pouvant être interprétée différemment par un proxy en amont. La correction compte explicitement les occurrences de `role=` dans la chaîne de requête brute (`req.query_string()`) et rejette la requête avec un code 400 en cas de duplication, conformément à la remédiation CWE-235.
