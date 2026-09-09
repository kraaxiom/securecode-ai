---
id: mime-bypass
category: upload
cwe: CWE-434
owasp: A04:2021-Insecure Design
severity_default: high
languages: [php, java, csharp, python, js, go]
---

# Bypass de validation MIME type

## Description
Cette vulnérabilité survient quand une application détermine si un fichier téléversé est autorisé en se basant sur l'en-tête `Content-Type` envoyé par le client dans la requête multipart, ou sur l'extension déclarée, plutôt que sur une analyse serveur du contenu réel. Ces deux valeurs sont entièrement contrôlées par l'attaquant et peuvent être falsifiées librement, rendant ce type de contrôle inefficace en tant que barrière de sécurité isolée.

## Où ça apparaît typiquement
- Code serveur lisant `$_FILES['file']['type']` (PHP), `request.files['file'].content_type` (Python), ou l'en-tête `Content-Type` d'une partie multipart, et l'utilisant comme unique critère de validation.
- Middleware d'upload configuré avec une liste blanche de types MIME comparée à la valeur déclarée par le client sans vérification côté serveur.
- API REST acceptant un champ `mimetype` dans le corps JSON de la requête et le stockant sans re-vérification.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Utilisation de la propriété MIME fournie par le client (`$_FILES[...]['type']`, `file.content_type`, en-tête HTTP) comme seule condition de validation avant écriture sur disque.
- Absence d'appel à une fonction de détection de type basée sur le contenu réel (ex: `finfo_file()`, `python-magic`, détection par signature côté serveur).
- Aucune corroboration entre extension déclarée, type MIME déclaré et contenu réel du fichier.

## Remédiation
- Déterminer le type réel du fichier côté serveur à partir de son contenu (détection par signature ou bibliothèque dédiée), jamais à partir de l'en-tête `Content-Type` du client.
- Combiner cette détection avec une liste blanche stricte d'extensions et un stockage hors webroot ou non exécutable.
- Rejeter toute incohérence entre extension déclarée, type MIME déclaré et type détecté côté serveur.
- Voir `rules/remediation/mime-bypass.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/mime-bypass/` (et équivalents pour les autres langages listés).

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-434: Unrestricted Upload of File with Dangerous Type
- CWE-20: Improper Input Validation
