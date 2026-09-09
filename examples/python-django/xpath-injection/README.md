# XPath Injection — CWE-643

Le code vulnérable authentifie l'utilisateur en concaténant `user` et `password` dans une expression XPath, ce qui permet un contournement d'authentification via des opérateurs XPath injectés. La correction utilise les variables XPath liées de `lxml` (`$u`, `$p`), séparant la structure de l'expression des valeurs utilisateur ; il est aussi recommandé de ne pas utiliser XPath comme mécanisme d'authentification. Voir `rules/remediation/xpath-injection.md` pour d'autres langages.
