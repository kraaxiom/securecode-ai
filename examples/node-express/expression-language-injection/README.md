## Expression Language Injection (CWE-917)

Le code vulnérable construit la source du template EJS par concaténation directe de l'entrée utilisateur, ce qui permet d'injecter des balises `<% %>` exécutées comme du code JavaScript lors du rendu. La correction fige la source du template comme une chaîne statique et transmet la donnée utilisateur uniquement comme variable de contexte via `<%= %>`, qui l'échappe automatiquement.
