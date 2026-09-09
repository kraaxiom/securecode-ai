# Command Injection

La version vulnérable construit une commande shell par f-string et l'exécute avec `subprocess.check_output(..., shell=True)`, ce qui correspond à CWE-78 : un attaquant peut injecter des métacaractères shell (`;`, `&&`, `|`) pour exécuter des commandes arbitraires. La version corrigée valide l'entrée comme adresse IP via `ipaddress.ip_address()`, puis exécute la commande sans shell avec les arguments passés en liste distincte, rendant l'injection impossible.
