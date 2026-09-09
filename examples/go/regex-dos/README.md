# regex-dos (CWE-1333)

La version vulnerable applique une regex a groupes repetes imbriques sur une entree de taille non bornee : meme avec le moteur RE2 de Go (garanti lineaire, sans backtracking catastrophique), l'absence de limite de taille permet a un attaquant d'envoyer une entree massive pour consommer du CPU et de la memoire de facon disproportionnee. La version corrigee impose une limite stricte de longueur avant tout traitement et simplifie le motif pour supprimer les groupes repetes imbriques inutiles.
