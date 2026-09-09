# Server-Side Includes Injection (CWE-97)

Le code vulnérable écrit le commentaire utilisateur tel quel dans un fichier `.shtml` servi par un serveur web avec SSI activé, ce qui permet à un attaquant d'injecter une directive `<!--#exec cmd="..."-->` exécutée côté serveur à l'affichage de la page. La version corrigée échappe le contenu HTML puis neutralise toute séquence `<!--#` résiduelle avant écriture, et recommande de désactiver SSI (en particulier la directive `exec`) sur le répertoire concerné au profit d'un moteur de templates applicatif comme Thymeleaf.
