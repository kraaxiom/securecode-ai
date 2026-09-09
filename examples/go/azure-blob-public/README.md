# Conteneur Azure Blob Storage public (CWE-284)

## Description de la vulnerabilite
Le fichier `vulnerable.go` illustre une application Go qui, via le SDK Azure (`azure-sdk-for-go`), cree un conteneur Blob Storage avec l'option `Access: container.PublicAccessTypeContainer`. Ce niveau d'acces public permet a n'importe qui de lire, et meme de lister, tous les blobs du conteneur sans authentification, y compris des fichiers uploades ulterieurement sans controle supplementaire.

## CWE reel utilise
**CWE-284 : Improper Access Control** (source : `knowledge/cloud/azure-blob-public.md`).

## Pourquoi c'est dangereux
Un conteneur cree en "Public access level: Container" ou "Blob" reste accessible anonymement tant que le compte de stockage autorise `allowBlobPublicAccess`. Cette configuration, souvent choisie par erreur ou heritee d'un besoin ponctuel (hebergement de site statique), finit frequemment par exposer des sauvegardes, des exports de donnees ou des documents utilisateurs sensibles.

## Comment le correctif fonctionne
Le fichier `fixed.go` applique la remediation recommandee :
- Le conteneur est cree sans niveau d'acces public (`PublicAccessType("")`), ce qui le rend prive par defaut — combine a `allowBlobPublicAccess = false` au niveau du compte de stockage (a configurer via Azure CLI/Terraform, voir `rules/remediation/azure-blob-public.md`).
- Le partage ponctuel de fichiers se fait via une URL SAS (Shared Access Signature) generee avec `GetSASURL`, scopee en lecture seule (`sas.BlobPermissions{Read: true}`) et avec une expiration courte (1 heure) au lieu d'un acces public permanent.
- L'upload de fichiers n'herite plus d'aucune exposition publique implicite.

## References
- OWASP Cloud Security Cheat Sheet
- Microsoft Learn: Prevent anonymous public read access to containers and blobs
- CWE-284: Improper Access Control — https://cwe.mitre.org/data/definitions/284.html
- Voir egalement `rules/remediation/azure-blob-public.md`
