# blind-sql-injection (CWE-89)

La version vulnerable concatene l'entree utilisateur dans une clause booleenne, permettant a un attaquant d'inferer des donnees par injections vrai/faux repetees. La version corrigee valide l'entree comme entier puis utilise une requete parametree, empechant toute alteration de la logique SQL.
