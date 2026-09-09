# LDAP Injection (CWE-90)

Le code vulnérable insère l'identifiant utilisateur directement dans le filtre LDAP `(uid=...)` par concaténation, permettant à un attaquant de modifier la logique du filtre via les métacaractères `* ( ) \`. La correction ajoute une fonction `escape_ldap_filter` qui échappe chacun de ces caractères spéciaux selon la notation hexadécimale LDAP avant construction du filtre, conformément à la remédiation CWE-90.
