# formula-injection (CWE-1236)

La version vulnerable ecrit un montant de depense saisi par l'utilisateur dans un rapport tableur sans validation ni neutralisation, permettant l'injection d'une formule active (`=HYPERLINK(...)`, `=cmd|...`) executee a l'ouverture du fichier par la victime. La version corrigee impose un format numerique strict pour le montant et neutralise, en defense en profondeur, tout champ texte commencant par un caractere declencheur de formule (`= + - @`).
