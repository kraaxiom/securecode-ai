## SSRF aveugle (CWE-918)

Le code vulnérable déclenche une requête sortante vers une URL de callback fournie par l'utilisateur dans un traitement asynchrone dont le résultat n'est jamais renvoyé au client, ce qui ne supprime pas le risque puisqu'aucune whitelist ni validation de destination n'est appliquée. La correction applique les mêmes contrôles qu'une SSRF classique (whitelist de domaines, résolution DNS avec rejet des IP privées, timeout) et journalise les tentatives rejetées pour permettre la détection d'un scan interne.
