// CWE-284 : correction — le bucket est créé privé, "Block Public Access" est
// activé explicitement, et le partage ponctuel passe par une URL pré-signée
// à durée de vie limitée plutôt que par un accès public permanent.

use aws_sdk_s3::Client;
use aws_sdk_s3::types::{BucketCannedAcl, PublicAccessBlockConfiguration};
use aws_sdk_s3::presigning::PresigningConfig;
use std::time::Duration;

async fn provisionner_bucket_exports(client: &Client, bucket: &str) -> Result<(), aws_sdk_s3::Error> {
    // Sécurisé : ACL privée par défaut (aucun accès public implicite).
    client
        .create_bucket()
        .bucket(bucket)
        .acl(BucketCannedAcl::Private)
        .send()
        .await?;

    // Sécurisé : "Block Public Access" activé intégralement au niveau du
    // bucket, ce qui neutralise toute ACL ou policy publique appliquée par
    // erreur ultérieurement.
    client
        .put_public_access_block()
        .bucket(bucket)
        .public_access_block_configuration(
            PublicAccessBlockConfiguration::builder()
                .block_public_acls(true)
                .ignore_public_acls(true)
                .block_public_policy(true)
                .restrict_public_buckets(true)
                .build(),
        )
        .send()
        .await?;

    // Sécurisé : bucket policy restreinte à un rôle applicatif précis,
    // avec exigence de transport chiffré.
    let policy = format!(
        r#"{{
            "Version": "2012-10-17",
            "Statement": [{{
                "Effect": "Allow",
                "Principal": {{ "AWS": "arn:aws:iam::123456789012:role/app-read-role" }},
                "Action": ["s3:GetObject"],
                "Resource": ["arn:aws:s3:::{bucket}/*"],
                "Condition": {{ "Bool": {{ "aws:SecureTransport": "true" }} }}
            }}]
        }}"#
    );

    client
        .put_bucket_policy()
        .bucket(bucket)
        .policy(policy)
        .send()
        .await?;

    Ok(())
}

async fn generer_url_partage_temporaire(
    client: &Client,
    bucket: &str,
    key: &str,
) -> Result<String, aws_sdk_s3::Error> {
    // Sécurisé : partage ponctuel via URL pré-signée expirant après 15 min,
    // au lieu d'un accès public permanent sur le bucket.
    let presigning_config = PresigningConfig::expires_in(Duration::from_secs(15 * 60))
        .expect("configuration de pré-signature invalide");

    let presigned = client
        .get_object()
        .bucket(bucket)
        .key(key)
        .presigned(presigning_config)
        .await?;

    Ok(presigned.uri().to_string())
}
