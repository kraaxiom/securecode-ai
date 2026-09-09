# default-credentials (CWE-1392)

La version vulnerable cree le compte administrateur avec un identifiant et un mot de passe fixes et connus (`admin`/`admin123`), jamais invalides apres l'installation. La version corrigee genere un mot de passe temporaire aleatoire unique, force son changement des la premiere connexion (`must_change_password`) et ne le stocke jamais en clair. C'est la remediation standard recommandee contre CWE-1392.
