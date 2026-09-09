# Mauvaise configuration IAM (CWE-269)

## Description de la vulnerabilite
Le fichier `vulnerable.go` illustre une application Go qui, via le SDK AWS (`aws-sdk-go-v2/service/iam`), attache a un role applicatif une policy IAM combinant `"Action": "*"` et `"Resource": "*"`, accordant l'integralite des permissions sur l'integralite des ressources du compte. Le meme fichier configure aussi une trust policy avec `"Principal": {"AWS": "*"}` sans aucune condition, permettant a n'importe quel compte AWS d'assumer ce role.

## CWE reel utilise
**CWE-269 : Improper Privilege Management** (source : `knowledge/cloud/iam-misconfiguration.md`).

## Pourquoi c'est dangereux
Une policy avec des wildcards sur `Action` et `Resource` viole le principe du moindre privilege : si les identifiants du role sont compromis (via une autre vulnerabilite applicative, une fuite de secret, ou un service tiers compromis), l'attaquant obtient un controle total sur le compte cloud. Une trust policy ouverte a `Principal: "*"` aggrave encore le risque en permettant a n'importe quel tiers d'assumer directement le role sans jamais avoir besoin de voler d'identifiants.

## Comment le correctif fonctionne
Le fichier `fixed.go` :
- Remplace la policy par une version granulaire n'autorisant que `s3:GetObject` et `s3:PutObject` sur un prefixe de ressource precis (`.../uploads/*`), avec une condition supplementaire sur la region (`aws:RequestedRegion`).
- Restreint la trust policy a un compte AWS precis (`trustedAcctID`), exige un `ExternalId` unique (protection contre le "confused deputy") et impose la presence du MFA (`aws:MultiFactorAuthPresent`).
- Applique ainsi le principe du moindre privilege recommande dans `rules/remediation/iam-misconfiguration.md`, limitant fortement l'impact d'une compromission eventuelle du role.

## References
- OWASP Cloud Security Cheat Sheet
- CIS Benchmarks (AWS/Azure/GCP) — section IAM
- CWE-269: Improper Privilege Management — https://cwe.mitre.org/data/definitions/269.html
- Voir egalement `rules/remediation/iam-misconfiguration.md`
