# sql-injection (CWE-89)

La version vulnerable construit la requete SQL par concatenation de chaines avec l'entree utilisateur (`username`), permettant l'injection de SQL arbitraire. La version corrigee utilise une requete parametree (`?` lie), ce qui force le pilote SQL a traiter la valeur comme une donnee et non comme du code. C'est la remediation standard recommandee contre CWE-89.
