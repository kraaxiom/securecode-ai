# Password Spraying — Node/Express

`vulnerable.js` ne limite les échecs d'authentification que par compte, sans jamais agréger le volume de tentatives par IP/source (CWE-307, Improper Restriction of Excessive Authentication Attempts). Un attaquant testant un mot de passe courant contre des milliers de comptes distincts échappe donc totalement à cette limitation.

`fixed.js` ajoute une agrégation globale par IP (`getFailedAttemptsByIp`) qui déclenche un blocage et une alerte de sécurité (`flagSuspiciousSource`) au-delà d'un seuil, en complément de la limite existante par compte.
