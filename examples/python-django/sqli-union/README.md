# UNION-based SQL Injection — CWE-89

Le code vulnérable concatène `id` (attendu numérique) directement dans une requête SQL sans typage ni paramètre lié, permettant un `UNION SELECT` pour extraire des données d'autres tables. La correction caste strictement `id` en entier et utilise un paramètre lié (`%s`), éliminant toute possibilité d'injecter une clause SQL supplémentaire. Voir `rules/remediation/sqli-union.md` pour d'autres langages.
