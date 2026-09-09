## Boolean-based SQL Injection (CWE-89)

Le code vulnérable insère le paramètre `name` par concaténation dans la clause `WHERE`, ce qui permet à un attaquant d'injecter une expression logique modifiant la véracité de la condition et donc le nombre de résultats retournés. La correction utilise une requête préparée avec paramètre lié (`?`), garantissant que la valeur utilisateur est toujours traitée comme une donnée et jamais comme une expression SQL.
