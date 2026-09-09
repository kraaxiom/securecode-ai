# Code Injection (CWE-94)

Le code vulnérable assemble une chaîne d'expression à partir de valeurs utilisateur et la transmet à un évaluateur d'expression générique, donnant à un attaquant un contrôle potentiel sur la logique exécutée. La correction remplace cette évaluation dynamique par un `match` explicite sur une liste blanche d'opérations (`add`, `sub`, `mul`, `div`), rejetant toute valeur non prévue. Cela élimine le vecteur d'exécution de code arbitraire décrit par CWE-94.
