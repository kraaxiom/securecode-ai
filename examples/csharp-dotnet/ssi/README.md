# Server-Side Includes (SSI) Injection (CWE-97)

La version vulnérable écrit le commentaire utilisateur tel quel dans un fichier `.shtml` interprété par le serveur web, permettant à un attaquant d'injecter une directive SSI (`<!--#exec cmd="..." -->`) exécutée côté serveur. La correction échappe le HTML via `WebUtility.HtmlEncode` puis neutralise explicitement toute séquence `<!--#` résiduelle avant l'écriture ; l'approche préférable reste de ne jamais écrire de contenu utilisateur dans un fichier interprété par SSI et d'utiliser un moteur de templates applicatif.
