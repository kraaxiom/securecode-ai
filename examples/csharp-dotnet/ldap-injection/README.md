# LDAP Injection (CWE-90)

La version vulnérable construit le filtre de recherche LDAP `(uid={uid})` par concaténation directe de l'entrée utilisateur, ce qui permet d'injecter des méta-caractères LDAP (`*`, `(`, `)`, `\`) pour altérer la logique du filtre et, par exemple, contourner une restriction de recherche. La correction valide d'abord le format attendu de l'identifiant (alphanumérique), puis échappe explicitement les caractères spéciaux du filtre selon la RFC 4515 avant insertion, empêchant toute modification de la structure du filtre.
