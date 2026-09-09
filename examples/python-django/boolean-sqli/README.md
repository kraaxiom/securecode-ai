# Boolean-based SQL Injection

La version vulnérable construit une requête brute via `Product.objects.raw()` en insérant `name` par f-string, ce qui correspond à CWE-89 : un attaquant peut altérer la valeur de vérité de la clause `WHERE` pour changer le nombre de résultats retournés. La version corrigée remplace la requête brute par `Product.objects.filter(name=name)`, qui s'appuie sur l'ORM Django et une requête paramétrée en interne, empêchant toute injection logique.
