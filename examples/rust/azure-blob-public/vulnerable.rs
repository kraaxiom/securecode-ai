// CWE-284 : Improper Access Control
// Le conteneur Azure Blob Storage est créé avec un niveau d'accès public
// "Blob", et un jeton SAS est généré sans expiration ni restriction de
// permissions, exposant le contenu à Internet.

use azure_storage::prelude::*;
use azure_storage_blobs::prelude::*;
use time::{Duration, OffsetDateTime};

async fn provisionner_conteneur_documents(
    service_client: &BlobServiceClient,
    conteneur: &str,
) -> azure_core::Result<()> {
    let container_client = service_client.container_client(conteneur);

    // Vulnérable : accès public au niveau "Blob" — tout objet est lisible
    // par une URL directe sans authentification.
    container_client
        .create()
        .public_access(PublicAccess::Blob)
        .await?;

    Ok(())
}

fn generer_sas_partage(blob_client: &BlobClient) -> azure_core::Result<String> {
    // Vulnérable : SAS avec permissions larges (lecture + écriture + suppression)
    // et une expiration extrêmement lointaine, équivalent à un accès permanent.
    let expiry = OffsetDateTime::now_utc() + Duration::days(3650);
    let permissions = BlobSasPermissions {
        read: true,
        write: true,
        delete: true,
        ..Default::default()
    };

    let sas = blob_client.shared_access_signature(permissions, expiry)?;
    Ok(sas.token())
}
