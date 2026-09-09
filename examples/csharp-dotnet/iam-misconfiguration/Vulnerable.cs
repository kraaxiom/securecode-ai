// Vulnérable : Mauvaise configuration IAM (CWE-269)
// La policy créée accorde toutes les actions sur toutes les ressources
// ("Action": "*", "Resource": "*"), violant le principe du moindre privilège.
using Amazon.IdentityManagement;
using Amazon.IdentityManagement.Model;

public class IamProvisioningService
{
    private readonly IAmazonIdentityManagementService _iamClient;

    public IamProvisioningService(IAmazonIdentityManagementService iamClient)
        => _iamClient = iamClient;

    public async Task ProvisionAppRoleAsync()
    {
        // VULNÉRABLE : wildcard total sur les actions et les ressources.
        const string policyDocument = """
        {
          "Version": "2012-10-17",
          "Statement": [{
            "Effect": "Allow",
            "Action": "*",
            "Resource": "*"
          }]
        }
        """;

        // VULNÉRABLE : trust policy sans restriction de principal,
        // n'importe quel compte AWS peut assumer ce rôle.
        const string trustPolicy = """
        {
          "Version": "2012-10-17",
          "Statement": [{
            "Effect": "Allow",
            "Principal": { "AWS": "*" },
            "Action": "sts:AssumeRole"
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
            PolicyName = "app-full-access",
            PolicyDocument = policyDocument
        });
    }
}
