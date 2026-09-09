---
id: directory-traversal
category: file-inclusion
cwe: CWE-22
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Directory Traversal

## Description
Le directory traversal (ou path traversal) survient quand une application utilise une entrée utilisateur pour construire un chemin de fichier ou de répertoire sans neutraliser les séquences `../`. Un attaquant peut alors remonter dans l'arborescence du système de fichiers et accéder à des répertoires en dehors du périmètre prévu par l'application. C'est une des causes les plus fréquentes de divulgation de fichiers sensibles (configuration, code source, secrets).

## Où ça apparaît typiquement
- Endpoints de téléchargement/affichage de fichiers (`?file=`, `?path=`, `?dir=`).
- Fonctions de listing de répertoires utilisateur (dossiers partagés, gestionnaires de médias).
- Chemins construits à partir de paramètres de configuration ou de templates dynamiques.
- Middlewares de fichiers statiques mal configurés.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation directe d'un paramètre utilisateur dans un chemin (`basePath + userInput`) sans normalisation.
- Absence d'appel à une fonction de canonicalisation (`realpath`, `Path.resolve`, `os.path.abspath`) suivie d'une vérification que le résultat reste sous le répertoire racine autorisé.
- Absence de whitelist de noms de fichiers/répertoires autorisés.
- Utilisation de fonctions bas niveau d'accès fichier directement sur une entrée réseau.

## Remédiation
- Résoudre le chemin final de façon canonique puis vérifier qu'il est un sous-chemin strict du répertoire racine autorisé.
- Préférer une whitelist d'identifiants (ID logique mappé en interne vers un chemin) plutôt qu'un chemin fourni par le client.
- Rejeter toute entrée contenant des séquences `..`, des chemins absolus ou des caractères nuls.
- Voir `rules/remediation/directory-traversal.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/directory-traversal/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Path Traversal
- CWE-22: Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')
