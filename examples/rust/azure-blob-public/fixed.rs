// CWE-284 : correction — conteneur créé privé, et tout partage ponctuel
// passe par un SAS à permissions minimales (lecture seule) et à durée de
// vie courte plutôt qu'un accès public permanent.

use azure_storage::prelude::*;
use azure_storage_blobs::prelude::*;
use time::{Duration, OffsetDateTime};

async fn provisionner_conteneur_documents(
    service_client: &BlobServiceClient,
    conteneur: &str,
) -> azure_core::Result<()> {
    let container_client = service_client.container_client(conteneur);

    // Sécurisé : accès public désactivé (`PublicAccess::None`) — le
    // conteneur n'est accessible qu'via authentification (clé de compte,
    // Azure AD ou SAS scoped).
    container_client
        .create()
        .public_access(PublicAccess::None)
        .await?;

    Ok(())
}

fn generer_sas_partage(blob_client: &BlobClient) -> azure_core::Result<String> {
    // Sécurisé : SAS en lecture seule, expirant après 15 minutes — le strict
    // nécessaire pour un partage ponctuel.
    let expiry = OffsetDateTime::now_utc() + Duration::minutes(15);
    let permissions = BlobSasPermissions {
        read: true,
        ..Default::default()
    };

    let sas = blob_client.shared_access_signature(permissions, expiry)?;
    Ok(sas.token())
}
