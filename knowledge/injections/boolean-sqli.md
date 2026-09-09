---
id: boolean-sqli
category: injections
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Boolean-based SQL Injection

## Description
L'injection SQL booléenne est une sous-catégorie de l'injection aveugle où l'attaquant modifie la valeur de vérité d'une condition SQL (par exemple dans une clause `WHERE`) pour observer une différence binaire dans la réponse de l'application (contenu affiché, redirection, code HTTP). Elle permet d'extraire des données bit par bit ou caractère par caractère en posant une série de questions vrai/faux à la base de données.

## Où ça apparaît typiquement
- Formulaires de connexion ou de récupération de mot de passe où une entrée est comparée directement dans une requête SQL.
- Filtres de recherche dont le résultat conditionne l'affichage d'un bloc de contenu ("aucun résultat" vs liste de résultats).
- Paramètres d'URL utilisés tels quels dans une clause `WHERE id = ...` ou `WHERE name = '...'`.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Comparaison SQL construite par concaténation de chaîne avec une entrée utilisateur.
- Différence de comportement applicatif pilotée directement par le résultat brut d'une requête SQL non paramétrée.
- Absence de validation de type (ex: un identifiant attendu numérique accepté en tant que chaîne libre).

## Remédiation
- Utiliser des requêtes préparées avec liaison de paramètres pour toute condition impliquant une donnée utilisateur.
- Valider strictement le format et le type des entrées avant toute utilisation dans une requête.
- Limiter le nombre de tentatives et journaliser les motifs de requêtes anormaux (WAF applicatif, rate limiting).
- Voir `rules/remediation/boolean-sqli.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/boolean-sqli/`.

## Références
- OWASP Cheat Sheet: SQL Injection Prevention
- CWE-89: Improper Neutralization of Special Elements used in an SQL Command
