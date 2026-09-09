# bola (CWE-639)

La version vulnerable recupere l'objet directement a partir de l'identifiant fourni par le client, sans verifier que l'utilisateur authentifie en est le proprietaire, permettant d'acceder aux ressources d'autrui en changeant simplement l'ID (Broken Object Level Authorization). La version corrigee restreint la requete au proprietaire authentifie (`getInvoiceByIDForOwner`), garantissant qu'aucune ressource appartenant a un autre utilisateur ne peut etre retournee. C'est la remediation standard recommandee contre CWE-639.
