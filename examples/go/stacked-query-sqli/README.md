# stacked-query-sqli (CWE-89)

La version vulnerable concatene deux valeurs utilisateur dans une requete `UPDATE`, autorisant l'empilement de requetes supplementaires via un point-virgule. La version corrigee utilise des parametres lies pour chaque valeur, ce qui empeche toute instruction SQL additionnelle d'etre executee.
