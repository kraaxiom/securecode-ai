# command-injection (CWE-78)

La version vulnerable concatene un nom de fichier utilisateur dans une commande shell (`sh -c`), permettant l'injection de commandes arbitraires. La version corrigee execute `tar` directement sans shell, restreint le chemin au repertoire d'upload via `filepath.Base`/`Join` et verifie le prefixe resultant.
