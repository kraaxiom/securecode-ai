# ssi (CWE-97)

La version vulnerable ecrit le commentaire utilisateur tel quel dans un fichier `.shtml` servi par un serveur avec SSI active : une valeur contenant `<!--#exec cmd="..." -->` est executee cote serveur a chaque affichage. La version corrigee echappe le HTML et neutralise explicitement la sequence `<!--#`, tout en recommandant en commentaire de ne jamais ecrire de contenu utilisateur dans un fichier interprete par SSI et d'utiliser un moteur de templates applicatif a la place.
