# IMAP Injection (CWE-93)

Le code vulnérable concatène le terme de recherche utilisateur dans une chaîne `SUBJECT "..."` avant de l'utiliser comme critère de recherche, ce qui permet d'injecter des guillemets ou des mots-clés IMAP supplémentaires. La correction supprime les guillemets et caractères de contrôle, limite la longueur du terme, et s'appuie sur l'API structurée `SearchQuery.SubjectContains` de MailKit qui encode correctement le littéral sans jamais construire de commande IMAP brute.
