# xpath-injection (CWE-643)

La version vulnerable construit l'expression XPath d'authentification par concatenation directe de l'entree utilisateur : une valeur comme `' or '1'='1` dans le mot de passe permet de contourner la verification. La version corrigee convertit chaque valeur en litteral XPath sur (via `concat()` quand la valeur contient une apostrophe), et recommande en commentaire de ne jamais utiliser XPath comme mecanisme d'authentification.
