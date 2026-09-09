# brute-force (CWE-307)

La version vulnerable n'impose aucune limite sur le nombre de tentatives d'authentification, permettant a un attaquant de tester un nombre illimite de mots de passe. La version corrigee ajoute une limitation de debit (rate limiting) par compte et par IP avec verrouillage progressif, journalise les echecs et renvoie un message d'erreur generique qui ne revele pas l'existence du compte. C'est la remediation standard recommandee contre CWE-307.
