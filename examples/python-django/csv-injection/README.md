# CSV Injection (Formula Injection)

La version vulnérable écrit les champs utilisateur `author_name` et `comment` directement dans le CSV exporté, ce qui correspond à CWE-1236 : si un champ commence par `=`, `+`, `-` ou `@`, le tableur peut l'interpréter comme une formule active à l'ouverture. La version corrigée ajoute `sanitize_csv_cell()`, qui préfixe d'une apostrophe toute cellule commençant par un caractère déclencheur, neutralisant l'interprétation en formule.
