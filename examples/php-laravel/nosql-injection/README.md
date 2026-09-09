# NoSQL Injection (CWE-943)

Le code vulnérable transmet directement les champs `username`/`password` du corps de requête HTTP comme filtre de la requête MongoDB, sans vérifier leur type. La correction ajoute une validation Laravel stricte (`string`) qui rejette tout objet ou opérateur (`$ne`, `$gt`, etc.) envoyé à la place d'une chaîne scalaire, empêchant ainsi la modification de la logique de la requête (CWE-943 : Improper Neutralization of Special Elements in Data Query Logic).
