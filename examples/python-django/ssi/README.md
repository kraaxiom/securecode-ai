# Server-Side Includes Injection — CWE-97

Le code vulnérable écrit un commentaire utilisateur non filtré dans un fichier `.shtml` interprété par SSI, ce qui permet d'injecter une directive telle que `<!--#exec-->` exécutée côté serveur. La correction échappe le HTML et neutralise toute séquence `<!--#` avant écriture ; la recommandation à terme reste de générer ce contenu via le moteur de templates Django plutôt que via un fichier SSI. Voir `rules/remediation/ssi.md`.
