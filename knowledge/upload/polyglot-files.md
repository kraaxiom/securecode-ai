---
id: polyglot-files
category: upload
cwe: CWE-434
owasp: A04:2021-Insecure Design
severity_default: high
languages: [php, java, csharp, python, js]
---

# Fichiers polyglots

## Description
Un fichier polyglot est un fichier valide simultanément dans plusieurs formats (ex: à la fois une image GIF/JPEG valide et un script/archive valide), construit en exploitant les différences entre la manière dont chaque format identifie son début et sa fin. Un tel fichier peut passer les contrôles de validation d'image (signature correcte, décodage réussi) tout en étant interprété comme un script exécutable si le contexte de traitement change (ex: inclusion via une fonction de template, extraction d'archive). Cette technique combine généralement une image légitime avec un fragment de code embarqué dans les métadonnées ou en fin de fichier.

## Où ça apparaît typiquement
- Fonctionnalités d'upload d'image dont le fichier résultant est ensuite inclus ou interprété par un autre composant de l'application (ex: `include()` PHP sur un fichier image, traitement par une bibliothèque de templating).
- Pipelines de traitement de fichiers combinant plusieurs bibliothèques (extraction d'archive, conversion d'image, OCR) où chacune peut interpréter une portion différente du même fichier.
- Systèmes acceptant des fichiers SVG (XML) qui peuvent embarquer du script ou des entités externes en plus d'être un format image valide.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichier accepté sur la seule base d'une validation de signature/format réussie, sans analyse de la totalité du contenu du fichier.
- Combinaison d'un point d'upload d'image avec un point d'inclusion ou d'exécution dynamique de fichiers ailleurs dans l'application (`include`, `require`, désérialisation).
- Absence de re-encodage de l'image après upload (le fichier stocké est un octet-pour-octet du fichier reçu).
- Acceptation de fichiers SVG sans neutralisation des balises de script ou des entités externes (XXE).

## Remédiation
- Toujours re-encoder/transformer les images après upload (redimensionnement, recompression) pour détruire tout contenu polyglot superflu.
- Ne jamais utiliser une fonction d'inclusion ou d'exécution dynamique sur un fichier provenant d'un répertoire d'upload utilisateur.
- Pour les SVG, utiliser un parseur strict qui neutralise scripts et entités externes, ou convertir vers un format raster.
- Isoler le traitement des fichiers uploadés dans un environnement sans capacité d'exécution.
- Voir `rules/remediation/polyglot-files.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/polyglot-files/` (et équivalents pour les autres langages listés).

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-434: Unrestricted Upload of File with Dangerous Type
- CWE-20: Improper Input Validation
