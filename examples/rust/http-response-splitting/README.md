# HTTP Response Splitting (CWE-113)

Le code vulnérable place directement le paramètre `next` dans l'en-tête `Location`, ce qui permettrait l'injection de CR/LF pour fractionner la réponse HTTP si aucun filtrage bas niveau n'est appliqué. La correction restreint la redirection à une liste blanche de chemins autorisés (`HashSet`), rejetant toute valeur hors de cet ensemble au profit d'une destination par défaut, éliminant tout vecteur d'injection selon CWE-113.
