---
id: lfi
category: file-inclusion
cwe: CWE-98
owasp: A03:2021-Injection
severity_default: critical
languages: [php, java, js, python]
---

# Local File Inclusion (LFI)

## Description
La Local File Inclusion survient quand une application inclut et exécute dynamiquement un fichier local dont le chemin est influencé par une entrée utilisateur (typiquement via `include`/`require` en PHP ou des mécanismes équivalents). Contrairement au simple path traversal en lecture, la LFI peut mener à l'exécution de code si le fichier inclus est interprété par le moteur applicatif, notamment en combinant la faille avec l'upload de fichiers, les logs applicatifs, ou les sessions PHP.

## Où ça apparaît typiquement
- Sélecteurs de page/langue/template (`?page=`, `?lang=`, `?module=`) passés à une fonction d'inclusion dynamique.
- Systèmes de plugins chargeant du code par nom fourni en configuration ou en requête.
- Frameworks maison avec routage basé sur l'inclusion directe de fichiers.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à une fonction d'inclusion dynamique (`include`, `require`, `include_once`, équivalents Java/Python de chargement dynamique de module) avec une variable non validée.
- Absence de whitelist stricte des valeurs autorisées pour le sélecteur de fichier/page.
- Concaténation d'extension (`.php`) après une entrée utilisateur, supposée empêcher l'inclusion d'autres types de fichiers (protection insuffisante).

## Remédiation
- Ne jamais faire dépendre le nom de fichier inclus d'une entrée utilisateur directe ; utiliser un mapping fermé (ID → chemin fixe côté serveur).
- Désactiver les options d'inclusion à distance au niveau de la configuration du runtime.
- Isoler les répertoires de contenu utilisateur (uploads) du chemin de recherche des fichiers inclus.
- Voir `rules/remediation/lfi.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/lfi/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: File Upload / Path Traversal (sections relatives à l'inclusion de fichiers)
- CWE-98: Improper Control of Filename for Include/Require Statement in PHP Program ('PHP Remote File Inclusion')
