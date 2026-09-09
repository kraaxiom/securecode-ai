# imap-injection (CWE-93)

La version vulnerable insere le terme de recherche utilisateur par concatenation directe dans une commande IMAP brute : une entree contenant CRLF permet d'injecter une commande IMAP arbitraire supplementaire. La version corrigee rejette toute entree contenant CR/LF et transmet la valeur via le mecanisme de "literal" IMAP (`{N}` suivi de N octets), qui empeche toute reinterpretation du contenu comme une nouvelle commande.
