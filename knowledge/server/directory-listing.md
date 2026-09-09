---
id: directory-listing
category: server
cwe: CWE-548
owasp: A05:2021-Security Misconfiguration
severity_default: medium
languages: []
---

# Listing de répertoire activé

## Description
Le listing de répertoire (directory listing) activé sur un serveur web affiche automatiquement la liste de tous les fichiers d'un dossier lorsqu'aucun fichier index n'y est présent. Cela permet à un attaquant de découvrir des fichiers non liés depuis l'application (backups, fichiers de configuration, scripts oubliés, archives) qui n'auraient jamais dû être accessibles publiquement.

## Où ça apparaît typiquement
- Configuration Apache avec `Options +Indexes` activée sur un répertoire servant du contenu statique.
- Configuration Nginx avec `autoindex on;` sur un bloc `location`.
- Dossiers d'upload, de cache ou de logs exposés directement sous le webroot sans fichier index ni règle de blocage.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Directive `Options +Indexes` (Apache) ou `autoindex on;` (Nginx) présente dans la configuration serveur.
- Requête HTTP sur un répertoire sans fichier index retournant une page de listing générée par le serveur plutôt qu'une erreur 403/404.
- Absence de fichier `index.html`/`index.php` vide de protection dans les dossiers sensibles du webroot.

## Remédiation
- Désactiver explicitement le listing de répertoire (`Options -Indexes` sous Apache, retirer/mettre `autoindex off;` sous Nginx).
- Placer les fichiers non destinés au public en dehors du webroot.
- Ajouter un fichier index vide ou une règle de refus explicite dans chaque répertoire sensible en défense en profondeur.
- Voir `rules/remediation/directory-listing.md`.

## Exemple avant/après
Voir `examples/nginx/directory-listing/`.

## Références
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Testing Guide: Test Directory Traversal / Directory Listing
- CWE-548: Exposure of Information Through Directory Listing
