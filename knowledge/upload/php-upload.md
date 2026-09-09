---
id: php-upload
category: upload
cwe: CWE-434
owasp: A04:2021-Insecure Design
severity_default: critical
languages: [php]
---

# Upload de fichier PHP exécutable

## Description
Cette vulnérabilité survient quand une application permet à un utilisateur de téléverser un fichier avec une extension exécutable par l'interpréteur PHP (`.php`, `.phtml`, `.php5`, `.phar`, etc.) dans un répertoire accessible via le web. Si le serveur exécute ce fichier au lieu de le servir comme contenu statique, l'attaquant obtient l'exécution de code arbitraire côté serveur. C'est l'une des causes les plus fréquentes de compromission totale d'une application PHP.

## Où ça apparaît typiquement
- Formulaires d'upload d'avatar, de logo, de pièce jointe ou de document stockant le fichier dans un répertoire du webroot (`public/uploads/`, `wp-content/uploads/`).
- Fonctions `move_uploaded_file()` utilisées sans validation stricte de l'extension ou du type réel du fichier.
- Configurations de serveur web (Apache/Nginx) où le dossier d'upload n'a pas de directive désactivant l'exécution de scripts.
- Champs de nom de fichier réutilisant directement `$_FILES['file']['name']` sans normalisation ni renommage.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à `move_uploaded_file()` ou `copy()` sans liste blanche stricte d'extensions autorisées.
- Validation basée uniquement sur `$_FILES['file']['type']` (MIME déclaré par le client, falsifiable) au lieu du contenu réel.
- Absence de renommage du fichier téléversé (le nom original, potentiellement contrôlé par l'attaquant, est conservé tel quel).
- Répertoire de destination situé dans le webroot sans fichier `.htaccess` ou configuration serveur interdisant l'exécution de scripts.
- Absence de vérification de la taille et du nombre d'extensions dans le nom de fichier.

## Remédiation
- Stocker les fichiers téléversés en dehors du webroot, ou dans un répertoire configuré pour ne jamais exécuter de scripts.
- Utiliser une liste blanche stricte d'extensions et de types MIME vérifiés côté serveur à partir du contenu réel du fichier (pas de l'en-tête déclaré par le client).
- Générer un nom de fichier aléatoire côté serveur, sans réutiliser le nom fourni par l'utilisateur.
- Configurer le serveur web pour désactiver l'exécution de scripts dans le répertoire d'upload (`php_admin_flag engine off`, directives Nginx `location` dédiées).
- Voir `rules/remediation/php-upload.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/php-upload/`.

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-434: Unrestricted Upload of File with Dangerous Type
