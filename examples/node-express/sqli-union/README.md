## UNION-based SQL Injection (CWE-89)

Le code vulnérable concatène directement le paramètre `id` dans la requête SQL, permettant à un attaquant d'ajouter une clause `UNION SELECT` pour extraire des données d'autres tables (ex: identifiants ou mots de passe de la table `users`). La correction caste l'identifiant en entier, le valide, puis utilise une requête préparée avec paramètre lié, empêchant toute injection de clause supplémentaire.
