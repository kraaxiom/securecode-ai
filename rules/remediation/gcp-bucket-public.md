# Remédiation — Bucket Google Cloud Storage public

## Principe
Retirer `allUsers`/`allAuthenticatedUsers` des bindings IAM du bucket, activer "Public Access Prevention" et `uniform_bucket_level_access`, et utiliser des URLs signées à durée limitée pour tout partage ponctuel.

## gcloud CLI
```bash
# Avant — vulnérable : accès public en lecture accordé à tous
gcloud storage buckets add-iam-policy-binding gs://my-app-bucket \
  --member=allUsers --role=roles/storage.objectViewer

# Après — sécurisé : retrait du binding public + activation de la prévention d'accès public
gcloud storage buckets remove-iam-policy-binding gs://my-app-bucket \
  --member=allUsers --role=roles/storage.objectViewer

gcloud storage buckets update gs://my-app-bucket \
  --public-access-prevention

gcloud storage buckets update gs://my-app-bucket \
  --uniform-bucket-level-access
```

## Terraform
```hcl
# Avant — vulnérable
resource "google_storage_bucket_iam_member" "public_read" {
  bucket = google_storage_bucket.my_app_bucket.name
  role   = "roles/storage.objectViewer"
  member = "allUsers"
}

# Après — sécurisé
resource "google_storage_bucket" "my_app_bucket" {
  name                        = "my-app-bucket"
  location                    = "EU"
  uniform_bucket_level_access = true
  public_access_prevention    = "enforced"
}

resource "google_storage_bucket_iam_member" "app_read" {
  bucket = google_storage_bucket.my_app_bucket.name
  role   = "roles/storage.objectViewer"
  member = "serviceAccount:app-reader@my-project.iam.gserviceaccount.com"
}
```

## Partage ponctuel sécurisé (URL signée)
```bash
# Après — sécurisé : URL signée valable 30 minutes au lieu d'un accès public permanent
gcloud storage sign-url gs://my-app-bucket/rapport.pdf \
  --private-key-file=service-account-key.json --duration=30m
```

## Checklist de vérification post-patch
- [ ] Plus aucun binding IAM du bucket ne contient `allUsers` ou `allAuthenticatedUsers`.
- [ ] "Public Access Prevention" est activé (`enforced`) sur le bucket, idéalement au niveau organisation.
- [ ] `uniform_bucket_level_access` est activé pour éviter les ACL d'objet incohérentes.
- [ ] Les rôles IAM accordés sont scoped à des service accounts précis, pas à des rôles larges.
- [ ] Un accès public anonyme à un objet du bucket échoue (403 AccessDenied).
- [ ] Les partages ponctuels utilisent des URLs signées avec une durée de vie courte documentée.
