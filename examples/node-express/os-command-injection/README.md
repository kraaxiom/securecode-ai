## OS Command Injection (CWE-78)

Le code vulnérable concatène le paramètre `filename` dans une commande shell exécutée via `exec()`, permettant à un attaquant d'injecter des méta-caractères shell (`;`, `|`, `&&`, backticks) pour exécuter des commandes arbitraires sur le serveur. La correction valide le nom de fichier avec une liste blanche stricte, vérifie que le chemin résolu reste dans le dossier autorisé, puis remplace `exec` par `execFile` avec des arguments passés en tableau (aucun shell impliqué).
