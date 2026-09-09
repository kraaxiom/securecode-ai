# Bucket Google Cloud Storage public (CWE-284)

## Description de la vulnerabilite
Le fichier `vulnerable.go` illustre une application Go qui, via le SDK GCP (`cloud.google.com/go/storage`), accorde le role `roles/storage.objectViewer` au principal `allUsers` sur un bucket GCS. Combine a un `uniform bucket-level access` desactive, cela rend tous les objets du bucket lisibles par n'importe qui sur Internet, sans authentification, et laisse en plus la possibilite d'ACL d'objet individuelles encore plus permissives.

## CWE reel utilise
**CWE-284 : Improper Access Control** (source : `knowledge/cloud/gcp-bucket-public.md`).

## Pourquoi c'est dangereux
`allUsers` (ou `allAuthenticatedUsers`) dans un binding IAM GCS signifie un acces anonyme (ou tout compte Google authentifie) a l'ensemble des objets du bucket. C'est une configuration frequemment introduite par erreur lors d'un besoin ponctuel (site statique) mais qui persiste ensuite, exposant potentiellement des donnees sensibles ajoutees ulterieurement au meme bucket.

## Comment le correctif fonctionne
Le fichier `fixed.go` :
- Retire explicitement les bindings `allUsers` et `allAuthenticatedUsers` de la policy IAM du bucket.
- Accorde `roles/storage.objectViewer` uniquement a un service account applicatif precis (`app-reader@...`), respectant le principe du moindre privilege.
- Active `UniformBucketLevelAccess` et `PublicAccessPreventionEnforced` pour empecher toute incoherence via des ACL d'objet individuelles et bloquer tout acces public futur au niveau du bucket.
- Fournit `GenerateSignedURL` pour tout partage ponctuel de fichier, avec une URL signee valable seulement 30 minutes, au lieu d'un acces public permanent.

## References
- OWASP Cloud Security Cheat Sheet
- Google Cloud Docs: Public access prevention
- CWE-284: Improper Access Control — https://cwe.mitre.org/data/definitions/284.html
- Voir egalement `rules/remediation/gcp-bucket-public.md`
