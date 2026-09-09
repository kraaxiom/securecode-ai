# password-spray (CWE-307)

La version vulnerable ne limite les tentatives d'authentification que par compte, ce qui laisse un attaquant tester un seul mot de passe courant sur un grand nombre de comptes distincts depuis la meme source sans jamais declencher de blocage. La version corrigee ajoute une limitation globale par IP/source en complement de la limitation par compte, capable de detecter ce pattern reparti caracteristique du password spraying. C'est la remediation standard recommandee contre CWE-307.
