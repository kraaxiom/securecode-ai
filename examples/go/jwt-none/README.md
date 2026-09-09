# jwt-none (CWE-347)

La version vulnerable ne restreint pas les algorithmes de signature acceptes lors du parsing du JWT, permettant a un attaquant d'envoyer un token avec `alg: "none"` et une signature vide qui sera accepte comme valide. La version corrigee impose une liste blanche explicite d'algorithmes autorises (`jwt.WithValidMethods`), excluant de fait l'algorithme "none". C'est la remediation standard recommandee contre CWE-347.
