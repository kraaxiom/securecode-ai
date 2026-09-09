# jwt-algorithm-confusion (CWE-347)

La version vulnerable verifie la signature du JWT en faisant confiance a l'algorithme annonce par le token lui-meme, permettant a un attaquant de basculer de RS256 vers HS256 et de signer avec la cle publique RSA (connue) utilisee comme secret HMAC. La version corrigee force explicitement l'algorithme attendu (`jwt.WithValidMethods`) et rejette toute signature ne correspondant pas au type de cle attendu. C'est la remediation standard recommandee contre CWE-347.
