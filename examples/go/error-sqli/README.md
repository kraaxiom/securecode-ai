# error-sqli (CWE-89)

La version vulnerable concatene l'entree dans la requete et renvoie le message d'erreur SQL brut au client, ce qui permet une exfiltration via les erreurs de base de donnees. La version corrigee utilise une requete parametree et ne renvoie qu'un message generique, le detail technique restant dans les logs serveur.
