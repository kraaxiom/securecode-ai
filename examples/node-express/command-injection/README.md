## Command Injection (CWE-78)

Le code vulnérable concatène le paramètre `host` dans une chaîne exécutée par `exec()` via un shell, ce qui permet l'injection de métacaractères shell pour exécuter des commandes arbitraires. La correction remplace `exec` par `execFile` avec des arguments passés en tableau distinct (aucun shell impliqué) et ajoute une validation stricte via `net.isIP()` avant toute exécution.
