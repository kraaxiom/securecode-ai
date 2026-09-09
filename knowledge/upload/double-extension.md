---
id: double-extension
category: upload
cwe: CWE-434
owasp: A04:2021-Insecure Design
severity_default: high
languages: [php, java, csharp, python, js]
---

# Bypass par double extension

## Description
Cette technique de contournement exploite une validation d'upload qui ne vérifie que la présence d'une extension autorisée dans le nom de fichier, sans en extraire correctement la dernière extension réelle. Un fichier nommé `shell.php.jpg` ou `shell.jpg.php` peut ainsi tromper une validation naïve basée sur `strpos()`/`contains()` alors que le serveur, selon sa configuration, exécutera le fichier comme script. Elle est étroitement liée à la mauvaise confiance accordée au nom de fichier fourni par le client (CWE-646).

## Où ça apparaît typiquement
- Validations d'extension basées sur une recherche de sous-chaîne (`strpos($filename, '.jpg')`) plutôt que sur l'extraction de l'extension finale.
- Serveurs Apache mal configurés interprétant plusieurs extensions consécutives via `AddHandler`/`AddType`, exécutant `fichier.php.jpg` comme PHP.
- Frameworks utilisant une regex trop permissive pour valider les extensions (`.*\.jpg.*`).
- Systèmes de validation côté client (JavaScript) uniquement, sans re-vérification côté serveur de l'extension finale.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Validation d'extension par `strpos()`, `indexOf()` ou `contains()` sur le nom complet plutôt que par extraction de l'extension terminale (`pathinfo($f, PATHINFO_EXTENSION)` ou équivalent).
- Absence de vérification qu'un nom de fichier ne contient qu'une seule extension valide.
- Regex de validation autorisant des caractères après l'extension attendue.
- Directive `AddHandler` ou `AddType` Apache appliquée globalement sur un répertoire d'upload sans restriction stricte.

## Remédiation
- Extraire explicitement la dernière extension du nom de fichier avec une fonction dédiée, et la comparer à une liste blanche exacte.
- Rejeter tout nom de fichier contenant plusieurs points suivis de séquences d'extensions suspectes.
- Renommer systématiquement le fichier téléversé avec un nom généré côté serveur, extension unique et contrôlée.
- Corriger la configuration du serveur web pour ne jamais interpréter un fichier sur la base d'une extension intermédiaire.
- Voir `rules/remediation/double-extension.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/double-extension/` (et équivalents pour les autres langages listés).

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-434: Unrestricted Upload of File with Dangerous Type
- CWE-646: Reliance on File Name or Extension of Externally-Supplied File
