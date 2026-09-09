# os-command-injection (CWE-78)

La version vulnerable interpole l'entree utilisateur dans une commande executee via `sh -c`, permettant l'injection de metacaracteres shell. La version corrigee execute le binaire directement avec des arguments separes (sans interpretation shell) et valide l'entree avec une allowlist regex avant usage.
