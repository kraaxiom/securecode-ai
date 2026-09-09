# http-response-splitting (CWE-113)

La version vulnerable construit manuellement un en-tete `Set-Cookie` avec une valeur utilisateur non filtree, permettant l'injection de CR/LF pour scinder la reponse HTTP. La version corrigee valide l'entree par regex puis utilise `http.SetCookie`, qui echappe correctement la valeur et fixe l'attribut `HttpOnly`.
