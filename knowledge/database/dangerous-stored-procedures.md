---
id: dangerous-stored-procedures
category: database
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Dangerous Stored Procedures (procédures stockées dangereuses)

## Description
Ce pattern concerne les procédures stockées ou fonctions de base de données qui construisent et exécutent du SQL dynamique en concaténant des paramètres non neutralisés (SQL injection au niveau de la procédure elle-même), ou qui exposent des fonctionnalités système dangereuses (exécution de commandes shell, accès au système de fichiers, appels réseau depuis la base). Une procédure stockée mal conçue peut devenir un point d'injection même si le code applicatif appelant utilise correctement des requêtes préparées.

## Où ça apparaît typiquement
- Procédures stockées construisant une requête via `EXEC`/`EXECUTE IMMEDIATE`/concaténation de chaînes à partir de paramètres d'entrée.
- Procédures utilisant des fonctionnalités étendues du SGBD permettant l'exécution de commandes système (ex: fonctionnalités d'exécution shell natives).
- Procédures accordant des privilèges élevés (`SECURITY DEFINER`, exécution avec les droits du propriétaire) sans validation stricte des entrées.
- Triggers de base de données déclenchant des actions sensibles à partir de données utilisateur non validées.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- SQL dynamique construit par concaténation de paramètres à l'intérieur d'une procédure stockée plutôt que par liaison de paramètres.
- Procédure stockée invoquant des fonctionnalités d'exécution système ou d'accès fichier natives du SGBD.
- Procédure exécutée avec des privilèges élevés (droits du propriétaire) alors que ses paramètres proviennent directement de l'entrée utilisateur.
- Absence de revue de sécurité dédiée aux procédures stockées lors des audits de code (souvent hors du périmètre des scanners applicatifs classiques).

## Remédiation
- Utiliser la liaison de paramètres même à l'intérieur des procédures stockées lors de la construction de SQL dynamique.
- Désactiver ou restreindre strictement l'accès aux fonctionnalités d'exécution système natives du SGBD lorsque non indispensables.
- Limiter les privilèges d'exécution des procédures au strict nécessaire et éviter `SECURITY DEFINER`/équivalent sans validation stricte des entrées.
- Inclure les procédures stockées dans le périmètre des revues de sécurité et des scans de code, pas uniquement le code applicatif.
- Voir `rules/remediation/dangerous-stored-procedures.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/dangerous-stored-procedures/`.

## Références
- OWASP Top 10: A03:2021 – Injection
- CWE-89: Improper Neutralization of Special Elements used in an SQL Command
