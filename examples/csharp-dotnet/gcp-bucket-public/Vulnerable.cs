// Vulnérable : Bucket Google Cloud Storage public (CWE-284)
// Le binding IAM du bucket accorde le rôle "storage.objectViewer" au
// principal "allUsers", rendant tous les objets lisibles publiquement
// sans authentification.
using Google.Apis.Auth.OAuth2;
using Google.Cloud.Storage.V1;
using Microsoft.AspNetCore.Mvc;

[ApiController]
[Route("api/storage")]
public class BucketProvisioningController : ControllerBase
{
    private readonly StorageClient _storageClient;

    public BucketProvisioningController(StorageClient storageClient)
    {
        _storageClient = storageClient;
    }

    [HttpPost("buckets")]
    public IActionResult CreateBucket([FromQuery] string bucketName = "my-app-bucket", string projectId = "my-project")
    {
        _storageClient.CreateBucket(projectId, bucketName);

        // Ajout d'un accès public en lecture à "allUsers" au niveau du bucket.
        var bucket = _storageClient.GetBucket(bucketName);
        bucket.IamConfiguration ??= new Google.Apis.Storage.v1.Data.Bucket.IamConfigurationData();

        var policy = _storageClient.GetBucketIamPolicy(bucketName);
        policy.Bindings.Add(new Google.Apis.Storage.v1.Data.Policy.BindingsData
        {
            Role = "roles/storage.objectViewer",
            Members = new List<string> { "allUsers" }
        });
        _storageClient.SetBucketIamPolicy(bucketName, policy);

        return Ok(new { bucketName, publicUrl = $"https://storage.googleapis.com/{bucketName}/" });
    }
}
