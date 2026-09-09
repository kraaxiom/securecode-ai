---
id: aws-s3-public
category: cloud
cwe: CWE-284
owasp: A01:2021-Broken Access Control
severity_default: high
languages: []
---

# Bucket S3 public

## Description
Un bucket AWS S3 configuré en accès public permet à quiconque connaît son nom (ou le découvre par énumération) de lister, lire, voire écrire des objets sans authentification. Cette mauvaise configuration provient généralement d'une politique de bucket (bucket policy) ou d'ACL trop permissive, ou de la désactivation des paramètres "Block Public Access". C'est l'une des causes les plus fréquentes de fuites massives de données dans le cloud.

## Où ça apparaît typiquement
- Bucket policy avec `"Principal": "*"` et une action `s3:GetObject`/`s3:ListBucket` sans condition restrictive.
- ACL du bucket ou des objets réglée sur "public-read" ou "public-read-write".
- Paramètres "Block Public Access" désactivés au niveau du compte ou du bucket.
- Buckets utilisés pour du contenu statique (sites web, assets) mais contenant aussi des fichiers sensibles (backups, logs, exports de données).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Bucket policy JSON contenant `"Principal": "*"` combiné à une action de lecture/écriture sans condition IAM.
- Configuration Terraform/CloudFormation avec `block_public_acls = false`, `acl = "public-read"` ou équivalent.
- Absence de chiffrement et de politique de blocage d'accès public dans l'infra-as-code.
- Nom de bucket prévisible (nom de domaine, environnement) facilitant l'énumération.

## Remédiation
- Activer "Block Public Access" au niveau du compte et de chaque bucket, sauf besoin explicite documenté.
- Restreindre les bucket policies au principe du moindre privilège (rôles/IAM spécifiques, pas `Principal: *`).
- Utiliser des URLs pré-signées à durée limitée pour partager des objets ponctuellement plutôt qu'un accès public permanent.
- Auditer régulièrement les buckets via AWS Config / Trusted Advisor / Access Analyzer.
- Voir `rules/remediation/aws-s3-public.md`.

## Exemple avant/après
Voir `examples/terraform/aws-s3-public/`.

## Références
- OWASP Cloud Security Cheat Sheet
- AWS Well-Architected Framework: Security Pillar
- CWE-284: Improper Access Control
