// CWE-284 : Improper Access Control
// L'application crée/configure le bucket S3 avec une ACL publique et une
// bucket policy ouverte à `Principal: "*"`, exposant tous les objets à
// Internet sans authentification.

use aws_sdk_s3::Client;
use aws_sdk_s3::types::BucketCannedAcl;

async fn provisionner_bucket_exports(client: &Client, bucket: &str) -> Result<(), aws_sdk_s3::Error> {
    // Vulnérable : ACL "public-read" appliquée directement au bucket qui
    // contient des exports de données applicatives (potentiellement sensibles).
    client
        .create_bucket()
        .bucket(bucket)
        .acl(BucketCannedAcl::PublicRead)
        .send()
        .await?;

    // Vulnérable : bucket policy autorisant n'importe quel principal (`*`)
    // à lister et lire tous les objets, sans condition restrictive.
    let policy = format!(
        r#"{{
            "Version": "2012-10-17",
            "Statement": [{{
                "Effect": "Allow",
                "Principal": "*",
                "Action": ["s3:GetObject", "s3:ListBucket"],
                "Resource": ["arn:aws:s3:::{bucket}", "arn:aws:s3:::{bucket}/*"]
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

async fn uploader_export(client: &Client, bucket: &str, key: &str, contenu: Vec<u8>) -> Result<(), aws_sdk_s3::Error> {
    // Vulnérable : l'objet hérite de l'ACL publique du bucket, aucun contrôle
    // additionnel n'est appliqué à l'upload.
    client
        .put_object()
        .bucket(bucket)
        .key(key)
        .body(contenu.into())
        .send()
        .await?;
    Ok(())
}
