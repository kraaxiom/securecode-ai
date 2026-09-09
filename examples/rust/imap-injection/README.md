# IMAP Injection (CWE-93)

Le code vulnérable concatène le critère de recherche utilisateur dans une commande IMAP texte sans échapper les guillemets ni les caractères de contrôle, permettant l'injection de commandes IMAP supplémentaires. La correction introduit `sanitize_imap_criteria`, qui retire les guillemets, `\r`, `\n` et borne la longueur du critère avant assemblage, réduisant la surface d'injection définie par CWE-93 (les intégrations réelles devraient en outre privilégier l'API structurée du client IMAP).
