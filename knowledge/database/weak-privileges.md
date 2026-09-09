---
id: weak-privileges
category: database
cwe: CWE-269
owasp: A01:2021-Broken-Access-Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Weak Privileges (privilèges excessifs en base de données)

## Description
Ce pattern décrit les configurations où le compte technique utilisé par une application pour se connecter à la base de données dispose de privilèges bien supérieurs à ceux réellement nécessaires (droits DDL, accès à des schémas non utilisés, compte administrateur partagé entre services). Le non-respect du principe du moindre privilège amplifie considérablement l'impact d'une injection SQL ou d'une fuite d'identifiants, transformant une lecture non autorisée en compromission totale de l'instance.

## Où ça apparaît typiquement
- Compte de connexion applicatif utilisant le rôle administrateur/superutilisateur de la base de données par simplicité.
- Un même compte technique partagé entre plusieurs microservices ayant des besoins d'accès très différents.
- Comptes techniques disposant de droits DDL (CREATE, DROP, ALTER) alors que l'application ne fait que du CRUD applicatif.
- Absence de séparation entre compte de migration/déploiement (droits élevés, usage ponctuel) et compte d'exécution applicatif quotidien.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Chaîne de connexion applicative utilisant un rôle disposant de droits d'administration complets sur l'instance.
- Absence de rôles distincts par service/schéma, tous les composants applicatifs utilisant le même compte de base de données.
- Compte applicatif disposant de droits DDL ou d'accès à des schémas systèmes non nécessaires à son fonctionnement.
- Pas de revue périodique documentée des privilèges accordés aux comptes techniques de base de données.

## Remédiation
- Créer un rôle applicatif dédié par service, limité strictement aux opérations CRUD sur les schémas/tables nécessaires.
- Séparer le compte utilisé pour les migrations (droits DDL, usage ponctuel et audité) du compte d'exécution applicatif quotidien.
- Réaliser une revue périodique des privilèges accordés et révoquer systématiquement les droits non utilisés.
- Documenter et automatiser la création des rôles avec des scripts versionnés plutôt que des attributions manuelles ad hoc.
- Voir `rules/remediation/weak-privileges.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/weak-privileges/`.

## Références
- OWASP Top 10: A01:2021 – Broken Access Control
- CWE-269: Improper Privilege Management
