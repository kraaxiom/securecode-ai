# sqli-union (CWE-89)

La version vulnerable concatene l'entree utilisateur dans une clause `WHERE`, permettant l'injection d'une clause `UNION SELECT` pour extraire des donnees d'autres tables. La version corrigee utilise une requete parametree, ce qui empeche toute modification de la structure de la requete SQL d'origine.
