# Mauvaise configuration IAM — Python/Django

`vulnerable.py` provisionne un utilisateur IAM de service destiné uniquement à téléverser des justificatifs dans un bucket S3, mais lui attache une policy inline accordant `"Action": "*"` sur `"Resource": "*"` via `put_user_policy` — un accès administrateur complet sur l'ensemble du compte AWS (CWE-269, Improper Privilege Management).

`fixed.py` restreint la policy aux seules actions nécessaires (`s3:PutObject`, `s3:GetObject`) sur le préfixe et le bucket précis utilisés par l'application, avec une condition supplémentaire sur la région, en suivant le principe du moindre privilège.

## Pourquoi c'est dangereux
- Une clé d'accès associée à une policy `Action: *` / `Resource: *` équivaut à un accès administrateur complet : si ces identifiants fuient (log, dépôt de code, variable d'environnement exposée), l'attaquant peut lire/modifier/supprimer toutes les ressources du compte AWS, créer de nouveaux utilisateurs IAM, ou faire escalader ses propres privilèges.
- L'impact d'une compromission est totalement disproportionné par rapport au besoin réel du compte de service (uploader des fichiers dans un bucket).
- Ce type de policy rend également l'audit et la détection d'anomalies plus difficiles, car toute action devient "normale" pour ce compte.

## Explication du correctif
- La policy liste explicitement les actions nécessaires (`s3:PutObject`, `s3:GetObject`) au lieu du wildcard `*`.
- La ressource est scoped à un ARN précis (`arn:aws:s3:::my-app-documents-bucket/uploads/*`) plutôt qu'à `*`.
- Une condition (`aws:RequestedRegion`) réduit encore la surface d'utilisation valide de ces identifiants.
- Un commentaire signale la recommandation additionnelle de préférer des identités temporaires (STS, Workload Identity Federation) à une clé d'accès statique.

## Notes résiduelles
- Même avec une policy scoped, préférer à terme des rôles temporaires (`sts:AssumeRole`) plutôt que des clés d'accès statiques à durée de vie illimitée.
- Mettre en place un audit périodique (AWS IAM Access Analyzer) pour détecter les permissions accordées mais jamais utilisées et les révoquer.
