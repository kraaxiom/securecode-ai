# Code Injection

La version vulnérable exécute directement `eval(formula)` sur une entrée POST utilisateur, ce qui correspond à CWE-94 (Improper Control of Generation of Code) : l'attaquant contrôle totalement l'interpréteur Python. La version corrigée remplace l'évaluation dynamique par un mapping déclaratif `ALLOWED_OPERATIONS`, où l'entrée utilisateur ne sert qu'à sélectionner une opération dans une liste blanche, jamais à fournir du code exécutable.
