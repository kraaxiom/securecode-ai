# Error-based SQL Injection (CWE-89)

Le code vulnérable concatène l'identifiant dans la requête SQL et renvoie `ex.Message` brut au client en cas d'erreur, exposant des informations exploitables pour de l'injection basée sur les erreurs. La correction contraint `id` à un entier via la route (`{id:int}`), utilise une requête paramétrée avec `@Id`, et remplace le message d'erreur détaillé par un message générique tout en journalisant l'exception côté serveur via `ILogger`.
