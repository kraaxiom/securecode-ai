---
id: nosql-injection
category: injections
cwe: CWE-943
owasp: A03:2021-Injection
severity_default: high
languages: [js, python, php, java, csharp, go]
---

# NoSQL Injection

## Description
L'injection NoSQL se produit lorsqu'une entrée utilisateur non neutralisée est intégrée dans une requête destinée à une base de données NoSQL (MongoDB, CouchDB, etc.), permettant à un attaquant de modifier la structure logique de la requête plutôt que sa syntaxe textuelle. Contrairement au SQL, l'injection passe souvent par l'injection d'objets ou d'opérateurs (ex: `$where`, `$ne`, `$gt`) plutôt que par des caractères d'échappement, ce qui peut permettre un contournement d'authentification ou une extraction de données.

## Où ça apparaît typiquement
- Endpoints d'API acceptant du JSON brut passé directement comme critère de requête à une base NoSQL sans validation de schéma.
- Formulaires de connexion où les champs utilisateur/mot de passe sont passés tels quels comme objet de requête.
- Utilisation d'opérateurs de requête natifs (`$where`, `$regex`) construits à partir d'entrées utilisateur.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Passage direct du corps de requête HTTP (`req.body`) comme filtre de requête à la base de données, sans validation de type/schéma.
- Absence de restriction empêchant l'utilisateur d'injecter des objets là où une valeur scalaire est attendue.
- Utilisation de `$where` ou d'évaluation de code côté base de données avec des données non fiables.

## Remédiation
- Valider strictement le schéma et le type de chaque champ attendu (rejeter tout objet/opérateur là où une valeur scalaire est attendue).
- Utiliser les mécanismes de requête paramétrée/typée fournis par le driver plutôt que de transmettre le payload brut.
- Désactiver ou éviter les opérateurs d'évaluation de code côté serveur (`$where`) avec des données non fiables.
- Voir `rules/remediation/nosql-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/nodejs-express/nosql-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: NoSQL Injection Prevention (OWASP Testing Guide, section NoSQL)
- CWE-943: Improper Neutralization of Special Elements in Data Query Logic
