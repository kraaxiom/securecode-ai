// CWE-269 : correction — trust policy restreinte à un principal précis avec
// `ExternalId`, et policy de permissions granulaire limitée aux actions
// réellement nécessaires au job CI (moindre privilège).

use aws_sdk_iam::Client;

async fn provisionner_role_ci(client: &Client, role_name: &str, external_id: &str) -> Result<(), aws_sdk_iam::Error> {
    // Sécurisé : trust policy limitée au provider OIDC du système CI
    // (GitHub Actions), avec un ExternalId pour empêcher le "confused
    // deputy" et restreindre l'usage au dépôt attendu.
    let trust_policy = format!(
        r#"{{
            "Version": "2012-10-17",
            "Statement": [{{
                "Effect": "Allow",
                "Principal": {{ "Federated": "arn:aws:iam::123456789012:oidc-provider/token.actions.githubusercontent.com" }},
                "Action": "sts:AssumeRoleWithWebIdentity",
                "Condition": {{
                    "StringEquals": {{
                        "token.actions.githubusercontent.com:sub": "repo:my-org/my-app:ref:refs/heads/main",
                        "sts:ExternalId": "{external_id}"
                    }}
                }}
            }}]
        }}"#
    );

    client
        .create_role()
        .role_name(role_name)
        .assume_role_policy_document(trust_policy)
        .send()
        .await?;

    // Sécurisé : policy granulaire limitée aux actions strictement
    // nécessaires au job de build (pousser une image, rien d'autre).
    let permissions_policy = r#"{
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Action": ["ecr:GetAuthorizationToken", "ecr:BatchCheckLayerAvailability", "ecr:PutImage"],
            "Resource": "arn:aws:ecr:eu-west-3:123456789012:repository/my-app"
        }]
    }"#;

    client
        .put_role_policy()
        .role_name(role_name)
        .policy_name("ci-ecr-push-only")
        .policy_document(permissions_policy)
        .send()
        .await?;

    Ok(())
}
