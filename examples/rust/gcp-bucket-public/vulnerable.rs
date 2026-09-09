// CWE-284 : Improper Access Control
// L'application accorde le rôle `roles/storage.objectViewer` au principal
// `allUsers` sur le bucket GCS, rendant tous les objets lisibles par
// n'importe qui sur Internet.

use google_cloud_storage::client::Client;
use google_cloud_storage::http::buckets::iam_configuration::UniformBucketLevelAccess;
use google_cloud_storage::http::objects::get_iam_policy::GetIamPolicyRequest;
use google_cloud_storage::http::objects::set_iam_policy::SetIamPolicyRequest;

async fn rendre_bucket_public(client: &Client, bucket: &str) -> anyhow::Result<()> {
    // Vulnérable : binding IAM ajoutant `allUsers` avec un rôle de lecture,
    // sans passer par une URL signée ni restreindre l'accès.
    let mut policy = client
        .get_iam_policy(&GetIamPolicyRequest {
            resource: bucket.to_string(),
            ..Default::default()
        })
        .await?;

    policy.bindings.push(google_cloud_storage::http::Binding {
        role: "roles/storage.objectViewer".to_string(),
        members: vec!["allUsers".to_string()],
        condition: None,
    });

    client
        .set_iam_policy(&SetIamPolicyRequest {
            resource: bucket.to_string(),
            policy,
            ..Default::default()
        })
        .await?;

    Ok(())
}
