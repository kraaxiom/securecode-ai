# Command Injection (CWE-78)

La version vulnérable concatène l'hôte utilisateur dans une chaîne exécutée via `sh -c`, permettant l'injection de métacaractères shell pour exécuter des commandes arbitraires. La correction valide d'abord l'entrée comme adresse IP stricte (`IpAddr::from_str`), puis exécute `ping` directement avec des arguments passés en tableau distinct, sans jamais invoquer de shell. Cette approche élimine le vecteur décrit par CWE-78.
