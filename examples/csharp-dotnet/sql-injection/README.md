# SQL Injection générique (CWE-89)

La version vulnérable construit la requête `SELECT ... WHERE name = '{name}'` par interpolation de chaîne, permettant à un attaquant d'injecter un caractère `'` ou un mot-clé SQL pour modifier la logique de la requête. La correction utilise un `SqlCommand` avec un paramètre lié (`@name`) typé (`NVarChar`), de sorte que la valeur utilisateur est transmise séparément du texte SQL et ne peut jamais en altérer la structure.
