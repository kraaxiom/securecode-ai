# http-parameter-pollution (CWE-235)

La version vulnerable lit implicitement la premiere valeur d'un parametre pouvant etre duplique, alors que d'autres composants de la chaine (WAF, proxy, backend) peuvent interpreter une autre occurrence, causant une desynchronisation exploitable. La version corrigee rejette explicitement toute requete contenant plusieurs occurrences du meme parametre.
