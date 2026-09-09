## Server-Side Template Injection - SSTI (CWE-1336)

Le code vulnérable concatène l'entrée utilisateur directement dans la chaîne de template avant compilation par Handlebars, au lieu de la passer comme variable de contexte. Le moteur interprète alors la syntaxe injectée comme du code de template légitime, pouvant mener à une divulgation d'information voire une exécution de code arbitraire. La correction sépare strictement le template statique (compilé une seule fois) des données utilisateur, transmises uniquement via le contexte de rendu.
