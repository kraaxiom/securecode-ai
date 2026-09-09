# Formula Injection

La version vulnérable exporte `ticket.title` et `ticket.description` sans neutralisation dans un fichier CSV destiné à être ouvert dans un tableur, ce qui correspond à CWE-1236 : une valeur commençant par `=`, `+`, `-` ou `@` peut déclencher l'exécution d'une formule (exfiltration via `WEBSERVICE`, DDE). La version corrigée introduit `neutraliser_formule()`, qui préfixe d'une apostrophe toute cellule à risque avant écriture.
