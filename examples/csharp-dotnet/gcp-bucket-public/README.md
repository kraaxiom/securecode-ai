# Bucket Google Cloud Storage public (CWE-284)

La version vulnérable accorde le rôle `roles/storage.objectViewer` au principal `allUsers` via l'API IAM du bucket, rendant tous les objets lisibles publiquement sans authentification. La correction retire ce binding public, active `uniform_bucket_level_access` ainsi que la "Public Access Prevention", n'accorde l'accès qu'à un compte de service applicatif précis, et génère une URL signée à durée limitée (30 minutes) pour tout partage ponctuel.
