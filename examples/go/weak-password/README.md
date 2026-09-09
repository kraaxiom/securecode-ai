# weak-password (CWE-521)

La version vulnerable n'impose aucune exigence de robustesse sur le mot de passe choisi a l'inscription, permettant l'utilisation de mots de passe triviaux. La version corrigee applique une longueur minimale de 12 caracteres (conforme aux recommandations NIST SP 800-63B) et rejette les mots de passe presents dans des listes de fuites connues, plutot que d'imposer des regles de complexite arbitraires peu efficaces. C'est la remediation standard recommandee contre CWE-521.
