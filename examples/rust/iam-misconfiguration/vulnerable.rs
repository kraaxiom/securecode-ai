// CWE-269 : Improper Privilege Management
// Le service applicatif crée un rôle IAM avec des permissions wildcard
// (`Action: "*"`, `Resource: "*"`) et une trust policy ouverte à n'importe
// quel compte AWS, violant le principe du moindre privilège.

use aws_sdk_iam::Client;

async fn provisionner_role_ci(client: &Client, role_name: &str) -> Result<(), aws_sdk_iam::Error> {
    // Vulnérable : trust policy sans restriction — n'importe quel principal
    // AWS peut assumer ce rôle, sans ExternalId ni condition MFA.
    let trust_policy = r#"{
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Principal": { "AWS": "*" },
            "Action": "sts:AssumeRole"
        }]
    }"#;

    client
        .create_role()
        .role_name(role_name)
        .assume_role_policy_document(trust_policy)
        .send()
        .await?;

    // Vulnérable : policy attachée donnant un accès administrateur complet
    // à un rôle utilisé uniquement pour des tâches de build CI.
    let permissions_policy = r#"{
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Action": "*",
            "Resource": "*"
        }]
    }"#;

    client
        .put_role_policy()
        .role_name(role_name)
        .policy_name("ci-full-access")
        .policy_document(permissions_policy)
        .send()
        .await?;

    Ok(())
}
