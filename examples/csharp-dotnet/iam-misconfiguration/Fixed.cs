// Corrigé : Mauvaise configuration IAM (CWE-269)
// La policy liste précisément les actions et la ressource nécessaires
// (principe du moindre privilège), et la trust policy restreint
// l'assomption du rôle à un principal précis avec ExternalId.
using Amazon.IdentityManagement;
using Amazon.IdentityManagement.Model;

public class IamProvisioningService
{
    private readonly IAmazonIdentityManagementService _iamClient;

    public IamProvisioningService(IAmazonIdentityManagementService iamClient)
        => _iamClient = iamClient;

    public async Task ProvisionAppRoleAsync()
    {
        // Policy scoped : uniquement les actions S3 nécessaires, sur le bucket
        // applicatif dédié — aucun accès aux autres services ou ressources.
        const string policyDocument = """
        {
          "Version": "2012-10-17",
          "Statement": [{
            "Effect": "Allow",
            "Action": [
              "s3:GetObject",
              "s3:PutObject"
            ],
            "Resource": "arn:aws:s3:::my-app-bucket/uploads/*",
            "Condition": {
              "StringEquals": { "aws:RequestedRegion": "eu-west-1" }
            }
          }]
        }
        """;

        // Trust policy restreinte : un seul compte AWS précis peut assumer
        // le rôle, avec ExternalId et MFA requis.
        const string trustPolicy = """
        {
          "Version": "2012-10-17",
          "Statement": [{
            "Effect": "Allow",
            "Principal": { "AWS": "arn:aws:iam::123456789012:root" },
            "Action": "sts:AssumeRole",
            "Condition": {
              "StringEquals": { "sts:ExternalId": "REPLACE_ME_UNIQUE_EXTERNAL_ID" },
              "Bool": { "aws:MultiFactorAuthPresent": "true" }
            }
          }]
        }
        """;

        await _iamClient.CreateRoleAsync(new CreateRoleRequest
        {
            RoleName = "app-service-role",
            AssumeRolePolicyDocument = trustPolicy
        });

        await _iamClient.PutRolePolicyAsync(new PutRolePolicyRequest
        {
            RoleName = "app-service-role",
            PolicyName = "app-scoped-s3-access",
            PolicyDocument = policyDocument
        });
    }
}
