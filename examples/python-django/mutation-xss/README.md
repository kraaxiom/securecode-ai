# Mutation XSS (CWE-79)

Le code vulnérable implémentait sa propre sanitisation HTML avec `lxml`, en supprimant seulement quelques balises connues, une approche fragile face aux quirks de reparsing du navigateur qui peuvent « muter » un balisage apparemment neutralisé en balisage exécutable. Le correctif remplace cette logique maison par `bleach`, une bibliothèque de sanitisation activement maintenue et testée contre les vecteurs de mutation connus, avec une liste blanche explicite de balises et d'attributs.
