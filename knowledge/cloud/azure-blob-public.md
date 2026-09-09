---
id: azure-blob-public
category: cloud
cwe: CWE-284
owasp: A01:2021-Broken Access Control
severity_default: high
languages: []
---

# Conteneur Azure Blob Storage public

## Description
Un conteneur Azure Blob Storage configuré avec un niveau d'accès public ("Blob" ou "Container") permet à quiconque de lire, et parfois lister, les objets sans authentification. Cette exposition résulte typiquement d'un niveau d'accès anonyme mal choisi lors de la création du conteneur ou d'un compte de stockage dont le paramètre "Allow Blob public access" reste activé par défaut.

## Où ça apparaît typiquement
- Conteneur créé avec "Public access level" réglé sur "Blob" ou "Container" au lieu de "Private".
- Paramètre de compte de stockage `allowBlobPublicAccess` laissé à `true` sans nécessité.
- Génération de SAS (Shared Access Signature) avec permissions trop larges ou sans date d'expiration.
- Contenu sensible (backups, exports, logs) stocké dans un conteneur destiné à l'origine à du contenu statique public.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Propriété `publicAccess` ou `allowBlobPublicAccess` à `true`/`Container`/`Blob` dans un template ARM/Bicep/Terraform.
- Absence de restriction réseau (firewall de compte de stockage ouvert à "All networks").
- SAS token généré sans expiration courte ni permissions scoped (lecture seule).
- Nom de conteneur prévisible facilitant l'énumération de fichiers sensibles.

## Remédiation
- Désactiver `allowBlobPublicAccess` au niveau du compte de stockage, sauf besoin explicite documenté.
- Fixer le niveau d'accès du conteneur sur "Private" et utiliser des SAS à durée de vie courte et permissions minimales pour le partage.
- Restreindre l'accès réseau au compte de stockage (VNet, IP allowlist, Private Endpoint).
- Auditer régulièrement via Azure Defender for Storage / Azure Policy.
- Voir `rules/remediation/azure-blob-public.md`.

## Exemple avant/après
Voir `examples/terraform/azure-blob-public/`.

## Références
- OWASP Cloud Security Cheat Sheet
- Microsoft Learn: Prevent anonymous public read access to containers and blobs
- CWE-284: Improper Access Control
