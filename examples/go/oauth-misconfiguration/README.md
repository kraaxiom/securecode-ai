# oauth-misconfiguration (CWE-287)

La version vulnerable accepte n'importe quel `redirect_uri` fourni par le client et ne verifie aucun parametre `state`, exposant le flux OAuth au vol de code d'autorisation et au CSRF sur le login. La version corrigee valide `redirect_uri` contre une liste blanche stricte, verifie le `state` genere serveur avant l'echange du code, et securise le cookie de session (`HttpOnly`, `Secure`, `SameSite`). C'est la remediation standard recommandee contre CWE-287.
