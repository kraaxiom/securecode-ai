# Boolean-based SQL Injection (CWE-89)

La version vulnérable insère la valeur du paramètre `name` directement dans la clause `WHERE`, permettant à un attaquant d'injecter une expression logique (ex: `' OR '1'='1`) qui modifie le résultat retourné. La correction utilise un paramètre lié `@Name` via `SqlCommand.Parameters`, garantissant que la valeur est traitée comme une donnée et non comme du code SQL. La logique métier reste identique, mais le canal d'inférence booléen disparaît.
