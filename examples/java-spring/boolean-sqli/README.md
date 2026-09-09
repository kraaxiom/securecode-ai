# Boolean-based SQL Injection (CWE-89)

Le code vulnérable construit la clause `WHERE` par concaténation directe du paramètre `name`, ce qui permet à un attaquant d'injecter une expression logique (par exemple `' OR '1'='1`) et d'altérer le nombre ou la nature des résultats retournés. La version corrigée utilise une requête paramétrée (`?` lié via `JdbcTemplate`), garantissant que la valeur utilisateur est traitée comme une donnée et non comme du code SQL, quelle que soit son contenu.
