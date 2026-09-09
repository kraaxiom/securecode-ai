## Code Injection (CWE-94)

Le code vulnérable transmet directement la formule fournie par le client à `eval()`, permettant l'exécution de code JavaScript arbitraire côté serveur. La correction supprime toute évaluation dynamique et la remplace par une liste blanche de fonctions autorisées (`ALLOWED_OPERATIONS`), le client ne pouvant plus fournir que le nom d'une opération prédéfinie ainsi que des opérandes numériques validés.
