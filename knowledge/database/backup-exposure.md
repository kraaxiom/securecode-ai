---
id: backup-exposure
category: database
cwe: CWE-538
owasp: A05:2021-Security-Misconfiguration
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Backup Exposure (exposition de sauvegardes)

## Description
Ce pattern couvre les cas où des sauvegardes de base de données (dumps SQL, exports, snapshots) sont accessibles à des personnes non autorisées, que ce soit via un chemin web public prévisible, un bucket de stockage cloud mal configuré, ou des permissions de fichiers trop permissives sur le serveur. Une sauvegarde contient généralement l'intégralité des données de production, rendant sa fuite aussi grave qu'une compromission directe de la base.

## Où ça apparaît typiquement
- Fichiers de sauvegarde (`.sql`, `.dump`, `.bak`) déposés dans un répertoire accessible depuis le webroot de l'application.
- Buckets de stockage cloud (S3 ou équivalent) contenant des sauvegardes configurés en accès public par erreur.
- Scripts de sauvegarde automatisés écrivant dans un emplacement partagé sans restriction de permissions adaptée.
- Sauvegardes transmises ou stockées sans chiffrement, y compris en transit vers un stockage externe.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Répertoire de sauvegarde situé sous le webroot ou accessible via une route HTTP sans authentification.
- Configuration de bucket de stockage cloud sans politique d'accès restrictive explicite pour les objets de sauvegarde.
- Permissions de fichiers de sauvegarde sur le serveur lisibles par d'autres utilisateurs/services que le processus de sauvegarde.
- Absence de chiffrement au repos et en transit pour les fichiers de sauvegarde contenant des données de production.

## Remédiation
- Stocker les sauvegardes en dehors du webroot et dans un emplacement non accessible directement via HTTP.
- Appliquer une politique d'accès restrictive explicite (liste blanche) sur tout stockage cloud contenant des sauvegardes.
- Chiffrer les sauvegardes au repos et en transit, avec gestion des clés séparée de l'accès aux données elles-mêmes.
- Restreindre les permissions du système de fichiers aux seuls comptes techniques nécessaires au processus de sauvegarde/restauration.
- Voir `rules/remediation/backup-exposure.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/backup-exposure/`.

## Références
- OWASP Top 10: A05:2021 – Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
