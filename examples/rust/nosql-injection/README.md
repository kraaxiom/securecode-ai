# NoSQL Injection (CWE-943)

Le code vulnérable convertit le corps JSON brut de la requête directement en document de filtre MongoDB, ce qui permettrait à un attaquant d'injecter un opérateur (`{"$ne": null}`) là où une valeur scalaire est attendue et de contourner l'authentification. La correction définit une structure `LoginRequest` typée en `String` pour `username`/`password`, garantissant via serde qu'aucun objet ne peut être désérialisé à leur place, puis reconstruit le filtre explicitement avec `doc!`, conformément à CWE-943.
