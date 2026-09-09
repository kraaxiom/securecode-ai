---
id: symlink-attack
category: file-inclusion
cwe: CWE-61
owasp: A01:2021-Broken Access Control
severity_default: medium
languages: [php, python, java, go, rust, csharp]
---

# Symlink Attack (Symbolic Link Following)

## Description
Une attaque par lien symbolique exploite le fait qu'une application accède à un fichier via un chemin qui peut, au moment de l'accès, être un lien symbolique pointant vers une autre ressource. Si l'application ne vérifie pas la nature du fichier avant de le lire, l'écrire ou en changer les permissions, un attaquant capable de créer ce lien peut rediriger l'opération vers un fichier sensible du système (race condition de type TOCTOU).

## Où ça apparaît typiquement
- Répertoires d'upload ou de fichiers temporaires partagés entre plusieurs utilisateurs/processus.
- Scripts d'administration ou de déploiement manipulant des fichiers dans des répertoires accessibles en écriture par des tiers.
- Opérations d'archivage/décompression écrivant dans un répertoire cible sans vérifier chaque entrée.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Vérification de l'existence/des permissions d'un fichier (`stat`) suivie plus tard d'une opération séparée sur le même chemin (fenêtre TOCTOU) sans verrouillage.
- Écriture de fichiers dans un répertoire partagé en mode monde-inscriptible sans création atomique exclusive (`O_EXCL`/équivalent).
- Absence de vérification explicite qu'un chemin n'est pas un lien symbolique avant une opération sensible.

## Remédiation
- Utiliser des opérations atomiques d'ouverture exclusive de fichier (drapeau équivalent à `O_EXCL`/`O_NOFOLLOW`) plutôt qu'un `stat` suivi d'une action séparée.
- Éviter les répertoires partagés en écriture pour les fichiers temporaires ; utiliser des répertoires dédiés par utilisateur/processus avec permissions restrictives.
- Vérifier explicitement qu'un chemin résolu n'est pas un lien symbolique avant toute opération sensible.
- Voir `rules/remediation/symlink-attack.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/symlink-attack/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: File Upload / Path Traversal
- CWE-61: UNIX Symbolic Link (Symlink) Following
