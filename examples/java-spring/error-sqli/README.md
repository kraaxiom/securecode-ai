# Error-based SQL Injection (CWE-89)

Le code vulnérable concatène l'identifiant `id` dans la requête SQL et renvoie le message d'exception détaillé du driver au client en cas d'erreur, permettant à un attaquant d'extraire des informations sur le schéma via des erreurs provoquées volontairement. La version corrigée valide et caste l'identifiant en entier, utilise une requête préparée (`?` lié via `JdbcTemplate`), et remplace la réponse d'erreur par un message générique, le détail complet étant journalisé uniquement côté serveur.
