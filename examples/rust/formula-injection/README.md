# Formula Injection (CWE-1236)

Le code vulnérable exporte les champs utilisateur bruts dans un rapport CSV sans vérifier s'ils commencent par un caractère déclencheur de formule, exposant les utilisateurs du tableur à une exécution de formule à l'ouverture du fichier. La correction ajoute la fonction `neutraliser_formule`, qui préfixe d'une apostrophe toute cellule commençant par `=`, `+`, `-`, `@`, tabulation ou retour chariot, en application directe de la remédiation CWE-1236.
