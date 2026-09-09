// CWE-284 : correction — aucun binding `allUsers`/`allAuthenticatedUsers`.
// L'accès reste restreint aux comptes de service applicatifs, et le partage
// ponctuel passe par une URL signée à durée limitée.

use google_cloud_storage::client::Client;
use google_cloud_storage::sign::{SignedURLOptions, SignedURLMethod};
use std::time::Duration;

async fn accorder_acces_service_applicatif(client: &Client, bucket: &str, service_account: &str) -> anyhow::Result<()> {
    // Sécurisé : rôle de lecture accordé uniquement au compte de service
    // applicatif précis, jamais à `allUsers`/`allAuthenticatedUsers`.
    let mut policy = client
        .get_iam_policy(&google_cloud_storage::http::objects::get_iam_policy::GetIamPolicyRequest {
            resource: bucket.to_string(),
            ..Default::default()
        })
        .await?;

    policy.bindings.push(google_cloud_storage::http::Binding {
        role: "roles/storage.objectViewer".to_string(),
        members: vec![format!("serviceAccount:{service_account}")],
        condition: None,
    });

    client
        .set_iam_policy(&google_cloud_storage::http::objects::set_iam_policy::SetIamPolicyRequest {
            resource: bucket.to_string(),
            policy,
            ..Default::default()
        })
        .await?;

    Ok(())
}

async fn generer_url_partage_temporaire(
    client: &Client,
    bucket: &str,
    object: &str,
) -> anyhow::Result<String> {
    // Sécurisé : URL signée expirant après 15 minutes pour un partage
    // ponctuel, au lieu d'un accès public permanent sur le bucket.
    let options = SignedURLOptions {
        method: SignedURLMethod::GET,
        expires: Duration::from_secs(15 * 60),
        ..Default::default()
    };

    let url = client.signed_url(bucket, object, None, None, options).await?;
    Ok(url)
}
