---
id: path-traversal
category: file-inclusion
cwe: CWE-22
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Path Traversal

## Description
Le path traversal désigne la capacité d'un attaquant à manipuler un chemin de fichier fourni à une application pour accéder à des ressources situées hors du répertoire prévu. Il s'agit du même mécanisme fondamental que le directory traversal, mais s'applique plus largement à toute opération sur fichier (lecture, écriture, suppression, archivage) où un chemin partiellement contrôlé par l'utilisateur est utilisé sans validation stricte.

## Où ça apparaît typiquement
- API de gestion de fichiers (upload, download, rename, delete) acceptant un nom ou chemin en paramètre.
- Génération de rapports/exports où le nom de fichier de sortie dérive d'une entrée utilisateur.
- Systèmes de templates ou de thèmes chargeant des fichiers par nom logique.
- Traitement d'archives (voir aussi `zip-slip.md`) où les entrées définissent leur propre chemin d'extraction.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Paramètre utilisateur utilisé tel quel dans un appel d'ouverture/écriture de fichier.
- Absence de normalisation (`normalize`, `realpath`) avant toute opération fichier.
- Absence de contrôle "le chemin résolu commence par le répertoire de base autorisé".
- Filtrage naïf basé uniquement sur un remplacement de chaîne `../` (contournable par encodage ou doublons).

## Remédiation
- Toujours résoudre le chemin en absolu puis vérifier son appartenance stricte au répertoire racine autorisé.
- Générer les noms de fichiers côté serveur (UUID) plutôt que d'utiliser le nom fourni par le client.
- Appliquer des permissions systèmes restrictives sur les répertoires de stockage.
- Voir `rules/remediation/path-traversal.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/path-traversal/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Path Traversal
- CWE-22: Improper Limitation of a Pathname to a Restricted Directory ('Path Traversal')
