# Conteneur Azure Blob Storage public (CWE-284)

La version vulnérable crée un conteneur Azure Blob Storage avec `PublicAccessType.BlobContainer`, rendant tous les blobs listables et lisibles sans authentification. La correction crée le conteneur en accès privé (`PublicAccessType.None`) et utilise un SAS (Shared Access Signature) en lecture seule, à permissions minimales et expirant après une heure, pour tout partage ponctuel de blob.
