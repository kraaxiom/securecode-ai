---
id: web-shell-upload
category: upload
cwe: CWE-434
owasp: A04:2021-Insecure Design
severity_default: critical
languages: [php, java, csharp, python, js, go]
---

# Upload de web shell

## Description
Un web shell est un script malveillant déposé sur un serveur via une fonctionnalité d'upload puis exécuté à distance, donnant à l'attaquant une interface pour lancer des commandes système, lire/écrire des fichiers ou pivoter dans l'infrastructure. Cette vulnérabilité est la conséquence directe de contrôles d'upload insuffisants combinés à un répertoire de destination exécutable, quel que soit le langage serveur utilisé. C'est souvent l'étape finale d'une chaîne d'exploitation ayant commencé par un bypass d'extension, de MIME ou de magic bytes.

## Où ça apparaît typiquement
- Tout point d'upload de fichier (avatar, document, plugin, thème, pièce jointe) dont le résultat est stocké dans un répertoire accessible et exécutable par le serveur web.
- Fonctionnalités d'import/export de plugins ou de thèmes dans des CMS (WordPress, Joomla) permettant le dépôt d'archives contenant des scripts.
- Interfaces d'administration ou de déploiement (FTP, panneaux de gestion de fichiers) mal cloisonnées vis-à-vis du webroot.
- Endpoints d'API permettant l'écriture de fichiers arbitraires côté serveur sans validation de destination.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Combinaison d'un point d'upload sans liste blanche stricte d'extensions et d'un répertoire de destination situé dans le webroot.
- Absence de séparation entre stockage des fichiers utilisateurs et arborescence servie publiquement.
- Permissions de fichiers/dossiers d'upload trop larges (écriture puis exécution possibles par le même processus serveur).
- Absence de scan antivirus/heuristique ou de sandboxing sur les fichiers déposés avant mise à disposition.
- Journalisation absente ou insuffisante des accès aux fichiers nouvellement créés dans les répertoires d'upload.

## Remédiation
- Séparer strictement le stockage des fichiers uploadés du répertoire servi comme code exécutable.
- Désactiver l'exécution de scripts dans tout répertoire accessible en écriture par les utilisateurs (configuration serveur web).
- Appliquer une liste blanche d'extensions et de types de contenu, validée sur le contenu réel du fichier.
- Mettre en place une revue périodique et une détection d'anomalies sur les répertoires d'upload (nouveaux fichiers avec extensions inattendues).
- Voir `rules/remediation/web-shell-upload.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/web-shell-upload/` (et équivalents pour les autres langages listés).

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-434: Unrestricted Upload of File with Dangerous Type
