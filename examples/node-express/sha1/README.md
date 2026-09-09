## Utilisation de SHA-1 (CWE-328)

Le code vulnérable hache le mot de passe utilisateur avec `crypto.createHash('sha1')`, une fonction dont la résistance aux collisions est cassée depuis la démonstration pratique de l'attaque SHAttered en 2017, et qui n'offre par ailleurs aucun sel ni coût adaptatif. La correction remplace SHA-1 par Argon2id, algorithme dédié au hachage de mots de passe avec sel intégré et coût configurable. Pour des usages non liés aux mots de passe (intégrité de fichier, signature), SHA-256 ou supérieur doit être utilisé à la place de SHA-1.
