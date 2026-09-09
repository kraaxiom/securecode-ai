# Code Injection (CWE-94)

Le code vulnérable utilise `CSharpScript.EvaluateAsync` pour exécuter dynamiquement une expression fournie par le client, donnant à un attaquant un contrôle total sur le code exécuté côté serveur. La correction remplace ce mécanisme par un dictionnaire de fonctions autorisées (`add`, `sub`, `mul`, `div`) sélectionnées via une liste blanche, sans jamais interpréter de code arbitraire. Toute opération non reconnue est rejetée explicitement avec un code 400.
