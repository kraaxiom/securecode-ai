---
id: sqli-union
category: injections
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# UNION-based SQL Injection

## Description
Une injection SQL de type UNION survient quand une entrée utilisateur non neutralisée est intégrée directement dans une requête SQL, permettant à un attaquant d'ajouter une clause `UNION SELECT` pour extraire des données d'autres tables. C'est l'une des variantes les plus documentées de l'injection SQL.

## Où ça apparaît typiquement
- Concaténation de chaînes pour construire une requête SQL à partir de paramètres HTTP (query string, body, headers).
- Utilisation de requêtes "brutes" d'un ORM (`raw()`, `query()`) sans liaison de paramètres.
- Requêtes construites dynamiquement pour des filtres/tris (`ORDER BY $champ`) où `$champ` vient de l'utilisateur.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de variable non échappée dans une chaîne SQL (`"SELECT * FROM x WHERE id=" + input`).
- Appel `raw query` d'un ORM avec une valeur non liée en paramètre.
- Absence de couche de validation de type/format avant la requête (ex: un ID censé être numérique n'est jamais casté/validé).

## Remédiation
- Utiliser systématiquement des requêtes préparées avec paramètres liés (jamais de concaténation).
- Valider et typer strictement les entrées (ex: caster en entier un ID).
- Appliquer le principe du moindre privilège sur le compte de connexion à la base de données.
- Voir `rules/remediation/sqli-union.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/sqli-union/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: SQL Injection Prevention
- CWE-89: Improper Neutralization of Special Elements used in an SQL Command
