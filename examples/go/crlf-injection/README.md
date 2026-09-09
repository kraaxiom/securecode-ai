# crlf-injection (CWE-93)

La version vulnerable place l'entree utilisateur brute dans l'en-tete `Location`, permettant l'injection de sequences CRLF pour ajouter des en-tetes ou scinder la reponse. La version corrigee rejette tout caractere de controle et restreint la redirection aux chemins relatifs internes.
