## Blind SQL Injection (CWE-89)

Le code vulnérable concatène directement le paramètre `user` dans une requête SQL, ce qui permet à un attaquant de manipuler la condition `WHERE` et d'inférer des informations via le comportement de l'application (booléen `exists`), même sans affichage direct des données. La correction remplace la concaténation par une requête préparée avec paramètre lié (`?`), empêchant toute interprétation de l'entrée comme du code SQL, et uniformise le message d'erreur renvoyé au client pour limiter les canaux d'inférence.
