# time-based-sqli (CWE-89)

La version vulnerable concatene l'entree sans limite d'execution, exposant l'application aux injections a base de delai (SLEEP/WAITFOR) pour de l'exfiltration a l'aveugle. La version corrigee utilise une requete parametree et applique un timeout de contexte strict, neutralisant l'injection et bornant l'impact d'un delai malicieux.
