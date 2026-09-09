# ldap-injection (CWE-90)

La version vulnerable construit le filtre de recherche LDAP par concatenation directe de l'entree utilisateur (`fmt.Sprintf("(uid=%s)", uid)`), permettant de modifier la logique du filtre pour contourner l'authentification ou enumerer des entrees. La version corrigee passe systematiquement l'entree utilisateur par `ldap.EscapeFilter` (go-ldap), qui echappe les metacaracteres du langage de filtre LDAP definis par la RFC 4515.
