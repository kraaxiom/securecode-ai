# Formula Injection (CWE-1236)

Le code vulnérable écrit `ticket.Title` et `ticket.Comment` directement dans les cellules d'un classeur ClosedXML exporté, permettant à un attaquant d'y insérer une formule active exécutée à l'ouverture par la victime. La correction ajoute une fonction `NeutraliserFormule` qui préfixe d'une apostrophe toute valeur commençant par `=`, `+`, `-`, `@`, tabulation ou retour chariot, neutralisant l'interprétation comme formule tout en conservant le texte affiché.
