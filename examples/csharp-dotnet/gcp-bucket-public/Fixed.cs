// Corrigé : Bucket Google Cloud Storage public (CWE-284)
// Aucun binding IAM n'accorde d'accès à "allUsers"/"allAuthenticatedUsers" :
// l'accès en lecture est scoped à un compte de service applicatif précis,
// "uniform_bucket_level_access" et la Public Access Prevention sont activés,
// et le partage ponctuel passe par une URL signée à durée limitée.
using Google.Apis.Storage.v1.Data;
using Google.Cloud.Storage.V1;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/storage")]
public class BucketProvisioningController : ControllerBase
{
    private readonly StorageClient _storageClient;
    private readonly UrlSigner _urlSigner;

    public BucketProvisioningController(StorageClient storageClient, UrlSigner urlSigner)
    {
        _storageClient = storageClient;
        _urlSigner = urlSigner;
    }

    [HttpPost("buckets")]
    public IActionResult CreateBucket([FromQuery] string bucketName = "my-app-bucket", string projectId = "my-project")
    {
        var bucket = new Bucket
        {
            Name = bucketName,
            IamConfiguration = new Bucket.IamConfigurationData
            {
                UniformBucketLevelAccess = new Bucket.IamConfigurationData.UniformBucketLevelAccessData
                {
                    Enabled = true
                },
                PublicAccessPrevention = "enforced"
            }
        };

        _storageClient.CreateBucket(projectId, bucket);

        // Accès en lecture restreint à un compte de service applicatif précis,
        // jamais à "allUsers" ou "allAuthenticatedUsers".
        var policy = _storageClient.GetBucketIamPolicy(bucketName);
        policy.Bindings.Add(new Policy.BindingsData
        {
            Role = "roles/storage.objectViewer",
            Members = new List<string> { "serviceAccount:app-reader@my-project.iam.gserviceaccount.com" }
        });
        _storageClient.SetBucketIamPolicy(bucketName, policy);

        return Ok(new { bucketName, message = "Bucket créé en accès privé avec Public Access Prevention." });
    }

    [HttpGet("buckets/{bucketName}/objects/{objectName}/share-link")]
    public IActionResult GetSignedUrl(string bucketName, string objectName)
    {
        // Partage ponctuel sécurisé via une URL signée expirant dans 30 minutes.
        string url = _urlSigner.Sign(
            bucketName,
            objectName,
            TimeSpan.FromMinutes(30),
            HttpMethod.Get);

        return Ok(new { url, expiresInMinutes = 30 });
    }
}
