---
id: asp-upload
category: upload
cwe: CWE-434
owasp: A04:2021-Insecure Design
severity_default: critical
languages: [csharp]
---

# Upload de fichier ASP/ASP.NET exécutable

## Description
Cette vulnérabilité survient quand une application ASP classique ou ASP.NET accepte le téléversement d'un fichier avec une extension interprétable par IIS (`.asp`, `.aspx`, `.ashx`, `.asmx`, `.cshtml`) dans un répertoire servi par le serveur web. Si IIS exécute ce fichier au lieu de le renvoyer comme contenu statique, l'attaquant obtient l'exécution de code arbitraire sur le serveur. Le risque est aggravé par les extensions alternatives (`.asp;.jpg`, `.cer`, `.asa`) parfois encore interprétées selon la configuration IIS.

## Où ça apparaît typiquement
- Formulaires de dépôt de documents ou d'images dans des applications ASP.NET WebForms ou MVC utilisant `HttpPostedFile.SaveAs()`.
- Répertoires d'upload situés sous le site IIS sans mapping de handler restreint.
- API REST ASP.NET Core exposant un endpoint d'upload sans validation stricte du type de contenu.
- Configurations IIS héritées où les extensions historiques (`.asa`, `.cer`, `.cdx`) restent mappées à l'interpréteur ASP.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à `SaveAs()` ou `File.WriteAllBytes()` sans liste blanche d'extensions vérifiée côté serveur.
- Validation reposant uniquement sur `ContentType` fourni par le client dans la requête multipart.
- Nom de fichier original conservé tel quel pour construire le chemin de destination.
- Répertoire de destination situé sous le site IIS sans `web.config` dédié interdisant l'exécution de scripts.
- Absence de vérification de la signature binaire réelle du fichier avant écriture sur disque.

## Remédiation
- Stocker les fichiers en dehors de l'arborescence servie par IIS, ou dans un dossier avec un `web.config` retirant tous les handlers d'exécution.
- Utiliser une liste blanche stricte d'extensions et valider le contenu réel du fichier (signature binaire) plutôt que le `ContentType` déclaré.
- Générer un nom de fichier aléatoire côté serveur.
- Retirer explicitement les mappings d'extensions exécutables hérités (`.asa`, `.cer`, `.cdx`) dans la configuration IIS du dossier d'upload.
- Voir `rules/remediation/asp-upload.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/csharp/asp-upload/`.

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-434: Unrestricted Upload of File with Dangerous Type
