# CSV Injection (CWE-1236)

Le code vulnérable écrit les valeurs utilisateur telles quelles dans l'export CSV, sans échappement, ce qui permet à un attaquant d'injecter une formule (par exemple `=cmd|'/c calc'!A1`) qui sera exécutée par le tableur du destinataire à l'ouverture du fichier. La version corrigée préfixe systématiquement d'une apostrophe toute cellule commençant par `=`, `+`, `-`, `@`, une tabulation ou un retour chariot, neutralisant son interprétation comme formule tout en préservant le contenu affiché.
