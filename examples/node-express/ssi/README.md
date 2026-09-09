## Server-Side Includes (SSI) Injection (CWE-97)

Le code vulnérable écrit le commentaire utilisateur tel quel dans un fichier `.shtml` servi par un serveur web avec SSI activé, permettant à un attaquant d'injecter une directive `<!--#exec cmd="..." -->` interprétée côté serveur et menant potentiellement à une exécution de commande. La correction échappe le contenu HTML puis neutralise toute séquence `<!--#` résiduelle avant écriture ; à terme, ce contenu devrait être rendu par un moteur de templates applicatif plutôt que via SSI.
