# csv-injection (CWE-1236)

La version vulnerable ecrit les champs utilisateur dans un export CSV sans verification. Si un champ commence par `=`, `+`, `-` ou `@`, le tableur qui ouvre le fichier peut l'interpreter comme une formule active (execution de commande via DDE, exfiltration de donnees). La version corrigee detecte ces caracteres declencheurs en debut de champ et prefixe le champ avec une apostrophe pour forcer une interpretation en texte brut.
