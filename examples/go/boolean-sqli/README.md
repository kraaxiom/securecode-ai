# boolean-sqli (CWE-89)

La version vulnerable insere l'entree utilisateur directement dans une clause LIKE, ce qui autorise l'injection de conditions booleennes arbitraires. La version corrigee construit le motif `%...%` en Go puis le transmet comme parametre lie a une requete preparee, eliminant l'injection.
