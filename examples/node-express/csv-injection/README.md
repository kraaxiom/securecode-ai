## CSV Injection (CWE-1236)

Le code vulnérable écrit les champs utilisateur directement dans un export CSV sans vérifier leur premier caractère, ce qui permet d'y injecter une formule active (`=`, `+`, `-`, `@`) exécutée à l'ouverture dans un tableur. La correction ajoute une fonction `sanitizeCsvCell` qui préfixe d'une apostrophe toute cellule commençant par un caractère déclencheur, et force le téléchargement via `Content-Disposition` plutôt qu'une ouverture directe non maîtrisée.
