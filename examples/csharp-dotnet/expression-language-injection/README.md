# Expression Language Injection (CWE-917)

La version vulnérable transmet directement `request.Expression` à `DataTable.Compute`, un moteur d'évaluation d'expression, ce qui permet à un attaquant d'exécuter une expression arbitraire côté serveur. La correction supprime toute évaluation dynamique et remplace la logique par un dictionnaire d'opérateurs autorisés (`gt`, `lt`, `eq`) sélectionné via une liste blanche, avec des opérandes numériques typés jamais interprétés comme du code.
