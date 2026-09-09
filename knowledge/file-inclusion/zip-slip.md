---
id: zip-slip
category: file-inclusion
cwe: CWE-22
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: [java, js, python, go, csharp, php, rust]
---

# Zip Slip

## Description
Zip Slip est une vulnérabilité de path traversal qui se produit lors de la décompression d'une archive (ZIP, TAR, etc.) dont les entrées peuvent contenir des séquences `../` dans leur nom de fichier. Si le code d'extraction concatène naïvement le répertoire de destination avec le nom de chaque entrée sans validation, un attaquant qui contrôle l'archive peut écrire des fichiers en dehors du répertoire cible, écrasant potentiellement des fichiers système ou applicatifs critiques.

## Où ça apparaît typiquement
- Fonctionnalités d'import/restauration de sauvegardes acceptant une archive fournie par l'utilisateur.
- Traitement de packages, plugins ou thèmes distribués sous forme d'archive.
- Pipelines CI/CD ou outils de déploiement décompressant des artefacts téléchargés.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Boucle d'extraction d'archive qui construit le chemin de sortie par simple concaténation (`destDir + entry.getName()`) sans normalisation ni vérification.
- Absence de contrôle que le chemin résolu de chaque entrée reste sous le répertoire de destination.
- Utilisation d'une bibliothèque de décompression bas niveau sans passer par une fonction d'extraction sécurisée fournie par le framework.

## Remédiation
- Pour chaque entrée d'archive, résoudre le chemin final de façon canonique et vérifier qu'il est un sous-chemin strict du répertoire de destination avant écriture.
- Rejeter toute entrée dont le nom contient des séquences `..`, un chemin absolu ou un lien symbolique.
- Privilégier les bibliothèques de décompression qui appliquent nativement ces contrôles.
- Voir `rules/remediation/zip-slip.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/zip-slip/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Path Traversal
- CWE-22: Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')
