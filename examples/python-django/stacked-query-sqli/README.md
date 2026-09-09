# Stacked Query SQL Injection — CWE-89

La vue vulnérable construit une requête `UPDATE` par concaténation de `name` et `user_id`, ce qui permet, selon la configuration du driver, d'ajouter une instruction SQL distincte après un point-virgule. La correction type strictement `user_id` en entier et utilise des paramètres liés (`%s`) pour toutes les valeurs, empêchant toute exécution multi-instructions. Voir `rules/remediation/stacked-query-sqli.md` pour d'autres langages.
