# Injection SQL (CWE-89)

Le code vulnérable concatène directement le paramètre `name` dans le texte de la requête SQL exécutée via `JdbcTemplate.queryForList`, ce qui permet à un attaquant d'injecter une clause telle que `' OR '1'='1` pour contourner le filtre de recherche et exposer l'ensemble des utilisateurs. La version corrigée utilise le placeholder paramétré `?` de `JdbcTemplate`, qui lie la valeur utilisateur en tant que donnée typée séparée du texte SQL, empêchant toute modification de la structure de la requête.
