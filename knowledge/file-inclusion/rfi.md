---
id: rfi
category: file-inclusion
cwe: CWE-98
owasp: A03:2021-Injection
severity_default: critical
languages: [php]
---

# Remote File Inclusion (RFI)

## Description
La Remote File Inclusion est une variante de la LFI où l'application inclut un fichier situé sur un serveur distant contrôlé par l'attaquant, via une URL passée en paramètre. Si le moteur d'inclusion suit les wrappers réseau (`http://`, `ftp://`), l'attaquant peut faire exécuter du code arbitraire hébergé sur son propre serveur, ce qui en fait historiquement l'une des vulnérabilités PHP les plus critiques.

## Où ça apparaît typiquement
- Applications PHP legacy avec `allow_url_include` activé et sélecteurs de module/page dynamiques.
- Systèmes de templates permettant de spécifier une URL comme source de contenu.
- Intégrations tierces chargeant dynamiquement des scripts par URL configurable.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fonction d'inclusion dynamique recevant une valeur pouvant contenir un schéma d'URL (`http://`, `https://`, `ftp://`).
- Configuration runtime autorisant l'inclusion de fichiers distants (à vérifier dans les fichiers de config, ex. `allow_url_include`/`allow_url_fopen`).
- Absence de validation de schéma/format sur les entrées destinées à une fonction d'inclusion.

## Remédiation
- Désactiver `allow_url_include` (et restreindre `allow_url_fopen`) au niveau de la configuration du runtime.
- Ne jamais construire un chemin d'inclusion à partir d'une entrée pouvant contenir un schéma d'URL.
- Utiliser un mapping fermé de valeurs autorisées plutôt qu'un chemin/URL fourni par le client.
- Voir `rules/remediation/rfi.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/rfi/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: File Upload / Path Traversal (sections relatives à l'inclusion de fichiers)
- CWE-98: Improper Control of Filename for Include/Require Statement in PHP Program ('PHP Remote File Inclusion')
