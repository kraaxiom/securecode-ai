## XPath Injection (CWE-643)

Le code vulnérable concatène directement les champs `user` et `pass` dans une expression XPath utilisée pour authentifier l'utilisateur contre un document XML, permettant à un attaquant de modifier la logique de sélection de nœuds et de contourner l'authentification. La correction transforme chaque valeur en littéral XPath sûr (fonction `xpathLiteral` gérant les apostrophes via `concat()`) avant insertion dans l'expression ; à terme, XPath ne devrait pas servir de mécanisme d'authentification.
