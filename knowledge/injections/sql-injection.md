---
id: sql-injection
category: injections
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# SQL Injection (générique)

## Description
L'injection SQL est une vulnérabilité qui survient lorsqu'une entrée utilisateur non neutralisée est intégrée directement dans une requête SQL, permettant à un attaquant d'altérer la logique de la requête exécutée par la base de données. Selon le contexte, elle peut permettre la lecture, la modification ou la suppression de données, le contournement d'authentification, voire l'exécution de commandes sur le système hôte. Cette fiche couvre le cas générique; voir les fiches spécialisées (UNION, blind, time-based, stacked query, error-based) pour les variantes techniques.

## Où ça apparaît typiquement
- Concaténation de chaînes pour construire une requête SQL à partir de paramètres HTTP (query string, body, cookies, headers).
- Requêtes "brutes" d'un ORM (`raw()`, `query()`, `createNativeQuery`) sans liaison de paramètres.
- Requêtes dynamiques pour filtres/tris/pagination où le nom de colonne ou la valeur vient de l'utilisateur.
- Procédures stockées construites dynamiquement par concaténation de chaînes côté base de données.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation ou interpolation de variable non échappée dans une chaîne SQL.
- Appel "raw query" d'un ORM avec une valeur non liée en paramètre.
- Absence de couche de validation de type/format avant la requête (ex: un ID censé être numérique jamais casté/validé).
- Construction dynamique de clauses `ORDER BY`, `LIMIT` ou de noms de table/colonne à partir d'entrée utilisateur.

## Remédiation
- Utiliser systématiquement des requêtes préparées avec paramètres liés (jamais de concaténation de valeurs).
- Pour les identifiants dynamiques (noms de colonnes/tables), utiliser une liste blanche stricte plutôt que la valeur brute.
- Valider et typer strictement les entrées avant la requête.
- Appliquer le principe du moindre privilège sur le compte de connexion à la base de données.
- Voir `rules/remediation/sql-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/sql-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: SQL Injection Prevention
- CWE-89: Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')
