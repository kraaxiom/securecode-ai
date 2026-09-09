# OS Command Injection — CWE-78

Le code vulnérable transmet le paramètre `host` non filtré à `subprocess.run(..., shell=True)`, ce qui permet à un attaquant d'injecter des méta-caractères shell pour exécuter des commandes arbitraires. La correction valide `host` avec une liste blanche stricte (regex alphanumérique/point/tiret) et remplace l'appel shell par un appel à tableau d'arguments (`shell=False`), qui n'interprète jamais de méta-caractères. Voir `rules/remediation/os-command-injection.md` pour d'autres langages.
