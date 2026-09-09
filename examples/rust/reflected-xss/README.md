# Reflected XSS

Le handler `search` réinjectait le paramètre de requête `q` directement dans le HTML de réponse sans encodage, permettant l'exécution de script via un lien piégé (CWE-79). La version corrigée applique un échappement HTML systématique à la valeur avant insertion, neutralisant toute tentative de rupture du contexte HTML tout en conservant l'affichage du terme de recherche légitime.
