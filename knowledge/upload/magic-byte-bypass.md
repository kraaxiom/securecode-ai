---
id: magic-byte-bypass
category: upload
cwe: CWE-434
owasp: A04:2021-Insecure Design
severity_default: high
languages: [php, java, csharp, python, js, go]
---

# Bypass par falsification des magic bytes

## Description
Certaines applications valident le type d'un fichier téléversé en inspectant uniquement ses premiers octets (les "magic bytes" ou "file signature") pour vérifier qu'il correspond à un format attendu (ex: `FFD8FF` pour JPEG). Un attaquant peut préfixer un script malveillant avec les octets de signature d'un format image valide pour passer ce contrôle, tout en gardant un contenu exécutable plus loin dans le fichier. Ce contournement fonctionne car la vérification de signature seule ne garantit ni l'intégrité ni l'innocuité du reste du contenu du fichier.

## Où ça apparaît typiquement
- Bibliothèques de validation d'upload qui lisent les N premiers octets du fichier et les comparent à une table de signatures connues, sans validation structurelle complète du format.
- Fonctions de vérification faites "maison" (`fread($handle, 4)` puis comparaison) au lieu d'un décodage complet via une bibliothèque d'image dédiée.
- Combinaison d'une validation de magic bytes avec un stockage dans un répertoire exécutable, permettant l'exploitation du contenu situé après la signature falsifiée.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Lecture des seuls premiers octets du fichier pour la validation, sans traitement complet ni ré-encodage du contenu.
- Absence d'appel à une bibliothèque de décodage d'image complète (ex: `getimagesize()`, `Image.open()` avec chargement effectif des pixels) qui échouerait sur un fichier corrompu/hybride.
- Validation de signature utilisée comme unique contrôle, sans liste blanche d'extension corroborée ni isolation du répertoire de destination.
- Fichier accepté et stocké tel quel (octet pour octet) sans transformation (recompression, redimensionnement) qui neutraliserait un contenu superflu.

## Remédiation
- Ne jamais se fier à la seule signature binaire ; décoder complètement le fichier avec une bibliothèque d'image robuste et re-encoder l'image dans un nouveau fichier propre.
- Combiner plusieurs contrôles indépendants : extension, type MIME serveur, structure complète du fichier, et destination non exécutable.
- Rejeter tout fichier dont le décodage complet échoue ou produit des avertissements/erreurs de format.
- Voir `rules/remediation/magic-byte-bypass.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/magic-byte-bypass/` (et équivalents pour les autres langages listés).

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-434: Unrestricted Upload of File with Dangerous Type
- CWE-20: Improper Input Validation
