---
id: stacked-query-sqli
category: injections
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: critical
languages: [php, js, python, java, csharp, go]
---

# Stacked Query SQL Injection

## Description
L'injection SQL par requêtes empilées (stacked queries) exploite les pilotes de base de données qui permettent l'exécution de plusieurs instructions SQL séparées par un point-virgule dans un seul appel. Lorsqu'une entrée utilisateur non neutralisée atteint ce type d'appel, un attaquant peut ajouter une instruction SQL complètement distincte (ex: `INSERT`, `UPDATE`, `DROP`) à la suite de la requête légitime, avec un impact potentiellement plus large qu'une injection classique limitée à une seule instruction `SELECT`.

## Où ça apparaît typiquement
- Pilotes/API de base de données autorisant l'exécution multi-instructions en un seul appel (certains drivers PHP PDO, SQL Server, PostgreSQL selon configuration).
- Requêtes construites par concaténation dans des contextes où le driver ne limite pas à une seule instruction par appel.
- Interfaces d'administration ou d'export exécutant des requêtes utilisateur avec un driver multi-statement activé.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de variable non échappée dans une requête exécutée via une API supportant nativement le multi-statement.
- Configuration de driver avec l'option multi-requêtes activée sans nécessité fonctionnelle (ex: `MULTI_STATEMENTS` MySQL).
- Absence de requêtes préparées alors que le point d'entrée accepte des entrées non fiables.

## Remédiation
- Désactiver l'exécution multi-instructions au niveau du driver lorsque l'application n'en a pas besoin.
- Utiliser systématiquement des requêtes préparées avec paramètres liés.
- Appliquer le principe du moindre privilège sur le compte de connexion à la base de données (restreindre les droits DDL/DML non nécessaires).
- Voir `rules/remediation/stacked-query-sqli.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/csharp-dotnet/stacked-query-sqli/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: SQL Injection Prevention
- CWE-89: Improper Neutralization of Special Elements used in an SQL Command ('SQL Injection')
