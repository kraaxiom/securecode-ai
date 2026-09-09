## IMAP Injection (CWE-93)

Le code vulnérable construit une commande IMAP brute par concaténation du terme de recherche utilisateur, permettant l'injection de guillemets ou de mots-clés de commande IMAP supplémentaires. La correction abandonne la construction de commande texte au profit de l'API structurée du client IMAP (`client.search({ subject })`), qui encode nativement les littéraux, avec en complément une limite de longueur sur l'entrée.
