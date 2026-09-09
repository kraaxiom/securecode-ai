# OS Command Injection (CWE-78)

Le code vulnérable concatène directement le paramètre `host` fourni par l'utilisateur dans une commande shell passée à `system()`, permettant l'injection de méta-caractères shell pour exécuter des commandes arbitraires. La correction impose une liste blanche stricte (adresse IP ou nom d'hôte alphanumérique) et échappe la valeur avec `escapeshellarg()` avant l'appel shell, neutralisant l'injection (CWE-78 : Improper Neutralization of Special Elements used in an OS Command).
