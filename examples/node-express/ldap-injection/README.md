## LDAP Injection (CWE-90)

Le code vulnérable concatène le paramètre `uid` directement dans un filtre de recherche LDAP (`ldapjs`), ce qui permet à un attaquant d'injecter des méta-caractères de filtre (`* ( ) \` ou NUL) pour altérer la logique de la recherche, contourner une authentification ou extraire des informations non autorisées de l'annuaire. La correction échappe la valeur utilisateur selon les règles d'échappement de filtre LDAP (RFC 4515) et valide son format avant construction du filtre.
