# session-prediction (CWE-330)

La version vulnerable derive l'identifiant de session d'un compteur sequentiel et d'un horodatage, des valeurs previsibles permettant a un attaquant de deviner ou d'enumerer des sessions valides d'autres utilisateurs. La version corrigee utilise `crypto/rand` pour generer un identifiant avec 256 bits d'entropie cryptographiquement sure, rendant toute prediction infaisable en pratique. C'est la remediation standard recommandee contre CWE-330.
