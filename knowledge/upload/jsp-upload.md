---
id: jsp-upload
category: upload
cwe: CWE-434
owasp: A04:2021-Insecure Design
severity_default: critical
languages: [java]
---

# Upload de fichier JSP exécutable

## Description
Cette vulnérabilité survient quand une application Java web (Servlet/JSP, Spring, Tomcat) accepte le téléversement d'un fichier avec une extension interprétable par le conteneur de servlets (`.jsp`, `.jspx`, `.jsw`, `.jsv`, `.war`) dans un répertoire déployé. Si le conteneur compile et exécute ce fichier comme une page JSP au lieu de le servir en statique, l'attaquant obtient l'exécution de code arbitraire côté serveur. Le déploiement d'un `.war` malveillant via une interface d'administration mal protégée constitue une variante particulièrement critique.

## Où ça apparaît typiquement
- Formulaires d'upload multipart traités avec `Part.write()` ou `commons-fileupload` sans validation de l'extension réelle.
- Répertoires d'upload situés dans le webapp déployé (`webapps/ROOT/uploads/`) plutôt que hors du contexte web.
- Interfaces d'administration Tomcat (Manager App) exposées permettant le déploiement direct d'archives `.war`.
- Frameworks Spring MVC utilisant `MultipartFile.transferTo()` sans liste blanche d'extensions.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à `transferTo()`, `Part.write()` ou équivalent sans validation stricte de l'extension et du contenu.
- Validation basée uniquement sur `Content-Type` déclaré par le client dans la requête multipart.
- Répertoire de destination situé sous le répertoire `webapps` du conteneur servlet, donc potentiellement exécutable.
- Absence de renommage du fichier téléversé côté serveur.
- Endpoint d'administration/déploiement (Manager App Tomcat) accessible sans authentification forte ou restriction réseau.

## Remédiation
- Stocker les fichiers téléversés en dehors du répertoire `webapps`, dans un emplacement non servi par le conteneur.
- Utiliser une liste blanche stricte d'extensions autorisées et vérifier le contenu réel du fichier (signature binaire), jamais le `Content-Type` déclaré.
- Générer un nom de fichier aléatoire côté serveur.
- Restreindre ou désactiver l'accès au Manager App Tomcat en production, et le protéger par authentification forte et filtrage IP.
- Voir `rules/remediation/jsp-upload.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java/jsp-upload/`.

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-434: Unrestricted Upload of File with Dangerous Type
