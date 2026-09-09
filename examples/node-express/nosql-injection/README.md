## NoSQL Injection (CWE-943)

Le code vulnérable transmet directement `req.body` comme filtre de requête à MongoDB via Mongoose (`findOne({ username, password })`). Un attaquant peut envoyer `{"username": {"$ne": null}, "password": {"$ne": null}}` au lieu de chaînes, ce qui modifie la structure logique de la requête et permet un contournement d'authentification. La correction valide strictement le type (`string`) de chaque champ avant de l'insérer dans le filtre, rejetant tout objet ou opérateur injecté.
