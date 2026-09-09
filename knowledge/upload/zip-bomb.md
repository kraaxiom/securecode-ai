---
id: zip-bomb
category: upload
cwe: CWE-409
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: [php, java, csharp, python, js, go]
---

# ZIP Bomb (bombe de décompression)

## Description
Une bombe de décompression est un fichier archive (ZIP, GZIP, XML, etc.) conçu pour se décompresser en une taille disproportionnellement plus grande que sa taille compressée, parfois de plusieurs ordres de grandeur, dans le but d'épuiser la mémoire, l'espace disque ou le CPU du serveur qui la traite. Lorsqu'une application décompresse automatiquement des fichiers uploadés sans limite de taille ni de ratio de compression, un attaquant peut provoquer un déni de service avec un fichier initial de quelques kilo-octets seulement.

## Où ça apparaît typiquement
- Fonctionnalités d'extraction automatique d'archives uploadées (import de plugins, de thèmes, de packages, d'archives de sauvegarde).
- Traitement de documents Office (`.docx`, `.xlsx`), qui sont eux-mêmes des archives ZIP, sans limite sur la décompression.
- Parsers XML activant l'expansion d'entités sans limite (risque connexe de type "billion laughs").
- Middlewares de décompression HTTP (gzip) appliqués sur le corps de requêtes sans limite de taille décompressée.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à une fonction d'extraction d'archive (`ZipArchive::extractTo()`, `zipfile.extractall()`, `java.util.zip`) sans vérification préalable de la taille totale décompressée annoncée dans les métadonnées de l'archive.
- Absence de limite sur le ratio taille décompressée / taille compressée avant traitement complet.
- Absence de limite de temps ou de mémoire allouée au processus de décompression.
- Décompression récursive d'archives imbriquées sans limite de profondeur.

## Remédiation
- Vérifier la taille décompressée annoncée dans les en-têtes de l'archive avant extraction, et rejeter au-delà d'un seuil raisonnable.
- Limiter le ratio de compression accepté et le nombre/profondeur de fichiers imbriqués.
- Décompresser dans un environnement isolé avec quotas stricts de mémoire, disque et CPU, et un timeout.
- Traiter les fichiers de manière incrémentale (streaming) avec arrêt dès dépassement des limites, plutôt que tout charger en mémoire.
- Voir `rules/remediation/zip-bomb.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/zip-bomb/` (et équivalents pour les autres langages listés).

## Références
- OWASP Cheat Sheet: Denial of Service Cheat Sheet
- CWE-409: Improper Handling of Highly Compressed Data (Data Amplification)
