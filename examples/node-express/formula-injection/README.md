## Formula Injection (CWE-1236)

Le code vulnérable écrit les champs `nom` et `commentaire` fournis par l'utilisateur dans un fichier CSV exporté sans vérifier leur premier caractère, permettant l'insertion d'une formule active interprétée par le tableur à l'ouverture. La correction ajoute la fonction `neutraliserFormule`, qui préfixe d'une apostrophe toute valeur commençant par `=`, `+`, `-`, `@`, tabulation ou retour chariot avant écriture.
