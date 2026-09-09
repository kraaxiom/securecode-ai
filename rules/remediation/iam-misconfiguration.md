# Remédiation — Mauvaise configuration IAM (permissions excessives)

## Principe
Remplacer toute policy IAM utilisant des wildcards larges (`Action: *`, `Resource: *`) par une policy granulaire listant précisément les actions et ressources nécessaires, appliquer le principe du moindre privilège, et préférer les identités temporaires aux clés statiques.

## AWS IAM (JSON)
```json
// Avant — vulnérable : policy avec wildcard total
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Action": "*",
    "Resource": "*"
  }]
}

// Après — sécurisé : permissions scoped à l'action et la ressource nécessaires
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Action": [
      "s3:GetObject",
      "s3:PutObject"
    ],
    "Resource": "arn:aws:s3:::my-app-bucket/uploads/*",
    "Condition": {
      "StringEquals": { "aws:RequestedRegion": "eu-west-1" }
    }
  }]
}
```

## Trust policy (AssumeRole)
```json
// Avant — vulnérable : n'importe quel compte peut assumer le rôle
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": { "AWS": "*" },
    "Action": "sts:AssumeRole"
  }]
}

// Après — sécurisé : compte précis + ExternalId + MFA requis
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": { "AWS": "arn:aws:iam::123456789012:root" },
    "Action": "sts:AssumeRole",
    "Condition": {
      "StringEquals": { "sts:ExternalId": "unique-partner-id-2026" },
      "Bool": { "aws:MultiFactorAuthPresent": "true" }
    }
  }]
}
```

## Terraform (rôle applicatif au lieu d'un rôle Owner/Admin)
```hcl
# Avant — vulnérable
resource "google_project_iam_member" "ci_sa" {
  project = "my-project"
  role    = "roles/owner"
  member  = "serviceAccount:ci-deploy@my-project.iam.gserviceaccount.com"
}

# Après — sécurisé : rôle granulaire limité au besoin réel du CI
resource "google_project_iam_member" "ci_sa" {
  project = "my-project"
  role    = "roles/cloudbuild.builds.editor"
  member  = "serviceAccount:ci-deploy@my-project.iam.gserviceaccount.com"
}
```

## Checklist de vérification post-patch
- [ ] Plus aucune policy IAM ne combine `Action: *` et `Resource: *`.
- [ ] Les rôles applicatifs/service accounts n'ont plus de rôle "Owner"/"Admin"/"*Contributor" par défaut.
- [ ] Les trust policies (`AssumeRole`) incluent une condition (`ExternalId`, MFA, source IP) et un principal précis.
- [ ] Les clés d'accès statiques ont été remplacées par des rôles temporaires (STS, Workload Identity Federation) quand possible.
- [ ] Un audit des permissions inutilisées (IAM Access Analyzer / Policy Analyzer) confirme l'absence de droits superflus.
- [ ] Les changements ont été testés pour vérifier que les fonctionnalités légitimes continuent de fonctionner avec les permissions réduites.
