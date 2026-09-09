# UNION-based SQL Injection (CWE-89)

Le code vulnérable concatène le paramètre `id` de la query string dans une requête SQL sans validation ni liaison de paramètre, ouvrant la voie à l'ajout d'une clause `UNION SELECT` pour extraire des données d'autres tables. La correction caste `id` en entier via `$request->integer()` et utilise un paramètre lié (`?`) dans la requête préparée, empêchant toute injection de clause supplémentaire (CWE-89 : Improper Neutralization of Special Elements used in an SQL Command).
