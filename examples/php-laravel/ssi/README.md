# Server-Side Includes Injection (CWE-97)

Le code vulnérable écrit le commentaire utilisateur sans filtrage dans un fichier `.shtml` interprété par le serveur web, permettant l'injection d'une directive SSI comme `<!--#exec -->`. La correction échappe le contenu en HTML puis neutralise explicitement toute séquence `<!--#` restante avant l'écriture, empêchant l'interprétation d'une directive SSI (CWE-97 : Improper Neutralization of Server-Side Includes (SSI) Within a Web Page).
