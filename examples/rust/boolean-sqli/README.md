# Boolean-based SQL Injection (CWE-89)

La version vulnérable construit la clause `WHERE name = '...'` par concaténation directe de l'entrée utilisateur, permettant à un attaquant de modifier la valeur de vérité de la condition et d'en déduire des données via une réponse binaire (résultats vs vide). La correction utilise `sqlx::query_as` avec un paramètre lié `$1`, garantissant que la valeur est toujours traitée comme une donnée et jamais comme un fragment de logique SQL, conformément à CWE-89.
