# Code Injection (CWE-94)

Le code vulnérable transmet directement le paramètre `formula` fourni par l'utilisateur à un moteur de script (`ScriptEngine.eval`), permettant l'exécution de code arbitraire sur le serveur. La version corrigée supprime toute évaluation dynamique et remplace le mécanisme par une liste blanche déclarative d'opérations autorisées (`add`, `sub`, `mul`, `div`), sélectionnée uniquement par nom, avec rejet explicite de toute valeur non prévue.
