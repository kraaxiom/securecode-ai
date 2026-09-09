# Bucket S3 public (CWE-284)

## Description de la vulnerabilite
Le fichier `vulnerable.go` illustre une application Go qui, en utilisant le SDK AWS (`aws-sdk-go-v2`), cree un bucket S3 avec une ACL `public-read`, desactive completement le "Block Public Access" et applique une bucket policy dont le `Principal` est `"*"`. Resultat : n'importe qui sur Internet connaissant (ou devinant) le nom du bucket peut lister et lire son contenu sans authentification.

## CWE reel utilise
**CWE-284 : Improper Access Control** (source : `knowledge/cloud/aws-s3-public.md`).

## Pourquoi c'est dangereux
Les buckets S3 exposes publiquement sont l'une des causes les plus frequentes de fuites massives de donnees dans le cloud. Un bucket pense pour heberger du contenu statique (site web, assets) finit souvent par contenir aussi des sauvegardes, des logs ou des exports de donnees sensibles, tous accessibles sans aucun controle une fois l'acces public active.

## Comment le correctif fonctionne
Le fichier `fixed.go` applique le principe du moindre privilege :
- Le bucket est cree avec l'ACL `private` (par defaut) au lieu de `public-read`.
- Les quatre parametres de "Block Public Access" sont actives (`BlockPublicAcls`, `IgnorePublicAcls`, `BlockPublicPolicy`, `RestrictPublicBuckets`).
- Le chiffrement au repos (SSE-KMS) est active sur le bucket.
- La bucket policy restreint l'acces a un role IAM applicatif precis (`app-read-role`) au lieu de `Principal: "*"`, avec une condition exigeant TLS (`aws:SecureTransport`).
- Pour tout partage ponctuel, une URL pre-signee a duree de vie courte (15 minutes) est generee via `s3.NewPresignClient`, plutot que de rendre le bucket public de facon permanente.

## References
- OWASP Cloud Security Cheat Sheet
- AWS Well-Architected Framework: Security Pillar
- CWE-284: Improper Access Control — https://cwe.mitre.org/data/definitions/284.html
- Voir egalement `rules/remediation/aws-s3-public.md`
