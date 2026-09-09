# Time-Based Blind SQL Injection (CWE-89)

Le code vulnérable concatène le paramètre `id` dans la requête SQL sans timeout ni liaison de paramètre, ce qui permet à un attaquant d'injecter une fonction de pause et de déduire des informations en mesurant le délai de réponse, même si aucune donnée n'est directement renvoyée. La version corrigée utilise une requête paramétrée (`?`) qui élimine toute la classe de vulnérabilité, complétée par un timeout d'exécution strict (`setQueryTimeout`) qui limite l'impact d'une éventuelle injection résiduelle ailleurs dans l'application.
