# Injection NoSQL (CWE-943)

Le code vulnérable insère directement les valeurs du corps JSON reçu (`Map<String, Object>`) dans le filtre de requête MongoDB, ce qui permet à un attaquant d'envoyer un objet contenant un opérateur MongoDB (par exemple `{"$ne": null}`) à la place d'une chaîne attendue et ainsi de contourner la vérification d'identifiants. La version corrigée valide strictement que `username` et `password` sont bien de type `String` avant de construire le filtre, rejetant tout objet ou opérateur qui altérerait la logique de la requête.
