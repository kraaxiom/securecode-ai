# XPath Injection (CWE-643)

Le code vulnérable concatène directement les champs `user` et `pass` dans une expression XPath utilisée pour l'authentification, permettant à un attaquant d'injecter une apostrophe et un opérateur XPath pour contourner la vérification. La correction utilise `DOMXPath::quote()` pour échapper strictement chaque valeur avant insertion dans l'expression, neutralisant l'injection (CWE-643 : Improper Neutralization of Data within XPath Expressions).
