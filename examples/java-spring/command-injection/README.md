# Command Injection (CWE-78)

Le code vulnérable concatène le paramètre `host` dans une commande shell exécutée via `/bin/sh -c`, ce qui permet à un attaquant d'injecter des métacaractères (`;`, `|`, `&&`) pour exécuter des commandes arbitraires sur le serveur. La version corrigée valide strictement que `host` est une adresse IP littérale, puis lance le processus via `ProcessBuilder` avec un tableau d'arguments distincts, sans jamais passer par un interpréteur shell — éliminant tout canal d'injection de commande.
