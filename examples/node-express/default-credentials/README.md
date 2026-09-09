# Default Credentials — Node/Express

`vulnerable.js` code en dur un identifiant et un mot de passe administrateur par défaut (`admin`/`admin`), jamais renouvelés (CWE-1392, Use of Default Credentials). Ces valeurs étant publiquement devinables, un attaquant obtient un accès administrateur immédiat sans avoir besoin de casser quoi que ce soit.

`fixed.js` génère un mot de passe temporaire aléatoire à chaque provisioning via `crypto.randomBytes`, le stocke haché avec argon2, et impose un flag `mustChangePassword` bloquant l'accès tant que le mot de passe temporaire n'a pas été changé.
