# Bucket S3 public (CWE-284)

La version vulnérable crée un bucket S3 via le SDK AWS et lui applique une ACL `public-read` ainsi qu'une bucket policy avec `Principal: *`, rendant tous les objets accessibles sans authentification. La correction supprime toute ACL/policy publique, active le `PutPublicAccessBlock` avec les quatre restrictions à `true`, et génère une URL pré-signée à durée limitée (15 minutes) pour tout partage ponctuel d'objet.
