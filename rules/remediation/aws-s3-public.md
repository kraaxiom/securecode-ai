# Remédiation — Bucket S3 public

## Principe
Bloquer tout accès public par défaut au niveau du compte et de chaque bucket, restreindre la bucket policy au principe du moindre privilège, et utiliser des URLs pré-signées à durée limitée pour tout partage ponctuel.

## AWS CLI
```bash
# Avant — vulnérable : Block Public Access désactivé
aws s3api put-public-access-block \
  --bucket my-app-bucket \
  --public-access-block-configuration \
  BlockPublicAcls=false,IgnorePublicAcls=false,BlockPublicPolicy=false,RestrictPublicBuckets=false

# Après — sécurisé : Block Public Access activé intégralement
aws s3api put-public-access-block \
  --bucket my-app-bucket \
  --public-access-block-configuration \
  BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true
```

## Bucket policy (JSON)
```json
// Avant — vulnérable
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": "*",
    "Action": ["s3:GetObject", "s3:ListBucket"],
    "Resource": ["arn:aws:s3:::my-app-bucket", "arn:aws:s3:::my-app-bucket/*"]
  }]
}

// Après — sécurisé : accès restreint à un rôle applicatif précis
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": { "AWS": "arn:aws:iam::123456789012:role/app-read-role" },
    "Action": ["s3:GetObject"],
    "Resource": ["arn:aws:s3:::my-app-bucket/*"],
    "Condition": { "Bool": { "aws:SecureTransport": "true" } }
  }]
}
```

## Terraform
```hcl
# Avant — vulnérable
resource "aws_s3_bucket_acl" "bucket_acl" {
  bucket = aws_s3_bucket.my_app_bucket.id
  acl    = "public-read"
}

# Après — sécurisé
resource "aws_s3_bucket_public_access_block" "my_app_bucket" {
  bucket                  = aws_s3_bucket.my_app_bucket.id
  block_public_acls       = true
  ignore_public_acls      = true
  block_public_policy     = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_server_side_encryption_configuration" "my_app_bucket" {
  bucket = aws_s3_bucket.my_app_bucket.id
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "aws:kms"
    }
  }
}
```

## Partage ponctuel sécurisé (URL pré-signée)
```bash
# Après — sécurisé : accès temporaire de 15 minutes au lieu d'un accès public permanent
aws s3 presign s3://my-app-bucket/rapport.pdf --expires-in 900
```

## Checklist de vérification post-patch
- [ ] "Block Public Access" est activé sur les 4 paramètres, au niveau du compte ET du bucket.
- [ ] La bucket policy ne contient plus `"Principal": "*"` sans condition restrictive.
- [ ] Aucune ACL de bucket ou d'objet n'est réglée sur `public-read`/`public-read-write`.
- [ ] Le chiffrement au repos (SSE-KMS ou SSE-S3) est activé sur le bucket.
- [ ] Un accès public tenté depuis un navigateur anonyme échoue (403/Access Denied).
- [ ] Les partages ponctuels utilisent des URLs pré-signées avec une expiration courte documentée.
- [ ] AWS Config / Access Analyzer ne remonte plus d'alerte "bucket public" sur ce bucket.
