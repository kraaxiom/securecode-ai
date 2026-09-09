# UNION-based SQL Injection (CWE-89)

La version vulnérable insère `id` directement dans `SELECT ... WHERE id = {id}` sans validation ni paramétrage, ce qui permet à un attaquant d'ajouter une clause `UNION SELECT` pour extraire des colonnes d'autres tables (ex: identifiants et mots de passe utilisateurs). La correction caste `id` en entier via `int.TryParse` (rejetant tout ce qui n'est pas numérique) et utilise un paramètre lié (`@id`), éliminant toute possibilité d'ajouter une instruction SQL supplémentaire.
