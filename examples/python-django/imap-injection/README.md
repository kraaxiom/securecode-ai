# IMAP Injection

La version vulnérable insère le critère de recherche `q` par f-string dans la commande IMAP `SEARCH SUBJECT "..."`, ce qui correspond à CWE-93 : un guillemet ou une séquence `\r\n` dans l'entrée permet d'injecter des commandes IMAP supplémentaires. La version corrigée neutralise les guillemets et caractères de contrôle, limite la longueur du critère, et le transmet comme argument séparé à `mail.search()` plutôt que concaténé dans une chaîne de commande.
