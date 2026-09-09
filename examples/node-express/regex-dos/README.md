## ReDoS - Regular Expression Denial of Service (CWE-1333)

Le code vulnérable valide un email avec la regex `/^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$/`, dont les quantificateurs imbriqués provoquent une complexité de correspondance exponentielle sur certaines entrées, permettant à un attaquant de bloquer le thread Node.js avec une seule requête. La correction utilise une regex non ambiguë (sans groupe répété imbriqué) et impose une limite de longueur sur l'entrée avant toute application de l'expression régulière.
