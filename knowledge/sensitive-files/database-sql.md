---
id: database-sql
category: sensitive-files
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition d'un dump de base de données (.sql)

## Description
Un fichier de dump SQL (`database.sql`, `backup.sql`, `dump.sql`) contient l'intégralité du schéma et souvent des données réelles de la base de données, y compris des mots de passe hashés, des données personnelles et des secrets applicatifs stockés en base. Sa présence dans un répertoire accessible publiquement constitue une fuite de données massive et immédiate, souvent plus grave qu'un accès direct à la base elle-même car le dump peut inclure l'historique complet des données.

## Où ça apparaît typiquement
- Sauvegardes générées par un script cron et écrites par erreur dans le webroot.
- Exports manuels effectués pour du débogage ou une migration, oubliés après usage.
- Outils d'administration de base de données (phpMyAdmin, Adminer) configurés pour exporter dans un répertoire public.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichiers `.sql` présents dans le répertoire servi publiquement par le serveur web.
- Scripts de sauvegarde dont le répertoire de destination est identique ou sous le webroot.
- Absence de règle serveur bloquant l'extension `.sql` en téléchargement direct.

## Remédiation
- Écrire les sauvegardes dans un répertoire strictement hors webroot, avec des permissions restreintes.
- Bloquer explicitement au niveau serveur toute extension `.sql` dans les répertoires servis publiquement.
- Chiffrer les sauvegardes au repos et limiter leur durée de rétention/accès.
- Voir `rules/remediation/database-sql.md` pour les diffs par configuration serveur.

## Exemple avant/après
Voir `examples/config/database-sql/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
