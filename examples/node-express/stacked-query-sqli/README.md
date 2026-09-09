## Stacked Query SQL Injection (CWE-89)

Le code vulnérable configure la connexion `mysql2` avec `multipleStatements: true` et concatène l'entrée utilisateur dans la requête, permettant à un attaquant d'ajouter un point-virgule suivi d'une instruction SQL complètement distincte (`DROP`, `INSERT`...), avec un impact potentiellement plus large qu'une injection classique. La correction désactive l'option multi-instructions du driver (non nécessaire fonctionnellement) et utilise systématiquement des paramètres liés.
