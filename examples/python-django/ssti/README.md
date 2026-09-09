# Server-Side Template Injection — CWE-1336

Le code vulnérable construit le texte du template en concaténant l'entrée utilisateur avant compilation (`engine.from_string("Bonjour " + name)`), ce qui permet à un attaquant d'injecter de la syntaxe de template interprétée par le moteur. La correction fixe le template comme chaîne statique versionnée (`"Bonjour {{ name }}"`) et passe la donnée utilisateur uniquement via le contexte de rendu. Voir `rules/remediation/ssti.md` pour d'autres langages.
