# credential-stuffing (CWE-307)

La version vulnerable ne detecte aucun pattern anormal de tentatives d'authentification (multiples comptes testes depuis une meme source), ce qui permet la reutilisation automatisee d'identifiants voles sur d'autres services. La version corrigee detecte la vitesse anormale de connexion par IP, rejette les mots de passe connus comme compromis, et impose un defi CAPTCHA/MFA additionnel en cas d'anomalie. C'est la remediation standard recommandee contre CWE-307.
