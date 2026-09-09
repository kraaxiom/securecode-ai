# Expression Language Injection (CWE-917)

Le code vulnérable concatène l'entrée utilisateur directement dans la source du template avant compilation, permettant l'injection de directives de template évaluées côté serveur. La correction fixe le template comme chaîne statique (`"Bonjour {{ nom }}"`) et transmet la valeur utilisateur uniquement via le contexte de rendu (`context.insert`), qui l'échappe automatiquement, conformément à la remédiation CWE-917.
