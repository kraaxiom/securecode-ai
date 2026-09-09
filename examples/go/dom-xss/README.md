# dom-xss (CWE-79)

La version vulnerable sert une page dont le script inline lit `location.search` cote client et insere la valeur via `innerHTML`, executant tout HTML/script injecte par l'URL sans jamais transiter par le serveur. La version corrigee remplace `innerHTML` par `textContent` dans le script servi et ajoute un en-tete `Content-Security-Policy` restreignant les sources de script, en defense en profondeur.
