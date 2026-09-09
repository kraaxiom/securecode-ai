# idor (CWE-639)

La version vulnerable renvoie le document identifie par l'ID fourni dans l'URL sans jamais verifier que l'utilisateur authentifie a le droit d'y acceder, permettant de lire les documents d'autrui en incrementant simplement l'identifiant (Insecure Direct Object Reference). La version corrigee ajoute une verification explicite du droit d'acces (propriete ou partage) avant de renvoyer le contenu, et repond de facon identique (404) que le document soit inexistant ou inaccessible, pour ne pas confirmer son existence. C'est la remediation standard recommandee contre CWE-639.
