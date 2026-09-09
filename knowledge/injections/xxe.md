---
id: xxe
category: injections
cwe: CWE-611
owasp: A05:2021-Security_Misconfiguration
severity_default: critical
languages: [php, js, python, java, csharp, go]
---

# XML External Entity (XXE) Injection

## Description
L'injection d'entité externe XML (XXE) survient lorsqu'un parseur XML mal configuré traite des déclarations d'entités externes présentes dans un document XML fourni par l'utilisateur. Un attaquant peut définir une entité pointant vers un fichier local ou une URL, ce qui peut mener à une divulgation de fichiers sensibles sur le serveur, une falsification de requête côté serveur (SSRF), voire un déni de service (entity expansion / "billion laughs").

## Où ça apparaît typiquement
- Endpoints acceptant de l'XML en entrée (API SOAP, upload de fichiers XML/SVG/DOCX/XLSX, flux de données B2B).
- Parseurs XML utilisés avec leur configuration par défaut, qui autorise historiquement le traitement des entités externes et des DTD dans de nombreuses bibliothèques.
- Bibliothèques tierces de traitement de documents (parsers SVG, PDF, Office) s'appuyant en interne sur un parseur XML non durci.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Utilisation d'un parseur XML sans désactivation explicite du traitement des DTD et des entités externes (`DOCTYPE`, `ENTITY`).
- Absence de configuration `disallow-doctype-decl`, `external-general-entities`, `external-parameter-entities` à `false`/désactivé selon la bibliothèque.
- Traitement d'uploads de formats basés sur XML (SVG, DOCX, XLSX) sans validation ni durcissement du parseur sous-jacent.

## Remédiation
- Désactiver explicitement le traitement des DTD et des entités externes sur tout parseur XML utilisé avec des entrées non fiables.
- Utiliser des bibliothèques ou des modes de parsing "sûrs par défaut" lorsqu'ils sont disponibles.
- Valider les formats de fichiers uploadés à un niveau plus strict que la simple extension (type MIME, structure attendue).
- Voir `rules/remediation/xxe.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/xxe/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: XML External Entity Prevention
- CWE-611: Improper Restriction of XML External Entity Reference
