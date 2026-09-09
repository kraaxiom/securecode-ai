# CSV Injection (CWE-1236)

La version vulnérable écrit les champs `Name` et `Comment` directement dans le fichier CSV exporté, permettant à un attaquant d'insérer une formule (ex: `=HYPERLINK(...)`) exécutée à l'ouverture par le tableur de la victime. La correction ajoute une fonction `SanitizeCsvCell` qui préfixe d'une apostrophe toute valeur commençant par `=`, `+`, `-`, `@`, tabulation ou retour chariot, forçant son interprétation comme texte brut.
