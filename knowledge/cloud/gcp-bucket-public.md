---
id: gcp-bucket-public
category: cloud
cwe: CWE-284
owasp: A01:2021-Broken Access Control
severity_default: high
languages: []
---

# Bucket Google Cloud Storage public

## Description
Un bucket Google Cloud Storage devient public quand son IAM policy ou son ACL accorde le rôle `roles/storage.objectViewer` (ou équivalent) aux principals `allUsers` ou `allAuthenticatedUsers`. Cela permet à n'importe qui sur Internet (ou tout compte Google authentifié) de lire, voire lister, les objets du bucket sans autorisation spécifique.

## Où ça apparaît typiquement
- Binding IAM du bucket incluant `allUsers` ou `allAuthenticatedUsers` avec un rôle de lecture ou d'écriture.
- ACL d'objet/bucket réglée sur "publicRead" via l'API ou la console.
- Configuration Terraform (`google_storage_bucket_iam_member`) avec `member = "allUsers"`.
- Contrainte d'organisation "Domain restricted sharing" / "Public Access Prevention" non activée.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence de `allUsers` ou `allAuthenticatedUsers` dans une policy IAM ou un fichier IaC de bucket.
- `uniform_bucket_level_access` désactivé, laissant les ACL par objet potentiellement plus permissives que la policy du bucket.
- Absence de la contrainte d'organisation "Public Access Prevention" (`enforcedPublicAccessPrevention`).
- Nom de bucket prévisible facilitant la découverte.

## Remédiation
- Activer "Public Access Prevention" au niveau organisation/projet pour bloquer tout accès public par défaut.
- Retirer `allUsers`/`allAuthenticatedUsers` des bindings IAM sauf besoin explicite documenté (ex: site statique public volontaire).
- Activer `uniform_bucket_level_access` pour éviter les ACL d'objet incohérentes avec la policy du bucket.
- Utiliser des URLs signées à durée limitée pour le partage ponctuel de fichiers.
- Voir `rules/remediation/gcp-bucket-public.md`.

## Exemple avant/après
Voir `examples/terraform/gcp-bucket-public/`.

## Références
- OWASP Cloud Security Cheat Sheet
- Google Cloud Docs: Public access prevention
- CWE-284: Improper Access Control
