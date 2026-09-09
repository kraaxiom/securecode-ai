## SQL Injection (CWE-89)

Le code vulnérable interpole le paramètre `name` directement dans une requête SQL via un template literal (`mysql2`), permettant à un attaquant d'injecter un caractère `'` ou un mot-clé SQL pour altérer la logique de la requête (lecture, modification, contournement d'authentification). La correction remplace l'interpolation par une requête préparée avec paramètre lié (`?`), qui sépare strictement le code SQL des données utilisateur.
